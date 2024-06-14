package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db     *mongo.Database
	client *mongo.Client
)

func Init() error {
	var err error
	mongoEndpoint := os.Getenv("MONGODB_URI")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		return err
	}
	db = client.Database("sendit")
	return nil
}

func Collection(col string) *mongo.Collection {
	return db.Collection(col)
}
