package user

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"

	"service_user_go/api/presenter"
	"service_user_go/pkg/entities"

)

type Repository interface {
	IndexRepository() ([]presenter.User, error)
	CreateRepository(user *entities.User) (*entities.User, error)
	ShowByUsername(username string) (*entities.User, error)
	SendQueue(email string) error
}

type repository struct {
	Database *gorm.DB
	Rabbit   *amqp.Connection
}

func NewRepo(database *gorm.DB, rabbitmq *amqp.Connection) Repository {
	return &repository{
		Database: database,
		Rabbit:   rabbitmq,
	}
}

func (r *repository) IndexRepository() ([]presenter.User, error) {
	var users []presenter.User
	r.Database.Find(&users)

	return users, nil
}

func (r *repository) CreateRepository(user *entities.User) (*entities.User, error) {
	err := r.Database.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) ShowByUsername(username string) (*entities.User, error) {
	var user *entities.User

	r.Database.Where("username = ?", username).First(&user)
	return user, nil
}

func (r *repository) SendQueue(phoneNumber string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rabbitmq, err := r.Rabbit.Channel()
	if err != nil {
		return err
	}
	defer rabbitmq.Close()

	_, err = rabbitmq.QueueDeclare(
		"WAQueue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	err = rabbitmq.PublishWithContext(
		ctx,
		"",
		"WAQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(phoneNumber),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
