package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
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

func IRRELEVANT() {
	// Capture connection properties.
	// Get a database handle
	cfg := getSqlConfig()
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	dailyValues, err := dailyQuery("1/1/2025")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("daily Values: %v\n", dailyValues)
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

func saveToDatabase_BodyResponse(data BodyResponse, nutritionInfo Food) error {
	db := helper_getDatabase()

	nutritionKey, err := helper_saveNutritionInfo(nutritionInfo)
	if handleError("saveToDatabase_BodyResponse/Error saving nutrition info to DB: ", err) {
		return err
	}

	_, err = db.Exec("INSERT INTO daily(date, food_string, nutrition_id) VALUES(?, ?, ?)", data.Date, data.FoodListString, nutritionKey)
	if handleError("Error inserting body values into database: ", err) {
		return err
	}

	return nil
}

func saveToDatabase_RecipeResponse(data RecipeResponse, nutritionInfo Food) error {
	db := helper_getDatabase()

	//TODO inserting into the Nutrition table is going to be cumbersome and frequent
	//some function should be made to automate that
	nutritionKey, err := helper_saveNutritionInfo(nutritionInfo)
	if handleError("saveToDatabase_RecipeResponse/Error saving nutrition info to DB: ", err) {
		return err
	}

	_, err = db.Exec("INSERT INTO recipe(recipe_name, food_string, serving_size, nutrition_id) VALUES(?, ?, ?, ?)", data.RecipeName, data.FoodString, 1, nutritionKey)
	if handleError("Error inserting into Recipe Table from Recipe Response", err) {
		return err
	}

	return nil
}

func getFromDatabase_Recipes() ([]Recipe, error) {
	db := helper_getDatabase()
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

func helper_getNutrient(nutritionInfo Food, nutritionId int64) float64 {
	for _, n := range nutritionInfo.FullNutrients {
		if n.AttrID == nutritionId {
			return n.Value
		}
	}

	return 0
}

func helper_saveNutritionInfo(nutritionInfo Food) (int64, error) {
	response, err := db.Exec("INSERT INTO nutrition_info(calories, protein, carbs, fiber, cholesterol, sugar, phosphorus, sodium, total_fat, saturated_fat, poly_fat, mono_fat, potassium) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		nutritionInfo.Calories,
		nutritionInfo.Protein,
		nutritionInfo.TotalCarbohydrate,
		nutritionInfo.DietaryFiber,
		nutritionInfo.Cholesterol,
		nutritionInfo.Sugars,
		nutritionInfo.Phosphorus,
		nutritionInfo.Sodium,
		nutritionInfo.TotalFat,
		nutritionInfo.SaturatedFat,
		helper_getNutrient(nutritionInfo, NutrientPolyunsaturatedFat),
		helper_getNutrient(nutritionInfo, NutrientMonounsaturatedFat),
		helper_getNutrient(nutritionInfo, NutrientPotassium),
	)

	if handleError("Error inserting into Nutrition Table from RecipeResponse", err) {
		return -1, err
	}

	nutritionKey, err := response.LastInsertId()
	if handleError("Error getting nutritionKey in Recipe Response", err) {
		return -1, err
	}

	return nutritionKey, nil
}
