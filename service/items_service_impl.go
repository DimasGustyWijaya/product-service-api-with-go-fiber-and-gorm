package service

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"product-service/helper"
	"product-service/model"
	"product-service/model/web"
	"product-service/repository"
)

type ItemsServiceImpl struct {
	DB         *gorm.DB
	repository repository.ItemsRepository
	validate   *validator.Validate
}

func NewItemsService(DB *gorm.DB, itemsRepository repository.ItemsRepository, validate *validator.Validate) ItemsService {
	return &ItemsServiceImpl{
		DB:         DB,
		repository: itemsRepository,
		validate:   validate,
	}
}

func (i ItemsServiceImpl) Save(ctx *fiber.Ctx, request web.ItemsRequest) web.ItemsResponse {
	errRequest := i.validate.Struct(request)
	if errRequest != nil {
		panic(errRequest)
	}

	items := model.Items{
		Name:  request.Name,
		Price: request.Price,
		Qty:   request.Qty,
	}

	result, err := i.repository.Save(ctx, items)
	if err != nil {
		panic(err)
	}

	return helper.ToItemsResponse(result)

}

func (i *ItemsServiceImpl) Update(ctx *fiber.Ctx, request web.ItemsUpdateRequest) web.ItemsResponse {
	if errUpdateRequest := i.validate.Struct(request); errUpdateRequest != nil {
		panic(errUpdateRequest)
	}

	itemsUpdate := model.Items{
		Id:    request.Id,
		Name:  request.Name,
		Qty:   request.Qty,
		Price: request.Price,
	}

	result, err := i.repository.Update(ctx, itemsUpdate)
	if err != nil {
		panic(err)
	}

	return helper.ToItemsResponse(result)

}

func (i *ItemsServiceImpl) FindById(ctx *fiber.Ctx, id int) web.ItemsResponse {

	result, err := i.repository.FindById(ctx, id)
	if err != nil {
		panic(err)
	}

	findProduct := web.ItemsResponse{
		Id:    int32(id),
		Name:  result.Name,
		Price: result.Price,
		Qty:   result.Qty,
	}

	return findProduct
}

func (i *ItemsServiceImpl) Delete(ctx *fiber.Ctx, id int) error {

	err := i.repository.Delete(ctx, id)
	if err != nil {
		panic(err)
	}

	return nil
}

func (i *ItemsServiceImpl) FindAll(ctx *fiber.Ctx) []web.ItemsResponse {

	result := i.repository.FindAll(ctx)

	return helper.ToItemsResponses(result)
}
