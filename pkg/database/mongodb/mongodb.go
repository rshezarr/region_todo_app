package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context) (*mongo.Database, error) {
	cliOptions := options.Client().ApplyURI("mongodb://lists_db:27017")

	client, err := mongo.Connect(ctx, cliOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database("todo_list_db"), nil
}
