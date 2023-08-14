package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"todo_list/internal/model"
)

type TodoList interface {
	Create(list model.List) (int, error)
	Get(id int) (model.List, error)
	Delete(id int) error
	Update(id int, newList model.List) error
}

type TodoListRepo struct {
	db *mongo.Database
}

func NewTodoListRepo(db *mongo.Database) TodoList {
	return &TodoListRepo{
		db: db,
	}
}

func (t TodoListRepo) Create(list model.List) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepo) Get(id int) (model.List, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepo) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepo) Update(id int, newList model.List) error {
	//TODO implement me
	panic("implement me")
}
