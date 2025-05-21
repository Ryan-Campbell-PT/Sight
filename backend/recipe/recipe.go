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

	nutritionInfo, err := nutrition.MakeTotalNutritionData(recipeObj.FoodListString)
	if util.HandleError(functionName+"Error getting total nutrition info from food string: ", err) {
		return
	}

	nutritionId, err := nutrition.SaveNutritionInfo(nutritionInfo)
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
