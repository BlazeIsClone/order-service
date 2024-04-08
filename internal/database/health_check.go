package database

import (
	"context"
	"database/sql"
)

func Ping(db *sql.DB) error {
	return db.PingContext(context.Background())
}
