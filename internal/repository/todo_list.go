package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"todo_list/internal/model"
)

type TodoList interface {
	CreateList(list model.List) (int, error)
	GetList(id int) (model.List, error)
	DeleteList(id int) error
	UpdateList(id int, newList model.List) error
}

type TodoListRepo struct {
	db *mongo.Database
}

func NewTodoListRepo(db *mongo.Database) TodoList {
	return &TodoListRepo{
		db: db,
	}
}

func (t TodoListRepo) CreateList(list model.List) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepo) GetList(id int) (model.List, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepo) DeleteList(id int) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepo) UpdateList(id int, newList model.List) error {
	//TODO implement me
	panic("implement me")
}
