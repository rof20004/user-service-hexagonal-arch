package main

import (
	"log"

	userRepository "user-service/adapters/database/memory/user"
	userApp "user-service/application/user"
	userPort "user-service/ports/user"

	"github.com/gofiber/fiber/v2"
)

const (
	messageUserIdRequired = "user id is required"
)

func main() {
	var (
		userDb = userRepository.NewUserMemoryDatabase()
		userSv = userPort.NewUserService(userDb)
	)

	app := fiber.New()

	app.Post("/users", saveUser(userSv))
	app.Get("/users/:id", getUser(userSv))
	app.Put("/users/:id", updateUser(userSv))

	log.Fatal(app.Listen(":8080"))
}

func saveUser(service userPort.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dto userApp.CreateUserDto
		if err := c.BodyParser(&dto); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		result, err := service.Create(dto)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(&result)
	}
}

func getUser(service userPort.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id = c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).SendString(messageUserIdRequired)
		}

		result, err := service.GetById(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(&result)
	}
}

func updateUser(service userPort.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id = c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).SendString(messageUserIdRequired)
		}

		var dto userApp.UpdateUserDto
		if err := c.BodyParser(&dto); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		result, err := service.Update(id, dto)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(&result)
	}
}
