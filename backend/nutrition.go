package main

import (
	"encoding/json"
	"math"
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

type Food struct {
	FoodName           string       `json:"food_name"`
	BrandName          *string      `json:"brand_name"`
	ServingQty         float64      `json:"serving_qty"`
	ServingUnit        string       `json:"serving_unit"`
	ServingWeightGrams float64      `json:"serving_weight_grams"`
	Calories           float64      `json:"nf_calories"`
	TotalFat           float64      `json:"nf_total_fat"`
	SaturatedFat       float64      `json:"nf_saturated_fat"`
	Cholesterol        float64      `json:"nf_cholesterol"`
	Sodium             float64      `json:"nf_sodium"`
	TotalCarbohydrate  float64      `json:"nf_total_carbohydrate"`
	DietaryFiber       float64      `json:"nf_dietary_fiber"`
	Sugars             float64      `json:"nf_sugars"`
	Protein            float64      `json:"nf_protein"`
	Potassium          float64      `json:"nf_potassium"`
	Phosphorus         float64      `json:"nf_p"`
	FullNutrients      []Nutrient   `json:"full_nutrients"`
	AltMeasures        []AltMeasure `json:"alt_measures"`
	Photo              Photo        `json:"photo"`
}

type FoodResponse struct {
	Foods []Food `json:"foods"`
}

type DailyNutrition struct {
	AllInformation  FoodResponse
	NutritionValues Food
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
	FoodString string `json:"foodString"`
}

type NutritionResponseObject struct {
	FoodInfo []Food                 `json:"foods"`
	Errors   []NutritionErrorObject `json:"errors"`
}

const (
	NutrientTotalFat           = 204
	NutrientSaturatedFat       = 606
	NutrientTransFat           = 605
	NutrientPolyunsaturatedFat = 646
	NutrientMonounsaturatedFat = 645
	NutrientProtein            = 203
	NutrientSugar              = 269
	NutrientSodium             = 307
	NutrientDietaryFiber       = 291
	NutrientCholesterol        = 601
	NutrientPotassium          = 306
	NutrientIron               = 303
	NutrientCaffeine           = 262
)

// id values from https://docx.syndigo.com/developers/docs/list-of-all-nutrients-and-nutrient-ids-from-api
// daily values taken from https://www.fda.gov/food/nutrition-facts-label/how-understand-and-use-nutrition-facts-label
// when making updates, be sure to update NutritionData.ts/NutritionLabelContent
var NutritionLabelContent = []NutritionixNutrient{
	{ID: 204, MacroName: "Total Fat", Unit: "g", DailyValue: intPtr(78)},
	{ID: 606, MacroName: "Saturated Fat", Unit: "g", DailyValue: nil},
	{ID: 605, MacroName: "Trans Fat", Unit: "g", DailyValue: nil},
	{ID: 646, MacroName: "Polyunsaturated Fat", Unit: "g", DailyValue: nil},
	{ID: 645, MacroName: "Monounsaturated Fat", Unit: "g", DailyValue: nil},
	{ID: 203, MacroName: "Protein", Unit: "g", DailyValue: nil},
	{ID: 269, MacroName: "Sugar", Unit: "g", DailyValue: intPtr(50)},
	{ID: 307, MacroName: "Sodium", Unit: "mg", DailyValue: intPtr(2300)},
	{ID: 291, MacroName: "Dietary Fiber", Unit: "g", DailyValue: intPtr(28)},
	{ID: 601, MacroName: "Cholesterol", Unit: "mg", DailyValue: intPtr(300)},
	{ID: 306, MacroName: "Potassium", Unit: "mg", DailyValue: nil},
	{ID: 303, MacroName: "Iron", Unit: "mg", DailyValue: nil},
	{ID: 262, MacroName: "Caffine", Unit: "mg", DailyValue: nil},
}

func intPtr(i int) *int {
	return &i
}

// TODO im not sure what the variable type needed for this is
func makeFoodResponse(responseBody string) FoodResponse {
	var response FoodResponse
	// TODO may need an additional value besides just err
	err := json.Unmarshal([]byte(responseBody), &response)
	if handleError("Error parsing Food Response JSON: ", err) {
		return response
	}

	return response
}

// TODO the work done on the front end should be done instead the back end,
// maybe adding an additional property to the Response object with total info
func makeTotalNutritionData(foodList []Food) Food {
	ret := Food{}

	for _, food := range foodList {
		ret.Calories = ret.Calories + math.Round(food.Calories)
		ret.Cholesterol = ret.Cholesterol + math.Round(food.Cholesterol)
		ret.DietaryFiber = ret.DietaryFiber + math.Round(food.DietaryFiber)
		ret.Phosphorus = ret.Phosphorus + math.Round(food.Phosphorus)
		ret.Potassium = ret.Potassium + math.Round(food.Potassium)
		ret.Protein = ret.Protein + math.Round(food.Protein)
		ret.SaturatedFat = ret.SaturatedFat + math.Round(food.SaturatedFat)
		ret.Sodium = ret.Sodium + math.Round(food.Sodium)
		ret.Sugars = ret.Sugars + math.Round(food.Sugars)
		ret.TotalCarbohydrate = ret.TotalCarbohydrate + math.Round(food.TotalCarbohydrate)
		ret.TotalFat = ret.TotalFat + math.Round(food.TotalFat)

		fullNutrientList := []Nutrient{}
		for _, n := range food.FullNutrients {
			retNut := 0.0
			for _, m := range ret.FullNutrients {
				if n.AttrID == m.AttrID {
					retNut = m.Value
					break
				}
			}
			fullNutrientList = append(fullNutrientList, Nutrient{AttrID: n.AttrID, Value: retNut + n.Value})
		}

		ret.FullNutrients = fullNutrientList
	}

	return ret
}

func makeDailyNutrition(foodList FoodResponse, date string) DailyNutrition {
	parseDate, err := time.Parse(time.DateOnly, date)
	handleError("Error parsing date for Daily Nutrition: ", err)
	ret := DailyNutrition{
		// AllInformation: foodList,
		Date:            parseDate,
		NutritionValues: makeTotalNutritionData(foodList.Foods),
	}

	return ret
}
