package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/felipemalacarne/lumina/logger/proto"
	"google.golang.org/grpc"
)

func main() {
	LoadConfig()
	log.SetOutput(os.Stdout)
	log.SetPrefix(fmt.Sprintf("[%s] ", time.Now().Format("2006-01-02 15:04:05")))

	client, err := Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer client.Disconnect(context.Background())

	collection := GetCollection(client, "logs")
	logger := New(collection)

	grpcServer := grpc.NewServer()
	pb.RegisterLoggerServer(grpcServer, NewLoggerServer(logger))

	lis, err := net.Listen("tcp", GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening on port: " + GrpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
