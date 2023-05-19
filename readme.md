# Go Ethereum
The Go Ethereum provides API and blocks indexer services for interacting with the Ethereum blockchain. The application enables users to retrieve a single block, m limit blocks, and transactions with event logs by REST API. Additionally, the application includes a block indexer service that scans blocks into the database concurrently via Web 3 RPC API, allowing for constant updates to the state of Ethereum blocks.

## Installation and Setup

### Prerequisite
- Go 1.16+
- Docker, the project is based on `docker-compose.yml` to boot up all runnable services

### How to use

1. Clone the repository to your local machine.
2. Ensure that Docker is installed and running on your machine.
3. Create a .env file and do some settings.
  - For example, provide the Ethereum client URL to connect the Ethereum node.
  - Please follow the prompt instruction!

```bash
make gen.env
```

4. Run `make app.start`
5. Test the following API endpoints
  - Make sure you have started all services, before sending requests!
  - And please wait for a few secs to let the indexer do its job.

```bash
# Endpoint: /blocks?limit=n
curl --url http://localhost:8080/blocks
curl --url http://localhost:8080/blocks?limit=5

# Endpoint: /blocks/:id
curl --url http://localhost:8080/blocks/1234..

# Endpoint: /tansaction/:txHash
curl --url http://localhost:8080/transaction/0xf754c..
```
6. Run `make app.stop` to clean up the containers

End! To know more about other executable commands, please check out the Makefile.

## System design
### Architecture
The system consists of the following components:

- API service: Handles 3 kinds of use cases.
  - GET /blocks?limit=n
  - GET /blocks/:id
  - GET /transaction/:txHash

- Ethereum block indexer service: Scan blocks into the database concurrently.
  - Use 4 flags (start, window, end, sleep) to regulate the rate of block scanning.
  - The start flag determines where the block starts.
  - The window flag controls the number of concurrent scans.
  - The end flag is an optional flag to decide if the indexer should keep searching for new blocks.
  - The sleep flag sets the rate of block scan to prevent the services from being throttled or seen as DOS.

- DB Service: Use Postgres as the relational database.

### Key Features
- The system is designed as a command line base structure.
  - The `http` command to run the HTTP server and schema migrations.
  - The `index` command to run the block indexer service.

- API service is an HTTP service.
  - Mainly based on Gin, HTTP framework and Gorm lib as ORM.

- Ethereum block indexer service:
  - Rely on `sync.WaitGroup` and `goroutine` to achieve concurrently scan.
  - Use `go-ethereum/ethclient` to interact with Web3 API RPC endpoints.
  - RPC endpoints are provided by INFURA.

## Contributors
- Weilong Fan (IgnacioFan): developer and maintainer
