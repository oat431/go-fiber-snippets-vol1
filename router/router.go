package router

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func init() {
	log.Println("Creating router")
}

func StartServer() {
	app := fiber.New()
	err := app.Listen("8080")
	if err != nil {
		return
	}
}
