package nutrition

import (
	"time"
)

// FullNutrientMap is my modified data structure from NutritionixFoodItem.FullNutrient
// to turn the regular Nutrient object into a Map collection for easier use

// any manipulation done with a FoodItem should be done with this struct
// this contains the map, and should be used in any instances a macro needs to be accessed
type CustomFoodItem struct {
	FoodName        string            `json:"food_name"`
	ServingQty      float64           `json:"serving_qty"`
	ServingUnit     string            `json:"serving_unit"`
	FullNutrientMap map[int64]float64 `json:"full_nutrient_map"`
	Photo           NutritionixPhoto  `json:"photo"`
}

type DailyNutrition struct {
	// AllInformation  Nutritionix_NaturalLanguageResponse
	NutritionValues CustomFoodItem
	FoodListString  string
	Date            time.Time
}

type NutritionLabelNutrient struct {
	ID         int    `json:"id"`
	MacroName  string `json:"macro_name"`
	Unit       string `json:"unit"`
	DailyValue *int   `json:"daily_value"`
}

type NutritionErrorObject struct {
	ErrorString string `json:"errorString"`
}

// TODO does not align with anything yet
type GetNutritionRequestBody struct {
	FoodListString string `json:"foodListString"`
	Date           string `json:"date"`
	SaveToDb       bool   `json:"saveToDb"`
}

// enriched type used to pass to the front end, using the data from Nutritionix
// aligns with NaturalLanguageResponseObject in NutritionData.ts
type NaturalLanguageResponse struct {
	Foods                     []CustomFoodItem       `json:"foods"`
	TotalNutritionInformation CustomFoodItem         `json:"total_nutrition_information"`
	Errors                    []NutritionErrorObject `json:"errors"`
}

// the response directly from the API
type NutritionixAPINaturalLanguageResponse struct {
	Foods []NutritionixFoodItem `json:"foods"`
}
