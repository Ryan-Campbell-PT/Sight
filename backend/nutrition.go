package main

import (
	"encoding/json"
	"math"
	"time"
)

type Nutrient struct {
	AttrID int     `json:"attr_id"`
	Value  float64 `json:"value"`
}

type AltMeasure struct {
	ServingWeight float64 `json:"serving_weight"`
	Measure       string  `json:"measure"`
	Seq           int     `json:"seq"`
	Qty           int     `json:"qty"`
}

type Photo struct {
	Thumb          string `json:"thumb"`
	HighRes        string `json:"highres"`
	IsUserUploaded bool   `json:"is_user_uploaded"`
}

type Food struct {
	FoodName           string      `json:"food_name"`
	BrandName         *string      `json:"brand_name"`
	ServingQty         float64     `json:"serving_qty"`
	ServingUnit        string      `json:"serving_unit"`
	ServingWeightGrams float64     `json:"serving_weight_grams"`
	Calories           float64     `json:"nf_calories"`
	TotalFat           float64     `json:"nf_total_fat"`
	SaturatedFat       float64     `json:"nf_saturated_fat"`
	Cholesterol        float64     `json:"nf_cholesterol"`
	Sodium             float64     `json:"nf_sodium"`
	TotalCarbohydrate  float64     `json:"nf_total_carbohydrate"`
	DietaryFiber       float64     `json:"nf_dietary_fiber"`
	Sugars             float64     `json:"nf_sugars"`
	Protein            float64     `json:"nf_protein"`
	Potassium          float64     `json:"nf_potassium"`
	Phosphorus         float64     `json:"nf_p"`
	FullNutrients      []Nutrient  `json:"full_nutrients"`
	AltMeasures        []AltMeasure`json:"alt_measures"`
	Photo              Photo       `json:"photo"`
}

type FoodResponse struct {
	Foods []Food `json:"foods"`
}

type DailyNutrition struct {
	AllInformation FoodResponse // save response for reference if needed
	NutritionValues Food
	FoodListString string
	Date time.Time
}
// TODO im not sure what the variable type needed for this is
func makeFoodResponse(responseBody string) FoodResponse {
	var response FoodResponse
	// TODO may need an additional value besides just err
	err := json.Unmarshal([]byte(responseBody), &response)
	if handleError("Error parsing Food Response JSON: ", err) { return response }

	return response
}

func makeDailyNutrition(foodList FoodResponse, date string) DailyNutrition {
	parseDate, err := time.Parse(time.DateOnly, date)
	handleError("Error parsing time for Daily Nutrition: ", err)
	ret := DailyNutrition{
		// AllInformation: foodList,
		Date: parseDate,
		NutritionValues: Food{},
	}
	for _, food := range foodList.Foods {
		ret.NutritionValues.Calories += math.Round(food.Calories)
		ret.NutritionValues.Cholesterol += math.Round(food.Cholesterol)
		ret.NutritionValues.DietaryFiber += math.Round(food.DietaryFiber)
		ret.NutritionValues.Phosphorus += math.Round(food.Phosphorus)
		ret.NutritionValues.Potassium += math.Round(food.Potassium)
		ret.NutritionValues.Protein += math.Round(food.Protein)
		ret.NutritionValues.SaturatedFat += math.Round(food.SaturatedFat)
		ret.NutritionValues.Sodium += math.Round(food.Sodium)
		ret.NutritionValues.Sugars += math.Round(food.Sugars)
		ret.NutritionValues.TotalCarbohydrate += math.Round(food.TotalCarbohydrate)
		ret.NutritionValues.TotalFat += math.Round(food.TotalFat)
	}
	return ret
}