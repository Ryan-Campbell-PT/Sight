package recipe

// this file will handle any work done on the database in regards to CustomRecipes

import (
	"database/sql"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/Ryan-Campbell-PT/Sight/backend/database"
	"github.com/Ryan-Campbell-PT/Sight/backend/util"
)

var (
	c                    *cache.Cache
	userRecipeListString = "UsersRecipeList"
)

// Interface that matches both *sql.Row and *sql.Rows
type scanner interface {
	Scan(dest ...any) error
}

func initCache() {
	c = cache.New(5*time.Minute, 10*time.Minute)
	cacheAllRecipes()
}

func scanRecipeItem(row *sql.Row) (*database.CustomRecipe, error) {
	return scanRecipeFromScanner(row)
}

func scanRecipeFromScanner(s scanner) (*database.CustomRecipe, error) {
	functionName := "scanRecipeItem/"
	var recipeObj database.CustomRecipe

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

func scanRecipeRows(rows *sql.Rows) ([]database.CustomRecipe, error) {
	functionName := "scanRecipeRows/"
	var recipeList []database.CustomRecipe

	for rows.Next() {
		recipeObj, err := scanRecipeFromScanner(rows)
		if util.HandleError(functionName+"Error scanning recipe from scanner: ", err) {
			return nil, err
		}
		recipeList = append(recipeList, *recipeObj)
	}

	return recipeList, nil
}

func getUsersRecipes(isActive bool) []database.CustomRecipe {
	initCache()
	var retRecipeList []database.CustomRecipe
	if cacheRecipeList, found := c.Get(userRecipeListString); found {
		rList := cacheRecipeList.([]database.CustomRecipe)
		for _, recipe := range rList {
			if recipe.Active == isActive {
				retRecipeList = append(retRecipeList, recipe)
			}
		}

		return retRecipeList
	}

	return nil
}

func GetUsersInactiveRecipes() []database.CustomRecipe {
	return getUsersRecipes(false)
}

func GetUsersActiveRecipes() []database.CustomRecipe {
	return getUsersRecipes(true)
}

// returns the Id of the CustomRecipe if successful
func IsActiveRecipeItem(foodString string) int {
	// functionName := "IsActiveRecipeItem/"
	return 0
}

func cacheAllRecipes() error {
	functionName := "cacheAllRecipes/"
	db := database.GetDatabase()
	// TODO will need to add a clause for UserId when thats completed
	rows, err := db.Query(`SELECT * FROM custom_recipe`)
	if util.HandleError(functionName+"Error getting all recipes: ", err) {
		return err
	}
	defer rows.Close()

	recipeList, err := scanRecipeRows(rows)
	if util.HandleError(functionName+"Error scanning recipe rows: ", err) {
		return err
	}

	// TODO dont think this works
	c.Set(userRecipeListString, recipeList, cache.DefaultExpiration)
	return nil
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
