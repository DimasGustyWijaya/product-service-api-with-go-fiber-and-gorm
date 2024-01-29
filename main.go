package main

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"product-service/app"
	"product-service/controller"
	"product-service/repository"
	"product-service/service"
)

func main() {
	router := fiber.New()
	DB := app.GetConnect()
	validate := validator.New()

	itemRepository := repository.NewItemsRepository()
	itemService := service.NewItemsService(DB, itemRepository, validate)
	itemController := controller.NewItemsController(itemService)

	router.Post("/api/addItems", itemController.Create)
	router.Put("/api/updateItems/:id", itemController.Update)
	router.Get("/api/getItems/:id", itemController.FindByID)
	router.Get("/api/getItems", itemController.FindAll)
	router.Delete("/api/deleteItems/:id", itemController.Delete)

	err := router.Listen("localhost:5600")
	if err != nil {
		panic(err)
	}
}
