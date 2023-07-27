package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"service_user_go/api/routes"
	"service_user_go/config"
	"service_user_go/pkg/entities"
	"service_user_go/pkg/user"

)

func main() {
	app := fiber.New()

	database := ConnectDB()
	rabbitmq := ConnectRabbitmq()
	userRepo := user.NewRepo(database, rabbitmq)
	userService := user.NewService(userRepo)

	api := app.Group("/api", logger.New())
	routes.Routes(api, userService)

	app.Listen(":3001")
}

func ConnectDB() *gorm.DB {
	var DB *gorm.DB

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error parsing DB port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to open database!")
	}

	fmt.Println("Connection Opened to Database: ")

	DB.AutoMigrate(&entities.User{})
	fmt.Println("Migration table users completed")

	return DB
}

func ConnectRabbitmq() *amqp.Connection {
	connect, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}

	return connect
}
