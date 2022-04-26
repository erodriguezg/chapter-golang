package sqlutils

import "database/sql"

type SqlTemplate[T any] interface {
	QueryForArray(query string, params []interface{}) ([]T, error)
	QueryForOne(query string, params []interface{}) (*T, error)
	Update(sql string, params []interface{}) (interface{}, error)
}

type defaultImpl[T any] struct {
	db *sql.DB
}

func NewSqlTemplate[T any](db *sql.DB) SqlTemplate[T] {
	return &defaultImpl[T]{db}
}

func (impl *defaultImpl[T]) QueryForArray(query string, params []interface{}) ([]T, error) {

	// aqui cierra conexion

	return nil, nil
}
func (impl *defaultImpl[T]) QueryForOne(query string, params []interface{}) (*T, error) {
	return nil, nil
}
func (impl *defaultImpl[T]) Update(sql string, params []interface{}) (interface{}, error) {
	return nil, nil
}
