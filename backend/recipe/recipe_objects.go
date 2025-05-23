package recipe

import (
	"github.com/Ryan-Campbell-PT/Sight/backend/database"
)

type GetRecipe_RequestBody struct {
	RecipeList []database.CustomRecipe `json:"recipe_list"`
}

// aligns with RecipeRequestObject
type SaveRecipeRequestBody struct {
	RecipeName             string   `json:"recipeName"`
	AlternativeRecipeNames []string `json:"alternativeRecipeNames"`
	FoodListString         string   `json:"foodListString"`
	NumServings            int64    `json:"numServings"`
}

type RecipeResponse struct {
	RecipeName        string `json:"recipe_name"`
	FoodString        string `json:"food_string"`
	Servings          int64  `json:"serving_size"`
	NutritionValuesId int64  `json:"nutrition_id"`
}

// this object is created from the natural language string
// a user types in representing their recipe
// ex: "1.5 servings of moms chocolate cake"
type CustomRecipeParse struct {
	// CustomRecipeId int64 `json:"custom_recipe_id"`
	// full string that was used to represent this item
	FoodString  string  `json:"food_string"`
	RecipeName  string  `json:"recipe_name"`
	NumServings float64 `json:"num_servings"`
	// TotalNutritionInfo nutrition.CustomFoodItem `json:"total_nutrition_information"`
}
