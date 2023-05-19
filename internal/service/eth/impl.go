package eth

import (
	"context"
	"encoding/json"
	"go-ethereum/internal/entity"
	"go-ethereum/internal/repository"
	"go-ethereum/internal/repository/eth_repo"
	"go-ethereum/internal/service"
	"go-ethereum/pkg/postgres"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "test"
)

type Impl struct {
	Client *ethclient.Client
	Repo   repository.Eth
}

func NewService(db *postgres.Postgres) (service.Eth, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Impl{
		Client: client,
		Repo:   eth_repo.NewRepo(db),
	}, nil
}

func (i *Impl) FormatBlock(input *entity.Block) *service.Block {
	txhashes := []string{}

	for _, tx := range input.Transactions {
		txhashes = append(txhashes, tx.Hash)
	}

	return &service.Block{
		Number:       input.Number,
		Hash:         input.Hash,
		Timestamp:    input.Timestamp,
		ParentHash:   input.ParentHash,
		Transactions: txhashes,
	}
}

func (i *Impl) FormatTransaction(tx *entity.Transaction, logs []*service.Log) *service.Transaction {
	return &service.Transaction{
		Hash:  tx.Hash,
		From:  tx.From,
		To:    tx.To,
		Nonce: tx.Nonce,
		Data:  tx.Data,
		Value: tx.Value,
		Logs:  logs,
	}
}

func (i *Impl) GetBlocks(ctx context.Context) (*service.Blocks, error) {
	blocks, err := i.Repo.GetBlocks()
	if err != nil {
		return nil, err
	}
	res := &service.Blocks{}
	for _, block := range blocks {
		res.Blocks = append(res.Blocks, i.FormatBlock(block))
	}
	return res, nil
}

func (i *Impl) GetBlock(ctx context.Context, number int64) (*service.Block, error) {
	block, err := i.Repo.GetBlockByNumber(uint64(number))
	if err != nil {
		return nil, err
	}
	return i.FormatBlock(block), err
}

func (i *Impl) GetTransaction(ctx context.Context, hash string) (*service.Transaction, error) {
	tx, err := i.Repo.GetTransaction(hash)
	if err != nil {
		return nil, err
	}

	if len(tx.Logs) == 0 {
		logs, err := i.TransactionLogsRPC(ctx, tx.Hash)
		if err != nil {
			return nil, err
		}
		logStr, _ := json.Marshal(logs)
		if err := i.Repo.SaveTransactionLogs(tx, string(logStr)); err != nil {
			return nil, err
		}
		return i.FormatTransaction(tx, logs), nil
	}

	var logs []*service.Log
	if err := json.Unmarshal([]byte(tx.Logs), &logs); err != nil {
		return nil, err
	}
	return i.FormatTransaction(tx, logs), nil
}

func (i *Impl) TransactionLogsRPC(ctx context.Context, hash string) ([]*service.Log, error) {
	commonHash := common.Hash{}
	if err := commonHash.UnmarshalText([]byte(hash)); err != nil {
		return nil, err
	}
	receipt, err := i.Client.TransactionReceipt(ctx, commonHash)
	if err != nil {
		return nil, err
	}

	logs := []*service.Log{}
	for _, log := range receipt.Logs {
		logs = append(logs, &service.Log{
			Index: log.Index,
			Data:  common.Bytes2Hex(log.Data),
		})
	}
	return logs, nil
}

func (i *Impl) BlocksExist(ctx context.Context, startNumber, endNumber int64) (bool, error) {
	return i.Repo.BlocksExist(startNumber, endNumber)
}

func (i *Impl) BlockExist(ctx context.Context, number int64) (bool, error) {
	return i.Repo.BlockExist(number)
}

func (i *Impl) NetworkIDRPC(ctx context.Context) (int64, error) {
	chainId, err := i.Client.NetworkID(ctx)
	return chainId.Int64(), err
}

func (i *Impl) SaveBlockRPC(ctx context.Context, number, chainId int64) error {
	exist, err := i.BlockExist(ctx, number)
	if exist {
		return nil
	} else if err != nil {
		return err
	}

	block, err := i.Client.BlockByNumber(ctx, big.NewInt(number))
	if err != nil {
		return err
	}

	transactions := make([]entity.Transaction, 0)
	signer := types.LatestSignerForChainID(big.NewInt(chainId))

	for _, tx := range block.Transactions() {
		sender, err := types.Sender(signer, tx)
		if err != nil {
			return err
		}
		var val int64
		if tx.Value() != nil {
			val = tx.Value().Int64()
		}
		var to string
		if tx.To() != nil {
			to = tx.To().String()
		}

		transactions = append(transactions, entity.Transaction{
			Hash:  tx.Hash().String(),
			From:  sender.String(),
			To:    to,
			Nonce: tx.Nonce(),
			Data:  common.Bytes2Hex(tx.Data()),
			Value: val,
		})
	}

	return i.Repo.SaveBlock(&entity.Block{
		Number:       block.NumberU64(),
		Hash:         block.Hash().String(),
		Timestamp:    block.Time(),
		ParentHash:   block.ParentHash().String(),
		Transactions: transactions,
	})
}
