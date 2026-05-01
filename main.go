package main

import (
	"digital-product-store-api/database"
	"digital-product-store-api/handlers"
	"digital-product-store-api/middleware"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func setupGorillaMuxRouter() *mux.Router {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/mux/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"message": "Hello from Gorilla Mux",
			"router":  "gorilla/mux",
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	muxRouter.HandleFunc("/mux/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"status":  "running",
			"project": "Luzmuerta Studio",
			"router":  "gorilla/mux",
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	return muxRouter
}

func main() {
	database.Connect()

	// Gorilla Mux server
	go func() {
		muxRouter := setupGorillaMuxRouter()

		log.Println("Gorilla Mux server is running on http://localhost:8081")

		if err := http.ListenAndServe(":8082", muxRouter); err != nil {
			log.Fatal("Failed to start Gorilla Mux server:", err)
		}
	}()

	// Gin server
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
			"router":  "gin",
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
	r.GET("/products/:id/recommendations", handlers.GetProductRecommendations)
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

	log.Println("Gin server is running on http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start Gin server:", err)
	}
}
