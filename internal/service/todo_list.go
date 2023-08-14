package service

import (
	"context"
	"todo_list/internal/model"
	"todo_list/internal/repository"
)

type TodoList interface {
	CreateList(ctx context.Context, list model.List) (int, error)
	GetList(ctx context.Context, status string) (model.List, error)
	UpdateList(ctx context.Context, newList model.List) (int, error)
	DeleteList(ctx context.Context, id int) error
}

type TodoListService struct {
	repo *repository.TodoListRepo
}

func NewTodoListService(repo *repository.TodoListRepo) TodoList {
	return &TodoListService{
		repo: repo,
	}
}

func (t TodoListService) CreateList(ctx context.Context, list model.List) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListService) GetList(ctx context.Context, status string) (model.List, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListService) UpdateList(ctx context.Context, newList model.List) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListService) DeleteList(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
