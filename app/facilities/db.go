package facilities

import (
	"fmt"
)

const (
	defaultPostgresPort = "5432"
)

// DO NOT create DBConfig directly without constructor.
type DBConfig struct {
	host        string
	port        string
	username    string
	password    string
	hasPassword bool
	dbName      string
}

func newDBConfig(host, database, username string) DBConfig {
	cfg := DBConfig{}

	cfg.host = host
	cfg.port = defaultPostgresPort

	cfg.username = username

	cfg.dbName = database

	return cfg
}

func (c *DBConfig) setPort(port string) {
	c.port = port
}

func (c *DBConfig) setPassword(password string) {
	c.password = password
	c.hasPassword = true
}

// String generates Postgres database connection URI.
//
//  Pattern: "postgres://username[:password]@host:port/dbname"
func (c DBConfig) String() string {
	userCreds := c.username

	if c.hasPassword {
		userCreds = fmt.Sprintf("%s:%s", c.username, c.password)
	}

	return fmt.Sprintf("postgres://%s@%s:%s/%s", userCreds, c.host, c.port, c.dbName)
}
