package handler

import (
	"github.com/gofiber/fiber/v3"
)

// Redirect handlers — pattern for 301 permanent redirects.
// Use 302 (c.Redirect().To(...)) for temporary redirects.

func RedirectLinkedIn(c fiber.Ctx) error {
	return c.Redirect().Status(301).To("https://www.linkedin.com/in/sahachan-tippimwong/")
}

func RedirectGitHub(c fiber.Ctx) error {
	return c.Redirect().Status(301).To("https://github.com/oat431")
}

func RedirectFacebook(c fiber.Ctx) error {
	return c.Redirect().Status(301).To("https://www.facebook.com/sahachan.tippimwong/")
}
