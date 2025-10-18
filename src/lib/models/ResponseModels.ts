import type { ListOfFoods, FoodItem } from "./FoodQueryModels";
import type { NixFoodItem } from "./NutritionixModels";

interface RecipeResponseObject {
    recipeList: Recipe[];
}

export interface SaveRecipeResponse extends DefaultErrorResponse {
    recipe_id: number;
    error_list: AnalysisErrorObject[]
}

export interface GetActiveRecipes extends DefaultErrorResponse {
    recipe_list: Recipe[];
}

export interface NaturalLanguageResponse extends DefaultErrorResponse {
    list_of_foods: ListOfFoods
    total_nutrition_data: FoodItem
    error_list: AnalysisErrorObject[]
}