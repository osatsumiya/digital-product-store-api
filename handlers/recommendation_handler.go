package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func getRecommendationURL() string {
	url := os.Getenv("RECOMMENDATION_SERVICE_URL")
	if url == "" {
		return "http://localhost:8083"
	}
	return url
}

func GetProductRecommendations(c *gin.Context) {
	productId := c.Param("id")

	log.Println("Calling recommendation-service for product:", productId)

	client := resty.New()

	client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		log.Println("[Resty] Request:", r.Method, r.URL)
		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		log.Println("[Resty] Response status:", r.Status())
		return nil
	})

	url := getRecommendationURL() + "/recommendations/" + productId

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to call recommendation service",
		})
		return
	}

	c.Data(http.StatusOK, "application/json", resp.Body())
}
