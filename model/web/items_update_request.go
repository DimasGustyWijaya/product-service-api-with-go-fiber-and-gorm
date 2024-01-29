package web

type ItemsUpdateRequest struct {
	Id, Qty int32
	Name    string
	Price   int
}
