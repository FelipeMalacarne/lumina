package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/felipemalacarne/lumina/logger/api/config"
	"github.com/felipemalacarne/lumina/logger/api/database"
	"github.com/felipemalacarne/lumina/logger/api/models"
	"github.com/felipemalacarne/lumina/logger/api/server"
	pb "github.com/felipemalacarne/lumina/logger/proto"
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
	logger := models.New(collection)

	grpcServer := grpc.NewServer()
	pb.RegisterLoggerServer(grpcServer, server.NewLoggerServer(logger))

	lis, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening on port: " + config.GrpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
