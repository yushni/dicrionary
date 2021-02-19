package facilities

import (
	"os"
)

type Config struct {
	DB DBConfig
}

func NewConfig() *Config {
	port := defaultPostgresPort

	if p, ok := os.LookupEnv("db_port"); ok {
		port = p
	}

	conf := &Config{
		DB: newDBConfig(
			os.Getenv("db_host"),
			os.Getenv("db_name"),
			port,
			os.Getenv("db_user"),
			os.Getenv("db_password"),
		),
	}

	return conf
}
