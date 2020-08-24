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

	if pass, hasPass := os.LookupEnv("db_password"); hasPass {
		conf.DB.setPassword(pass)
	}

	if port, hasPort := os.LookupEnv("db_port"); hasPort {
		conf.DB.setPort(port)
	}

	return conf
}
