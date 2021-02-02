package facilities

import (
	"fmt"
	"os"
)

type Config struct {
	DB DBConfig
}

func NewConfig() (*Config, error) {
	conf := &Config{
		DB: newDBConfig(
			os.Getenv("db_host"),
			os.Getenv("db_name"),
			os.Getenv("db_user"),
		),
	}

	if err := conf.DB.setPassword(os.Getenv("db_password")); err != nil {
		return nil, fmt.Errorf("cannot create Config: %e", err)
	}
	conf.DB.setPort(os.Getenv("db_port"))

	return conf, nil
}
