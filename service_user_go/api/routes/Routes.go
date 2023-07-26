package routes

import (
	"github.com/gofiber/fiber/v2"

	"service_user_go/api/handlers"
	"service_user_go/pkg/user"
)

func Routes(app fiber.Router, service user.Service) {
	app.Get("/users", handlers.IndexHandler(service))
	app.Post("/user", handlers.CreateHandler(service))
}
