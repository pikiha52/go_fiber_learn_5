package user

import (
	"gorm.io/gorm"

	"service_user_go/api/presenter"
	"service_user_go/pkg/entities"

)

type Repository interface {
	IndexRepository() (*[]presenter.User, error)
	CreateRepository(user *entities.User) (*entities.User, error)
}

type repository struct {
	Database *gorm.DB
}

func NewRepo(database *gorm.DB) Repository {
	return &repository{
		Database: database,
	}
}

func (r *repository) IndexRepository() (*[]presenter.User, error) {
	var users []presenter.User
	r.Database.Find(&users)

	return &users, nil
}

func (r *repository) CreateRepository(user *entities.User) (*entities.User, error) {
	err := r.Database.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
