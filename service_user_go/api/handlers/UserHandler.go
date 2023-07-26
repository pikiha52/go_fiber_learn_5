package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"service_user_go/api/presenter"
	"service_user_go/pkg/entities"
	"service_user_go/pkg/user"
)

func IndexHandler(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		results, err := service.IndexService()
		if err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UsersSuccessResponse(results))
	}
}

func CreateHandler(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body entities.User

		err := c.BodyParser(&body)
		if err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}

		validate := validator.New()
		errValidate := validate.Struct(body)

		if errValidate != nil {
			c.Status(400)
			return c.JSON(presenter.UserErrorResponse(errValidate))
		}

		_, errService := service.CreateService(&body)
		if errService != nil {
			return c.JSON(presenter.UserErrorResponse(errService))
		}

		c.Status(201)
		return c.JSON(fiber.Map{"status": true, "message": "User successfully created!"})
	}
}

func FindUserByUsername(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Query("username")
		result, err := service.FindUserByUsernameService(username)

		if err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UserSuccessResponse(result))
	}
}
