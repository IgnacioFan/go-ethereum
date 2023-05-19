# Go Ethereum
The Go Ethereum provides API and blocks indexer services for interacting with the Ethereum blockchain. The application enables users to retrieve a single block, m limit blocks, and transactions with event logs by REST API. Additionally, the application includes a block indexer service that scans blocks into the database concurrently via Web 3 RPC API, allowing for constant updates to the state of Ethereum blocks.

## Installation and Setup

### Prerequisite
- Go 1.16+
- Docker, the project is based on `docker-compose.yml` to boot up all runnable services

### Run all services

1. Clone the repository to your local machine.
2. Ensure that Docker is installed and running on your machine.
3. Create a .env file

```bash
make gen.env
```

4. Pickup an operaton that you want
```bash
# to start all services
make app.start

# to close all services
make app.close

# to access db container
make db.cli
```

5. Test the following API endpoints
  - Make sure you have started out the services, before sending requests!
  - Endpoint: /blocks?limit=n
  ```bash
  curl --url http://localhost:8080/blocks
  curl --url http://localhost:8080/blocks?limit=10
  ```

  - Endpoint: /blocks/:block_number
  ```bash
  curl --url http://localhost:8080/blocks/12345
  ```

  - Endpoint: /tansaction/:txHash
  ```bash
  curl --url http://localhost:8080/transaction/0x9419ac0239693bd86619906da85e921fe8f1bf4a21fb53274daea5e73e64f669
  ```

## Contributors
- Weilong Fan (IgnacioFan): developer and maintainer
