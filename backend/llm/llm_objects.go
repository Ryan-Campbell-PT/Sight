package llm

import "github.com/Ryan-Campbell-PT/Sight/backend/nutrition"

type LLMReturnResponse struct {
	OriginalUserInput string                           `json:"originalUserInput"`
	ParsedUserInput   []UserInputParse                 `json:"parsedUserInput"`
	ListOfFoods       []nutrition.CustomFoodItem       `json:"listOfFoods"`
	ListOfRecipes     []CustomRecipeParse              `json:"listOfRecipes"`
	ListOfErrors      []nutrition.NutritionErrorObject `json:"listOfErrors"`
}

// this is the object representation of something like "2 servings of moms chocolate cake"
type UserInputParse struct {
	FoodString string `json:"foodString"`
	// contains the index of where the food was listed in the original string
	OriginalIndex int64 `json:"originalIndex"`
	// this value is only populated if there is a recipe, otherwise -1
	RecipeId int64 `json:"recipeId"`
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
