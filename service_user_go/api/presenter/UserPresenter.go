package presenter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"service_user_go/pkg/entities"

)

type User struct {
	ID          uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Bod         time.Time `json:"bod"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func UserSuccessResponse(data *entities.User) *fiber.Map {
	user := User{
		ID:          data.ID,
		Name:        data.Name,
		Username:    data.Username,
		Gender:      data.Gender,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		Bod:         data.Bod,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}

	return &fiber.Map{
		"status": true,
		"result": user,
	}
}

func UsersSuccessResponse(data *[]User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"result": data,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"errors": err.Error(),
	}
}
