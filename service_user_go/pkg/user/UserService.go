package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"service_user_go/api/presenter"
	"service_user_go/pkg/entities"
)

type Service interface {
	IndexService() (*[]presenter.User, error)
	CreateService(user *entities.User) (*entities.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) IndexService() (*[]presenter.User, error) {
	return s.repository.IndexRepository()
}

func (s *service) CreateService(user *entities.User) (*entities.User, error) {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.ID = uuid.New()
	user.Password = hash

	return s.repository.CreateRepository(user)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
