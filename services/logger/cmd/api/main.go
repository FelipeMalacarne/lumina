package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/felipemalacarne/lumina/logger/internal/config"
	"github.com/felipemalacarne/lumina/logger/internal/database"
	"github.com/felipemalacarne/lumina/logger/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
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

	l := logger.New(collection)

	l.Log(logger.LogEntry{
		Level:   logger.INFO,
		Message: "Log created successfully",
		Service: logger.LOGGER,
		Data: bson.M{
			"key1": "value1",
		},
	})

	log.Println("Log created successfully")
}
