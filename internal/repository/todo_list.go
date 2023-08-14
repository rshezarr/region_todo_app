package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"todo_list/internal/model"
)

type TodoList interface {
	Create(ctx context.Context, list model.List) (int, error)
	Get(ctx context.Context, id int) (model.List, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, newList model.List) error
}

type TodoListRepo struct {
	db *mongo.Collection
}

func NewTodoListRepo(db *mongo.Collection) TodoList {
	return &TodoListRepo{
		db: db,
	}
}

func (t TodoListRepo) Create(ctx context.Context, list model.List) (int, error) {
	result, err := t.db.InsertOne(ctx, list)
	if err != nil {
		return 0, err
	}

	id := result.InsertedID.(int)

	return id, nil
}

func (t TodoListRepo) Get(ctx context.Context, id int) (model.List, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepo) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepo) Update(ctx context.Context, id int, newList model.List) error {
	//TODO implement me
	panic("implement me")
}
