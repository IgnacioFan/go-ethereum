package postgres

import (
	"fmt"
	"go-ethereum/deployment/migration"

	"github.com/go-gormigrate/gormigrate/v2"
)

func (p *Postgres) NewMirgate() error {
	if err := gormigrate.New(p.DB, gormigrate.DefaultOptions, migration.Migrations).Migrate(); err != nil {
		return err
	}
	fmt.Println("migration success")
	return nil
}
