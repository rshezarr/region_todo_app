package service

import (
	"todo_list/internal/repository"
)

type Service struct {
	TodoList TodoList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		TodoList: NewTodoListService(repo.TodoList),
	}
}
