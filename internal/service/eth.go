package service

import (
	"context"
)

type Eth interface {
	GetBlocks(ctx context.Context) (*Blocks, error)
	GetBlock(ctx context.Context, number int64) (*Block, error)
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

