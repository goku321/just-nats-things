package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"

	pb "github.com/nutanix-beam/cpaas-nats-accounts/grpc/go/account"
)

const (
	accLookupReqSubj = "$SYS.REQ.ACCOUNT.%s.CLAIMS.LOOKUP"
	accListReqSubj   = "$SYS.REQ.CLAIMS.LIST"
	accClaimsReqSubj = "$SYS.REQ.CLAIMS.UPDATE"
)

type accountMap map[string]*jwt.AccountClaims

func getAccountsList(logger *log.Logger, nc *nats.Conn) []string {
	ib := nats.NewInbox()
	sub, err := nc.SubscribeSync(ib)
	if err != nil {
		logger.Fatalf("get_account_list: failed to subscribe to %s: %s", ib, err)
	}
	if err := nc.PublishRequest(accListReqSubj, ib, nil); err != nil {
		logger.Fatalf("failed to pull accounts: %s", err)
	}

	if resp, err := sub.NextMsg(time.Second * 5); err != nil {
		logger.Fatalf("failed to get response to pull: %s", err)
	} else if msg := string(resp.Data); msg == "" { // empty response means end
	} else if tk := strings.Split(string(resp.Data), "|"); len(tk) != 2 {
		var natsResp struct {
			Data []string `json:"data"`
		}
		if err = json.Unmarshal(resp.Data, &natsResp); err != nil {
			logger.Fatalf("error unmarshalling accounts list: %s", err)
		}

		return natsResp.Data
	}
	return nil
}

func getAccountClaims(logger *log.Logger, nc *nats.Conn, name string) *jwt.AccountClaims {
	ib := nats.NewInbox()
	sub, err := nc.SubscribeSync(ib)
	if err != nil {
		logger.Fatalf("get_account_claims: failed to subscribe: %s", err)
	}

	accountLookupSubject := fmt.Sprintf(accLookupReqSubj, name)
	if err := nc.PublishRequest(accountLookupSubject, ib, nil); err != nil {
		logger.Fatalf("get_account_claims: failed to pull accounts: %s", err)
	}

	if resp, err := sub.NextMsg(time.Second * 5); err != nil {
		logger.Printf("get_account_claims: [%s] failed to get response to pull: %s", name, err)
	} else if msg := string(resp.Data); msg == "" { // empty response means end
	} else if tk := strings.Split(string(resp.Data), "|"); len(tk) != 2 {
		ac, err := jwt.DecodeAccountClaims(string(resp.Data))
		if err != nil {
			logger.Printf("error decoding account claims for %s: %s", name, err)
		}
		return ac
	}
	return nil
}

func updateAccountClaims(logger *log.Logger, nc *nats.Conn, claims *jwt.AccountClaims) error {
	// Load operator
	kp, err := nkeys.FromSeed([]byte(""))
	if err != nil {
		logger.Fatalf("Could not decode signing key: %v", err)
	}
	encodedJWT, err := claims.Encode(kp)
	if err != nil {
		return fmt.Errorf("error encoding jwt: %s", err)
	}

	if err = nc.Publish(accClaimsReqSubj, []byte(encodedJWT)); err != nil {
		return fmt.Errorf("error publishing to %s: %s", accClaimsReqSubj, err)
	}
	return nil
}

// deleteImports calls DeleteImports API to delete imports
func deleteImports(natsManagerClient pb.AccountAPIClient, accountName string, toDelete []string) error {
	req := &pb.DeleteImportsRequest{
		Cluster:     "<cluster-name>",
		AccountName: accountName,
		Names:       toDelete,
	}
	_, err := natsManagerClient.DeleteImports(context.Background(), req)
	return err
}
