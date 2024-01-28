package service

import (
	"github.com/go-playground/validator"
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

func (i ItemsServiceImpl) Save(request web.ItemsRequest) web.ItemsResponse {
	errRequest := i.validate.Struct(request)
	if errRequest != nil {
		panic(errRequest)
	}

	items := model.Items{
		Name:  request.Name,
		Price: request.Price,
		Qty:   request.Qty,
	}

	result, err := i.repository.Save(items)
	if err != nil {
		panic(err)
	}

	return helper.ToItemsResponse(result)

}
