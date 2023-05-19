package migration

import (
	"go-ethereum/internal/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var v20230516 = &gormigrate.Migration{
	ID: "20230516",
	Migrate: func(tx *gorm.DB) error {
		// when table already exists, it just adds fields as columns
		type Block struct {
			Number     uint64 `gorm:"primaryKey;unique"`
			Hash       string `gorm:"size:128;uniqueIndex;not null"`
			Timestamp  uint64 `gorm:"not null"`
			ParentHash string `gorm:"size:128;not null"`
		}
		return tx.AutoMigrate(&Block{})
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable(&entity.Block{}); err != nil {
			return err
		}
		return nil
	},
}
