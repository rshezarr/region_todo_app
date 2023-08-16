package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todo_list/internal/model"
)

type TodoList interface {
	Create(ctx context.Context, list model.List) (string, error)
	Get(ctx context.Context, status string) ([]model.List, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, newList model.List) (string, error)
}

type TodoListRepo struct {
	db *mongo.Collection
}

func NewTodoListRepo(db *mongo.Database) TodoList {
	db.CreateCollection(context.Background(), "lists")
	return &TodoListRepo{
		db: db.Collection("lists"),
	}
}

func (t *TodoListRepo) Create(ctx context.Context, list model.List) (string, error) {
	result, err := t.db.InsertOne(ctx, list)
	if err != nil {
		return "", err
	}

	objectIDValue, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		fmt.Println("Interface does not hold a primitive.ObjectID value")
		return "", err
	}

	id := objectIDValue.String()

	return id, nil
}

func (t *TodoListRepo) Get(ctx context.Context, status string) ([]model.List, error) {
	filter := bson.M{
		"title": "Купить книгу 2",
	}

	var list []model.List

	cur, err := t.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var l model.List
		if err := cur.Decode(&l); err != nil {
			return nil, err
		}
		list = append(list, l)
	}

	return list, nil
}

func (t *TodoListRepo) Delete(ctx context.Context, id string) error {
	_, err := t.db.DeleteOne(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoListRepo) Update(ctx context.Context, id string, newList model.List) (string, error) {
	result, err := t.db.UpdateByID(ctx, id, newList)
	if err != nil {
		return "", err
	}

	objectIDValue, ok := result.UpsertedID.(primitive.ObjectID)
	if !ok {
		fmt.Println("Interface does not hold a primitive.ObjectID value")
		return "", err
	}

	resId := objectIDValue.String()

	return resId, nil
}
