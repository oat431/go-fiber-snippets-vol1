package controller

import (
	"go-fiber-snippets/internal/domain"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func HelloWorld(c fiber.Ctx) error {
	response := domain.ResponseDTO[string]{
		Data:   "Hello, World!",
		Status: domain.SUCCESS,
		Error:  nil,
	}
	log.Info(c.Method() + " " + c.Path() + "\n")
	return c.JSON(response)
}
