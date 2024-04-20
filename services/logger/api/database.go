package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(MongoURI)
	clientOptions.SetAuth(options.Credential{
		Username: MongoRootUsername,
		Password: MongoRootPassword,
	})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(MongoDBName).Collection(collectionName)
}
