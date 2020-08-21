package migration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDBConfig(t *testing.T) {
	t.Run("Is default port set", func(t *testing.T) {
		cfg := NewDBConfig(
			"qwe",
			"",
			"",
		)

		assert.Equal(t, fmt.Sprintf("qwe:%s", defaultPostgresPort), cfg.address)
	})

	t.Run("Multiple passwords passed", func(t *testing.T) {
		cfg := NewDBConfig(
			"qwe",
			"",
			"",
		)
		cfg.SetPort("asd")

		assert.Equal(t, "qwe:asd", cfg.address)
	})

	t.Run("Is passed password used", func(t *testing.T) {
		cfg := NewDBConfig(
			"",
			"",
			"qwe",
		)
		cfg.SetPassword("asd")

		assert.Equal(t, "qwe:asd", cfg.userCreds)
	})
}

// Temporary unused
/*func failUnexpected(t *testing.T, err error) {
	t.Logf("Expected no error but got: %s", err.Error())
	t.FailNow()
}*/
