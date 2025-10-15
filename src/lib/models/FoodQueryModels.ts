// Counterpart to NutritionixFood, this holds all the necessary data for the app,
// condensed down and given more functionality
export class FoodItem {
    food_name: string;
    brand_name?: string;
    serving_qty: number;
    serving_unit: string;
    is_recipe: boolean;
    full_nutrient_dict: Map<number, number>;

    constructor(
        food_name: string = "",
        brand_name: string = "",
        serving_qty: number = 0,
        serving_unit: string = "",
        is_recipe: boolean = false,
        full_nutrients: NutritionixNutrient[] = []
    ) {
        this.food_name = food_name;
        this.brand_name = brand_name;
        this.serving_qty = serving_qty;
        this.serving_unit = serving_unit;
        this.is_recipe = is_recipe;

        const dict = new Map<number, number>();
        full_nutrients.forEach((nut) => {
            dict.set(nut.attr_id, Number(nut.value));
        });
        this.full_nutrient_dict = dict;
    }
}

// Helper interface for nutrient objects
export interface NutritionixNutrient {
    attr_id: number;
    value: number;
}

// Counterpart to the NutritionixNaturalLanguageResponse, a container for FoodItems
export class ListOfFoods {
    food_list: FoodItem[];

    constructor(jsonResponse: NutritionixNaturalLanguageResponse) {
        this.food_list = jsonResponse.foods.map(
            (food) =>
                new FoodItem(
                    food.food_name,
                    food.brand_name,
                    food.serving_qty,
                    food.serving_unit,
                    false,
                    food.full_nutrients
                )
        );
    }

    get_total_nutrition_data(): FoodItem {
        const ret = new FoodItem();

        this.food_list.forEach((food) => {
            food.full_nutrient_dict.forEach((value, key) => {
                const prev = ret.full_nutrient_dict.get(key) ?? 0.0;
                ret.full_nutrient_dict.set(key, parseFloat((prev + value).toFixed(2)));
            });
        });

        return ret;
    }
}

// Mock type for the Nutritionix response
export interface NutritionixNaturalLanguageResponse {
    foods: Array<{
        food_name: string;
        brand_name?: string;
        serving_qty: number;
        serving_unit: string;
        full_nutrients: NutritionixNutrient[];
    }>;
}
