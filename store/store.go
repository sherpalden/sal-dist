package store

type Store interface {
	Create(item interface{}) error
	FindAll(items interface{}) error
}
