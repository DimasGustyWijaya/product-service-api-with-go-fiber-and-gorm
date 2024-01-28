package service

import "product-service/model/web"

type ItemsService interface {
	Save(request web.ItemsRequest) web.ItemsResponse
}
