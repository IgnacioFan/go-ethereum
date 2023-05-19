package repository

import "go-ethereum/internal/entity"

type Eth interface {
	GetBlocks() ([]*entity.Block, error)
	GetBlockByNumber(blockId uint64) (*entity.Block, error)
	BlocksExist(start, end int64) (bool, error)
	BlockExist(number int64) (bool, error)
	SaveBlock(block *entity.Block) error
}
