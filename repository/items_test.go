package repository

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestRepositoryItems(t *testing.T) {
	x := NewItemsRepository()

	ctx := fiber.Ctx{}

	res := x.FindAll(&ctx)

	fmt.Println(res)
}

func TestRepositoryItemsUpdate(t *testing.T) {
	//x := NewItemsRepository()
	//
	//
	//pr, err := x.FindById(c, 1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(pr)

}
