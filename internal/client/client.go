package client

import (
	"context"
	"go-ethereum/internal/entity"
)

type Eth interface {
	BlockByNumber(ctx context.Context, blockId int64) (*entity.Block, error)
}
