package migration

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"strings"
)

func RunMigrate(upOrDown string) error {
	m, err := migrate.New(
		"file://migrate/migrations",
		"postgres://admin:admin@localhost:5432/dictionary",
	)
	if err != nil {
		return err
	}

	switch strings.ToLower(upOrDown) {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		return errors.New("unspecified weather upgrade or downgrade migration version")
	}
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
