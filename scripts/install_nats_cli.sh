#!/bin/sh

# Usage
# bash install_nats_cli.sh <nats_cli_version>

VERSION=0.0.34

wget https://github.com/nats-io/natscli/releases/download/v${VERSION}/nats-${VERSION}-linux-amd64.zip

unzip nats-${VERSION}-linux-amd64.zip

cp nats-${VERSION}-linux-amd64/nats /usr/local/bin

# remove downloaded artifacts
rm -rf nats-${VERSION}-linux-amd64
rm -rf nats-${VERSION}-linux-amd64.zip
