package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"todo_list/internal/model"
)

type TodoList interface {
	Create(ctx context.Context, list model.List) (int, error)
	Get(ctx context.Context, status string) (model.List, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, newList model.List) error
}

type TodoListRepo struct {
	db *mongo.Collection
}

func NewTodoListRepo(db *mongo.Database) TodoList {
	return &TodoListRepo{
		db: db.Collection("lists"),
	}
}

func (t *TodoListRepo) Create(ctx context.Context, list model.List) (int, error) {
	result, err := t.db.InsertOne(ctx, list)
	if err != nil {
		return 0, err
	}

	id := result.InsertedID.(int)

	return id, nil
}

func (t *TodoListRepo) Get(ctx context.Context, status string) (model.List, error) {
	filter := bson.M{
		"activeAt": bson.M{"$lte": time.Now()},
	}

	list := model.List{}

	cur, err := t.db.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.List{}, errors.New("list not found")
		}

		return model.List{}, err
	}

	if err := cur.Decode(&list); err != nil {
		return model.List{}, err
	}

	return list, nil
}

func (t *TodoListRepo) Delete(ctx context.Context, id int) error {
	_, err := t.db.DeleteOne(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoListRepo) Update(ctx context.Context, id int, newList model.List) (int, error) {
	result, err := t.db.UpdateByID(ctx, id, newList)
	if err != nil {
		return 0, err
	}

	resId := result.UpsertedID.(int)

	return resId, nil
}
