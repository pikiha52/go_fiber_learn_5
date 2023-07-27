package presenter

import (
	"github.com/gofiber/fiber/v2"

	"service_auth_go/pkg/entities"

)

type AuthRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

func AuthResponseSuccess(data *entities.Response) *fiber.Map {
	response := entities.Response{
		Username:    data.Username,
		AccessToken: data.AccessToken,
	}

	return &fiber.Map{
		"status": true,
		"data":   response,
	}
}

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  err.Error(),
	}
}
