package main

import (
	"fmt"
	"log"
	"os"

	"github.com/felipemalacarne/lumina/broker/src/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	Rabbit *amqp.Connection
}

func main() {
	config.Load()
	rabbitConn, err := connectToRabbit()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// app := App{Rabbit: rabbitConn}

	fmt.Println("Connected to RabbitMQ")
}

func connectToRabbit() (*amqp.Connection, error) {
	conn, err := amqp.Dial(config.RabbitURI)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
