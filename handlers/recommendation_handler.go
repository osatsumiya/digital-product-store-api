package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func GetProductRecommendations(c *gin.Context) {
	productId := c.Param("id")

	// 🔹 ЛОГ (из лекции)
	log.Println("Calling recommendation-service for product:", productId)

	client := resty.New()

	// 🔹 Resty middleware (hooks)
	client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		log.Println("[Resty] Request:", r.Method, r.URL)
		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		log.Println("[Resty] Response status:", r.Status())
		return nil
	})

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get("http://localhost:8083/recommendations/" + productId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to call recommendation service",
		})
		return
	}

	c.Data(http.StatusOK, "application/json", resp.Body())
}
