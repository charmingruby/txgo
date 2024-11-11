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
	connectionString := BuildConnectionString(false, in)

	db, err := sql.Open(DATASOURCE, connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func BuildConnectionString(hasDatasourcePrefix bool, in MySQLConnectionInput) string {
	if !hasDatasourcePrefix {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", in.Username, in.Password, in.Host, in.Port, in.DatabaseName)
	}

	return fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s?parseTime=true", DATASOURCE, in.Username, in.Password, in.Host, in.Port, in.DatabaseName)
}
