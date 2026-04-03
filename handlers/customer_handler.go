package handlers

import (
	"digital-product-store-api/database"
	"digital-product-store-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	database.DB.Find(&customers)
	c.JSON(http.StatusOK, customers)
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if customer.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "full_name is required"})
		return
	}

	if customer.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	database.DB.Create(&customer)
	c.JSON(http.StatusCreated, customer)
}
