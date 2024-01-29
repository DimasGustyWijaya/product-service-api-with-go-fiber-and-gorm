package controller

import (
	"github.com/gofiber/fiber/v2"
	"product-service/model/web"
	"product-service/service"
	"strconv"
)

type ItemsControllerImpl struct {
	service service.ItemsService
}

func NewItemsController(itemsService service.ItemsService) ItemsController {
	return &ItemsControllerImpl{
		service: itemsService,
	}
}

func (i *ItemsControllerImpl) Create(c *fiber.Ctx) error {
	itemReq := web.ItemsRequest{}
	err := c.BodyParser(&itemReq)
	if err != nil {
		panic(err)
	}

	itemsResponse := i.service.Save(c, itemReq)

	return c.Status(200).JSON(itemsResponse)

}

func (i *ItemsControllerImpl) Update(c *fiber.Ctx) error {
	itemUpdateReq := web.ItemsUpdateRequest{}

	err := c.BodyParser(&itemUpdateReq)
	if err != nil {
		panic(err)
	}

	id, err := strconv.Atoi(c.Params("id"))

	itemUpdateReq.Id = int32(id)

	itemsResponse := i.service.Update(c, itemUpdateReq)

	return c.Status(200).JSON(itemsResponse)

}

func (i *ItemsControllerImpl) FindByID(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}

	result := i.service.FindById(c, id)

	return c.Status(200).JSON(result)

}

func (i *ItemsControllerImpl) Delete(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}

	err = i.service.Delete(c, id)

	return c.Status(200).SendStatus(200)

}

func (i *ItemsControllerImpl) FindAll(c *fiber.Ctx) error {

	result := i.service.FindAll(c)

	return c.Status(200).JSON(result)

}
