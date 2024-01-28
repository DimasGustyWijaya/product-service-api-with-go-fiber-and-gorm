package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"product-service/model/web"
	"product-service/service"
)

type ItemsControllerImpl struct {
	service service.ItemsService
}

func NewItemsController(itemsService service.ItemsService) ItemsController {
	return &ItemsControllerImpl{
		service: itemsService,
	}
}

func (i ItemsControllerImpl) Create(c *fiber.Ctx) error {
	itemReq := web.ItemsRequest{}
	bodyReq := c.Body()

	errDecode := json.Unmarshal(bodyReq, &itemReq)
	if errDecode != nil {
		panic(errDecode)
	}

	itemsResponse := i.service.Save(itemReq)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   itemsResponse,
	}

	result, errEncode := json.Marshal(webResponse)
	if errEncode != nil {
		panic(errEncode)
	}

	return c.Send(result)

}
