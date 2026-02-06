package controller

import (
	"go-fiber-snippets/common"

	"github.com/gofiber/fiber/v3"
)

func HelloWorld(c fiber.Ctx) error {
	response := common.ResponseDTO[string]{
		Data:   "Hello, World!",
		Status: common.SUCCESS,
		Error:  nil,
	}
	return c.JSON(response)
}
