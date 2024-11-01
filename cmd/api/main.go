package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/txgo/config"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	_, err := config.New()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %v", err))
		os.Exit(1)
	}
}
