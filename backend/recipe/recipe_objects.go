package recipe

import (
	"github.com/Ryan-Campbell-PT/Sight/backend/database"
)

type GetRecipe_RequestBody struct {
	RecipeList []database.Recipe `json:"recipe_list"`
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
