package facilities

import "os"

type Config struct {
	DB DBConfig
}

func NewConfig() Config {
	conf := Config{
		DB: newDBConfig(
			os.Getenv("db_host"),
			os.Getenv("db_name"),
			os.Getenv("db_user"),
		),
	}

	conf.DB.setPassword(os.Getenv("db_password"))
	conf.DB.setPort(os.Getenv("db_port"))

	return conf
}
