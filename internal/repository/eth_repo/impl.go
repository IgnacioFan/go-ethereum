package eth_repo

import (
	"go-ethereum/internal/entity"
	"go-ethereum/internal/repository"
	"go-ethereum/pkg/postgres"

	"gorm.io/gorm"
)

type Impl struct {
	postgres.Postgres
}

func NewRepo(db *postgres.Postgres) repository.Eth {
	return &Impl{
		*db,
	}
}

func (i *Impl) GetBlocks() ([]*entity.Block, error) {
	var blocks []*entity.Block
	if err := i.DB.Limit(5).Order("number desc").Find(&blocks).Error; err != nil {
		return nil, err
	}
	return blocks, nil
}

func (i *Impl) GetBlockByNumber(number uint64) (*entity.Block, error) {
	var block *entity.Block
	if err := i.DB.Preload("Transactions", func(db *gorm.DB) *gorm.DB {
		return db.Select("Hash", "BlockHash")
	}).First(&block, "number = ?", number).Error; err != nil {
		return nil, err
	}
	return block, nil
}

func (i *Impl) GetTransaction(hash string) (*entity.Transaction, error) {
	var tx *entity.Transaction
	if err := i.DB.First(&tx, "hash", hash).Error; err != nil {
		return nil, err
	}
	return tx, nil
}

func (i *Impl) SaveTransactionLogs(tx *entity.Transaction, logs string) error {
	return i.DB.Model(&tx).Update("name", logs).Error
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
