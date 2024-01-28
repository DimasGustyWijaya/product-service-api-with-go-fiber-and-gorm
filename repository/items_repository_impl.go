package repository

import (
	"gorm.io/gorm"
	"product-service/app"
	"product-service/model"
)

type ItemsRepositoryImpl struct {
}

func NewItemsRepository() ItemsRepository {
	return &ItemsRepositoryImpl{}
}

type LastId struct {
	Id int32
}

func (i ItemsRepositoryImpl) Save(items model.Items) (model.Items, error) {
	DB := app.GetConnect()

	itemProduct := model.Items{
		Name:  items.Name,
		Price: items.Price,
		Qty:   items.Qty,
	}

	err := DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&itemProduct)
		return nil
	})
	if err != nil {
		DB.Rollback()
	} else {
		DB.Commit()
	}

	return itemProduct, nil

}
