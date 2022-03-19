package handlers

import (
	"user-service/ports"

	"github.com/gofiber/fiber/v2"
)

type HttpHandler struct {
	userService ports.Usecase
}

func NewHttpHandler(userService ports.Usecase) *HttpHandler {
	return &HttpHandler{userService}
}

func (handler *HttpHandler) Create(c *fiber.Ctx) error {
	var dto createUserDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	user, err := handler.userService.Create(dto.Name, dto.Email, dto.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	view := viewUserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.Status(fiber.StatusOK).JSON(&view)
}

func (handler *HttpHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := handler.userService.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	view := viewUserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.Status(fiber.StatusOK).JSON(&view)
}
