package repository

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"product-service/app"
	"product-service/model"
)

type ItemsRepositoryImpl struct {
}

func NewItemsRepository() ItemsRepository {
	return &ItemsRepositoryImpl{}
}

func (i *ItemsRepositoryImpl) Save(ctx *fiber.Ctx, items model.Items) (model.Items, error) {
	DB := app.GetConnect()

	itemProduct := model.Items{
		Name:  items.Name,
		Price: items.Price,
		Qty:   items.Qty,
	}

	err := DB.Transaction(func(tx *gorm.DB) error {
		tx.WithContext(ctx.Context()).Create(&itemProduct)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return itemProduct, nil

}

func (i *ItemsRepositoryImpl) Update(ctx *fiber.Ctx, items model.Items) (model.Items, error) {
	DB := app.GetConnect()

	prodUpdate := model.Items{}

	err := DB.Transaction(func(tx *gorm.DB) error {
		tx.WithContext(ctx.Context()).Model(&prodUpdate).Where("id = ?", items.Id).Updates(model.Items{
			Id:    items.Id,
			Name:  items.Name,
			Price: items.Price,
			Qty:   items.Qty,
		})

		return nil
	})
	if err != nil {
		DB.Rollback()
	} else {
		DB.Commit()
	}

	return prodUpdate, nil
}

func (i *ItemsRepositoryImpl) FindById(ctx *fiber.Ctx, id int) (model.Items, error) {
	DB := app.GetConnect()

	items := model.Items{}

	err := DB.Transaction(func(tx *gorm.DB) error {
		tx.WithContext(ctx.Context()).First(&items, id)

		return nil
	})
	if err != nil {
		panic(err)
	}

	return items, nil
}

func (i *ItemsRepositoryImpl) Delete(ctx *fiber.Ctx, id int) error {
	DB := app.GetConnect()

	items := model.Items{}

	err := DB.Transaction(func(tx *gorm.DB) error {
		tx.WithContext(ctx.Context()).Delete(&items, id)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func (i *ItemsRepositoryImpl) FindAll(ctx *fiber.Ctx) []model.Items {
	DB := app.GetConnect()

	items := []model.Items{}

	err := DB.Transaction(func(tx *gorm.DB) error {
		tx.Find(&items)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return items

}
