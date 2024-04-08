package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/felipemalacarne/lumina/logger/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Logger struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) *Logger {
	return &Logger{collection: collection}
}

func (l *Logger) Log(level string, message string, data bson.M) error {
	logEntry := bson.M{
		"timestamp": time.Now(),
		"level":     level,
		"message":   message,
		"data":      data,
	}
	_, err := l.collection.InsertOne(context.Background(), logEntry)
	return err
}

func (l *Logger) RotateLogs() error {
	file, err := os.OpenFile(filepath.Join(config.LogFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	now := time.Now()
	weekAgo := now.AddDate(0, 0, -config.LogFileMaxAge)
	cutoff := time.Date(weekAgo.Year(), weekAgo.Month(), weekAgo.Day(), 0, 0, 0, 0, time.UTC)

	var logs []bson.M
	filter := bson.M{"timestamp": bson.M{"$lt": cutoff}}
	cursor, err := l.collection.Find(context.Background(), filter)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var log bson.M
		err := cursor.Decode(&log)
		if err != nil {
			return err
		}
		logs = append(logs, log)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	for _, log := range logs {
		_, err := fmt.Fprintf(file, "%v %v %v %v\n", log["timestamp"], log["level"], log["message"], log["data"])
		if err != nil {
			return err
		}
	}

	return nil
}
