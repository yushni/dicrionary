package migration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type migrateTest struct {
	name      string
	useConfig bool
	wantError bool
	downgrade bool
	errorMsg  string
	config    DBConfig
}

func initMigrationTest(t *testing.T, tests []migrateTest, opts ...MigrateOption) {

	// tc means "test case"

}

func TestDBConfig(t *testing.T) {
	cases := []migrateTest{
		{
			name:      "No credentials error",
			wantError: true,
			errorMsg:  ErrNoCredentials.Error(),
		},
		{
			name:      "No host and user",
			wantError: true,
			errorMsg:  "migration: missing fields: Host, User",
			useConfig: true,
			config:    DBConfig{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewMigrate()

			if tc.useConfig {
				_, err = NewMigrate(
					WithConfig(tc.config),
				)
			}

			if tc.downgrade {
				_, err = NewMigrate(
					Downgrade(),
				)
			}

			if err != nil && tc.wantError {
				assert.EqualError(t, err, tc.errorMsg)
			} else if err != nil {
				failUnexpected(t, err)
			}
		})
	}

	t.Run("Use default port", func(t *testing.T) {
		cfg := DBConfig{
			User: "qwe",
			Host: "qwe",
		}

		m, err := NewMigrate(
			WithConfig(cfg),
		)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, m.address, fmt.Sprintf("%s:%s", cfg.Host, defaultPostgresPort))
	})

	t.Run("User creds without password", func(t *testing.T) {
		cfg := DBConfig{
			User: "qwe",
			Host: "qwe",
		}

		m, err := NewMigrate(
			WithConfig(cfg),
		)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, m.userCreds, fmt.Sprintf("%s", cfg.User))
	})

	t.Run("User creds with password", func(t *testing.T) {
		cfg := DBConfig{
			User:     "qwe",
			Password: "qwe",
			Host:     "qwe",
		}

		m, err := NewMigrate(
			WithConfig(cfg),
		)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, m.userCreds, fmt.Sprintf("%s:%s", cfg.User, cfg.Password))
	})
}

func failUnexpected(t *testing.T, err error) {
	t.Logf("Expected no error but got: %s", err.Error())
	t.FailNow()
}
