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

func postFoodList(c *gin.Context) {
    // Http request would be formatted like /postFoodList?foodListString=
    // foodListString := c.Query("foodListString")
    // TODO this could be an error point, if requst.body doesnt have the data needed
    cfg := getEnvironmentVariables()
    foodListString := io.ReadAll(c.Request.Body)
    foodListQuery := map[string]string {"query": foodListString}
    foodListJsonByteArray, _ := json.Marshal(foodListQuery)
    url := cfg.Nutritionix_domain + cfg.Nutritionix_naturalLanguage
    request, err := http.NewRequest("POST", url, bytes.NewBuffer(foodListJsonByteArray))
    if handleError(err, "Error creating request: ") return

    request.Header.Set("Content-Type", cfg.Nutritionix_contentType)
    request.Header.Set("x-app-id", cfg.Nutritionix_appid)
    request.Header.Set("x-app-key", cfg.Nutritionix_appkey)

    client := &http.Client{}
    response, err := client.Do(request)
    if handleError(err, "Error sending request: ") return
    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)
    if handleError(err, "Error reading response body: ") return

    dailyNutrition := makeFoodResponse(body)
    if dailyNutrition == nil {
        // TODO Might be an error here due to needing multiple parameters or bad StatusError
        return c.JSON(http.StatusError)
    }
    fmt.Println("Response body:", string(body))
    // TODO changed this from body to dailyNutrition, could cause error
    c.JSON(http.StatusOK, dailyNutrition)
}

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



    router.GET("/nutrition", func(c *gin.Context) {
        data, err := fetchData()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
            return
        }
        c.JSON(http.StatusOK, data)
    })

    router.POST("/postFoodList", postFoodList)

    router.Run(":8080")
}