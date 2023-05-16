package eth_client

import (
	"context"
	"go-ethereum/internal/entity"
)

type Client interface {
	BlockByNumber(ctx context.Context, blockId int64) (*entity.Block, error)
}
