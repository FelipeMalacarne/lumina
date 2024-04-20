package config

import (
	"os"
)

var (
	MongoURI          string
	MongoDBName       string
	MongoRootUsername string
	MongoRootPassword string
	RpcPort           string
	GrpcPort          string
)

func Load() {
	MongoURI = os.Getenv("MONGO_URI")
	MongoDBName = os.Getenv("MONGO_DB_NAME")
	MongoRootUsername = os.Getenv("MONGO_ROOT_USERNAME")
	MongoRootPassword = os.Getenv("MONGO_ROOT_PASSWORD")
	RpcPort = os.Getenv("RPC_PORT")
	GrpcPort = os.Getenv("GRPC_PORT")
}
