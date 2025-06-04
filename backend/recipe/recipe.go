package recipe

import (
	"encoding/json"
	"net/http"

	"github.com/Ryan-Campbell-PT/Sight/backend/database"
	"github.com/Ryan-Campbell-PT/Sight/backend/nutrition"
	"github.com/Ryan-Campbell-PT/Sight/backend/util"
	"github.com/gin-gonic/gin"
)

func saveRecipe(c *gin.Context) {
	functionName := "saveRecipe/"
	body, err := util.ReadRequestBody(c.Request.Body)
	if util.HandleError(functionName+"Error reading recipe request body: ", err) {
		return
	}

	var recipeObj SaveRecipeRequestBody
	err = json.Unmarshal(body, &recipeObj)
	if util.HandleError(functionName+"Error reading body from recipe request: ", err) {
		return
	}

	// get the nutrition information from the food string
	nutritionInfo := nutrition.GetNaturalLanguageResponse(recipeObj.FoodListString)
	// nutritionInfo, err := nutrition.GetNutritionInfoResponse(recipeObj.FoodListString)
	// if util.HandleError(functionName+"Error getting total nutrition info from food string: ", err) {
	// 	return
	// }

	nutritionId, err := nutrition.SaveNutritionInfo(nutritionInfo.TotalNutritionInformation)
	if util.HandleError(functionName+"Error saving nutrition information to database: ", err) {
		return
	}

	var recipe database.CustomRecipe
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

func GetUserRecipesJson(c *gin.Context) {
	activeRecipes := GetUsersActiveRecipes()
	inactiveRecipes := GetUsersInactiveRecipes()

	// if util.HandleError("Server.go/Error getting recipes from database: ", err) {
	// 	c.JSON(http.StatusBadRequest, "")
	// }

	ret := GetUserRecipesResponseObject{RecipeList: append(activeRecipes, inactiveRecipes...)}
	recipeJson, err := json.Marshal(ret)
	if util.HandleError("Error marshaling recipeResponse into recipeJson: ", err) {
		return
	}

	c.JSON(http.StatusOK, string(recipeJson))
}
