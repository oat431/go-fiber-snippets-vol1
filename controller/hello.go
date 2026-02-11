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
	//log.Info("📝 LOG: %s %s | Status: %d \n",
	//	c.Method(),                // เช่น GET
	//	c.Path(),                  // เช่น /health
	//	c.Response().StatusCode(), // เช่น 200
	//)
	log.Info(c.Method() + " " + c.Path() + "\n")
	return c.JSON(response)
}
