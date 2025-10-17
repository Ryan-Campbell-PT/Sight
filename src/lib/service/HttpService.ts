// this file will handle contacting the backend via fetch commands

import type { FoodItem } from "$lib/models/FoodQueryModels";
import type { SaveRecipeRequest, NaturalLanguageRequest } from "$lib/models/RequestModels";
import type { UserFoodQueryResponse } from "$lib/models/ResponseModels"
import type { GetActiveRecipes, SaveRecipeResponse } from "$lib/models/ResponseModels";

const host = "http://localhost:8080/"


// TODO will eventually want to return an error class 
export let save_recipe = async (r: Recipe) => {
    if (!r) return;

    const request: SaveRecipeRequest = {
        recipe_name: r.recipe_name,
        recipe_servings: r.serving_size,
        user_food_query: r.food_string,
        recipe_id: r.id,
    };

    const res = await fetch(host + "save_recipe", {
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

export let post_user_food_query = async (userFoodQuery: string): Promise<UserFoodQueryResponse> => {
    // TODO the naming of these request and response objects need to be changed/the same
    const request: NaturalLanguageRequest = { user_food_query: userFoodQuery }
    const res = await fetch(host + "user_food_query", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(request)

    });

    if (res.ok) {
        const response = reviveUserFoodQueryResponse((await res.json()) as UserFoodQueryResponse)
        return response
    }

    // TODO figure this out
    return {} as UserFoodQueryResponse
}

// this is necessary when copying data from api/json, as the nutrition_info comes back as a plain object, not a dict
function reviveFoodItem(data: any): FoodItem {
    const item: FoodItem = data;
    Object.assign(item, data);

    // Convert plain object { "208": 250, "204": 10 } → Map<number, number>
    if (data.full_nutrient_dict && !(data.full_nutrient_dict instanceof Map)) {
        item.full_nutrient_dict = new Map(
            Object.entries(data.full_nutrient_dict).map(([k, v]) => [Number(k), Number(v)])
        );
    }

    return item;
}

function reviveUserFoodQueryResponse(raw: any): UserFoodQueryResponse {
    const response = raw as UserFoodQueryResponse;

    if (response.total_nutrition_data)
        response.total_nutrition_data = reviveFoodItem(response.total_nutrition_data);

    if (Array.isArray(response.food_list))
        // TOOD i never noticed this is a double naming. may want to change
        response.food_list.food_list = response.food_list.map((f) => reviveFoodItem(f));

    return response;
}
