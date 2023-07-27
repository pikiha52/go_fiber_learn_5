package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	app := fiber.New()

	connection := ConnectRabbit()

	go func() {
		ReceiveQueue(connection)
	}()

	app.Listen(":3003")
}

func ConnectRabbit() *amqp.Connection {
	connect, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}

	return connect
}

func ReceiveQueue(connection *amqp.Connection) error {
	channel, err := connection.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer channel.Close()

	messages, err := channel.Consume(
		"WAQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	for message := range messages {
		fmt.Printf("Send Whatsapp to:  %s\n", message.Body)
	}

	return nil
}
