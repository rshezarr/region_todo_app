package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	TodoList TodoList
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		TodoList: NewTodoListRepo(db),
	}
}
