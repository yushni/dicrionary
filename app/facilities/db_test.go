package facilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDBConfig(t *testing.T) {
	t.Run("Is default port set", func(t *testing.T) {
		cfg := newDBConfig(
			"",
			"",
			"",
		)

		assert.Equal(t, defaultPostgresPort, cfg.port)
	})

	t.Run("Is passed port used", func(t *testing.T) {
		cfg := newDBConfig(
			"",
			"",
			"",
		)

		cfg.setPort("1234")

		assert.Equal(t, "1234", cfg.port)
	})

	t.Run("Is passed password used", func(t *testing.T) {
		cfg := newDBConfig(
			"",
			"",
			"",
		)
		cfg.setPassword("asd")

		assert.Equal(t, "asd", cfg.password)
	})

	t.Run("Is hasPassword field changed", func(t *testing.T) {
		cfg := newDBConfig(
			"",
			"",
			"",
		)
		cfg.setPassword("asd")

		assert.Equal(t, true, cfg.hasPassword)
	})

	t.Run("String() returns correct value", func(t *testing.T) {
		cfg := newDBConfig(
			"host",
			"db",
			"user",
		)

		cfg.setPort("port")
		cfg.setPassword("pass")

		assert.Equal(t, "postgres://user:pass@host:port/db", cfg.String())
	})
}

// Temporary unused
/*func failUnexpected(t *testing.T, err error) {
	t.Logf("Expected no error but got: %s", err.Error())
	t.FailNow()
}*/
