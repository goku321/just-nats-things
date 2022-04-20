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

The above command will write the config to `server.conf`. `--nats-resolver` flag enables NATS based resolver for account lookup.

## Start the nats-server

nats-server can be started using above generated config:

``` bash
nats-server -c resolver.conf
```
