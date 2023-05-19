#!/bin/bash

# Prompt the user for environment variables
read -p "Enter an expose port number for API service: " EXPOSE_API_PORT
read -p "Enter an Ethereum client URL for scanning blocks: " ETH_CLIENT_URL

ENV_CONTENT=$(cat <<EOF
# DB
POSTGRES_HOST=go-ethereum-db
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=go-ethereum
POSTGRES_PORT=5432

# API
EXPOSE_API_PORT=$EXPOSE_API_PORT

# Eth client URL
ETH_CLIENT_URL=$ETH_CLIENT_URL

EOF
)
echo "$ENV_CONTENT" > .env

echo "Env variables and .env file generated successfully!"
echo "Now, you can edit the env variables to fit your needs!"
