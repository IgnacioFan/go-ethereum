package eth_repo

import (
	"go-ethereum/internal/entity"
	"go-ethereum/internal/repository"
	"go-ethereum/pkg/postgres"
)

type Impl struct {
	DB *postgres.Postgres
}

func NewRepo(db *postgres.Postgres) repository.Eth {
	return &Impl{
		DB: db,
	}
}

func (i *Impl) SaveBlock(block *entity.Block) error {
	return i.DB.DB.Create(block).Error
}
