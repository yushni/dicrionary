package migration

import (
	"fmt"
	"strings"
)

type DBConfig struct {
	// address is a combination of host and port.
	//
	//  Example: host[:port]
	address string

	// userCreds is a combination of username and password (if exists).
	//
	//  Example: username[:password]
	userCreds string

	dbName string
}

func NewDBConfig(host, database, username string) DBConfig {
	cfg := DBConfig{}

	// Combine ip address and default Postgres port
	cfg.address = fmt.Sprintf("%s:%s", host, defaultPostgresPort)

	cfg.userCreds = username

	cfg.dbName = database

	return cfg
}

// SetPort sets database port.
//
//  Default: "5432"
func (c *DBConfig) SetPort(port string) {
	host := strings.Split(c.address, ":")[0]

	c.address = fmt.Sprintf("%s:%s", host, port)
}

// SetPassword sets user's password to connect to database.
func (c *DBConfig) SetPassword(password string) {
	username := strings.Split(c.userCreds, ":")

	c.userCreds = fmt.Sprintf("%s:%s", username, password)
}
