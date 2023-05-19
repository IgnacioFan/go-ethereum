#!/bin/bash

echo "To start the Go Ethereum API service, you have to"
read -p "Enter a port number to expose the service to your local machine: " EXPOSE_API_PORT
echo "To run the Go Ethereum block indexer service, you have to"
read -p "Enter an Ethereum client URL to access the Ethereum node: " ETH_CLIENT_URL

ENV_CONTENT=$(cat <<EOF
# DB
POSTGRES_HOST=go-ethereum-db
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=go-ethereum
POSTGRES_PORT=5432

# API
API_HOST=go-ethereum-api
EXPOSE_API_PORT=$EXPOSE_API_PORT

# Indexer
INDEXER_HOST=go-ethereum-indexer
ETH_CLIENT_URL=$ETH_CLIENT_URL
INDEXER_START_BLOCK=17292710
INDEXER_SCAN_WINDOW=5
INDEXER_END_BLOCK=0
INDEXER_SLEEP_SECS=5

EOF
)
echo "$ENV_CONTENT" > .env

echo "Env variables and .env file generated successfully!"
echo "Now, you can edit the .env file to fit your needs!"
