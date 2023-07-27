package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	amqp "github.com/rabbitmq/amqp091-go"

	"service_auth_go/api/routes"
	"service_auth_go/pkg/auth"
)

func main() {
	app := fiber.New()

	base_url := "http://localhost:3001/api"
	authRepository := auth.NewRepo(base_url)
	authService := auth.NewService(authRepository)

	api := app.Group("/api", logger.New())
	routes.Routes(api, authService)

	app.Listen(":3002")
}

func ConnectRabbitmq() *amqp.Connection {
	connect, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}

	return connect
}
