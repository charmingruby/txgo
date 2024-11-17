package core

type TransactionalConsistencyProvider[T any] interface {
	Transact(func(params T) error) error
}
