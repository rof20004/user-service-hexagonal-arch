package main

import (
	"log"
	userService "user-service/application/services"

	"github.com/gofiber/fiber/v2"
	userRepository "user-service/adapters/database/memory/user"
	userHandler "user-service/adapters/handlers"
)

func main() {
	var (
		database = userRepository.NewUserMemoryDatabase()
		service  = userService.NewUserService(database)
		handler  = userHandler.NewHttpHandler(service)
	)

	app := fiber.New()

	app.Post("/users", handler.Create)
	app.Get("/users/:id", handler.GetById)

	log.Fatal(app.Listen(":8080"))
}
