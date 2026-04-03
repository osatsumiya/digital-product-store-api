package main

import (
	"digital-product-store-api/database"
	"digital-product-store-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Digital Product Store API is running",
		})
	})

	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProductByID)
	r.POST("/products", handlers.CreateProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	r.GET("/customers", handlers.GetCustomers)
	r.POST("/customers", handlers.CreateCustomer)

	r.GET("/orders", handlers.GetOrders)
	r.POST("/orders", handlers.CreateOrder)

	r.GET("/licenses", handlers.GetLicenses)
	r.POST("/licenses", handlers.CreateLicense)

	r.Run(":8080")
}
