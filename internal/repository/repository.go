package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Repository struct {
	TodoList TodoList
}

func NewRepository(db *mongo.Database) *Repository {
	db.CreateCollection(context.Background(), "lists")
	collection := db.Collection("lists")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "title", Value: 1}, {Key: "activeAt", Value: 1}},
		Options: options.Index().SetUnique(true).SetSparse(true),
	}

	// Create the unique compound index
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("error while setting index: %v\n", err)
	}

	return &Repository{
		TodoList: NewTodoListRepo(db),
	}
}
