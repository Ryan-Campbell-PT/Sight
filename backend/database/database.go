package database

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/Ryan-Campbell-PT/Sight/backend/util"

	_ "github.com/microsoft/go-mssqldb"
)

// this dbOnce variable makes it so no matter how many times you call the function getDatabase()
// the code inside will only run once
var (
	db     *sql.DB
	dbOnce sync.Once
)

func getMsSqlConnectionString() string {
	cfg := util.GetEnvConfig()

	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		cfg.Azure_Server, cfg.Azure_User, cfg.Azure_Password, cfg.Azure_Port, cfg.Azure_Database)
}

func GetDatabase() *sql.DB {
	return helper_getMsSqlDatabase()
}

func helper_getMsSqlDatabase() *sql.DB {
	functionName := "helper_getMsSqlDatabase/"
	dbOnce.Do(func() {
		dbObj, err := sql.Open("sqlserver", getMsSqlConnectionString())
		if util.HandleError(functionName+"Error connecting to MsSql db: ", err) {
			return
		}

		db = dbObj
	})

	return db
}

// get records from the Daily database
// startDate and endDate are INCLUSIVE
// if you only need one date, start and endDate should be the same value
func GetDailyRecord(startDate time.Time, endDate time.Time) ([]Daily, error) {
	functionName := "GetDailyRecord/"
	var records []Daily

	rows, err := db.Query("SELECT * FROM daily WHERE date BETWEEN @StartDate AND @EndDate",
		sql.Named("StartDate", util.GetDate(startDate)),
		sql.Named("EndDate", util.GetDate(endDate)),
	)

	if util.HandleError(functionName+"Error getting daily records from dates: ", err) {
		return records, err
	}
	defer rows.Close()

	for rows.Next() {
		var day Daily
		err := rows.Scan(&day.ID, &day.FoodString, &day.Date)
		if util.HandleError(functionName+"Error scanning db row into local variable: ", err) {
			return records, err
		}
		records = append(records, day)
	}

	err = rows.Err()
	if util.HandleError(functionName+"Error with rows returned from database: ", err) {
		return records, err
	}

	return records, nil
}

// TODO this is dumb cause SaveRecipe takes the nutritionId but SaveDailyRecord takes a foodItem
// and creates the nutritionId
// that needs to be standardized

// TODO SaveNutritionInfo could be un-exported and strictly used as
// a local function, so you always have to pass in the FoodItem
// and never the nutritionId
func SaveRecipe(recipe Recipe, nutritionId int64) error {
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

/*
	func SaveRecipeResponse(recipe Recipe, nutritionInfo nutrition.FoodItem) error {
		functionName := "SaveRecipeResponse/"
		db := GetDatabase()

		//TODO inserting into the Nutrition table is going to be cumbersome and frequent
		//some function should be made to automate that
		nutritionKey, err := saveNutritionInfo(nutritionInfo)
		if util.HandleError(functionName+"Error saving nutrition info to DB: ", err) {
			return err
		}

		_, err = db.Exec(`INSERT INTO recipe(nutrition_id, recipe_name, food_string, serving_size)
			VALUES(@NutritionKey, @RecipeName, @FoodString, @ServingSize)`,
			sql.Named("NutritionKey", nutritionKey),
			sql.Named("Calories", nutritionInfo.Calories),
			sql.Named("RecipeName", data.RecipeName),
			sql.Named("FoodString", data.FoodString),
		)

		if util.HandleError("Error inserting into Recipe Table from Recipe Response", err) {
			return err
		}

		return nil
	}
*/
func GetAllRecipes() ([]Recipe, error) {
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

func GetRecipes(active bool) ([]Recipe, error) {
	functionName := "getFromDatabase_Recipes/"
	db := GetDatabase()

	response, err := db.Query("SELECT * FROM recipe WHERE active = @Active", sql.Named("Active", active))

	if util.HandleError("Database.go/Error grabbing recipes: ", err) {
		return nil, err
	}
	defer response.Close()

	recipeList, err := createRecipeList(response)
	if util.HandleError(functionName+"Error getting recipe list from db query: ", err) {
		return nil, err
	}

	return recipeList, nil
}

func createRecipeList(dbQuery *sql.Rows) ([]Recipe, error) {
	var recipeList []Recipe
	for dbQuery.Next() {
		var recipe Recipe
		if err := dbQuery.Scan(&recipe.Id, &recipe.Name, &recipe.FoodListString, &recipe.ServingSize, &recipe.NutritionInfoId); err != nil {
			return nil, err
		}
		recipeList = append(recipeList, recipe)
	}

	return recipeList, nil
}
