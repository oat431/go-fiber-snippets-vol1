package controller

import (
	"go-fiber-snippets/internal/domain"

	"github.com/gofiber/fiber/v3"
)

func TriggerHookExample(c fiber.Ctx) error {
	res := domain.ResponseDTO[string]{
		Data:   "Hook triggered successfully",
		Status: domain.SUCCESS,
		Error:  nil,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
