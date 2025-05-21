package nutrition

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Ryan-Campbell-PT/Sight/backend/database"
	"github.com/Ryan-Campbell-PT/Sight/backend/util"
	"github.com/gin-gonic/gin"
)

func SaveNutritionInfo(nutritionInfo FoodItem) (int64, error) {
	db := database.GetDatabase()

	// SELECT at the bottom grabs the id of the row that was just created
	row := db.QueryRow(`
		INSERT INTO nutrition_info (
			calories, protein, carbs, fiber, cholesterol, sugar,
			phosphorus, sodium, total_fat, saturated_fat, poly_fat, mono_fat, potassium
		) VALUES (
			@Calories, @Protein, @Carbs, @Fiber, @Cholesterol, @Sugar,
			@Phosphorus, @Sodium, @TotalFat, @SaturatedFat, @PolyFat, @MonoFat, @Potassium
		);
		SELECT ID = CONVERT(BIGINT, SCOPE_IDENTITY());
	`,
		sql.Named("Calories", GetNutrient(nutritionInfo, util.CaloriesId)),
		sql.Named("Protein", GetNutrient(nutritionInfo, util.ProteinId)),
		sql.Named("Carbs", GetNutrient(nutritionInfo, util.TotalCarbohydrateId)),
		sql.Named("Fiber", GetNutrient(nutritionInfo, util.DietaryFiberId)),
		sql.Named("Cholesterol", GetNutrient(nutritionInfo, util.CholesterolId)),
		sql.Named("Sugar", GetNutrient(nutritionInfo, util.SugarId)),
		sql.Named("Phosphorus", GetNutrient(nutritionInfo, util.PhosphorusId)),
		sql.Named("Sodium", GetNutrient(nutritionInfo, util.SodiumId)),
		sql.Named("TotalFat", GetNutrient(nutritionInfo, util.TotalFatId)),
		sql.Named("SaturatedFat", GetNutrient(nutritionInfo, util.SaturatedFatId)),
		sql.Named("PolyFat", GetNutrient(nutritionInfo, util.PolyunsaturatedFatId)),
		sql.Named("MonoFat", GetNutrient(nutritionInfo, util.MonounsaturatedFatId)),
		sql.Named("Potassium", GetNutrient(nutritionInfo, util.PotassiumId)),
	)
	var nutritionKey int64
	err := row.Scan(&nutritionKey)
	if util.HandleError("Error getting nutritionKey from Recipe Response", err) {
		return -1, err
	}

	return nutritionKey, nil
}

// this function will take the list of foods provided by a user
// and handle all the work associated with that string:
// reaching out to api, marshaling/unmarshaling, building response object
func FetchNaturalLanguageResponse(foodListString string) (NutritionixAPINaturalLanguageResponse, error) {
	functionName := "handle_naturalLanguage_foodList/"
	var nutritionInfo NutritionixAPINaturalLanguageResponse

	request, err := buildNutritionixRequest(foodListString)
	if util.HandleError(functionName+"Error building Nutritionix request: ", err) {
		return nutritionInfo, err
	}

	responseByteArray, err := util.SendHttpRequest(request)
	if util.HandleError(functionName+"Error sending Nutritionix request: ", err) {
		return nutritionInfo, err
	}

	err = json.Unmarshal(responseByteArray, &nutritionInfo)
	if util.HandleError(functionName+"Error reading nutrition info from nutritionix response and unmarshaling to Food item: ", err) {
		return nutritionInfo, err
	}

	// since I have to create the whole map again on the front end, this part of code is no longer needed
	/*
		for i, food := range nutritionInfo.Foods {
			// create the map
			nMap := CreateNutrientMap(food.FullNutrients)
			// assign it
			food.FullNutrientMap = nMap
			// replace the old index with the updated index with the map
			nutritionInfo.Foods[i] = food
		}
	*/

	return nutritionInfo, nil
}

func CheckFoodArrayForErrors(foodListString string, foods []FoodItem) []NutritionErrorObject {
	errorList := []NutritionErrorObject{}
	splitByComma := strings.Split(foodListString, ",")
	if len(splitByComma) > len(foods) {
		responseArrayIndex := 0
		for _, inputString := range splitByComma {
			inputStringTrimmed := strings.ToLower(strings.TrimSpace(inputString))

			// If all known foods have been matched, everything else is an error
			if responseArrayIndex >= len(foods) {
				errorList = append(errorList, NutritionErrorObject{ErrorString: inputStringTrimmed})
				continue
			}

			// TODO this will probably need to be looked at, as what was typed may be slightly different
			// than what the foodName actually is

			// if the string typed by the user contains the food recognized by the api
			foodName := foods[responseArrayIndex].FoodName
			if strings.Contains(inputStringTrimmed, foodName) {
				// then there isnt an issue, and you can move futher along the array
				responseArrayIndex++
			} else {
				// if there is an issue, record the string and add it to the ErrorObject array
				errorList = append(errorList, NutritionErrorObject{ErrorString: inputStringTrimmed})
			}
		}
	}

	return errorList
}

func GetNutritionResponse(c *gin.Context) {
	functionName := "getNutritionResponse/"

	// read the request body
	bodyJson, err := util.ReadRequestBody(c.Request.Body)
	if util.HandleError(functionName+"Error reading query request body: ", err) {
		return
	}

	// put the request body into an object
	var bodyObj GetNutritionRequestBody
	err = json.Unmarshal(bodyJson, &bodyObj)
	if util.HandleError(functionName+"Error reading body from query request: ", err) {
		return
	}

	// pass in the foodListString, get back the information from the api
	naturalLanguageResponseObject, err := FetchNaturalLanguageResponse(bodyObj.FoodListString)
	if util.HandleError(functionName+"Error handling food list from request body: ", err) {
		return
	}

	ret := NutritionInfoResponse{
		Foods:                     naturalLanguageResponseObject.Foods,
		TotalNutritionInformation: MakeTotalNutritionData(naturalLanguageResponseObject.Foods),
		Errors:                    CheckFoodArrayForErrors(bodyObj.FoodListString, naturalLanguageResponseObject.Foods),
	}

	//i dont think this need to exist, it can be reworked
	/* if bodyObj.SaveToDb {
		// save this information to the Daily table
		err = saveToDatabase_NutritionInformation(bodyObj.FoodListString, bodyObj.Date, ret.TotalNutritionInformation)
		if handleError(functionName+"Error saving nutrition info to database: ", err) {
			return
		}
	}
	*/

	// create return object
	responseMarshal, err := json.Marshal(ret)
	if util.HandleError(functionName+"Error marshalling NutritionResponseObject", err) {
		return
	}

	c.JSON(http.StatusOK, string(responseMarshal))
}

func buildNutritionixRequest(foodList string) (*http.Request, error) {
	cfg := util.GetEnvConfig()
	foodQuery := map[string]string{"query": foodList}
	body, err := json.Marshal(foodQuery)
	if err != nil {
		return nil, err
	}

	url := cfg.Nutritionix_domain + cfg.Nutritionix_naturalLanguage
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", cfg.Nutritionix_contentType)
	request.Header.Set("x-app-id", cfg.Nutritionix_appid)
	request.Header.Set("x-app-key", cfg.Nutritionix_appkey)

	return request, nil
}

// id values from https://docx.syndigo.com/developers/docs/list-of-all-nutrients-and-nutrient-ids-from-api
// daily values taken from https://www.fda.gov/food/nutrition-facts-label/how-understand-and-use-nutrition-facts-label
// when making updates, be sure to update NutritionData.ts/NutritionLabelContent
var NutritionLabelContent = []NutritionixNutrient{
	{ID: util.CaloriesId, MacroName: util.CaloriesString, Unit: "kcal", DailyValue: intPtr(2000)},
	{ID: util.TotalCarbohydrateId, MacroName: util.TotalCarbohydrateString, Unit: "g", DailyValue: nil},
	{ID: util.TotalFatId, MacroName: util.TotalFatString, Unit: "g", DailyValue: intPtr(78)},
	{ID: util.SaturatedFatId, MacroName: util.SaturatedFatString, Unit: "g", DailyValue: nil},
	{ID: util.TransFatId, MacroName: util.TransFatString, Unit: "g", DailyValue: nil},
	{ID: util.PolyunsaturatedFatId, MacroName: util.PolyunsaturatedFatString, Unit: "g", DailyValue: nil},
	{ID: util.MonounsaturatedFatId, MacroName: util.MonounsaturatedFatString, Unit: "g", DailyValue: nil},
	{ID: util.ProteinId, MacroName: util.ProteinString, Unit: "g", DailyValue: nil},
	{ID: util.SugarId, MacroName: util.SugarString, Unit: "g", DailyValue: intPtr(50)},
	{ID: util.SodiumId, MacroName: util.SodiumString, Unit: "mg", DailyValue: intPtr(2300)},
	{ID: util.DietaryFiberId, MacroName: util.DietaryFiberString, Unit: "g", DailyValue: intPtr(28)},
	{ID: util.CholesterolId, MacroName: util.CholesterolString, Unit: "mg", DailyValue: intPtr(300)},
	{ID: util.PotassiumId, MacroName: util.PotassiumString, Unit: "mg", DailyValue: nil},
	{ID: util.IronId, MacroName: util.IronString, Unit: "mg", DailyValue: nil},
	{ID: util.CaffeineId, MacroName: util.CaffeineString, Unit: "mg", DailyValue: nil},
	{ID: util.PhosphorusId, MacroName: util.PhosphorusString, Unit: "mg", DailyValue: nil},
}

func intPtr(i int) *int {
	return &i
}

// TODO the work done on the front end should be done instead the back end,
// maybe adding an additional property to the Response object with total info
func MakeTotalNutritionData(foodList []FoodItem) FoodItem {
	ret := FoodItem{}

	for _, food := range foodList {
		ret.Calories = util.RoundToNearestDecimal(ret.Calories+food.Calories, 2)
		ret.Cholesterol = util.RoundToNearestDecimal(ret.Cholesterol+food.Cholesterol, 2)
		ret.DietaryFiber = util.RoundToNearestDecimal(ret.DietaryFiber+food.DietaryFiber, 2)
		ret.Phosphorus = util.RoundToNearestDecimal(ret.Phosphorus+food.Phosphorus, 2)
		ret.Potassium = util.RoundToNearestDecimal(ret.Potassium+food.Potassium, 2)
		ret.Protein = util.RoundToNearestDecimal(ret.Protein+food.Protein, 2)
		ret.SaturatedFat = util.RoundToNearestDecimal(ret.SaturatedFat+food.SaturatedFat, 2)
		ret.Sodium = util.RoundToNearestDecimal(ret.Sodium+food.Sodium, 2)
		ret.Sugars = util.RoundToNearestDecimal(ret.Sugars+food.Sugars, 2)
		ret.TotalCarbohydrate = util.RoundToNearestDecimal(ret.TotalCarbohydrate+food.TotalCarbohydrate, 2)
		ret.TotalFat = util.RoundToNearestDecimal(ret.TotalFat+food.TotalFat, 2)

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
			fullNutrientMap[key] = util.RoundToNearestDecimal(ret.FullNutrientMap[key]+value, 2)
		}

		ret.FullNutrientMap = fullNutrientMap
	}

	return ret
}

func CreateNutrientMap(nutrientList []Nutrient) map[int64]float64 {
	nutrientMap := make(map[int64]float64)

	for _, n := range nutrientList {
		nutrientMap[n.AttrID] = n.Value
	}

	return nutrientMap
}

func GetNutrient(nutritionInfo FoodItem, nutritionId int64) float64 {
	return nutritionInfo.FullNutrientMap[nutritionId]
}
