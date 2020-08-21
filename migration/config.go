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

// NewDBConfig returns DBConfig struct with properly formatted data you should
// pass to this constructor.
//
// `password` is optional (pass it only if database user you are using has password)
// and if you pass more than one optional argument, all of them will
// be ignored except the first.
func NewDBConfig(address, database, username string, password ...string) DBConfig {
	cfg := DBConfig{}

	// Check if pattern `host:port` has the `port` value
	if len(strings.Split(address, ":")) > 1 {
		cfg.address = address
	} else {
		// Combine ip address and default Postgres port
		cfg.address = fmt.Sprintf("%s:%s", address, defaultPostgresPort)
	}

	if len(password) > 0 {
		cfg.userCreds = fmt.Sprintf("%s:%s", username, password[0])
	} else {
		cfg.userCreds = username
	}

	cfg.dbName = database

	return cfg
}
