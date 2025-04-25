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

func getDatabase() *sql.DB {
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

func saveToDatabase_BodyResponse(data BodyResponse) error {
	db := getDatabase()

	_, err := db.Exec("INSERT INTO daily(Date, FoodString) VALUES(?, ?)", data.Date, data.FoodListString)
	if handleError("Error inserting body values into database: ", err) {
		return err
	}

	return nil
}

// func saveToDatabase_FoodCount(data FoodItem)

func saveToDatabase_RecipeResponse(data RecipeResponse, nutritionInfo Food) error {
	db := getDatabase()
	// TODO inserting into database is working correctly
	// however, nutrition information is not corrctly being assigned (always ends up as 0 atm)
	// and there are no recipe_name, or food_string in the recipe table
	// foreign_key stuff is working correctly

	//TODO inserting into the Nutrition table is going to be cumbersome and frequent
	//some function should be made to automate that
	response, err := db.Exec("INSERT INTO nutrition_info(calories, protein, carbs, fiber) VALUES(?, ?, ?, ?)", nutritionInfo.Calories, nutritionInfo.Protein, nutritionInfo.TotalCarbohydrate, nutritionInfo.DietaryFiber)
	if handleError("Error inserting into Nutrition Table from RecipeResponse", err) {
		return err
	}

	nutritionKey, err := response.LastInsertId()
	if handleError("Error getting nutritionKey in Recipe Response", err) {
		return err
	}

	response, err = db.Exec("INSERT INTO recipe(recipe_name, food_string, serving_size, nutrition_id) VALUES(?, ?, ?, ?)", data.RecipeName, data.FoodString, 1, nutritionKey)
	if handleError("Error inserting into Recipe Table from Recipe Response", err) {
		return err
	}

	return nil
}

func getFromDatabase_Recipes() ([]Recipe, error) {
	db := getDatabase()
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

// func createInsertStatementFromNutritionInfo(nutInfo Food) string {
// }
