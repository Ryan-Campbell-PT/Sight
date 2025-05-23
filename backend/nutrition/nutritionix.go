package nutrition

// this file contains objects
// that represent the schema returned back from the nutritionix api

type NutritionixNutrient struct {
	AttrID int64   `json:"attr_id"`
	Value  float64 `json:"value"`
}

type NutritionixAltMeasure struct {
	ServingWeight float64 `json:"serving_weight"`
	Measure       string  `json:"measure"`
	Seq           int64   `json:"seq"`
	Qty           float64 `json:"qty"`
}

type NutritionixPhoto struct {
	Thumb          string `json:"thumb"`
	HighRes        string `json:"highres"`
	IsUserUploaded bool   `json:"is_user_uploaded"`
}

// FullNutrients is how the Nutritionix API returns its data
type NutritionixFoodItem struct {
	FoodName           string                  `json:"food_name"`
	BrandName          *string                 `json:"brand_name"`
	ServingQty         float64                 `json:"serving_qty"`
	ServingUnit        string                  `json:"serving_unit"`
	ServingWeightGrams float64                 `json:"serving_weight_grams"`
	Calories           float64                 `json:"nf_calories"`
	TotalFat           float64                 `json:"nf_total_fat"`
	SaturatedFat       float64                 `json:"nf_saturated_fat"`
	Cholesterol        float64                 `json:"nf_cholesterol"`
	Sodium             float64                 `json:"nf_sodium"`
	TotalCarbohydrate  float64                 `json:"nf_total_carbohydrate"`
	DietaryFiber       float64                 `json:"nf_dietary_fiber"`
	Sugars             float64                 `json:"nf_sugars"`
	Protein            float64                 `json:"nf_protein"`
	Potassium          float64                 `json:"nf_potassium"`
	Phosphorus         float64                 `json:"nf_p"`
	FullNutrients      []NutritionixNutrient   `json:"full_nutrients"`
	AltMeasures        []NutritionixAltMeasure `json:"alt_measures"`
	Photo              NutritionixPhoto        `json:"photo"`
}
