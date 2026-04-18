package main

import (
	"digital-product-store-api/database"
	"digital-product-store-api/handlers"
	"digital-product-store-api/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:5174",
			"http://127.0.0.1:5173",
			"http://127.0.0.1:5174",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Luzmuerta Studio API is running",
		})
	})

	// Auth
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Users
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)

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

	// Favorites
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/favorites/products", handlers.GetFavoriteProducts)
		auth.PUT("/favorites/products/:productId", handlers.AddProductToFavorites)
		auth.DELETE("/favorites/products/:productId", handlers.RemoveProductFromFavorites)
	}

	r.Run(":8080")
}
