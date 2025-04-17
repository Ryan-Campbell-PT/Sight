package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type BodyResponse struct {
	FoodListString string `json:"foodListString"`
	Date           string `json:"date"`
	SaveToDb       bool   `json:"saveToDb"`
}

type RecipeResponse struct {
	RecipeName string `json:"recipeName"`
	FoodString string `json:"foodString"`
	Servings   int64  `json:"servings"`
}

func readRequestBody(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	data, err := io.ReadAll(body)
	return data, err
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
	bodyJson, err := readRequestBody(c.Request.Body)
	if handleError("Error reading query request body: ", err) {
		return
	}

	var bodyObj BodyResponse
	err = json.Unmarshal(bodyJson, &bodyObj)
	if handleError("Error reading body from query request: ", err) {
		return
	}

	request, err := buildNutritionixRequest(bodyObj.FoodListString)
	if handleError("Error building Nutritionix request: ", err) {
		return
	}

	responseByteArray, err := sendRequest(request)
	if handleError("Error sending Nutritionix request: ", err) {
		return
	}

	if bodyObj.SaveToDb {
		err := saveToDatabase_BodyResponse(bodyObj)
		if handleError("Error saving bodyObj to database: ", err) {
			return
		}
	}

	c.JSON(http.StatusOK, string(responseByteArray))
}

func post_saveRecipe(c *gin.Context) {
	body, err := readRequestBody(c.Request.Body)
	if handleError("Error reading recipe request body: ", err) {
		return
	}
	fmt.Println("Past request body")

	var recipeObj RecipeResponse
	err = json.Unmarshal(body, &recipeObj)
	if handleError("Error reading body from recipe request: ", err) {
		return
	}
	fmt.Println("past body recipe request")

	request, err := buildNutritionixRequest(recipeObj.FoodString)
	if handleError("Error building nutritionix request from recipe: ", err) {
		return
	}
	fmt.Println("past build nutritoin request")

	//this contains the nutrition information (should probably marshal that into a specified nutrition object)
	responseByteArray, err := sendRequest(request)
	if handleError("Error sending recipe nutritionix request: ", err) {
		return
	}
	fmt.Println("Past send request")

	var nutritionInfo Food
	err = json.Unmarshal(responseByteArray, &nutritionInfo)
	if handleError("Error reading nutrition info from nutritionix response and assigning to Food item: ", err) {
		return
	}
	fmt.Println("Past reading nutirtion info from response")

	err = saveToDatabase_RecipeResponse(recipeObj, nutritionInfo)
	if handleError("Error saving recipe information to database: ", err) {
		return
	}
	fmt.Println("Past saving ot recipe db")

	c.JSON(http.StatusOK, "")
}

func runServer() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	router.POST("/postFoodList", post_nutritionixQueryRequest)
	router.POST("/postRecipe", post_saveRecipe)

	router.Run(":8080")
}
