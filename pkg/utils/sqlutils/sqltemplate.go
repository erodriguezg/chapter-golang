package sqlutils

import (
	"context"
	"database/sql"

	"github.com/erodriguezg/chapter-golang/pkg/utils/transaction"
)

type SqlTemplate[T any] interface {
	QueryForArray(ctx context.Context, query string, params []interface{}, mapperFunc func(rows *sql.Rows) (T, error)) ([]T, error)
	QueryForOne(ctx context.Context, query string, params []interface{}, mapperFunc func(row *sql.Row) (T, error)) (*T, error)
	Exec(ctx context.Context, sql string, params []interface{}) (int64, error)
}

type defaultImpl[T any] struct {
	db        *sql.DB
	txManager transaction.TxManager[*sql.Tx]
}

func NewSqlTemplate[T any](db *sql.DB, txManager transaction.TxManager[*sql.Tx]) SqlTemplate[T] {
	return &defaultImpl[T]{db, txManager}
}

func (impl *defaultImpl[T]) QueryForArray(ctx context.Context, query string, params []interface{}, mapperFunc func(rows *sql.Rows) (T, error)) ([]T, error) {
	tx, err := impl.txManager.GetTx(ctx)
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows

	if tx != nil {
		rows, err = (*tx).Query(query, params...)
	} else {
		rows, err = impl.db.Query(query, params...)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var outputArray []T
	for rows.Next() {
		aux, err := mapperFunc(rows)
		if err != nil {
			return nil, err
		}
		outputArray = append(outputArray, aux)
	}
	return outputArray, nil
}

func (impl *defaultImpl[T]) QueryForOne(ctx context.Context, query string, params []interface{}, mapperFunc func(row *sql.Row) (T, error)) (*T, error) {
	tx, err := impl.txManager.GetTx(ctx)
	if err != nil {
		return nil, err
	}

	var row *sql.Row

	if tx != nil {
		row = (*tx).QueryRow(query, params...)
	} else {
		row = impl.db.QueryRow(query, params...)
	}

	aux, err := mapperFunc(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &aux, nil
}

func (impl *defaultImpl[T]) Exec(ctx context.Context, query string, params []interface{}) (int64, error) {

	tx, err := impl.txManager.GetTx(ctx)
	if err != nil {
		return 0, err
	}

	var result sql.Result

	if tx != nil {
		result, err = (*tx).Exec(query, params...)
	} else {
		result, err = impl.db.Exec(query, params...)
	}

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
