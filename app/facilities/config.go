package facilities

import (
	"os"
)

type Config struct {
	DB DBConfig
}

func NewConfig() (*Config, error) {
	var port string

	port, ok := os.LookupEnv("db_port")
	if !ok {
		port = defaultPostgresPort
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

	return conf, nil
}
