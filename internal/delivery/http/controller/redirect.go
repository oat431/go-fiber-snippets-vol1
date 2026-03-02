package controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func ToLinkedIn(c fiber.Ctx) error {
	log.Info(c.Method() + " " + c.Path() + "\n")
	return c.Redirect().Status(301).To("https://www.linkedin.com/in/sahachan-tippimwong/")
}

func ToGitHub(c fiber.Ctx) error {
	log.Info(c.Method() + " " + c.Path() + "\n")
	return c.Redirect().Status(301).To("https://github.com/oat431")
}

func ToFacebook(c fiber.Ctx) error {
	log.Info(c.Method() + " " + c.Path() + "\n")
	return c.Redirect().Status(301).To("https://www.facebook.com/sahachan.tippimwong/")
}
