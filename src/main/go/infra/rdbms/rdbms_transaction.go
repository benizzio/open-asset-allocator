package rdbms

import (
	"context"
	"database/sql"
)

const sqlTransactionContextKey = "TRANSACTION"

type SQLTransactionalContext struct {
	context.Context
}

func (transactionalContext *SQLTransactionalContext) GetTransaction() *sql.Tx {
	return transactionalContext.Context.Value(sqlTransactionContextKey).(*sql.Tx)
}

func withTransaction(db *sql.DB) (*SQLTransactionalContext, error) {

	var transaction, err = db.Begin()
	if err != nil {
		return nil, err
	}

	var parentContext = context.WithValue(context.Background(), sqlTransactionContextKey, transaction)
	return &SQLTransactionalContext{parentContext}, nil
}

type TransactionalContext interface {
	context.Context
	GetTransaction() *sql.Tx
}

func ToSQLTransactionalContext(transContext context.Context) (*SQLTransactionalContext, bool) {
	if transactionalContext, ok := transContext.(*SQLTransactionalContext); ok {
		return transactionalContext, true
	}
	return nil, false
}
