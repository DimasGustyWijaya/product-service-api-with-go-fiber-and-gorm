package app

import (
	"fmt"
	"product-service/model"
	"testing"
)

func TestDB(t *testing.T) {
	db := GetConnect()

	pc := &model.Items{
		Name:  "Asus ROG PHONE",
		Price: 10000000,
		Qty:   2,
	}

	result := db.Create(&pc)

	fmt.Println(result)

}
