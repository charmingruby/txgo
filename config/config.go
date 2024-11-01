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
	DatabaseSSL      string `env:"DATABASE_SSL,required"`
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
		PostgresConfig: postgresConfig{
			User:         environment.DatabaseUser,
			Password:     environment.DatabasePassword,
			Host:         environment.DatabaseHost,
			DatabaseName: environment.DatabaseName,
			SSL:          environment.DatabaseSSL,
		},
	}

	return cfg, nil
}

type Config struct {
	ServerConfig   serverConfig
	PostgresConfig postgresConfig
}

type serverConfig struct {
	Host string
	Port string
}

type postgresConfig struct {
	User         string
	Password     string
	Host         string
	DatabaseName string
	SSL          string
}
