package service

import "github.com/zhanibek05/golang-todo-app/pkg/database"

type Authorization interface{}

type TodoList interface {
}
type TodoItem interface {
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(db *database.Database) *Service {
	return &Service{}
}
