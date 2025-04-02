package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"io"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func fetchData() (map[string]interface{}, error) {
    url := "https://jsonplaceholder.typicode.com/todos/1"

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var data map[string]interface{}
    json.Unmarshal(body, &data)

    return data, nil
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


    cfg := getEnvironmentVariables()

    router.GET("/nutrition", func(c *gin.Context) {
        data, err := fetchData()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
            return
        }
        c.JSON(http.StatusOK, data)
    })

    router.POST("/postFoodList", func(ctx *gin.Context) {
        foodListQuery := map[string]string {"query": "1 banana"}
        foodListJsonByteArray, _ := json.Marshal(foodListQuery)
        url := cfg.Nutritionix_domain + cfg.Nutritionix_naturalLanguage
        request, err := http.NewRequest("POST", url, bytes.NewBuffer(foodListJsonByteArray))
        if err != nil {
            fmt.Println("Error creating request: ", err)
            return
        }
        request.Header.Set("Content-Type", cfg.Nutritionix_contentType)
        request.Header.Set("x-app-id", cfg.Nutritionix_appid)
        request.Header.Set("x-app-key", cfg.Nutritionix_appkey)

        client := &http.Client{}
        response, err := client.Do(request)
        if err != nil {
            fmt.Println("Error sending request: ", err)
            return
        }
        defer response.Body.Close()

        body, err := io.ReadAll(response.Body)
        if err != nil {
            fmt.Println("Error reading response body:", err)
            return
        }
            fmt.Println("Response body:", string(body))
        ctx.JSON(http.StatusOK, response.Body)
    });

    router.Run(":8080")
}