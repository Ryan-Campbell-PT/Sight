import type { NixFoodItem } from "./NutritionixModels";

interface NaturalLanguageResponseObject {
    foods: NixFoodItem[];
    totalNutritionInformation: NixFoodItem;
    errors: NutritionErrorObject[];
}

interface RecipeResponseObject {
    recipeList: Recipe[];
}

interface SaveRecipeResponse extends DefaultErrorResponse {
}

export interface GetActiveRecipes extends DefaultErrorResponse {
    recipe_list: Recipe[];
}