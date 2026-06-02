package middleware

import "github.com/gofiber/fiber/v3"

// WSGuard restricts WebSocket upgrades to allowed hosts.
// Usage: app.Use("/chat", middleware.WSGuard)
func WSGuard(c fiber.Ctx) error {
	host := c.Get("Host")
	if host == "localhost:8000" || host == "localhost:8080" {
		c.Locals("Host", host)
		return c.Next()
	}
	return c.Status(fiber.StatusForbidden).SendString("origin not allowed")
}
