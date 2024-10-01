package postgres

import (
	"context"
	"database/sql"
)

type sqlDB struct {
	ddb SqlDB

	// needed to beginTx
	db *sql.DB
}

type SqlDB interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func NewSqlDB(driver, connStr string) (repo *sqlDB, err error) {
	db, err := sql.Open(driver, connStr)

	return &sqlDB{
		db, db,
	}, err
}
