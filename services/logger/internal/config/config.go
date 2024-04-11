package config

import (
	"os"
)

var (
	MongoURI          string
	MongoDBName       string
	MongoRootUsername string
	MongoRootPassword string
	rpcPort           string
	gRpcPort          string
)

func Load() {
	MongoURI = os.Getenv("MONGO_URI")
	MongoDBName = os.Getenv("MONGO_DB_NAME")
	MongoRootUsername = os.Getenv("MONGO_ROOT_USERNAME")
	MongoRootPassword = os.Getenv("MONGO_ROOT_PASSWORD")
	rpcPort = "5001"
	gRpcPort = "50051"
}
