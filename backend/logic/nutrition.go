package logic

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Ryan-Campbell-PT/Sight/backend/models"
	"github.com/Ryan-Campbell-PT/Sight/backend/util"
	"github.com/gin-gonic/gin"
)

// this function saves nutrition information to the database
// in most cases, this nutritionInfo variable
// will be a TotalNutritionInformation variable
func SaveNutritionInfo(nutritionInfo models.CustomFoodItem) (int64, error) {
	db := GetDatabase()

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
func fetchNaturalLanguageResponse(foodListString string) (*models.NutritionixAPINaturalLanguageResponse, error) {
	functionName := "handle_naturalLanguage_foodList/"
	var nutritionInfo models.NutritionixAPINaturalLanguageResponse

	request, err := buildNutritionixRequest(foodListString)
	if util.HandleError(functionName+"Error building Nutritionix request: ", err) {
		return nil, err
	}

	responseByteArray, err := util.SendHttpRequest(request)
	if util.HandleError(functionName+"Error sending Nutritionix request: ", err) {
		return nil, err
	}

	err = json.Unmarshal(responseByteArray, &nutritionInfo)
	if util.HandleError(functionName+"Error reading nutrition info from nutritionix response and unmarshaling to Food item: ", err) {
		return nil, err
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

	return &nutritionInfo, nil
}

func CheckFoodArrayForErrors(foodListString string, foods []models.CustomFoodItem) []models.NutritionErrorObject {
	errorList := []models.NutritionErrorObject{}
	splitByComma := strings.Split(foodListString, ",")
	if len(splitByComma) > len(foods) {
		responseArrayIndex := 0
		for _, inputString := range splitByComma {
			inputStringTrimmed := strings.ToLower(strings.TrimSpace(inputString))

			// If all known foods have been matched, everything else is an error
			if responseArrayIndex >= len(foods) {
				errorList = append(errorList, models.NutritionErrorObject{ErrorString: inputStringTrimmed})
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
				errorList = append(errorList, models.NutritionErrorObject{ErrorString: inputStringTrimmed})
			}
		}
	}

	return errorList
}

// this function will take a food string
// handle contacting the api
// and return a response object of the information
// in most cases, this will be the most commonly called function
func GetNaturalLanguageResponse(foodString string) *models.NaturalLanguageResponse {
	functionName := "GetNutritionInfoResponse/"

	// pass in the foodListString, get back the information from the api
	naturalLanguageResponseObject, err := fetchNaturalLanguageResponse(foodString)
	if util.HandleError(functionName+"Error fetching natural language response: ", err) {
		return nil
	}

	ding := models.NaturalLanguageResponse{
		Foods:                     makeCustomFoodItemArray(naturalLanguageResponseObject.Foods),
		TotalNutritionInformation: MakeTotalNutritionData(naturalLanguageResponseObject.Foods),
		Errors:                    CheckFoodArrayForErrors(foodString, makeCustomFoodItemArray(naturalLanguageResponseObject.Foods)),
	}
	return &ding
}

func makeCustomFoodItemArray(foodItemArray []models.NutritionixFoodItem) []models.CustomFoodItem {
	var ret []models.CustomFoodItem
	for _, food := range foodItemArray {
		ret = append(ret, models.CustomFoodItem{
			FoodName:        food.FoodName,
			Photo:           food.Photo,
			ServingQty:      food.ServingQty,
			ServingUnit:     food.ServingUnit,
			FullNutrientMap: makeNutrientMap(food.FullNutrients),
		})
	}

	return ret
}

// this function will be called from the front end
// to reach out to the nutritionix api
// get the nutrition information
// populate the response object
// and return back to the front end
func GetNaturalLanguageJson(c *gin.Context) {
	functionName := "GetNaturalLanguageResponse/"

	// read the request body
	bodyJson, err := util.ReadRequestBody(c.Request.Body)
	if util.HandleError(functionName+"Error reading query request body: ", err) {
		return
	}

	// put the request body into an object
	var bodyObj models.GetNutritionRequestBody
	err = json.Unmarshal(bodyJson, &bodyObj)
	if util.HandleError(functionName+"Error reading body from query request: ", err) {
		return
	}

	ret := GetNaturalLanguageResponse(bodyObj.FoodListString)
	/*
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
	*/

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
	if util.HandleError(functionName+"Error marshalling NaturalLanguageResponse", err) {
		return
	}

	c.JSON(http.StatusOK, string(responseMarshal))
}

func MakeTotalNutritionData(foodList []models.NutritionixFoodItem) models.CustomFoodItem {
	ret := models.CustomFoodItem{}
	fullNutrientMap := make(map[int64]float64)

	for _, food := range foodList {
		// fullNutrientList := []NutritionixNutrient{}
		for _, n := range food.FullNutrients {

			fullNutrientMap[n.AttrID] = util.RoundToNearestDecimal(fullNutrientMap[n.AttrID]+n.Value, 2)
			/*
					retNut := 0.0
					for _, m := range ret.FullNutrient {
						if n.AttrID == m.AttrID {
							retNut = m.Value
							break
						}
					}
					fullNutrientList = append(fullNutrientList, Nutrient{AttrID: n.AttrID, Value: retNut + n.Value})
				}

				ret.FullNutrients = fullNutrientList

				for key, value := range food.FullNutrientMap {
					fullNutrientMap[key] = util.RoundToNearestDecimal(ret.FullNutrientMap[key]+value, 2)
				}
			*/
		}
	}
	ret.FullNutrientMap = fullNutrientMap

	return ret
}

func makeNutrientMap(nutrientList []models.NutritionixNutrient) map[int64]float64 {
	nutrientMap := make(map[int64]float64)

	for _, n := range nutrientList {
		nutrientMap[n.AttrID] = util.RoundToNearestDecimal(nutrientMap[n.AttrID]+n.Value, 2)
	}

	return nutrientMap
}

func GetNutrient(nutritionInfo models.CustomFoodItem, nutritionId int64) float64 {
	return nutritionInfo.FullNutrientMap[nutritionId]
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
