package routes

import (
	"github.com/gofiber/fiber/v2"

	"service_auth_go/api/handler"
	"service_auth_go/pkg/auth"
)

func Routes(app fiber.Router, service auth.Service) {
	app.Post("/auth/signin", handler.SigninHandler(service))
}
