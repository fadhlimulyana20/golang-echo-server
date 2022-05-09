package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()
var db *mongo.Database
var e error

func Connect() error {
	clientOptions := options.Client()
	clientOptions.ApplyURI(os.Getenv("MONGO_URL"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return err
	}

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	db = client.Database(os.Getenv("MONGO_DB_NAME"))
	return nil
}

func DbManager() *mongo.Database {
	return db
}
