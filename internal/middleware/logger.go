package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

// RequestLogger logs method + path for every request.
// Attach globally: app.Use(middleware.RequestLogger())
func RequestLogger() fiber.Handler {
	return func(c fiber.Ctx) error {
		log.Infof("%s %s", c.Method(), c.Path())
		return c.Next()
	}
}
