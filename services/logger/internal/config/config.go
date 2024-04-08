package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoURI      string
	MongoDBName   string
	LoggerLevel   string
	LogFile       string
	LogFileMaxAge int
	rpcPort       string
	gRpcPort      string
)

func Load() {
	MongoURI = os.Getenv("MONGO_URI")
	MongoDBName = os.Getenv("MONGO_DB_NAME")
	LoggerLevel = os.Getenv("LOGGER_LEVEL")
	LogFile = os.Getenv("LOG_FILE")
	LogFileMaxAge = 7
	rpcPort = "5001"
	gRpcPort = "50051"
}

func GetMongoClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(MongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}
