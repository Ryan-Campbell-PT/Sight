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