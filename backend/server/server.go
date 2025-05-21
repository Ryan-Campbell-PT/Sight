package server

import (
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

// HTTP POST
// given the body of a NutritionRequest

func RunServer() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// router.POST("/postFoodList", post_foodList)
	// router.POST("/postSaveRecipe", post_saveRecipe)
	// router.GET("/getActiveRecipes", get_activeRecipes)
	// router.GET("/getInactiveRecipes", get_inactiveRecipes)
	// router.GET("/getAllRecipes", get_allRecipes)

	router.Run(":8080")
}
