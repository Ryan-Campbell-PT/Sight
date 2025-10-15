import type { ListOfFoods } from "./FoodQueryModels";
import type { NixFoodItem } from "./NutritionixModels";

interface NaturalLanguageResponseObject {
    foods: NixFoodItem[];
    totalNutritionInformation: NixFoodItem;
    errors: NutritionErrorObject[];
}

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

export interface UserFoodQueryResponse extends DefaultErrorResponse {
    error_list: AnalysisErrorObject[]
    food_list: ListOfFoods
}