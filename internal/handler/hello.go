package handler

import (
	"go-fiber-snippets/pkg/common"

	"github.com/gofiber/fiber/v3"
)

// Hello returns a simple greeting — useful as a smoke-test endpoint.
func Hello(c fiber.Ctx) error {
	return c.JSON(common.OK("Hello, World!"))
}
