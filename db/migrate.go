package db

import (
	"errors"

	"dictionary/app/facilities"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	Up = iota + 1
	Down
)

type Migrate struct {
	direction int
	source    string
	config    facilities.DBConfig
}

func NewMigrate(cfg facilities.DBConfig) *Migrate {
	return &Migrate{
		direction: Up,
		source:    "migration/migrations",
		config:    cfg,
	}
}

func (o *Migrate) Run() error {
	m, err := migrate.New("file://"+o.source, o.config.String())
	if err != nil {
		return err
	}

	switch o.direction {
	case Up:
		err = m.Up()
	case Down:
		err = m.Down()
	default:
		return errors.New("migration: unspecified whether upgrade or downgrade migration version")
	}

	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
