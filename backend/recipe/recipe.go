package recipe

import (
	"database/sql"
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

func scanRecipeItem(row *sql.Row) (*database.CustomRecipe, error) {
	functionName := "scanRecipeItem/"
	var recipeObj database.CustomRecipe

	err := row.Scan(
		&recipeObj.Id,
		&recipeObj.Name,
		&recipeObj.AlternativeRecipeNames,
		&recipeObj.ServingSize,
		&recipeObj.Active,
		&recipeObj.NutritionInfoId,
		&recipeObj.LastModified,
	)

	if util.HandleError(functionName+"Error scanning CustomRecipe: ", err) {
		return nil, err
	}

	return &recipeObj, nil
}

// returns the Id of the CustomRecipe if successful
func IsRecipeItem(foodString string) int64 {
	functionName := "IsRecipeItem/"
	db := database.GetDatabase()
	parse, err := parseCustomRecipe(foodString)
	if util.HandleError(functionName+"Error parsing custom recipe: ", err) {
		return -1
	}

	// TODO CONTAINS needs to be checked
	sqlRow := db.QueryRow(`
		SELECT *
		FROM recipe
		WHERE food_string
		LIKE '@FoodString' OR alt_recipe_names CONTAINS '@FoodString'
	`,
		sql.Named("FoodString", parse.FoodString))

	if util.HandleError(functionName+"Error querying recipe from food_string: "+foodString, sqlRow.Err()) {
		return -1
	}

	recipeItem, err := scanRecipeItem(sqlRow)
	if util.HandleError(functionName+"Error scanning recipe item:", err) {
		return -1
	}

	return recipeItem.Id
}

// while there is a recipe.Recipe object type
// database schema objects are not to be manipulated, only used to make different objects
// so in this case, a CustomFoodItem is being returned, with data filled from Recipe
func GetRecipeItem(foodString string) *database.CustomRecipe {
	functionName := "GetRecipeItem/"
	db := database.GetDatabase()

	recipeId := IsRecipeItem(foodString)
	if recipeId == -1 {
		return nil
	}

	sqlRow := db.QueryRow(`SELECT * FROM recipe WHERE id=@RecipeId`, sql.Named("RecipeId", recipeId))

	recipeItem, err := scanRecipeItem(sqlRow)
	if util.HandleError(functionName+"Error scanning recipe item: ", err) {
		return nil
	}

	return recipeItem
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
	nutritionInfo := nutrition.GetNutritionInfoResponse(recipeObj.FoodListString)
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
