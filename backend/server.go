package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func readRequestBody(body io.ReadCloser) (string, error) {
	defer body.Close()
	data, err := io.ReadAll(body)
	return string(data), err
}

func buildNutritionixRequest(foodList string) (*http.Request, error) {
	cfg := getEnvironmentVariables()
	foodQuery := map[string]string{"query": foodList}
	body, err := json.Marshal(foodQuery)
	if err != nil {
		return nil, err
	}

	url := cfg.Nutritionix_domain + cfg.Nutritionix_naturalLanguage
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", cfg.Nutritionix_contentType)
	request.Header.Set("x-app-id", cfg.Nutritionix_appid)
	request.Header.Set("x-app-key", cfg.Nutritionix_appkey)

	return request, nil
}

func sendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func post_nutritionixQueryRequest(c *gin.Context) {
	foodListString, err := readRequestBody(c.Request.Body)
	if handleError("Error reading request body: ", err) {
		return
	}

	request, err := buildNutritionixRequest(foodListString)
	if handleError("Error building Nutritionix request: ", err) {
		return
	}

	responseByteArray, err := sendRequest(request)
	if handleError("Error sending Nutritionix request: ", err) {
		return
	}

	c.JSON(http.StatusOK, string(responseByteArray))
}

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	router.POST("/postFoodList", post_nutritionixQueryRequest)

	router.Run(":8080")
}
