package repository

import "go-ethereum/internal/entity"

type Eth interface {
	SaveBlock(block *entity.Block) error
}
