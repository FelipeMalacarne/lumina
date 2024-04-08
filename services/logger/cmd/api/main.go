package main

import "log"

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50051"
)

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
}
