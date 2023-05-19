package repository

import "go-ethereum/internal/entity"

type Eth interface {
	BlocksExist(start, end int64) (bool, error)
	BlockExist(number int64) (bool, error)
	SaveBlock(block *entity.Block) error
}
