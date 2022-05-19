package transaction

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	txContextKey = "txContextKey"
)

type TxManager[T any] interface {
	Begin(context.Context) (context.Context, error)
	GetTx(context.Context) (T, error)
	Commit(context.Context) error
	Rollback(context.Context) error
}

type sqlTxManager struct {
	db *sql.DB
}

func NewSqlTxManager(db *sql.DB) TxManager[*sql.Tx] {
	return &sqlTxManager{db}
}

func (m *sqlTxManager) Begin(ctx context.Context) (context.Context, error) {

	if ctx == nil {
		return nil, nil
	}

	// search for an existing transaction

	beforeTx, err := m.GetTx(ctx)
	if err != nil {
		return ctx, err
	}

	// use the actual transacction

	if beforeTx != nil {
		return ctx, nil
	}

	// no existing transaction, create new

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, txContextKey, tx), nil
}

func (m *sqlTxManager) GetTx(ctx context.Context) (*sql.Tx, error) {
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
	tx, err := m.GetTx(ctx)
	if err != nil {
		return err
	}
	if tx == nil {
		return fmt.Errorf("no transaction active!")
	}
	return tx.Commit()
}

func (m *sqlTxManager) Rollback(ctx context.Context) error {
	tx, err := m.GetTx(ctx)
	if err != nil {
		return err
	}
	if tx == nil {
		return fmt.Errorf("no transaction active!")
	}
	return tx.Rollback()
}
