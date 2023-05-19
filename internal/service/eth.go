package service

import (
	"context"
)

type Eth interface {
	BlocksExist(ctx context.Context, startNumber, endNumber int64) (bool, error)
	BlockExist(ctx context.Context, number int64) (bool, error)
	NetworkIDRPC(ctx context.Context) (int64, error)
	SaveBlockRPC(ctx context.Context, number, chainId int64) error
}
