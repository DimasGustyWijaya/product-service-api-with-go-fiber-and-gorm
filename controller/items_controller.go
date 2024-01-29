package controller

import (
	"github.com/gofiber/fiber/v2"
)

type ItemsController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}
