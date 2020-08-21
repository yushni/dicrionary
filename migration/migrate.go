package migration

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	// Up used to update database structure running all up migrations
	Up = iota

	// Down used to run all down migrations
	Down

	defaultPostgresPort = "5432"
)

var (
	ErrNoUpOrDown = errors.New("migration: unspecified whether upgrade or downgrade migration version")
)

type MigrateOption func(*Migrate)

// Downgrade is used to run down migrations - roll back changes made to
// database structure.
func Downgrade() MigrateOption {
	return func(m *Migrate) {
		m.direction = Down
	}
}

// WithDirectory sets a path to directory where .sql migration files
// are located.
//
// You should not start it with `file://` because it will
// be added automatically.
func WithDirectory(source string) MigrateOption {
	return func(m *Migrate) {
		m.source = source
	}
}

func NewMigrate(cfg DBConfig, opts ...MigrateOption) (*Migrate, error) {
	migration := &Migrate{
		direction: Up,
		source:    "migration/migrations",
		config:    cfg,
	}

	for _, opt := range opts {
		opt(migration)
	}

	return migration, nil
}

type Migrate struct {
	// direction used to define whether run up or down migrations
	direction int

	// source is a path to folder with .sql migration files
	// from root project directory.
	//
	// 	Default: "migration/migrations"
	source string

	config DBConfig
}

func (o *Migrate) Run() error {
	m, err := migrate.New(
		fmt.Sprintf("file://%s", o.source),
		fmt.Sprintf("postgres://%s@%s/%s", o.config.userCreds, o.config.address, o.config.dbName),
	)
	if err != nil {
		return err
	}

	switch o.direction {
	case Up:
		err = m.Up()
	case Down:
		err = m.Down()
	default:
		return ErrNoUpOrDown
	}
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
