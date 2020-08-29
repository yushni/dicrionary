package facilities

import (
	"fmt"
)

const (
	defaultPostgresPort = "5432"
)

type DBConfig struct {
	host     string
	port     string
	username string
	password string
	dbName   string
}

func newDBConfig(host, database, username string) DBConfig {
	return DBConfig{
		host:     host,
		port:     defaultPostgresPort,
		username: username,
		dbName:   database,
	}
}

func (c *DBConfig) setPort(port string) {
	c.port = port
}

func (c *DBConfig) setPassword(password string) {
	c.password = password
}

// String generates Postgres database connection URI.
//
//  Pattern: "postgres://username[:password]@host:port/dbname"
func (c DBConfig) String() string {
	userCreds := c.username

	if c.password != "" {
		userCreds = fmt.Sprintf("%s:%s", c.username, c.password)
	}

	return fmt.Sprintf("postgres://%s@%s:%s/%s", userCreds, c.host, c.port, c.dbName)
}
