package helper

import (
	"product-service/model"
	"product-service/model/web"
)

func ToItemsResponse(items model.Items) web.ItemsResponse {
	return web.ItemsResponse{
		Id:    items.Id,
		Name:  items.Name,
		Price: items.Price,
		Qty:   items.Qty,
	}
}
