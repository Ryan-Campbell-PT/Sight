package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Ryan-Campbell-PT/Sight/backend/database"
	"github.com/Ryan-Campbell-PT/Sight/backend/nutrition"
	"github.com/Ryan-Campbell-PT/Sight/backend/recipe"
	"github.com/Ryan-Campbell-PT/Sight/backend/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ReadRequestBody(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	data, err := io.ReadAll(body)
	return data, err
}

func SendHttpRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if util.HandleError("SendHttpRequest/Error sending http request: ", err) {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

/*

func getTotalNutritionInformation_fromFoodListString(foodListString string) (FoodItem, error) {
	functionName := "getTotalNutritionInformation_fromFoodListString/"
	nutritionInfo, err := buildNutritionixResponse_fromFoodListString(foodListString)
	if handleError(functionName+"Error getting response from foodliststring: ", err) {
		return FoodItem{}, err
	}

	return makeTotalNutritionData_fromFoodList(nutritionInfo.Foods), nil
}

*/

// HTTP POST
// given the body of a NutritionRequest

func getNutritionResponse(c *gin.Context) {
	functionName := "getNutritionResponse/"

	// read the request body
	bodyJson, err := ReadRequestBody(c.Request.Body)
	if util.HandleError(functionName+"Error reading query request body: ", err) {
		return
	}

	// put the request body into an object
	var bodyObj nutrition.GetNutritionRequestBody
	err = json.Unmarshal(bodyJson, &bodyObj)
	if util.HandleError(functionName+"Error reading body from query request: ", err) {
		return
	}

	// pass in the foodListString, get back the information from the api
	// TODO this i feel should be handled by the server? cause its a fetch?
	naturalLanguageResponseObject, err := nutrition.FetchNaturalLanguageResponse(bodyObj.FoodListString)
	if util.HandleError(functionName+"Error handling food list from request body: ", err) {
		return
	}

	ret := nutrition.NutritionInfoResponse{
		Foods:                     naturalLanguageResponseObject.Foods,
		TotalNutritionInformation: nutrition.MakeTotalNutritionData(naturalLanguageResponseObject.Foods),
		Errors:                    nutrition.CheckFoodArrayForErrors(bodyObj.FoodListString, naturalLanguageResponseObject.Foods),
	}

	//i dont think this need to exist, it can be reworked
	/* if bodyObj.SaveToDb {
		// save this information to the Daily table
		err = saveToDatabase_NutritionInformation(bodyObj.FoodListString, bodyObj.Date, ret.TotalNutritionInformation)
		if handleError(functionName+"Error saving nutrition info to database: ", err) {
			return
		}
	}
	*/

	// create return object
	responseMarshal, err := json.Marshal(ret)
	if util.HandleError(functionName+"Error marshalling NutritionResponseObject", err) {
		return
	}

	c.JSON(http.StatusOK, string(responseMarshal))
}

func saveRecipe(c *gin.Context) {
	functionName := "saveRecipe/"
	body, err := ReadRequestBody(c.Request.Body)
	if util.HandleError(functionName+"Error reading recipe request body: ", err) {
		return
	}

	var recipeObj recipe.SaveRecipeRequestBody
	err = json.Unmarshal(body, &recipeObj)
	if util.HandleError(functionName+"Error reading body from recipe request: ", err) {
		return
	}

	nutritionInfo, err := nutrition.GetTotalNutritionInformation(recipeObj.FoodListString)
	if util.HandleError(functionName+"Error getting total nutrition info from food string: ", err) {
		return
	}

	nutritionId, err := database.SaveNutritionInfo(nutritionInfo)
	if util.HandleError(functionName+"Error saving nutrition information to database: ", err) {
		return
	}

	var recipe database.Recipe
	recipe.Active = true
	recipe.FoodListString = recipeObj.FoodListString
	recipe.Name = recipeObj.RecipeName
	recipe.ServingSize = recipeObj.NumServings
	err = database.SaveRecipe(recipe, nutritionId)
	if util.HandleError(functionName+"Error saving recipe: ", err) {
		return
	}

	c.JSON(http.StatusOK, "")
}

func get_recipes(c *gin.Context) {
	recipes, err := database.GetAllRecipes()
	if util.HandleError("Server.go/Error getting recipes from database: ", err) {
		c.JSON(http.StatusBadRequest, "")
	}

	recipeJson, err := json.Marshal(recipes)
	if util.HandleError("Error marshaling recipeResponse into recipeJson: ", err) {
		return
	}

	c.JSON(http.StatusOK, string(recipeJson))
}

func RunServer() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	router.POST("/postFoodList", post_foodList)
	router.POST("/postSaveRecipe", post_saveRecipe)
	// router.GET("/getActiveRecipes", get_activeRecipes)
	// router.GET("/getInactiveRecipes", get_inactiveRecipes)
	// router.GET("/getAllRecipes", get_allRecipes)

	router.Run(":8080")
}
