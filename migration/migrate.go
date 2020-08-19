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
)

var (
	ErrNoCredentials = errors.New("migration: no database credentials specified")
	ErrNoUpOrDown    = errors.New("unspecified whether upgrade or downgrade migration version")
)

type MigrateOption func(*Migrate)

func Downgrade() MigrateOption {
	return func(m *Migrate) {
		m.direction = Down
	}
}

func WithUser(username string) MigrateOption {
	return func(m *Migrate) {
		m.user = username
	}
}

func WithPassword(password string) MigrateOption {
	return func(m *Migrate) {
		m.password = password
	}
}

func WithHost(host string) MigrateOption {
	return func(m *Migrate) {
		m.host = host
	}
}

func WithPort(port string) MigrateOption {
	return func(m *Migrate) {
		m.port = port
	}
}

func NewMigrate(opts ...MigrateOption) (*Migrate, error) {
	migration := &Migrate{
		direction: Up,
		dbName:    "dictionary",
		host:      "localhost",
		port:      "5432",
	}

	for _, opt := range opts {
		opt(migration)
	}

	if migration.user == "" || migration.password == "" {
		return nil, ErrNoCredentials
	}

	return migration, nil
}

type Migrate struct {
	// Direction used to define whether run up or down migrations
	direction int

	host string
	port string

	user     string
	password string

	dbName string
}

func (o *Migrate) Run() error {
	m, err := migrate.New(
		"file://migration/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s", o.user, o.password, o.host, o.port, o.dbName),
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
