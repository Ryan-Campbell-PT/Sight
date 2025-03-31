package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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

    router.GET("/", func(c *gin.Context) {
        data, err := fetchData()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
            return
        }
        c.JSON(http.StatusOK, data)
    })

    router.Run(":8080")
}