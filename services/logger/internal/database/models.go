package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Level string

const (
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
	DEBUG Level = "DEBUG"
	FATAL Level = "FATAL"
)

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
