package service

import (
	"context"
	"todo_list/internal/model"
	"todo_list/internal/repository"
)

type TodoList interface {
	CreateList(ctx context.Context, list model.List) (int, error)
	GetList(ctx context.Context, status string) ([]model.List, error)
	UpdateList(ctx context.Context, newList model.List) (int, error)
	DeleteList(ctx context.Context, id int) error
}

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) TodoList {
	return &TodoListService{
		repo: repo,
	}
}

func (t TodoListService) CreateList(ctx context.Context, list model.List) (int, error) {
	id, err := t.repo.Create(ctx, list)
	if err != nil {

	}
	return id, nil
}

func (t TodoListService) GetList(ctx context.Context, status string) ([]model.List, error) {
	list, err := t.repo.Get(ctx, status)
	if err != nil {
		return []model.List{}, err
	}

	return list, nil
}

func (t TodoListService) UpdateList(ctx context.Context, newList model.List) (int, error) {
	id, err := t.repo.Update(ctx, newList.ID, newList)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t TodoListService) DeleteList(ctx context.Context, id int) error {
	return t.repo.Delete(ctx, id)
}
