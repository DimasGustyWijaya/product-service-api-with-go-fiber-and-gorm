package service

import (
	"github.com/gofiber/fiber/v2"
	"product-service/model/web"
)

type ItemsService interface {
	Save(ctx *fiber.Ctx, request web.ItemsRequest) web.ItemsResponse
	Update(ctx *fiber.Ctx, request web.ItemsUpdateRequest) web.ItemsResponse
	FindById(ctx *fiber.Ctx, id int) web.ItemsResponse
	Delete(ctx *fiber.Ctx, id int) error
	FindAll(ctx *fiber.Ctx) []web.ItemsResponse
}
