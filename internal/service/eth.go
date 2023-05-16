package service

import "context"

type Eth interface {
	SaveBlock(ctx context.Context, blockId int64) error
}
