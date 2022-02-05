package facilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConfig_Sting(t *testing.T) {
	tests := map[string]struct {
		c      DBConfig
		expStr string
	}{
		"success: fully filled config": {
			c: DBConfig{
				Host:     "host",
				DBName:   "db",
				Port:     "port",
				Username: "user",
				Password: "pass",
			},
			expStr: "postgres://user:pass@host:port/db",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expStr, tc.c.String())
		})
	}
}
