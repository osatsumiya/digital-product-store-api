package handlers

import (
	"digital-product-store-api/database"
	"digital-product-store-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if order.CustomerID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer_id is required"})
		return
	}

	if order.ProductID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id is required"})
		return
	}

	if order.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity must be greater than 0"})
		return
	}

	var customer models.Customer
	if err := database.DB.First(&customer, order.CustomerID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, order.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	order.TotalPrice = product.Price * float64(order.Quantity)

	if order.Status == "" {
		order.Status = "created"
	}

	database.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}
