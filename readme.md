# Go Ethereum
The Go Ethereum provides API and blocks indexer services for interacting with the Ethereum blockchain. The application enables users to retrieve a single block, m limit blocks, and transactions with event logs by REST API. Additionally, the application includes a block indexer service that scans blocks into the database concurrently via Web 3 RPC API, allowing for constant updates to the state of Ethereum blocks.

[Demonstration]

## Installation and Setup

### Prerequisite
- Go 1.16+
- Docker, we use `docker-compose` to boot up all required services

### Boot up HTTP server

1. Clone the repository to your local machine.
2. Ensure that Docker is installed and running on your machine.

```
// start all services
docker compose up

// stop all services
docker compose down
```
3. Access the application at http://localhost

## How to use?

## System design

### APIs

### System Architecture

API service
Ethereum block indexer service

## Test

## What's next?

## Contributors
- Weilong Fan (IgnacioFan): developer and maintainer

## References
