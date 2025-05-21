package nutrition

import (
	"database/sql"

	"github.com/Ryan-Campbell-PT/Sight/backend/database"
	"github.com/Ryan-Campbell-PT/Sight/backend/util"
)

func SaveDailyRecord(daily database.Daily, nutritionInfo FoodItem) error {
	functionName := "SaveDailyRecord/"
	db := database.GetDatabase()

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
