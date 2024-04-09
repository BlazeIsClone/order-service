package database

import (
	"database/sql"
	"fmt"

	"context"

	driver "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Init() (*sql.DB, error) {
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

func Migrate(db *sql.DB, config *mysql.Config) (*migrate.Migrate, error) {
	driver, _ := mysql.WithInstance(db, config)

	return migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
}

func Ping(db *sql.DB) error {
	return db.PingContext(context.Background())
}
