package migration

import (
	"errors"
	"fmt"
	"strings"

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
	ErrNoCredentials = errors.New("migration: no database credentials specified")
	ErrNoUpOrDown    = errors.New("migration: unspecified whether upgrade or downgrade migration version")
)

type MigrateOption func(*Migrate) error

func Downgrade() MigrateOption {
	return func(m *Migrate) error {
		m.direction = Down
		return nil
	}
}

func WithConfig(config DBConfig) MigrateOption {
	return func(m *Migrate) error {
		var notPassed []string

		if config.Host == "" {
			notPassed = append(notPassed, "Host")
		}

		if config.User == "" {
			notPassed = append(notPassed, "User")
		}

		if len(notPassed) > 0 {
			return fmt.Errorf("migration: missing fields: %s", strings.Join(notPassed, ", "))
		}

		if config.Port == "" {
			// Set default Postgres port
			config.Port = defaultPostgresPort
		}

		// Combine Host and Port values
		m.address = fmt.Sprintf("%s:%s", config.Host, config.Port)

		// Add User value to full user credentials
		m.userCreds = config.User

		// Check if Password is not empty
		if config.Password != "" {
			// if so then add it to full user credentials
			m.userCreds += fmt.Sprintf(":%s", config.Password)
		}

		m.dbName = config.DBName

		return nil
	}
}

func NewMigrate(opts ...MigrateOption) (*Migrate, error) {
	migration := &Migrate{
		direction: Up,
	}

	for _, opt := range opts {
		err := opt(migration)

		if err != nil {
			return nil, err
		}
	}

	if migration.userCreds == "" {
		return nil, ErrNoCredentials
	}

	return migration, nil
}

type Migrate struct {
	// direction used to define whether run up or down migrations
	direction int

	// address is a combination of Host and Port fields
	//
	// Example: Host:Port
	address string

	// userCreds is a combination of User and Password (if exists) fields
	//
	// Example: User[:Password]
	userCreds string

	dbName string
}

func (o *Migrate) Run() error {
	m, err := migrate.New(
		"file://migration/migrations",
		fmt.Sprintf("postgres://%s@%s/%s", o.userCreds, o.address, o.dbName),
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
