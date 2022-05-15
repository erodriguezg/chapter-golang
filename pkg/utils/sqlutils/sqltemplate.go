package sqlutils

import (
	"database/sql"

	"github.com/erodriguezg/chapter-golang/pkg/utils/transaction"
)

type SqlTemplate[T any] interface {
	QueryForArray(query string, params []interface{}) ([]T, error)
	QueryForOne(query string, params []interface{}) (*T, error)
	Update(sql string, params []interface{}) (interface{}, error)
}

type defaultImpl[T any] struct {
	db        *sql.DB
	txManager transaction.TxManager[*sql.Tx]
}

func NewSqlTemplate[T any](db *sql.DB, txManager transaction.TxManager[*sql.Tx]) SqlTemplate[T] {
	return &defaultImpl[T]{db, txManager}
}

func (impl *defaultImpl[T]) QueryForArray(query string, params []interface{}) ([]T, error) {

	tx, err := impl.txManager.GetTx()
	if err != nil {
		return nil, err
	}

	if tx != nil {
		return nil, nil
	}

	// aqui cierra conexion

	return nil, nil
}
func (impl *defaultImpl[T]) QueryForOne(query string, params []interface{}) (*T, error) {
	return nil, nil
}
func (impl *defaultImpl[T]) Update(sql string, params []interface{}) (interface{}, error) {
	return nil, nil
}
