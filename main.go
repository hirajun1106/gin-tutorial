package main

import (
	"gin-tutorial/controllers"
	"gin-tutorial/models"
	"gin-tutorial/repositories"
	"gin-tutorial/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品名1", Price: 1000, Description: "商品説明1", SoldOut: false},
		{ID: 2, Name: "商品名2", Price: 2000, Description: "商品説明2", SoldOut: true},
		{ID: 3, Name: "商品名3", Price: 3000, Description: "商品説明3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)
	r.Run("localhost:8080")
}
