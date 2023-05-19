package migration

import (
	"go-ethereum/internal/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var v20230517 = &gormigrate.Migration{
	ID: "20230517",
	Migrate: func(tx *gorm.DB) error {
		// when table already exists, it just adds fields as columns
		type Transaction struct {
			Hash      string `gorm:"primaryKey;unique"`
			From      string `gorm:"size:128"`
			To        string `gorm:"size:128"`
			Nonce     uint64
			Data      string
			Value     int64
			BlockHash string `gorm:"size:128;not null"`
			Logs      string
		}

		type Block struct {
			Hash         string        `gorm:"size:128;uniqueIndex;not null"`
			Transactions []Transaction `gorm:"foreignKey:BlockHash;references:Hash"`
		}

		return tx.AutoMigrate(&Transaction{}, &Block{})
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable(&entity.Transaction{}); err != nil {
			return err
		}
		return nil
	},
}
