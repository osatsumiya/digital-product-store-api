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

	var existingCustomer models.Customer
	if err := database.DB.Where("email = ?", customer.Email).First(&existingCustomer).Error; err == nil {
		c.JSON(http.StatusOK, existingCustomer)
		return
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, customer)
}
