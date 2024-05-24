package main

import (
	"github.com/fauzanfl/belajar-go-restapi/controllers/productController"
	"github.com/fauzanfl/belajar-go-restapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := models.Config{}

	config.ConnectDB()

	r.GET("products", productController.Index)
	r.GET("products/:id", productController.Show)
	r.POST("products/create", productController.Create)
	r.PUT("products/update/:id", productController.Update)
	r.DELETE("products/delete/:id", productController.Delete)

	r.Run()
}
