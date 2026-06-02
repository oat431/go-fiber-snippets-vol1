package handler

import (
	"go-fiber-snippets/pkg/common"

	"github.com/gofiber/fiber/v3"
)

// Webhook is a simple inbound hook endpoint.
// In production: validate a signature header, parse the payload,
// enqueue a job, and return 200 quickly.
func Webhook(c fiber.Ctx) error {
	// Example: read raw body for signature verification
	// body := c.Body()
	// if !verifySignature(c.Get("X-Signature"), body) { ... }

	return c.Status(fiber.StatusOK).JSON(common.OK("hook accepted"))
}
