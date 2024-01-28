package controller

import (
	"github.com/gofiber/fiber/v2"
)

type ItemsController interface {
	Create(c *fiber.Ctx) error
}
