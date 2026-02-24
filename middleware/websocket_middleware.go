package middleware

import "github.com/gofiber/fiber/v3"

func WebSocketMiddleware(c fiber.Ctx) error {
	if c.Get("Host") == "localhost:8000" {
		c.Locals("Host", "localhost:8000")
		return c.Next()
	}
	return c.Status(fiber.StatusForbidden).SendString("Request origin not allowed")
}
