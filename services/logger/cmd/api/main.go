package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/felipemalacarne/lumina/logger/cmd/api/v1"
	"github.com/felipemalacarne/lumina/logger/internal/config"
	"github.com/felipemalacarne/lumina/logger/internal/database"
	"github.com/felipemalacarne/lumina/logger/internal/logger"
	"github.com/felipemalacarne/lumina/logger/internal/server"

	"google.golang.org/grpc"
)

func main() {
	config.Load()
	log.SetOutput(os.Stdout)
	log.SetPrefix(fmt.Sprintf("[%s] ", time.Now().Format("2006-01-02 15:04:05")))

	client, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer client.Disconnect(context.Background())

	collection := database.GetCollection(client, "logs")
	logger := logger.New(collection)

	// Create the gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterLoggerServer(grpcServer, server.NewLoggerServer(logger))

	// Start the server
	lis, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening on port: " + config.GrpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// l.Log(logger.LogEntry{
	// 	Level:   logger.INFO,
	// 	Message: "Log created successfully",
	// 	Service: logger.LOGGER,
	// 	Data: bson.M{
	// 		"key1": "value1",
	// 	},
	// })

	// log.Println("Log created successfully")
}
