package repository

import (
	"fmt"
	"product-service/model"
	"testing"
)

func TestRepositoryItems(t *testing.T) {
	x := NewItemsRepository()

	Helmet := model.Items{
		Name:  "Helmet KYT",
		Qty:   1,
		Price: 450000,
	}

	res, err := x.Save(Helmet)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
