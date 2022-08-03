# NATS Server Set Up Using Decentralised JWT

We will cover the following in this document:

- Operator managed deployment of nats-server

- NATS based resolver configuration

- Using `nsc` to create, push and pull accounts to and from the server

- Creating dynamic accounts and users using go client

## Prerequisites

- [nsc](https://github.com/nats-io/nsc)
- [nats-server](https://docs.nats.io/running-a-nats-service/introduction/installation)
- [nats CLI](https://github.com/nats-io/natscli)

## Set-up operator and system account

Create an operator "op" (with a SYS account) that will be the root of trust for our nats-server:

``` bash
nsc add operator --sys -n op
```

## Generate server config using nsc

We will use above created operator to generate server config:

``` bash
nsc generate config --nats-resolver --sys-account SYS > server.conf
```

The above command will write the config to `server.conf`. `--nats-resolver` flag enables NATS based resolver for account lookup.

## Start the nats-server

nats-server can be started using above generated config:

``` bash
nats-server -c server.conf
```

## Create an account and an user for pub/sub

If you try to do pub/sub at this point, you will get "Authorization Violation" error from the nats-server. Let's create an account and an user to do pub/sub.

- Edit operator to add nats-server URL (Default context will be the recently created operator):

    ``` bash
    nsc edit operator --account-jwt-server-url nats://0.0.0.0:4222
    ```

- Create account:

    ``` bash
    nsc add account --name a
    ```

- Push account to the nats-server:

    ``` bash
    nsc push -a a
    ```

- Create an user for account "a":

    ``` bash
    nsc add user --name u1 --account a
    ```

## Pub/Sub

When a user is created using nsc, it will print out the path for the creds file. We'll use those creds to publish a message.

- Publish a message

    ``` bash
    nats pub sub.test hello --creds <path/to/creds/file>
    ```
