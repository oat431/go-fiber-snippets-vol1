package controller

import (
	"go-fiber-snippets/pkg/common"

	"github.com/gofiber/fiber/v3"
)

func TriggerHookExample(c fiber.Ctx) error {
	res := common.ResponseDTO[string]{
		Data:   "Hook triggered successfully",
		Status: common.SUCCESS,
		Error:  nil,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
