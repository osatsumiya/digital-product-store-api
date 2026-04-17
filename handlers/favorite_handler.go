package handlers

import (
	"digital-product-store-api/database"
	"digital-product-store-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFavoriteProducts(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in token"})
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id type"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var total int64
	database.DB.Model(&models.FavoriteProduct{}).Where("user_id = ?", userID).Count(&total)

	var favorites []models.FavoriteProduct
	if err := database.DB.
		Preload("Product").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&favorites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch favorite products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"total": total,
		"items": favorites,
	})
}

func AddProductToFavorites(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in token"})
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id type"})
		return
	}

	productIDParam := c.Param("productId")
	productID64, err := strconv.ParseUint(productIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	productID := uint(productID64)

	var product models.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	var existingFavorite models.FavoriteProduct
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&existingFavorite).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message":  "product is already in favorites",
			"favorite": existingFavorite,
		})
		return
	}

	favorite := models.FavoriteProduct{
		UserID:    userID,
		ProductID: productID,
	}

	if err := database.DB.Create(&favorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add product to favorites"})
		return
	}

	if err := database.DB.Preload("Product").First(&favorite, favorite.ID).Error; err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "product added to favorites",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "product added to favorites",
		"favorite": favorite,
	})
}

func RemoveProductFromFavorites(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in token"})
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id type"})
		return
	}

	productIDParam := c.Param("productId")
	productID64, err := strconv.ParseUint(productIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	productID := uint(productID64)

	var favorite models.FavoriteProduct
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&favorite).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "favorite product not found"})
		return
	}

	if err := database.DB.Delete(&favorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove product from favorites"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product removed from favorites",
	})
}
