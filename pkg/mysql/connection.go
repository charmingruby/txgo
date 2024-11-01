package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DATASOURCE = "mysql"
)

type MySQLConnectionInput struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func New(in MySQLConnectionInput) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", in.Username, in.Password, in.Host, in.Port, in.DatabaseName)

	db, err := sql.Open(DATASOURCE, connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
