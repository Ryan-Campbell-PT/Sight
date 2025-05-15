package main

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

type Nutritionix_NaturalLanguageResponse struct {
	Foods []FoodItem `json:"foods"`
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

type NaturalLanguageResponseObject struct {
	ListOfFoods               []FoodItem             `json:"foods"`
	TotalNutritionInformation FoodItem               `json:"totalNutritionInformation"`
	Errors                    []NutritionErrorObject `json:"errors"`
}

const (
	CaloriesString           = "Calories"
	TotalCarbohydrateString  = "Total Carbohydrate"
	TotalFatString           = "Total Fat"
	SaturatedFatString       = "Saturated Fat"
	TransFatString           = "Trans Fat"
	PolyunsaturatedFatString = "Polyunsaturated Fat"
	MonounsaturatedFatString = "Monounsaturated Fat"
	ProteinString            = "Protein"
	SugarString              = "Sugar"
	SodiumString             = "Sodium"
	DietaryFiberString       = "Dietary Fiber"
	CholesterolString        = "Cholesterol"
	PotassiumString          = "Potassium"
	IronString               = "Iron"
	CaffeineString           = "Caffeine"
	PhosphorusString         = "Phosphorus"
)

const (
	CaloriesId           = 208
	TotalCarbohydrateId  = 205
	TotalFatId           = 204
	SaturatedFatId       = 606
	TransFatId           = 605
	PolyunsaturatedFatId = 646
	MonounsaturatedFatId = 645
	ProteinId            = 203
	SugarId              = 269
	SodiumId             = 307
	DietaryFiberId       = 291
	CholesterolId        = 601
	PotassiumId          = 306
	IronId               = 303
	CaffeineId           = 262
	PhosphorusId         = 305
)

// id values from https://docx.syndigo.com/developers/docs/list-of-all-nutrients-and-nutrient-ids-from-api
// daily values taken from https://www.fda.gov/food/nutrition-facts-label/how-understand-and-use-nutrition-facts-label
// when making updates, be sure to update NutritionData.ts/NutritionLabelContent
var NutritionLabelContent = []NutritionixNutrient{
	{ID: CaloriesId, MacroName: CaloriesString, Unit: "kcal", DailyValue: intPtr(2000)},
	{ID: TotalCarbohydrateId, MacroName: TotalCarbohydrateString, Unit: "g", DailyValue: nil},
	{ID: TotalFatId, MacroName: TotalFatString, Unit: "g", DailyValue: intPtr(78)},
	{ID: SaturatedFatId, MacroName: SaturatedFatString, Unit: "g", DailyValue: nil},
	{ID: TransFatId, MacroName: TransFatString, Unit: "g", DailyValue: nil},
	{ID: PolyunsaturatedFatId, MacroName: PolyunsaturatedFatString, Unit: "g", DailyValue: nil},
	{ID: MonounsaturatedFatId, MacroName: MonounsaturatedFatString, Unit: "g", DailyValue: nil},
	{ID: ProteinId, MacroName: ProteinString, Unit: "g", DailyValue: nil},
	{ID: SugarId, MacroName: SugarString, Unit: "g", DailyValue: intPtr(50)},
	{ID: SodiumId, MacroName: SodiumString, Unit: "mg", DailyValue: intPtr(2300)},
	{ID: DietaryFiberId, MacroName: DietaryFiberString, Unit: "g", DailyValue: intPtr(28)},
	{ID: CholesterolId, MacroName: CholesterolString, Unit: "mg", DailyValue: intPtr(300)},
	{ID: PotassiumId, MacroName: PotassiumString, Unit: "mg", DailyValue: nil},
	{ID: IronId, MacroName: IronString, Unit: "mg", DailyValue: nil},
	{ID: CaffeineId, MacroName: CaffeineString, Unit: "mg", DailyValue: nil},
	{ID: PhosphorusId, MacroName: PhosphorusString, Unit: "mg", DailyValue: nil},
}

func intPtr(i int) *int {
	return &i
}

// TODO the work done on the front end should be done instead the back end,
// maybe adding an additional property to the Response object with total info
func makeTotalNutritionData_fromFoodList(foodList []FoodItem) FoodItem {
	ret := FoodItem{}

	for _, food := range foodList {
		ret.Calories = roundToNearestDecimal(ret.Calories+food.Calories, 2)
		ret.Cholesterol = roundToNearestDecimal(ret.Cholesterol+food.Cholesterol, 2)
		ret.DietaryFiber = roundToNearestDecimal(ret.DietaryFiber+food.DietaryFiber, 2)
		ret.Phosphorus = roundToNearestDecimal(ret.Phosphorus+food.Phosphorus, 2)
		ret.Potassium = roundToNearestDecimal(ret.Potassium+food.Potassium, 2)
		ret.Protein = roundToNearestDecimal(ret.Protein+food.Protein, 2)
		ret.SaturatedFat = roundToNearestDecimal(ret.SaturatedFat+food.SaturatedFat, 2)
		ret.Sodium = roundToNearestDecimal(ret.Sodium+food.Sodium, 2)
		ret.Sugars = roundToNearestDecimal(ret.Sugars+food.Sugars, 2)
		ret.TotalCarbohydrate = roundToNearestDecimal(ret.TotalCarbohydrate+food.TotalCarbohydrate, 2)
		ret.TotalFat = roundToNearestDecimal(ret.TotalFat+food.TotalFat, 2)

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

		fullNutrientMap := make(map[int64]float64)
		for key, value := range food.FullNutrientMap {
			fullNutrientMap[key] = roundToNearestDecimal(ret.FullNutrientMap[key]+value, 2)
		}

		ret.FullNutrientMap = fullNutrientMap
	}

	return ret
}

func createNutrientMap(nutrientList []Nutrient) map[int64]float64 {
	nutrientMap := make(map[int64]float64)

	for _, n := range nutrientList {
		nutrientMap[n.AttrID] = n.Value
	}

	return nutrientMap
}
