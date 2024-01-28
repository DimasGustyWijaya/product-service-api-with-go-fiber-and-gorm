package web

type ItemsRequest struct {
	Name  string `validate:"min=1,max=100"`
	Price int
	Qty   int32
}
