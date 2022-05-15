package transaction

type TxManager[T any] interface {
	Begin() error
	GetTx() (T, error)
	Commit() error
	Rollback()
}
