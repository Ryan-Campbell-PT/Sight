package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/microsoft/go-mssqldb"
)

type Daily struct {
	ID         int64
	foodString string
	date       string
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

func dailyQuery(date string) ([]Daily, error) {
	var daily []Daily

	rows, err := db.Query("SELECT * FROM daily WHERE date = ?", date)
	if err != nil {
		return nil, fmt.Errorf("dailyQuery %q: %v", date, err)
	}
	defer rows.Close()

	for rows.Next() {
		var day Daily
		if err := rows.Scan(&day.ID, &day.foodString, &day.date); err != nil {
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

func helper_getDatabase() *sql.DB {
	dbOnce.Do(func() {
		cfg := getSqlConfig()
		dbObj, err := sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatal(err)
		}

		pingErr := dbObj.Ping()
		if pingErr != nil {
			log.Fatal(pingErr)
		}
		db = dbObj
	})

	return db
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

func saveToDatabase_NutritionInformation(foodListString string, date string, nutritionInfo FoodItem) error {
	db := helper_getMsSqlDatabase()

	nutritionKey, err := helper_saveNutritionInfo(nutritionInfo)
	if handleError("saveToDatabase_BodyResponse/Error saving nutrition info to DB: ", err) {
		return err
	}

	_, err = db.Exec(`INSERT INTO daily(food_string, date, nutrition_id) VALUES(@FoodListString, @Date, @NutritionKey)`,
		sql.Named("FoodListString", foodListString),
		sql.Named("Date", date),
		sql.Named("NutritionKey", nutritionKey),
	)

	if handleError("Error inserting body values into database: ", err) {
		return err
	}

	return nil
}

func saveToDatabase_RecipeResponse(data RecipeResponse, nutritionInfo FoodItem) error {
	db := helper_getMsSqlDatabase()

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
	for _, n := range nutritionInfo.FullNutrients {
		if n.AttrID == nutritionId {
			return n.Value
		}
	}

	return 0
}

func helper_saveNutritionInfo(nutritionInfo FoodItem) (int64, error) {
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
