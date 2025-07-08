package models

// aligns with NutritionData.ts/GetUserRecipesResponseObject
type GetUserRecipesResponseObject struct {
	RecipeList []CustomRecipe `json:"recipeList"`
}

// aligns with SaveRecipeRequestObject
type SaveRecipeRequestBody struct {
	/*
		RecipeName             string   `json:"recipeName"`
		AlternativeRecipeNames []string `json:"alternativeRecipeNames"`
		FoodListString         string   `json:"foodListString"`
		NumServings            int64    `json:"numServings"`
	*/
	Recipe CustomRecipe `json:"recipe"`
	//TODO need to implement
	// the user can either upload the food that the recipe contains,
	// or the full macro information instead
	IsMacroInfo bool `json:"isMacroInfo"`
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
