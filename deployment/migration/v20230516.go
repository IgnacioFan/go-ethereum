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
			Number     uint64 `gorm:"primaryKey,index"`
			Hash       string `gorm:"index"`
			Timestamp  uint64
			ParentHash string
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
