package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"service_auth_go/api/presenter"
	"service_auth_go/config"
	"service_auth_go/pkg/entities"
)

type Service interface {
	SigninService(request *presenter.AuthRequest) (*entities.Response, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) SigninService(request *presenter.AuthRequest) (*entities.Response, error) {
	data, err := s.repository.SigninRepository(request)
	if err != nil {
		return nil, err
	}

	if data.Status != true {
		return nil, errors.New("Username or password wrong!")
	}

	if !CheckPasswordHash(request.Password, data.Result.Password) {
		return nil, errors.New("Username or password wrong!")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = data.Result.Username
	claims["user_id"] = data.Result.Uuid
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return nil, err
	}

	response := entities.Response{
		Username:    data.Result.Username,
		AccessToken: t,
	}

	return &response, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
