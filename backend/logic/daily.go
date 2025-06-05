package logic

import (
	"database/sql"
	"time"

	"github.com/Ryan-Campbell-PT/Sight/backend/models"
	"github.com/Ryan-Campbell-PT/Sight/backend/util"
)

// get records from the Daily database
// startDate and endDate are INCLUSIVE
// if you only need one date, start and endDate should be the same value

// TODO this is dumb cause SaveRecipe takes the nutritionId but SaveDailyRecord takes a foodItem
// and creates the nutritionId
// that needs to be standardized
func GetDailyRecord(startDate time.Time, endDate time.Time) ([]models.Daily, error) {
	functionName := "GetDailyRecord/"
	var records []models.Daily

	rows, err := db.Query("SELECT * FROM daily WHERE date BETWEEN @StartDate AND @EndDate",
		sql.Named("StartDate", util.GetDate(startDate)),
		sql.Named("EndDate", util.GetDate(endDate)),
	)

	if util.HandleError(functionName+"Error getting daily records from dates: ", err) {
		return records, err
	}
	defer rows.Close()

	for rows.Next() {
		var day models.Daily
		err := rows.Scan(&day.ID, &day.FoodString, &day.Date)
		if util.HandleError(functionName+"Error scanning db row into local variable: ", err) {
			return records, err
		}
		records = append(records, day)
	}

	err = rows.Err()
	if util.HandleError(functionName+"Error with rows returned from database: ", err) {
		return records, err
	}

	return records, nil
}

func SaveDailyRecord(daily models.Daily, nutritionInfo models.CustomFoodItem) error {
	functionName := "SaveDailyRecord/"
	db := GetDatabase()

	nutritionId, err := SaveNutritionInfo(nutritionInfo)
	if util.HandleError(functionName+"Error saving nutrition info to DB: ", err) {
		return err
	}

	_, err = db.Exec(`INSERT INTO daily(food_string, date, nutrition_id) VALUES(@FoodListString, @Date, @NutritionKey)`,
		sql.Named("FoodListString", daily.FoodString),
		sql.Named("Date", daily.Date),
		sql.Named("NutritionKey", nutritionId),
	)

	if util.HandleError(functionName+"Error inserting body values into database: ", err) {
		return err
	}

	return nil
}
