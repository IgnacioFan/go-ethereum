package service

import (
	"context"
)

type Eth interface {
	GetBlocks(ctx context.Context, limit int) (*Blocks, error)
	GetBlock(ctx context.Context, number int64) (*Block, error)
	GetTransaction(ctx context.Context, hash string) (*Transaction, error)
	BlocksExist(ctx context.Context, startNumber, endNumber int64) (bool, error)
	BlockExist(ctx context.Context, number int64) (bool, error)
	NetworkIDRPC(ctx context.Context) (int64, error)
	SaveBlockRPC(ctx context.Context, number, chainId int64) error
}

type Block struct {
	Number       uint64   `json:"block_num"`
	Hash         string   `json:"block_hash"`
	Timestamp    uint64   `json:"block_time"`
	ParentHash   string   `json:"parent_hash"`
	Transactions []string `json:"transactions,omitempty"`
}

type Blocks struct {
	Blocks []*Block `json:"blocks"`
}

type Transaction struct {
	Hash  string `json:"tx_hash"`
	From  string `json:"from"`
	To    string `json:"to"`
	Nonce uint64 `json:"nonce"`
	Data  string `json:"data"`
	Value int64  `json:"value"`
	Logs  []*Log `json:"logs"`
}

type Log struct {
	Index uint   `json:"index"`
	Data  string `json:"data"`
}
