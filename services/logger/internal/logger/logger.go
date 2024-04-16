package logger

import (
	"context"
	"time"

	"github.com/felipemalacarne/lumina/logger/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Logger struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) *Logger {
	return &Logger{collection: collection}
}

func (l *Logger) NewEntry(level database.Level, message string, data bson.M, service database.Service) database.LogEntry {
	return database.LogEntry{
		CreatedAt: time.Now(),
		Data:      data,
		Message:   message,
		Level:     level,
		Service:   service,
	}
}

func (l *Logger) Log(entry database.LogEntry) error {
	_, err := l.collection.InsertOne(context.TODO(), entry)
	if err != nil {
		return err
	}

	return nil
}
