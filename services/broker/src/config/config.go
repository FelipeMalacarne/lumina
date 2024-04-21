package config

import "os"

var (
	RabbitUser string = "guest"
	RabbitPass string = "guest"
	RabbitHost string = "rabbitmq"
	RabbitPort string = "5672"
	RabbitURI  string
)

func Load() {
	user, ok := os.LookupEnv("RABBITMQ_DEFAULT_USER")
	if ok {
		RabbitUser = user
	}

	pass, ok := os.LookupEnv("RABBITMQ_DEFAULT_PASS")
	if ok {
		RabbitPass = pass
	}

	host, ok := os.LookupEnv("RABBITMQ_HOST")
	if ok {
		RabbitHost = host
	}

	port, ok := os.LookupEnv("RABBITMQ_PORT")
	if ok {
		RabbitPort = port
	}

	RabbitURI = "amqp://" + RabbitUser + ":" + RabbitPass + "@" + RabbitHost + ":" + RabbitPort + "/"
}
