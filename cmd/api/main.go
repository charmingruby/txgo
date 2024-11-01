package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/txgo/config"
	"github.com/charmingruby/txgo/pkg/mysql"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	config, err := config.New()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %v", err))
		os.Exit(1)
	}

	_, err = mysql.New(mysql.MySQLConnectionInput{
		Username:     config.MySQLConfig.User,
		Password:     config.MySQLConfig.Password,
		Host:         config.MySQLConfig.Host,
		Port:         config.MySQLConfig.Port,
		DatabaseName: config.MySQLConfig.DatabaseName,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("MYSQL: %v", err))
		os.Exit(1)
	}
}
