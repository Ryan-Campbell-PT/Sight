package recipe

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Ryan-Campbell-PT/Sight/backend/database"
	"github.com/Ryan-Campbell-PT/Sight/backend/nutrition"
	"github.com/Ryan-Campbell-PT/Sight/backend/util"
	"github.com/gin-gonic/gin"
)

// take a the food list string provided by the user
// and turn it into an object representing the important info
func parseCustomRecipe(foodString string) (*CustomRecipeParse, error) {
	trimmedFoodString := strings.ToLower(strings.TrimSpace(foodString))

	// string: 1.5 servings of moms chocolate cake
	// Match pattern: number + "servings of" OR "serving of" + the rest
	re := regexp.MustCompile(`(?i)^\s*([\d.]+)\s+servings?\s+of\s+(.+)$`)

	matches := re.FindStringSubmatch(trimmedFoodString)
	if len(matches) != 3 {
		// return 0, "", fmt.Errorf("input did not match expected format")
		return nil, gin.Error{}
	}

	servingsStr := matches[1]
	foodName := strings.TrimSpace(matches[2])

	servings, err := strconv.ParseFloat(servingsStr, 64)
	if err != nil {
		// return 0, "", fmt.Errorf("invalid serving number: %v", err)
		return nil, err
	}

	return &CustomRecipeParse{RecipeName: foodName, NumServings: servings, FoodString: trimmedFoodString}, nil
}

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
