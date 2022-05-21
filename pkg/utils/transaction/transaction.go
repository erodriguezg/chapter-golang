package transaction

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	txContextKey = "txContextKey"

	txErrorNoTransactionActive = "no transaction active!"

	txErrorNoSqlTransaction = "the transaction is not a sql type!"
)

type TxManager interface {
	Begin(context.Context) (context.Context, error)
	GetTx(context.Context) (any, error)
	Commit(context.Context) error
	Rollback(context.Context) error
}

type sqlTxManager struct {
	db *sql.DB
}

func NewSqlTxManager(db *sql.DB) TxManager {
	return &sqlTxManager{db}
}

func (m *sqlTxManager) Begin(ctx context.Context) (context.Context, error) {

	if ctx == nil {
		return nil, nil
	}

	// search for an existing transaction

	beforeTx, err := m.GetTx(ctx)
	if err != nil {
		return nil, err
	}

	if beforeTx != nil {
		// use the before transaction
		_, ok := beforeTx.(*sql.Tx)
		if !ok {
			return nil, fmt.Errorf(txErrorNoSqlTransaction)
		}
		return ctx, nil
	}

	// no existing transaction, create new

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, txContextKey, tx), nil
}

func (m *sqlTxManager) GetTx(ctx context.Context) (any, error) {
	if ctx == nil {
		return nil, nil
	}
	tx := ctx.Value(txContextKey)
	if tx == nil {
		return nil, nil
	}
	aux, ok := tx.(*sql.Tx)
	if !ok {
		return nil, fmt.Errorf("problem casting tx")
	}
	return aux, nil
}

func (m *sqlTxManager) Commit(ctx context.Context) error {
	sqlTx, err := m.getSqlTxManagerFromContext(ctx)
	if err != nil {
		return err
	}
	return sqlTx.Commit()
}

func (m *sqlTxManager) Rollback(ctx context.Context) error {
	sqlTx, err := m.getSqlTxManagerFromContext(ctx)
	if err != nil {
		return err
	}
	return sqlTx.Rollback()
}

// private

func (m *sqlTxManager) getSqlTxManagerFromContext(ctx context.Context) (*sql.Tx, error) {
	tx, err := m.GetTx(ctx)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, fmt.Errorf(txErrorNoTransactionActive)
	}
	sqlTx, ok := tx.(*sql.Tx)
	if !ok {
		return nil, fmt.Errorf(txErrorNoSqlTransaction)
	}
	return sqlTx, nil
}
