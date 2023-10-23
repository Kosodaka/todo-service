package repository

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	Dsn string
}

const (
	usersItemsTable = "users_items"
	todoItemTable   = "todo_items"
)

// SQl connection
func NewPostgresDb(cfg Config) (*bun.DB, error) {
	// Connecting to Postgres
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.Dsn)))
	// Create Bun connection
	db := bun.NewDB(sqldb, pgdialect.New())

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
