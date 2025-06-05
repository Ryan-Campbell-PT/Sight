package models

import "github.com/Ryan-Campbell-PT/Sight/backend/util"

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

// id values from https://docx.syndigo.com/developers/docs/list-of-all-nutrients-and-nutrient-ids-from-api
// daily values taken from https://www.fda.gov/food/nutrition-facts-label/how-understand-and-use-nutrition-facts-label
// when making updates, be sure to update NutritionData.ts/NutritionLabelContent
var NutritionLabelContent = []NutritionLabelNutrient{
	{ID: util.CaloriesId, MacroName: util.CaloriesString, Unit: "kcal", DailyValue: util.IntPtr(2000)},
	{ID: util.TotalCarbohydrateId, MacroName: util.TotalCarbohydrateString, Unit: "g", DailyValue: nil},
	{ID: util.TotalFatId, MacroName: util.TotalFatString, Unit: "g", DailyValue: util.IntPtr(78)},
	{ID: util.SaturatedFatId, MacroName: util.SaturatedFatString, Unit: "g", DailyValue: nil},
	{ID: util.TransFatId, MacroName: util.TransFatString, Unit: "g", DailyValue: nil},
	{ID: util.PolyunsaturatedFatId, MacroName: util.PolyunsaturatedFatString, Unit: "g", DailyValue: nil},
	{ID: util.MonounsaturatedFatId, MacroName: util.MonounsaturatedFatString, Unit: "g", DailyValue: nil},
	{ID: util.ProteinId, MacroName: util.ProteinString, Unit: "g", DailyValue: nil},
	{ID: util.SugarId, MacroName: util.SugarString, Unit: "g", DailyValue: util.IntPtr(50)},
	{ID: util.SodiumId, MacroName: util.SodiumString, Unit: "mg", DailyValue: util.IntPtr(2300)},
	{ID: util.DietaryFiberId, MacroName: util.DietaryFiberString, Unit: "g", DailyValue: util.IntPtr(28)},
	{ID: util.CholesterolId, MacroName: util.CholesterolString, Unit: "mg", DailyValue: util.IntPtr(300)},
	{ID: util.PotassiumId, MacroName: util.PotassiumString, Unit: "mg", DailyValue: nil},
	{ID: util.IronId, MacroName: util.IronString, Unit: "mg", DailyValue: nil},
	{ID: util.CaffeineId, MacroName: util.CaffeineString, Unit: "mg", DailyValue: nil},
	{ID: util.PhosphorusId, MacroName: util.PhosphorusString, Unit: "mg", DailyValue: nil},
}
