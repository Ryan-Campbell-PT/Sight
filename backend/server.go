package main

import (
	"bytes"
	"encoding/json"
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
	RecipeName        string `json:"recipe_name"`
	FoodString        string `json:"food_string"`
	Servings          int64  `json:"serving_size"`
	NutritionValuesId int64  `json:"nutrition_id"`
}

type GetRecipeResponse struct {
	RecipeList []Recipe
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

// TODO this should be seperated from the request field, and should instead of returning json back to the front end
// should just return the nutrition info, and let the function that calls it deal with that
// it should return the byte array, not the string version
func post_nutritionixQueryRequest(c *gin.Context) {
	bodyJson, err := readRequestBody(c.Request.Body)
	if handleError("post_nutritionixQueryRequest/Error reading query request body: ", err) {
		return
	}

	var bodyObj BodyResponse
	err = json.Unmarshal(bodyJson, &bodyObj)
	if handleError("post_nutritionixQueryRequest/Error reading body from query request: ", err) {
		return
	}

	request, err := buildNutritionixRequest(bodyObj.FoodListString)
	if handleError("post_nutritionixQueryRequest/Error building Nutritionix request: ", err) {
		return
	}

	responseByteArray, err := sendRequest(request)
	if handleError("post_nutritionixQueryRequest/Error sending Nutritionix request: ", err) {
		return
	}

	if bodyObj.SaveToDb {
		var nutritionInfo FoodResponse
		err = json.Unmarshal(responseByteArray, &nutritionInfo)
		if handleError("post_nutritionixQueryRequest/Error reading nutrition info from nutritionix response and assigning to Food item: ", err) {
			return
		}

		err := saveToDatabase_BodyResponse(bodyObj, makeTotalNutritionData(nutritionInfo.Foods))
		if handleError("post_nutritionixQueryRequest/Error saving bodyObj to database: ", err) {
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

	var recipeObj RecipeResponse
	err = json.Unmarshal(body, &recipeObj)
	if handleError("Error reading body from recipe request: ", err) {
		return
	}

	request, err := buildNutritionixRequest(recipeObj.FoodString)
	if handleError("Error building nutritionix request from recipe: ", err) {
		return
	}

	//this contains the nutrition information (should probably marshal that into a specified nutrition object)
	responseByteArray, err := sendRequest(request)
	if handleError("Error sending recipe nutritionix request: ", err) {
		return
	}

	var nutritionInfo Food
	err = json.Unmarshal(responseByteArray, &nutritionInfo)
	if handleError("Error reading nutrition info from nutritionix response and assigning to Food item: ", err) {
		return
	}

	err = saveToDatabase_RecipeResponse(recipeObj, nutritionInfo)
	if handleError("Error saving recipe information to database: ", err) {
		return
	}

	c.JSON(http.StatusOK, "")
}

func get_recipes(c *gin.Context) {
	recipes, err := getFromDatabase_Recipes()
	if handleError("Server.go/Error getting recipes from database: ", err) {
		c.JSON(http.StatusBadRequest, "")
	}

	recipeJson, err := json.Marshal(recipes)
	if handleError("Error marshaling recipeResponse into recipeJson: ", err) {
		return
	}

	c.JSON(http.StatusOK, string(recipeJson))
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
	router.GET("/getRecipes", get_recipes)

	router.Run(":8080")
}
