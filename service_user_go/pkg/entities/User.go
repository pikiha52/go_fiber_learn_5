package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;not null;unique"`
	Name        string    `validate:"required" gorm:"not null" json:"name"`
	Username    string    `validate:"required" gorm:"index;unique;not null" json:"username"`
	Gender      string    `validate:"required" gorm:"not null" json:"gender"`
	Address     string    `validate:"required" gorm:"not null" json:"address"`
	PhoneNumber string    `validate:"required" gorm:"not null" json:"phone_number"`
	Bod         time.Time `json:"bod"`
	Password    string    `validate:"required" gorm:"not null" json:"password"`
	IsDelete    bool      `gorm:"default:false" gorm:"not null" json:"is_delete"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	DeletedAT   time.Time `json:"deleted_at"`
}
