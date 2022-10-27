package main

import (
	"Go-MongoDb-Api/app"
	"Go-MongoDb-Api/config"
	"Go-MongoDb-Api/repository"
	"Go-MongoDb-Api/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	appRoute := fiber.New()

	config.ConnectDB()
	dbClient := config.GetCollection(config.DB, "users")

	UserRepositoryDb := repository.NewUserRepositoryDb(dbClient)
	td := app.UserHandler{Service: services.NewUserService(UserRepositoryDb)}

	appRoute.Post("/api/user/", td.CreateUser)
	appRoute.Get("/api/users/", td.GetAllUser)
	appRoute.Delete("/api/user/:id", td.DeleteUser)
	appRoute.Listen(":8080")

}
