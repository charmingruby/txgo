package container

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"

	mysqlConn "github.com/charmingruby/txgo/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	mysqlContainer "github.com/testcontainers/testcontainers-go/modules/mysql"
)

const (
	DATABASE_NAME     = "foo"
	DATABASE_USERNAME = "root"
	DATABASE_PASSWORD = "password"

	MIGRATIONS_PATH = "../../db/migration"
)

var (
	containerHost string
	containerPort string
)

type MySQL struct {
	DB        *sql.DB
	container *mysqlContainer.MySQLContainer
}

func NewMySQL() *MySQL {
	ctx := context.Background()

	container, err := mysqlContainer.Run(ctx,
		"mysql:8.0.36",
		mysqlContainer.WithDatabase(DATABASE_NAME),
		mysqlContainer.WithUsername(DATABASE_USERNAME),
		mysqlContainer.WithPassword(DATABASE_PASSWORD),
	)

	if err != nil {
		slog.Error(fmt.Sprintf("MYSQL TESTCONTAINER: failed to start container: %s", err))
		os.Exit(1)
	}

	containerHost, err = container.Host(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("MYSQL TESTCONTAINER: failed to get container host: %s", err))
		os.Exit(1)
	}

	natPort, err := container.MappedPort(ctx, "3306")
	if err != nil {
		slog.Error(fmt.Sprintf("MYSQL TESTCONTAINER: failed to get container port: %s", err))
		os.Exit(1)
	}

	containerPort = natPort.Port()

	db, err := mysqlConn.New(mysqlConn.MySQLConnectionInput{
		Username:     DATABASE_USERNAME,
		Password:     DATABASE_PASSWORD,
		Host:         containerHost,
		Port:         containerPort,
		DatabaseName: DATABASE_NAME,
	})

	if err := db.Ping(); err != nil {
		slog.Error(fmt.Sprintf("MYSQL TESTCONTAINER: failed to ping database: %s", err))
		os.Exit(1)
	}

	return &MySQL{
		DB:        db,
		container: container,
	}
}

func (m *MySQL) Teardown() error {
	if err := m.RollbackMigrations(); err != nil {
		slog.Error(fmt.Sprintf("MYSQL TESTCONTAINER: failed to rollback migrations: %s", err))
		return err
	}

	if err := m.DB.Close(); err != nil {
		slog.Error(fmt.Sprintf("MYSQL TESTCONTAINER: failed to close database connection: %s", err))
		return err
	}

	if err := m.container.Terminate(context.Background()); err != nil {
		slog.Error(fmt.Sprintf("MYSQL TESTCONTAINER: failed to terminate container: %s", err))
		return err
	}

	return nil
}

func (c *MySQL) RunMigrations() error {
	pathToMigrate := fmt.Sprintf("file://%s", MIGRATIONS_PATH)
	connectionString := mysqlConn.BuildConnectionString(true, mysqlConn.MySQLConnectionInput{
		Username:     DATABASE_USERNAME,
		Password:     DATABASE_PASSWORD,
		Host:         containerHost,
		Port:         containerPort,
		DatabaseName: DATABASE_NAME,
	})

	m, err := migrate.New(pathToMigrate, connectionString)
	if err != nil {
		slog.Error(fmt.Sprintf("MYSQL MIGRATION TESTCONTAINER: failed to create migrate up: %s", err))
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		slog.Error(fmt.Sprintf("MYSQL MIGRATION TESTCONTAINER: failed to run migrations: %s", err))
		return err
	}

	return nil
}

func (c *MySQL) RollbackMigrations() error {
	pathToMigrate := fmt.Sprintf("file://%s", MIGRATIONS_PATH)
	connectionString := mysqlConn.BuildConnectionString(true, mysqlConn.MySQLConnectionInput{
		Username:     DATABASE_USERNAME,
		Password:     DATABASE_PASSWORD,
		Host:         containerHost,
		Port:         containerPort,
		DatabaseName: DATABASE_NAME,
	})

	m, err := migrate.New(pathToMigrate, connectionString)
	if err != nil {
		slog.Error(fmt.Sprintf("MYSQL MIGRATION TESTCONTAINER: failed to create migrate down: %s", err))
		return err
	}
	defer m.Close()

	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		slog.Error(fmt.Sprintf("MYSQL MIGRATION TESTCONTAINER: failed to run migrations: %s", err))
		return err
	}

	return nil
}
