package database

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sql.DB, config *mysql.Config) (*migrate.Migrate, error) {
	driver, _ := mysql.WithInstance(db, config)

	return migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
}
