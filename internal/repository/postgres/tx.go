package postgres

import (
	"context"
	"database/sql"
	"go-file-upload/internal/repository"
)

type sqlDBTx struct {
	ddb SqlDB
	db  *sqlDB
}

func TxConn(tx *sql.Tx, db *sqlDB) sqlDBTx {
	adapter := &sqlDB{ddb: tx, db: db.db}
	return sqlDBTx{tx, adapter}
}

func (db *sqlDB) WrapTx(ctx context.Context, fn func(repository.FileRepo) error) (err error) {
	var tx *sql.Tx

	if tx, err = db.db.BeginTx(ctx, nil); err != nil {
		return
	}
	defer tx.Rollback()

	dbTX := TxConn(tx, db)

	if err = fn(dbTX.db); err != nil {
		return
	}

	return tx.Commit()
}
