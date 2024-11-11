package safesql

import (
	"context"
	"database/sql"
	_ "database/sql/driver"
)

type DB struct {
	db *sql.DB
}

type Rows = sql.Rows
type Result = sql.Result

func (db *DB) QueryContext(ctx context.Context, query TrustedSQL, args ...any) (*Rows, error) {
	r, err := db.db.QueryContext(ctx, query.s, args...)
	return r, err
}

func (db *DB) ExecContext(ctx context.Context, query TrustedSQL, args ...any) (Result, error) {
	r, err := db.db.ExecContext(ctx, query.s, args...)
	return r, err
}

func (db *DB) Open(driverName, dataSourceName string) (*DB, error) {
	db, err := db.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db: db.db}, nil
}
