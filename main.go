package main

import (
	"digital-product-store-api/database"
	"digital-product-store-api/handlers"
	"digital-product-store-api/middleware"

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

	// Auth
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Products
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProductByID)
	r.POST("/products", handlers.CreateProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	// Customers
	r.GET("/customers", handlers.GetCustomers)
	r.POST("/customers", handlers.CreateCustomer)

	// Orders
	r.GET("/orders", handlers.GetOrders)
	r.POST("/orders", handlers.CreateOrder)

	// Licenses
	r.GET("/licenses", handlers.GetLicenses)
	r.POST("/licenses", handlers.CreateLicense)

	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)

	// Favorites (protected routes)
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/favorites/products", handlers.GetFavoriteProducts)
		auth.PUT("/favorites/products/:productId", handlers.AddProductToFavorites)
		auth.DELETE("/favorites/products/:productId", handlers.RemoveProductFromFavorites)
	}

	r.Run(":8080")
}
