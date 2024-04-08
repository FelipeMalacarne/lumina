package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/felipemalacarne/lumina/logger/internal/config"
	"github.com/felipemalacarne/lumina/logger/internal/database"
	"github.com/felipemalacarne/lumina/logger/internal/logger"
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

	collection := client.Database(config.MongoDBName).Collection("logs")
	logger := logger.New(collection)

	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	for {
		select {
		case <-done:
			log.Println("Shutting down...")
			return
		case t := <-ticker.C:
			err := logger.RotateLogs()
			if err != nil {
				log.Printf("Failed to rotate logs: %v", err)
			}
			log.Printf("Rotated logs at %s", t.Format(time.RFC3339))
		}
	}
}
