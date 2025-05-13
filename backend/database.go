package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/microsoft/go-mssqldb"
)

type Daily struct {
	ID         int64
	FoodString string
	Date       string
}

type Recipe struct {
	Id              int64  `json:"id"`
	Name            string `json:"recipe_name"`
	FoodListString  string `json:"food_string"`
	ServingSize     int64  `json:"serving_size"`
	NutritionInfoId int64  `json:"nutrition_id"`
}

// this dbOnce variable makes it so no matter how many times you call the function getDatabase()
// the code inside will only run once
var (
	db     *sql.DB
	dbOnce sync.Once
)

func getMsSqlConnectionString() string {
	cfg := getConfig()

	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		cfg.Azure_Server, cfg.Azure_User, cfg.Azure_Password, cfg.Azure_Port, cfg.Azure_Database)
}

func getDatabase() *sql.DB {
	return helper_getMsSqlDatabase()
}

func helper_getMsSqlDatabase() *sql.DB {
	functionName := "helper_getMsSqlDatabase/"
	dbOnce.Do(func() {
		dbObj, err := sql.Open("sqlserver", getMsSqlConnectionString())
		if handleError(functionName+"Error connecting to MsSql db: ", err) {
			return
		}

		db = dbObj
	})

	return db
}

func dailyQuery(date string) ([]Daily, error) {
	var daily []Daily

	rows, err := db.Query("SELECT * FROM daily WHERE date = ?", date)
	if err != nil {
		return nil, fmt.Errorf("dailyQuery %q: %v", date, err)
	}
	defer rows.Close()

	for rows.Next() {
		var day Daily
		if err := rows.Scan(&day.ID, &day.FoodString, &day.Date); err != nil {
			return nil, fmt.Errorf("dailyQuery %q: %v", date, err)
		}
		daily = append(daily, day)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("dailyQuery %q: %v", date, err)
	}
	return daily, nil
}

/*
func visualizationTest_queryForDailyCalories() ([]Daily, error) {
	duh()
	var daily []Daily;
	rows, err := db.Query("select * from daily")
	if err != nil {
		return nil, fmt.Errorf("dailyQuery %q: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var day Daily
		if err := rows.Scan(&day.ID, &day.foodString, &day.date, &day.calories); err != nil {
			return nil, fmt.Errorf("dailyQuery %q: %v", err)
		}
		daily = append(daily, day)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("dailyQuery %q: %v", err)
	}
	return daily, nil
}
*/

func saveToDatabase_DailyRecord(foodListString string, date string, nutritionInfo FoodItem) error {
	db := getDatabase()

	nutritionId, err := helper_saveNutritionInfo(nutritionInfo)
	if handleError("saveToDatabase_BodyResponse/Error saving nutrition info to DB: ", err) {
		return err
	}

	_, err = db.Exec(`INSERT INTO daily(food_string, date, nutrition_id) VALUES(@FoodListString, @Date, @NutritionKey)`,
		sql.Named("FoodListString", foodListString),
		sql.Named("Date", date),
		sql.Named("NutritionKey", nutritionId),
	)

	if handleError("Error inserting body values into database: ", err) {
		return err
	}

	return nil
}

// TODO it makes more sense to have the whole class as the parameter,
// so you dont have to modify the function parameters if anything changes in the class
// but I dont like having that ugly class name as a function parameter,
// especially since you just recently hange all the functions to NOT include the Request/Response objs
func saveToDatabase_Recipe(data PostRecipe_RequestBody, nutritionId int64) error {
	functionName := "saveToDatabase_Recipe/"
	db := getDatabase()

	_, err := db.Exec(`INSERT INTO recipe(recipe_name, food_string, serving_size, active, nutrition_id)
			VALUES (@RecipeName, @FoodString, @ServingSize, @Active, @NutritionId)`,
		sql.Named("RecipeName", data.RecipeName),
		sql.Named("FoodString", data.FoodListString),
		sql.Named("ServingSize", data.NumServings),
		sql.Named("Active", true),
		sql.Named("NutritionId", nutritionId),
	)

	if handleError(functionName+"Error saving recipe information to db: ", err) {
		return err
	}

	return nil
}

func saveToDatabase_NutritionInformation(nutritionInfo FoodItem) (int64, error) {
	nutritionKey, err := helper_saveNutritionInfo(nutritionInfo)
	if handleError("saveToDatabase_BodyResponse/Error saving nutrition info to DB: ", err) {
		return -1, err
	}
	return nutritionKey, nil
}

func saveToDatabase_RecipeResponse(data RecipeResponse, nutritionInfo FoodItem) error {
	db := getDatabase()

	//TODO inserting into the Nutrition table is going to be cumbersome and frequent
	//some function should be made to automate that
	nutritionKey, err := helper_saveNutritionInfo(nutritionInfo)
	if handleError("saveToDatabase_RecipeResponse/Error saving nutrition info to DB: ", err) {
		return err
	}

	_, err = db.Exec(`INSERT INTO recipe(nutrition_id, recipe_name, food_string, serving_size)
		VALUES(@NutritionKey, @RecipeName, @FoodString, @ServingSize)`,
		sql.Named("NutritionKey", nutritionKey),
		sql.Named("Calories", nutritionInfo.Calories),
		sql.Named("RecipeName", data.RecipeName),
		sql.Named("FoodString", data.FoodString),
	)

	if handleError("Error inserting into Recipe Table from Recipe Response", err) {
		return err
	}

	return nil
}

func getFromDatabase_Recipes() ([]Recipe, error) {
	db := helper_getMsSqlDatabase()
	var recipeList []Recipe

	response, err := db.Query("SELECT * FROM recipe")
	if handleError("Database.go/Error grabbing recipes: ", err) {
		return nil, err
	}
	defer response.Close()

	for response.Next() {
		var recipe Recipe
		if err := response.Scan(&recipe.Id, &recipe.Name, &recipe.FoodListString, &recipe.ServingSize, &recipe.NutritionInfoId); err != nil {
			return nil, err
		}
		recipeList = append(recipeList, recipe)
	}

	return recipeList, nil
}

func helper_getNutrient(nutritionInfo FoodItem, nutritionId int64) float64 {
	return nutritionInfo.FullNutrientMap[nutritionId]
}

func helper_saveNutritionInfo(nutritionInfo FoodItem) (int64, error) {
	db := getDatabase()

	row := db.QueryRow(`
		INSERT INTO nutrition_info (
			calories, protein, carbs, fiber, cholesterol, sugar,
			phosphorus, sodium, total_fat, saturated_fat, poly_fat, mono_fat, potassium
		) VALUES (
			@Calories, @Protein, @Carbs, @Fiber, @Cholesterol, @Sugar,
			@Phosphorus, @Sodium, @TotalFat, @SaturatedFat, @PolyFat, @MonoFat, @Potassium
		);
		SELECT ID = CONVERT(BIGINT, SCOPE_IDENTITY());
	`,
		sql.Named("Calories", nutritionInfo.Calories),
		sql.Named("Protein", nutritionInfo.Protein),
		sql.Named("Carbs", nutritionInfo.TotalCarbohydrate),
		sql.Named("Fiber", nutritionInfo.DietaryFiber),
		sql.Named("Cholesterol", nutritionInfo.Cholesterol),
		sql.Named("Sugar", nutritionInfo.Sugars),
		sql.Named("Phosphorus", nutritionInfo.Phosphorus),
		sql.Named("Sodium", nutritionInfo.Sodium),
		sql.Named("TotalFat", nutritionInfo.TotalFat),
		sql.Named("SaturatedFat", nutritionInfo.SaturatedFat),
		sql.Named("PolyFat", helper_getNutrient(nutritionInfo, NutrientPolyunsaturatedFat)),
		sql.Named("MonoFat", helper_getNutrient(nutritionInfo, NutrientMonounsaturatedFat)),
		sql.Named("Potassium", helper_getNutrient(nutritionInfo, NutrientPotassium)),
	)
	var nutritionKey int64
	err := row.Scan(&nutritionKey)
	if handleError("Error getting nutritionKey from Recipe Response", err) {
		return -1, err
	}

	return nutritionKey, nil
}
