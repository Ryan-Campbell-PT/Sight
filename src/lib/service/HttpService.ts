// this file will handle contacting the backend via fetch commands

import type { SaveRecipeRequest } from "$lib/models/RequestModels";
import type { GetActiveRecipes, SaveRecipeResponse } from "$lib/models/ResponseModels";

const host = "http://localhost:8080/"


// TODO will eventuall want to return an error class 
export let save_recipe = async (r: Recipe) => {
    if (!r) return;

    const request: SaveRecipeRequest = {
        recipe_name: r.recipe_name,
        recipe_servings: r.serving_size,
        user_food_query: r.food_string,
        recipe_id: r.id,
    };

    const res = await fetch(host + "post_recipe", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(request),
    });

    if (res.ok) {
        const response: SaveRecipeResponse = await res.json();
        if (response && response.success) {
            // display some success alert
            console.log("Save Recipe success");
        } else {
            //display some error
            console.log("Save recipe errror");
        }
    }

}

export let get_active_recipes = async (): Promise<Recipe[]> => {
    const res = await fetch(host + "get_active_recipes", {
        method: "GET",
        headers: { "Content-Type": "application/json" },
    });

    if (res.ok) {
        const response = (await res.json()) as GetActiveRecipes;
        if (response.success)
            return response.recipe_list
    }

    return []
};
