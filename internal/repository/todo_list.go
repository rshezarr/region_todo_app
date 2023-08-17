package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"todo_list/internal/model"
)

type TodoList interface {
	Create(ctx context.Context, list model.List) (string, error)
	Get(ctx context.Context, status string) ([]model.List, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, newList model.List) error
	UpdateStatus(ctx context.Context, id string) error
}

type TodoListRepo struct {
	db *mongo.Collection
}

func NewTodoListRepo(db *mongo.Database) TodoList {

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

	id := objectIDValue.Hex()

	return id, nil
}

func (t *TodoListRepo) Get(ctx context.Context, status string) ([]model.List, error) {
	filter := bson.M{}
	if status == "active" {
		filter = bson.M{
			"activeat": bson.M{"$lte": time.Now()},
		}
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
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting ID: %w", err)
	}

	filter := bson.M{"_id": objectID}

	_, err = t.db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoListRepo) Update(ctx context.Context, id string, newList model.List) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting ID: %w", err)
	}

	update := bson.M{"$set": bson.M{"title": newList.Title, "activeat": newList.ActiveAt}}

	res, err := t.db.UpdateByID(ctx, objectID, update)
	if err != nil {
		return fmt.Errorf("repo: error while updating: %w", err)
	} else if res.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}

	return nil
}

func (t *TodoListRepo) UpdateStatus(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting ID: %w", err)
	}

	update := bson.M{"$set": bson.M{"status": "done"}}

	res, err := t.db.UpdateByID(ctx, objectID, update)
	if err != nil {
		return fmt.Errorf("repo: error while updating: %w", err)
	} else if res.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}

	return nil
}
