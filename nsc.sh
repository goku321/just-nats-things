nsc env --operator falcon

nsc add account a
nsc edit account --name a --js-consumer -1 --js-disk-storage -1 --js-mem-storage -1 --js-streams -1

nsc add account b
nsc edit account --name b --js-consumer -1 --js-disk-storage -1 --js-mem-storage -1 --js-streams -1

# add exports
for i in {1..10}
do
    nsc add export --name exp$i --subject a.test$i --account  a --service
done

# add imports
for i in {1..10}
do
    nsc add import --account b \
                --src-account $(nsc list accounts 2>&1 | awk '$2 == "a" {print $0}' | awk '{print $4}') \
                --remote-subject a.test$i --service --local-subject test$i
done

# create a jetstream enabled account
nsc add account js-account
nsc edit account --name js-account --js-consumer -1 --js-disk-storage -1 --js-mem-storage -1 --js-streams -1

# create a user
nsc add user js-account-user

# push the account to the server
nsc push

# create a stream
nats stream add stream1 --subjects "stream1.*" --ack --max-msgs=-1 --max-bytes=-1 --max-age=1y --storage file --retention limits --max-msg-size=-1 --discard=old

# publish some messages
nats pub stream1.test hellow --creds
# nats consumer add ORDERS NEW --filter ORDERS.received --ack explicit --pull --deliver all --max-deliver=-1 --sample 100
# nats consumer add ORDERS DISPATCH --filter ORDERS.processed --ack explicit --pull --deliver all --max-deliver=-1 --sample 100
# nats consumer add ORDERS MONITOR --filter '' --ack none --target monitor.ORDERS --deliver last --replay instant