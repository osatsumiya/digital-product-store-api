package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Recommendation struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Format      string  `json:"format"`
}

func main() {
	r := gin.Default()

	r.GET("/recommendations/:productId", func(c *gin.Context) {
		productId := c.Param("productId")

		recommendations := []Recommendation{
			{
				ID:          101,
				Title:       "Premium UI Kit Bundle",
				Description: "Recommended bundle of UI kits for modern web and mobile projects.",
				Price:       39.99,
				Format:      "Figma",
			},
			{
				ID:          102,
				Title:       "Design System Starter",
				Description: "A starter design system with typography, colors, and reusable components.",
				Price:       24.99,
				Format:      "Figma",
			},
			{
				ID:          103,
				Title:       "Portfolio Template Pack",
				Description: "A collection of clean portfolio layouts for designers and developers.",
				Price:       19.99,
				Format:      "HTML",
			},
		}

		c.JSON(http.StatusOK, gin.H{
			"product_id":      productId,
			"recommendations": recommendations,
			"service":         "recommendation-service",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "running",
			"service": "recommendation-service",
		})
	})

	r.Run(":8083")
}
