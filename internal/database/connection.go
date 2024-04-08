package database

import (
	"database/sql"
	"fmt"

	driver "github.com/go-sql-driver/mysql"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Connection() (*sql.DB, error) {
	cnf := driver.Config{
		User:                 "root",
		Passwd:               "password",
		DBName:               "db",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cnf.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("sql.Open %w", err)
	}

	if err := Ping(db); err != nil {
		return nil, fmt.Errorf("sql.error %w", err)
	}

	return db, nil
}
