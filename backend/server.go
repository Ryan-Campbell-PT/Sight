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

type PostFoodList_RequestBody struct {
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

func buildNutritionixRequest_fromFoodListString(foodList string) (*http.Request, error) {
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

func sendHttpRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// this function will take the list of foods provided by a user
// and handle all the work associated with that string:
// marshaling/unmarshaling, reaching out to api, building response object
func buildNutritionixResponse_fromFoodListString(foodListString string) (Nutritionix_NaturalLanguageResponse, error) {
	functionName := "handle_naturalLanguage_foodList/"
	var nutritionInfo Nutritionix_NaturalLanguageResponse

	request, err := buildNutritionixRequest_fromFoodListString(foodListString)
	if handleError(functionName+"Error building Nutritionix request: ", err) {
		return nutritionInfo, err
	}

	responseByteArray, err := sendHttpRequest(request)
	if handleError(functionName+"Error sending Nutritionix request: ", err) {
		return nutritionInfo, err
	}

	err = json.Unmarshal(responseByteArray, &nutritionInfo)
	if handleError(functionName+"Error reading nutrition info from nutritionix response and unmarshaling to Food item: ", err) {
		return nutritionInfo, err
	}

	for i, food := range nutritionInfo.Foods {
		nMap := createNutrientMap(food.FullNutrients)
		food.FullNutrientMap = nMap
		nutritionInfo.Foods[i] = food
	}

	return nutritionInfo, nil
}

func getTotalNutritionInformation_fromFoodListString(foodListString string) (FoodItem, error) {
	functionName := "getTotalNutritionInformation_fromFoodListString/"
	nutritionInfo, err := buildNutritionixResponse_fromFoodListString(foodListString)
	if handleError(functionName+"Error getting response from foodliststring: ", err) {
		return FoodItem{}, err
	}

	return makeTotalNutritionData_fromFoodList(nutritionInfo.Foods), nil
}

func helper_checkFoodArrayForErrors(foodListString string, foods []FoodItem) []NutritionErrorObject {
	errorList := []NutritionErrorObject{}
	splitByComma := strings.Split(foodListString, ",")
	if len(splitByComma) > len(foods) {
		responseArrayIndex := 0
		for _, inputString := range splitByComma {
			inputStringTrimmed := strings.ToLower(strings.TrimSpace(inputString))

			// If all known foods have been matched, everything else is an error
			if responseArrayIndex >= len(foods) {
				errorList = append(errorList, NutritionErrorObject{ErrorString: inputStringTrimmed})
				continue
			}

			// TODO this will probably need to be looked at, as what was typed may be slightly different
			// than what the foodName actually is

			// if the string typed by the user contains the food recognized by the api
			foodName := foods[responseArrayIndex].FoodName
			if strings.Contains(inputStringTrimmed, foodName) {
				// then there isnt an issue, and you can move futher along the array
				responseArrayIndex++
			} else {
				// if there is an issue, record the string and add it to the ErrorObject array
				errorList = append(errorList, NutritionErrorObject{ErrorString: inputStringTrimmed})
			}
		}
	}

	return errorList
}

func post_foodList(c *gin.Context) {
	functionName := "post_foodList/"

	// read the request body
	bodyJson, err := readRequestBody(c.Request.Body)
	if handleError(functionName+"Error reading query request body: ", err) {
		return
	}

	// put the request body into an object
	var bodyObj PostFoodList_RequestBody
	err = json.Unmarshal(bodyJson, &bodyObj)
	if handleError(functionName+"Error reading body from query request: ", err) {
		return
	}

	// pass in the foodListString, get back the information from the api
	naturalLanguageResponseObject, err := buildNutritionixResponse_fromFoodListString(bodyObj.FoodListString)
	if handleError(functionName+"Error handling food list from request body: ", err) {
		return
	}

	ret := NaturalLanguageResponseObject{
		ListOfFoods:               naturalLanguageResponseObject.Foods,
		TotalNutritionInformation: makeTotalNutritionData_fromFoodList(naturalLanguageResponseObject.Foods),
		Errors:                    helper_checkFoodArrayForErrors(bodyObj.FoodListString, naturalLanguageResponseObject.Foods),
		// TODO i dont like this
	}

	if bodyObj.SaveToDb {
		// save this information to the Daily table
		err = saveToDatabase_NutritionInformation(bodyObj.FoodListString, bodyObj.Date, ret.TotalNutritionInformation)
		if handleError(functionName+"Error saving nutrition info to database: ", err) {
			return
		}
	}

	// create return object
	responseMarshal, err := json.Marshal(ret)
	if handleError(functionName+"Error marshalling NutritionResponseObject", err) {
		return
	}

	c.JSON(http.StatusOK, string(responseMarshal))
}

func post_saveRecipe(c *gin.Context) {
	functionName := "post_saveRecipe/"
	body, err := readRequestBody(c.Request.Body)
	if handleError(functionName+"Error reading recipe request body: ", err) {
		return
	}

	var recipeObj RecipeResponse
	err = json.Unmarshal(body, &recipeObj)
	if handleError(functionName+"Error reading body from recipe request: ", err) {
		return
	}

	nutritionInfo, err := getTotalNutritionInformation_fromFoodListString(recipeObj.FoodString)
	if handleError(functionName+"Error getting total nutrition info from food string: ", err) {
		return
	}

	err = saveToDatabase_RecipeResponse(recipeObj, nutritionInfo)
	if handleError(functionName+"Error saving recipe information to database: ", err) {
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

	router.POST("/postFoodList", post_foodList)
	router.POST("/postRecipe", post_saveRecipe)
	router.GET("/getRecipes", get_recipes)

	router.Run(":8080")
}
