package handlers

import (
	"digital-product-store-api/database"
	"digital-product-store-api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLicenses(c *gin.Context) {
	var licenses []models.License
	database.DB.Preload("Product").Preload("Customer").Find(&licenses)
	c.JSON(http.StatusOK, licenses)
}

func CreateLicense(c *gin.Context) {
	var license models.License

	if err := c.ShouldBindJSON(&license); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if license.ProductID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id is required"})
		return
	}

	if license.CustomerID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer_id is required"})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, license.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	var customer models.Customer
	if err := database.DB.First(&customer, license.CustomerID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	if license.LicenseKey == "" {
		license.LicenseKey = fmt.Sprintf("LIC-%d-%d", license.ProductID, license.CustomerID)
	}

	license.IsActive = true

	database.DB.Create(&license)
	c.JSON(http.StatusCreated, license)
}
