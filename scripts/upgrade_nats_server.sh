#!/bin/sh

NATS_SERVER_VERSION=2.9.3

# Change the directory
cd /opt

# Download the binary
# wget https://github.com/nats-io/nats-server/releases/download/v2.9.3/nats-server-v2.9.3-linux-amd64.tar.gz
wget https://github.com/nats-io/nats-server/releases/download/v${NATS_SERVER_VERSION}/nats-server-v${NATS_SERVER_VERSION}-linux-amd64.tar.gz

# Untar
# tar -xvf nats-server-v2.9.3-linux-amd64.tar.gz
tar -xvf nats-server-v${NATS_SERVER_VERSION}-linux-amd64.tar.gz

# Stop the server
supervisorctl stop nats

# Delete the symlink
rm nats-server

# Create symlink
# ln -s nats-server-v2.9.3-linux-amd64/nats-server nats-server
ln -s nats-server-v${NATS_SERVER_VERSION}-linux-amd64/nats-server nats-server

# Start the server
supervisorctl start nats

# v2.9.4
cd /opt
wget https://github.com/nats-io/nats-server/releases/download/v2.9.4/nats-server-v2.9.4-linux-amd64.tar.gz
tar -xvf nats-server-v2.9.4-linux-amd64.tar.gz
supervisorctl stop nats
rm nats-server
ln -s nats-server-v2.9.4-linux-amd64/nats-server nats-server
supervisorctl start nats