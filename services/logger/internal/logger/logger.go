package logger

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Logger struct {
	collection *mongo.Collection
	LogEntry   LogEntry
}

type Service string

const (
	GATEWAY Service = "GATEWAY"
	BROKER  Service = "BROKER"
	LOGGER  Service = "LOGGER"
	MAILER  Service = "MAILER"
	AUTH    Service = "AUTH"
)

type LogEntry struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	Data      bson.M    `bson:"data" json:"data"`
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Message   string    `bson:"message" json:"message"`
	Level     Level     `bson:"level" json:"level"`
	Service   Service   `bson:"service" json:"service"`
}

func New(collection *mongo.Collection) *Logger {
	return &Logger{collection: collection}
}

func (l *Logger) Log(entry LogEntry) error {
	_, err := l.collection.InsertOne(context.TODO(), LogEntry{
		Level:     entry.Level,
		Message:   entry.Message,
		Data:      entry.Data,
		Service:   entry.Service,
		CreatedAt: time.Now(),
	})
	return err
}
