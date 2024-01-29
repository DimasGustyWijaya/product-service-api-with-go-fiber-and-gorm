package repository

import (
	"github.com/gofiber/fiber/v2"
	"product-service/model"
)

type ItemsRepository interface {
	Save(ctx *fiber.Ctx, items model.Items) (model.Items, error)
	Update(ctx *fiber.Ctx, items model.Items) (model.Items, error)
	FindById(ctx *fiber.Ctx, id int) (model.Items, error)
	Delete(ctx *fiber.Ctx, id int) error
	FindAll(ctx *fiber.Ctx) []model.Items
}
