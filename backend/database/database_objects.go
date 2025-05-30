package database

import "time"

// this file is direct schemas from the database
// these objects will be used but never manipulated
// if manipulated, theyd be manipulated into other custom objects

type Daily struct {
	ID              int64     `json:"id"`
	FoodString      string    `json:"food_string"`
	Date            time.Time `json:"date"`
	NutritionInfoId int64     `json:"nutrition_id"`
}

type CustomRecipe struct {
	Id                     int64     `json:"id"`
	Name                   string    `json:"recipe_name"`
	FoodListString         string    `json:"food_string"`
	AlternativeRecipeNames *string   `json:"alt_recipe_names,omitempty"`
	ServingSize            int64     `json:"serving_size"`
	Active                 bool      `json:"active"`
	NutritionInfoId        int64     `json:"nutrition_id"`
	LastModified           time.Time `json:"last_modified"`
}

type FoodRecord struct {
	Id    int64  `json:"id"`
	Food  string `json:"food"`
	Count int64  `json:"count"`
}

type NutritionData struct {
}
