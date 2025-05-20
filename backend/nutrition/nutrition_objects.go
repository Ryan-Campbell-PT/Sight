package nutrition

import (
	"time"
)

type Nutrient struct {
	AttrID int64   `json:"attr_id"`
	Value  float64 `json:"value"`
}

type AltMeasure struct {
	ServingWeight float64 `json:"serving_weight"`
	Measure       string  `json:"measure"`
	Seq           int64   `json:"seq"`
	Qty           float64 `json:"qty"`
}

type Photo struct {
	Thumb          string `json:"thumb"`
	HighRes        string `json:"highres"`
	IsUserUploaded bool   `json:"is_user_uploaded"`
}

// FullNutrients is how the Nutritionix API returns its data
// FullNutrientMap is my modified data structure to turn the regular Nutrient object into a Map collection
type FoodItem struct {
	FoodName           string            `json:"food_name"`
	BrandName          *string           `json:"brand_name"`
	ServingQty         float64           `json:"serving_qty"`
	ServingUnit        string            `json:"serving_unit"`
	ServingWeightGrams float64           `json:"serving_weight_grams"`
	Calories           float64           `json:"nf_calories"`
	TotalFat           float64           `json:"nf_total_fat"`
	SaturatedFat       float64           `json:"nf_saturated_fat"`
	Cholesterol        float64           `json:"nf_cholesterol"`
	Sodium             float64           `json:"nf_sodium"`
	TotalCarbohydrate  float64           `json:"nf_total_carbohydrate"`
	DietaryFiber       float64           `json:"nf_dietary_fiber"`
	Sugars             float64           `json:"nf_sugars"`
	Protein            float64           `json:"nf_protein"`
	Potassium          float64           `json:"nf_potassium"`
	Phosphorus         float64           `json:"nf_p"`
	FullNutrients      []Nutrient        `json:"full_nutrients"`
	FullNutrientMap    map[int64]float64 `json:"full_nutrient_map"`
	AltMeasures        []AltMeasure      `json:"alt_measures"`
	Photo              Photo             `json:"photo"`
}

type DailyNutrition struct {
	// AllInformation  Nutritionix_NaturalLanguageResponse
	NutritionValues FoodItem
	FoodListString  string
	Date            time.Time
}

type NutritionixNutrient struct {
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
type NutritionInfoResponse struct {
	Foods                     []FoodItem             `json:"foods"`
	TotalNutritionInformation FoodItem               `json:"total_nutrition_info"`
	Errors                    []NutritionErrorObject `json:"errors"`
}

// the response directly from the API
type NutritionixAPINaturalLanguageResponse struct {
	Foods []FoodItem `json:"foods"`
}
