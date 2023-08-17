package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todo_list/internal/dto"
	"todo_list/internal/model"
	"todo_list/internal/repository"
)

var (
	ErrInvalidTitle = errors.New("title length out of range")
	ErrInvalidDate  = errors.New("invalid date")
)

type TodoList interface {
	CreateList(ctx context.Context, list dto.List) (string, error)
	GetList(ctx context.Context, status string) ([]dto.List, error)
	UpdateList(ctx context.Context, id string, newList dto.List) error
	DeleteList(ctx context.Context, id string) error
	UpdateStatus(ctx context.Context, id string) error
}

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) TodoList {
	return &TodoListService{
		repo: repo,
	}
}

func checkList(list dto.List) error {
	if len(list.Title) > 200 {
		return ErrInvalidTitle
	}

	// Parse the date
	parsedTime, err := time.Parse("2006-01-02", list.ActiveAt)
	if err != nil {
		return ErrInvalidDate
	}

	// Check if the month and day are within valid ranges
	month := int(parsedTime.Month())
	day := parsedTime.Day()

	if !(month >= 1 && month <= 12 && day >= 1 && day <= 31) {
		return fmt.Errorf("error while checking date: %w", ErrInvalidDate)
	}

	return nil
}

func parseDtoList(list dto.List) (model.List, error) {
	parsedTime, err := time.Parse("2006-01-02", list.ActiveAt)
	if err != nil {
		return model.List{}, err
	}

	l := model.List{
		Title:    list.Title,
		ActiveAt: parsedTime,
		Status:   list.Status,
	}

	return l, nil
}

func (t TodoListService) CreateList(ctx context.Context, list dto.List) (string, error) {
	if err := checkList(list); err != nil {
		return "", err
	}

	l, err := parseDtoList(list)
	if err != nil {
		return "", err
	}

	id, err := t.repo.Create(ctx, l)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (t TodoListService) GetList(ctx context.Context, status string) ([]dto.List, error) {
	lists, err := t.repo.Get(ctx, status)
	if err != nil {
		return []dto.List{}, err
	}

	var respLists []dto.List

	for _, list := range lists {
		var respList dto.List
		respList.Title = list.Title
		respList.ActiveAt = list.ActiveAt.Format("2006-01-02")
		respList.Status = list.Status

		respLists = append(respLists, respList)
	}

	return respLists, nil
}

func (t TodoListService) UpdateList(ctx context.Context, id string, newList dto.List) error {
	if err := checkList(newList); err != nil {
		return err
	}

	l, err := parseDtoList(newList)
	if err != nil {
		return err
	}

	if err := t.repo.Update(ctx, id, l); err != nil {
		return err
	}

	return nil
}

func (t TodoListService) DeleteList(ctx context.Context, id string) error {
	return t.repo.Delete(ctx, id)
}

func (t *TodoListService) UpdateStatus(ctx context.Context, id string) error {
	return t.repo.UpdateStatus(ctx, id)
}
