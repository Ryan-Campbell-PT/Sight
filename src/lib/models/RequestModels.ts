export interface NaturalLanguageRequest {
    user_food_query: String
}


export interface SaveRecipeRequest {
    recipe_id?: number; // for if an already made recipe is being updated
    recipe_name: string;
    recipe_servings: number;
    // recipe_color: string; for later
    user_food_query: string;
}