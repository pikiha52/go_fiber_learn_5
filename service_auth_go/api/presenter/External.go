package presenter

import "time"

type ExternalRequest struct {
	Result Result `json:"result"`
	Status bool   `json:"status"`
}

type Result struct {
	Uuid        string    `json:"uuid"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	BOD         string    `json:"bod"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Password    string    `json:"password"`
}
