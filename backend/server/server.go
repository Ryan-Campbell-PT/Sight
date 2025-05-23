package server

import (
	"github.com/Ryan-Campbell-PT/Sight/backend/nutrition"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*

func getTotalNutritionInformation_fromFoodListString(foodListString string) (FoodItem, error) {
	functionName := "getTotalNutritionInformation_fromFoodListString/"
	nutritionInfo, err := buildNutritionixResponse_fromFoodListString(foodListString)
	if handleError(functionName+"Error getting response from foodliststring: ", err) {
		return FoodItem{}, err
	}

	return makeTotalNutritionData_fromFoodList(nutritionInfo.Foods), nil
}

*/

func RunServer() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	router.POST("/postNaturalLanguageRequest", nutrition.GetNaturalLanguageJson)
	// router.POST("/postSaveRecipe", post_saveRecipe)
	// router.GET("/getActiveRecipes", get_activeRecipes)
	// router.GET("/getInactiveRecipes", get_inactiveRecipes)
	// router.GET("/getAllRecipes", get_allRecipes)

	router.Run(":8080")
}
