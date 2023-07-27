package handler

import (
	"github.com/gofiber/fiber/v2"

	"service_auth_go/api/presenter"
	"service_auth_go/pkg/auth"
)

func SigninHandler(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request presenter.AuthRequest

		err := c.BodyParser(&request)
		if err != nil {
			return c.JSON(presenter.ErrorResponse(err))
		}

		result, err := service.SigninService(&request)
		if err != nil {
			return c.JSON(presenter.ErrorResponse(err))
		}

		return c.JSON(presenter.AuthResponseSuccess(result))
	}
}
