package logic

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Ryan-Campbell-PT/Sight/backend/models"
	"github.com/Ryan-Campbell-PT/Sight/backend/util"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var (
	c                    *cache.Cache
	userRecipeListString = "UsersRecipeList"
)

// Interface that matches both *sql.Row and *sql.Rows
type scanner interface {
	Scan(dest ...any) error
}

func initRecipeCache() {
	c = cache.New(5*time.Minute, 10*time.Minute)
	cacheAllRecipes()
}

func cacheAllRecipes() error {
	functionName := "cacheAllRecipes/"
	db := GetDatabase()
	// TODO will need to add a clause for UserId when thats completed
	rows, err := db.Query(`SELECT * FROM custom_recipe`)
	if util.HandleError(functionName+"Error getting all recipes: ", err) {
		return err
	}
	defer rows.Close()

	recipeList, err := scanRecipeList(rows)
	if util.HandleError(functionName+"Error scanning recipe rows: ", err) {
		return err
	}

	// TODO dont think this works
	c.Set(userRecipeListString, recipeList, cache.DefaultExpiration)
	return nil
}

func scanRecipe(s scanner) (*models.CustomRecipe, error) {
	functionName := "scanRecipeItem/"
	var recipeObj models.CustomRecipe

	err := s.Scan(
		&recipeObj.Id,
		&recipeObj.Name,
		&recipeObj.AlternativeRecipeNames,
		&recipeObj.FoodListString,
		&recipeObj.ServingSize,
		&recipeObj.Active,
		&recipeObj.NutritionInfoId,
		// &recipeObj.LastModified,
	)

	if util.HandleError(functionName+"Error scanning CustomRecipe: ", err) {
		return nil, err
	}

	return &recipeObj, nil
}

func scanRecipeList(rows *sql.Rows) ([]models.CustomRecipe, error) {
	functionName := "scanRecipeRows/"
	var recipeList []models.CustomRecipe

	for rows.Next() {
		recipeObj, err := scanRecipe(rows)
		if util.HandleError(functionName+"Error scanning recipe from scanner: ", err) {
			return nil, err
		}
		recipeList = append(recipeList, *recipeObj)
	}

	return recipeList, nil
}

// TODO SaveNutritionInfo could be un-exported and strictly used as
// a local function, so you always have to pass in the FoodItem
// and never the nutritionId
func SaveRecipe(recipe models.CustomRecipe, nutritionId int64) error {
	functionName := "saveToDatabase_Recipe/"
	db := GetDatabase()

	_, err := db.Exec(`INSERT INTO recipe(recipe_name, food_string, serving_size, active, nutrition_id)
			VALUES (@RecipeName, @FoodString, @ServingSize, @Active, @NutritionId)`,
		sql.Named("RecipeName", recipe.Name),
		sql.Named("FoodString", recipe.FoodListString),
		sql.Named("ServingSize", recipe.ServingSize),
		sql.Named("Active", true),
		sql.Named("NutritionId", nutritionId),
	)

	if util.HandleError(functionName+"Error saving recipe information to db: ", err) {
		return err
	}

	return nil
}

func GetAllRecipes() ([]models.CustomRecipe, error) {
	functionName := "getFromDatabase_AllRecipes/"

	inactiveResponse, err := GetRecipes(false)
	if util.HandleError(functionName+"Error getting inactive recipes: ", err) {
		return nil, err
	}
	activeResponse, err := GetRecipes(true)
	if util.HandleError(functionName+"Error getting active recipes: ", err) {
		return nil, err
	}

	return append(activeResponse, inactiveResponse...), nil
}

func GetRecipes(active bool) ([]models.CustomRecipe, error) {
	functionName := "getFromDatabase_Recipes/"
	db := GetDatabase()

	response, err := db.Query("SELECT * FROM recipe WHERE active = @Active", sql.Named("Active", active))

	if util.HandleError("Database.go/Error grabbing recipes: ", err) {
		return nil, err
	}
	defer response.Close()

	recipeList, err := scanRecipeList(response)
	if util.HandleError(functionName+"Error getting recipe list from db query: ", err) {
		return nil, err
	}

	return recipeList, nil
}

/*
func createRecipeList(dbQuery *sql.Rows) ([]models.CustomRecipe, error) {
	functionName := "createRecipeList/"
	var recipeList []models.CustomRecipe
	for dbQuery.Next() {
		// var recipe models.CustomRecipe
		recipe, err := scanRecipe(dbQuery)
		if util.HandleError(functionName, err) {
			return nil, err
		}
		recipeList = append(recipeList, *recipe)
	}

	return recipeList, nil
}
*/

// this function is no longer needed, but is a good reference for caching
/*
func getUsersRecipes(isActive bool) []models.CustomRecipe {
	initRecipeCache()
	var retRecipeList []models.CustomRecipe
	if cacheRecipeList, found := c.Get(userRecipeListString); found {
		rList := cacheRecipeList.([]models.CustomRecipe)
		for _, recipe := range rList {
			if recipe.Active == isActive {
				retRecipeList = append(retRecipeList, recipe)
			}
		}

		return retRecipeList
	}

	return nil
}
*/

func GetUsersInactiveRecipes() []models.CustomRecipe {
	functionName := "GetUsersActiveRecipes/"
	ret, err := GetRecipes(false)
	if util.HandleError(functionName, err) {
		return nil
	}
	return ret
}

func GetUsersActiveRecipes() []models.CustomRecipe {
	functionName := "GetUsersActiveRecipes/"
	ret, err := GetRecipes(true)
	if util.HandleError(functionName, err) {
		return nil
	}
	return ret
}

// returns the Id of the CustomRecipe if successful
func IsActiveRecipeItem(foodString string) int {
	// functionName := "IsActiveRecipeItem/"
	return 0
}

// returns the Id of the CustomRecipe if successful
func IsRecipeItem(foodString string) int64 {
	functionName := "IsRecipeItem/"
	db := GetDatabase()
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

	recipeItem, err := scanRecipe(sqlRow)
	if util.HandleError(functionName+"Error scanning recipe item:", err) {
		return -1
	}

	return recipeItem.Id
}

// while there is a recipe.Recipe object type
// database schema objects are not to be manipulated, only used to make different objects
// so in this case, a CustomFoodItem is being returned, with data filled from Recipe
func GetRecipeItem(foodString string) *models.CustomRecipe {
	functionName := "GetRecipeItem/"
	db := GetDatabase()

	recipeId := IsRecipeItem(foodString)
	if recipeId <= 0 {
		return nil
	}

	sqlRow := db.QueryRow(`SELECT * FROM recipe WHERE id=@RecipeId`, sql.Named("RecipeId", recipeId))

	recipeItem, err := scanRecipe(sqlRow)
	if util.HandleError(functionName+"Error scanning recipe item: ", err) {
		return nil
	}

	return recipeItem
}

/*
// this function will take the entire foodListString
// and return any recipes in the string
func CheckForRecipes(foodListString string) ([]models.CustomRecipe, error) {
	splitFoodString := strings.Split(strings.ToLower(foodListString), ",")

	for i, str := range splitFoodString {
		if IsRecipeItem(str) > 0 {


		}
	}
}

*/
// take a the individual food string provided by the user
// and turn it into an object representing the important info
func parseCustomRecipe(foodString string) (*models.CustomRecipeParse, error) {
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

	return &models.CustomRecipeParse{RecipeName: foodName, NumServings: servings, FoodString: trimmedFoodString}, nil
}

func GetUserRecipesJson(c *gin.Context) {
	activeRecipes := GetUsersActiveRecipes()
	inactiveRecipes := GetUsersInactiveRecipes()

	// if util.HandleError("Server.go/Error getting recipes from database: ", err) {
	// 	c.JSON(http.StatusBadRequest, "")
	// }

	ret := models.GetUserRecipesResponseObject{RecipeList: append(activeRecipes, inactiveRecipes...)}
	recipeJson, err := json.Marshal(ret)
	if util.HandleError("Error marshaling recipeResponse into recipeJson: ", err) {
		return
	}

	c.JSON(http.StatusOK, string(recipeJson))
}

// check the users input for any recipes
// and remove from the original userInputArray and recipes caught
// necessary cause some recipes may match names the api reads like "breakfast wrap"
// may be a recipe, but also something nutritionix understands
func CheckUserInputForRecipes(userInputArray []models.ParsedUserInput) []models.CustomRecipe {
	functionName := "CheckUserInputForRecipes/"
	var ret []models.CustomRecipe

	// ignore error cause we only want the object, an error will just not be added into array
	for i, input := range userInputArray {
		parsedRecipe := parseRecipeItem(input.FoodString)
		if parsedRecipe == nil {
			continue
		}

		recipeItem, err := GetRecipeItemFromString(parsedRecipe.RecipeName)
		if recipeItem == nil || util.HandleError(functionName, err) {
			continue
		}
		// arrays in golang are essentially pointers to all the content
		// so modifing the array is easy
		userInputArray[i].IsRecipe = true
		ret = append(ret, *recipeItem)
	}
	return ret
}

// space
// space
// space
// space
// space
// space
// space

// alt recipeNames can be empty (golang doesnt allow default values)
func GetRecipeItemFromString(recipeName string) (*models.CustomRecipe, error) {
	functionName := "GetRecipeItemFromString/"
	// TODO this will have to implement caching at some point
	db := GetDatabase()

	sqlRow := db.QueryRow(`
		SELECT * FROM custom_recipe
		WHERE (recipe_name LIKE ? OR alt_recipe_names LIKE ?)
		AND active = 1
	`, "%"+recipeName+"%", "%"+recipeName+"%")

	recipe, err := scanRecipe(sqlRow)
	if util.HandleError(functionName, err) {
		return nil, err
	}

	return recipe, nil
}

// TODO
func GetRecipeItemFromId(recipeId int64) (*models.CustomRecipe, error) {
	return nil, nil
}

// take a the individual food string provided by the user, as a ParsedUserInput array
// and turn it into an object representing the important info
func parseRecipeItem(foodString string) *models.CustomRecipeParse {
	// string: 1.5 servings of moms chocolate cake
	// Match pattern: number + "servings of" OR "serving of" + the rest
	re := regexp.MustCompile(`(?i)^\s*([\d.]+)\s+servings?\s+of\s+(.+)$`)

	trimmedString := strings.TrimSpace(foodString)
	matches := re.FindStringSubmatch(trimmedString)
	if len(matches) != 3 {
		// return 0, "", fmt.Errorf("input did not match expected format")
		return nil
	}

	servingsStr := matches[1]
	foodName := strings.TrimSpace(matches[2])
	servings, err := strconv.ParseFloat(servingsStr, 64)
	if err != nil {
		// return 0, "", fmt.Errorf("invalid serving number: %v", err)
		return nil
	}

	return &models.CustomRecipeParse{RecipeName: foodName, NumServings: servings, FoodString: trimmedString}
}

func removeRecipesFromUserInput(parsedUserInput []models.ParsedUserInput) string {
	ret := ""
	for _, input := range parsedUserInput {
		if input.IsRecipe {
			continue
		}
		ret += input.FoodString + ", "
	}

	return strings.Trim(ret, ", ")
}
