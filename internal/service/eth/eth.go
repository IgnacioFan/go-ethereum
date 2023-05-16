package eth

import "context"

type EthService interface {
	SaveBlock(ctx context.Context, blockId int64) error
}
