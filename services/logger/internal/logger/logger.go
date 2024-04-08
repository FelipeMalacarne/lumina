package logger

import (
	"context"
	"time"

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
