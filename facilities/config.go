package facilities

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DB DBConfig
}

func NewConfig() *Config {
	cfg := &Config{
		DB: DBConfig{},
	}
	if err := env.Parse(&cfg.DB); err != nil {
		panic(fmt.Sprintf("failed to parse db config: %s", err))
	}
	return cfg
}

type DBConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	Username string `env:"DB_USER_NAME" envDefault:"user"`
	Password string `env:"DB_USER_PASSWORD" envDefault:"example"`
	DBName   string `env:"DB_NAME" envDefault:"postgres"`
	SSL      string `env:"DB_SSL" envDefault:"disable"`
}

func (c DBConfig) String() string {
	credentials := c.Username
	if c.Password != "" {
		credentials = fmt.Sprintf("%s:%s", c.Username, c.Password)
	}
	fmt.Println(c.DBName)
	return fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s", credentials, c.Host, c.Port, c.DBName, c.SSL)
}
