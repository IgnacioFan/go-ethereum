package postgres

import (
	"go-ethereum/deployment/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres() (*Postgres, error) {
	dsn, err := env.DSN()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Postgres{DB: db}, nil
}
