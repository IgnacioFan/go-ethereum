package eth_repo

import (
	"go-ethereum/internal/entity"
	"go-ethereum/internal/repository"
	"go-ethereum/pkg/postgres"
)

type Impl struct {
	postgres.Postgres
}

func NewRepo(db *postgres.Postgres) repository.Eth {
	return &Impl{
		*db,
	}
}

func (i *Impl) BlocksExist(start, end int64) (bool, error) {
	var count int64
	result := i.DB.Model(&entity.Block{}).Where("number >= ? AND number <= ?", start, end).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	expectedCount := int64(end - start + 1)
	return count == expectedCount, nil
}

func (i *Impl) BlockExist(number int64) (bool, error) {
	var count int64
	result := i.DB.Model(&entity.Block{}).Where("number = ?", number).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count == 1, nil
}

func (i *Impl) SaveBlock(block *entity.Block) error {
	return i.DB.Create(block).Error
}
