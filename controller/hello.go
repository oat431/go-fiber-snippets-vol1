package controller

import (
	"go-fiber-snippets/common"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func HelloWorld(c fiber.Ctx) error {
	response := common.ResponseDTO[string]{
		Data:   "Hello, World!",
		Status: common.SUCCESS,
		Error:  nil,
	}
	log.Info(c.Method() + " " + c.Path() + "\n")
	return c.JSON(response)
}
