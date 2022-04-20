# NATS Server Set Up Using Decentralised JWT

We will cover the following in this document:

- Operator managed deployment of nats-server

- NATS based resolver configuration

- Using `nsc` to create, push and pull accounts to and from the server

- Creating dynamic accounts and users using go client

## Prerequisites

- [nsc](https://github.com/nats-io/nsc)
- [nats-server](https://docs.nats.io/running-a-nats-service/introduction/installation)

## Set-up operator, account and user

- Create an operator with a SYS account that will be the root of trust for our nats-server:

    ``` bash
    nsc add operator --sys -n op
    ```

- Create a couple of accounts:

    ``` bash
    nsc add account --name admin
    nsc add account --name restricted
    ```

- Create an user for each account:

    ``` bash
    nsc add user --name admin-user --account admin
    nsc add user --name restricted-user --account restricted
    # Create an user for SYS account as well (to be used later in go code):
    nsc add user --name sys-user --account SYS
    ```

## Generate server config using nsc

We will use above created operator to generate server config:

``` bash
nsc generate config --nats-resolver --sys-account SYS > server.conf
```

The above command will write the config to `server.conf`. We will also add resolver block to the config:

``` yaml
// Operator "op"
operator: eyJ0eXAiOiJKV1QiLCJhbGciOiJlZDI1NTE5LW5rZXkifQ.eyJqdGkiOiJCNTQ0U1BPQkEyR0I3Qk5XNUhCR0ZVRzI3NVRZV0xCRkRaRUdJQlNLNzdWVlE1QkRFNjRBIiwiaWF0IjoxNjQzMzA2OTc4LCJpc3MiOiJPQUhUS1RXN0M3TVFFNDZXMkg3STRIQkxUUFc1N1VPNUZBNEVHSE9UNFU1TUVFSzRNQzVWWUpRUyIsIm5hbWUiOiJzZXJ2ZXIxLW9wIiwic3ViIjoiT0FIVEtUVzdDN01RRTQ2VzJIN0k0SEJMVFBXNTdVTzVGQTRFR0hPVDRVNU1FRUs0TUM1VllKUVMiLCJuYXRzIjp7InN5c3RlbV9hY2NvdW50IjoiQUJNTlBMQk5GWVY3Vk1QUkdLSFdaTkZMNjZBWk1KSFVKSjJGRDRPT1FLRlJOTkNLM0lZQVNPWDQiLCJ0eXBlIjoib3BlcmF0b3IiLCJ2ZXJzaW9uIjoyfX0.q_SbINQUTZtLf6JlAiMT_J_V2oK057e4qV3riaJeW3e5mmqy0qAn_fPpK1RNXZPxKilG5G3QAF9JGpQJ6HF9BA

system_account: ABMNPLBNFYV7VMPRGKHWZNFL66AZMJHUJJ2FD4OOQKFRNNCK3IYASOX4

resolver: {
  # full means all the accounts JWTs will be stored on the server. The other type is cache which is based on LRU.
  type: full
  # Local directory to store JWTs. It is not shared b/w other servers.
  dir: './jwt'
  allow_delete: true
  interval: "2m"
  limit: 500
}
```

## Start the nats-server

nats-server can be started using above generated config:

``` bash
nats-server -c resolver.conf
```
