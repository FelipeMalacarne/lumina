package database

import (
	"context"

	"github.com/felipemalacarne/lumina/logger/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.MongoURI)
	clientOptions.SetAuth(options.Credential{
		Username: config.MongoRootUsername,
		Password: config.MongoRootPassword,
	})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(config.MongoDBName).Collection(collectionName)
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry interface{}
