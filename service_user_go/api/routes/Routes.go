package routes

import (
	"github.com/gofiber/fiber/v2"

	"service_user_go/api/handlers"
	"service_user_go/api/middleware"
	"service_user_go/pkg/user"

)

func Routes(app fiber.Router, service user.Service) {
	app.Get("/users", middleware.Protected(), handlers.IndexHandler(service))
	app.Post("/user", middleware.Protected(), handlers.CreateHandler(service))
	app.Get("/user", handlers.FindUserByUsername(service))
}
