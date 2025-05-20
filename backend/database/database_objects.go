package database

import "time"

type Daily struct {
	ID              int64     `json:"id"`
	FoodString      string    `json:"food_string"`
	Date            time.Time `json:"date"`
	NutritionInfoId int64     `json:"nutrition_id"`
}

type Recipe struct {
	Id              int64  `json:"id"`
	Name            string `json:"recipe_name"`
	FoodListString  string `json:"food_string"`
	ServingSize     int64  `json:"serving_size"`
	Active          bool   `json:"active"`
	NutritionInfoId int64  `json:"nutrition_id"`
}

type FoodRecord struct {
	Id    int64  `json:"id"`
	Food  string `json:"food"`
	Count int64  `json:"count"`
}

type NutritionData struct {
}
