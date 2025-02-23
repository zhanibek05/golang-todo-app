package database

type Authorization interface{}

type TodoList interface {
}
type TodoItem interface {
}
type Database struct {
	Authorization
	TodoList
	TodoItem
}

func NewDatabase() *Database {
	return &Database{}
}
