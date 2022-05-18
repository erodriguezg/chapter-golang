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
	GetTx(context.Context) (*T, error)
	Commit(context.Context) error
	Rollback(context.Context)
}

type defaultTxManager[T any] struct {
	db *sql.DB
}

func NewTxManager[T any](db *sql.DB) TxManager[T] {
	return &defaultTxManager[T]{db}
}

func (m *defaultTxManager[T]) Begin(ctx context.Context) (context.Context, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, txContextKey, tx), nil
}

func (m *defaultTxManager[T]) GetTx(ctx context.Context) (*T, error) {
	if ctx == nil {
		return nil, nil
	}
	tx := ctx.Value(txContextKey)
	if tx == nil {
		return nil, nil
	}
	aux, ok := tx.(T)
	if !ok {
		return nil, fmt.Errorf("problem casting tx")
	}
	return &aux, nil
}

func (m *defaultTxManager[T]) Commit(ctx context.Context) error {
	return nil
}

func (m *defaultTxManager[T]) Rollback(ctx context.Context) {

}
