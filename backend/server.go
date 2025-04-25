package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TODO these have to be organized. all the classes made are a mess, some are for json
// some are not. should probably make a seperate file for all these definitions, to organize
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

type NutritionResponse struct {
	FoodInfo FoodResponse           `json:"foodInfo"`
	Errors   []NutritionErrorObject `json:"errors"`
}

type NutritionErrorObject struct {
	FoodString string `json:"foodString"`
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

	// jsonData := string(responseByteArray)
	response := NutritionResponse{}
	err = json.Unmarshal(responseByteArray, &response.FoodInfo)
	handleError("Error unmarshaling json data: ", err)

	splitByComma := strings.Split(bodyObj.FoodListString, ",")
	// if the api return information that is less than that of what the user typed in, there was an error somewhere
	if len(splitByComma) > len(response.FoodInfo.Foods) {
		errorList := []NutritionErrorObject{}
		foodListIndex := 0

		for _, inputStr := range splitByComma {
			inputStrTrimmed := strings.TrimSpace(inputStr)

			// If all known foods have been matched, everything else is an error
			if foodListIndex >= len(response.FoodInfo.Foods) {
				errorList = append(errorList, NutritionErrorObject{FoodString: inputStrTrimmed})
				continue
			}

			// Match input with the current food from response
			foodName := response.FoodInfo.Foods[foodListIndex].FoodName
			if strings.Contains(inputStrTrimmed, foodName) {
				foodListIndex++
			} else {
				errorList = append(errorList, NutritionErrorObject{FoodString: inputStrTrimmed})
			}
		}

		response.Errors = errorList
	}

	jsonData, err := json.Marshal(response)
	if handleError("Error marshaling NutritionResponse: ", err) {
		return
	}

	c.JSON(http.StatusOK, string(jsonData))
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
