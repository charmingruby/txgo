package config

import (
	env "github.com/caarlos0/env/v6"
)

type environment struct {
	ServerPort       string `env:"SERVER_PORT,required"`
	DatabaseUser     string `env:"DATABASE_USER,required"`
	DatabasePassword string `env:"DATABASE_PASSWORD,required"`
	DatabaseHost     string `env:"DATABASE_HOST,required"`
	DatabaseName     string `env:"DATABASE_NAME,required"`
	DatabasePort     string `env:"DATABASE_PORT,required"`
}

func New() (Config, error) {
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return Config{}, err
	}

	cfg := Config{
		ServerConfig: serverConfig{
			Port: environment.ServerPort,
		},
		MySQLConfig: mysqlConfig{
			User:         environment.DatabaseUser,
			Password:     environment.DatabasePassword,
			Host:         environment.DatabaseHost,
			DatabaseName: environment.DatabaseName,
			Port:         environment.DatabasePort,
		},
	}

	return cfg, nil
}

type Config struct {
	ServerConfig serverConfig
	MySQLConfig  mysqlConfig
}

type serverConfig struct {
	Host string
	Port string
}

type mysqlConfig struct {
	User         string
	Password     string
	Host         string
	DatabaseName string
	Port         string
}
