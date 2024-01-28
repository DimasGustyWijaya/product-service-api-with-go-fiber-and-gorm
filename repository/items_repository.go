package repository

import (
	"product-service/model"
)

type ItemsRepository interface {
	Save(items model.Items) (model.Items, error)
}
