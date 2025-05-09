import { roundToDecimal } from "./util"

interface Nutrient {
    attr_id: number;
    value: number;
}

interface AltMeasure {
    serving_weight: number;
    measure: string;
    seq: number | null;
    qty: number;
}

interface Photo {
    thumb: string;
    highres: string;
    is_user_uploaded: boolean;
}

interface Metadata {
    is_raw_food: boolean;
}

interface Tags {
    item: string;
    measure: string | null;
    quantity: string;
    food_group: number;
    tag_id: number;
}

interface FoodItem {
    food_name: string;
    brand_name: string | null;
    serving_qty: number;
    serving_unit: string;
    serving_weight_grams: number;
    nf_calories: number;
    nf_total_fat: number;
    nf_saturated_fat: number;
    nf_cholesterol: number;
    nf_sodium: number;
    nf_total_carbohydrate: number;
    nf_dietary_fiber: number;
    nf_sugars: number;
    nf_protein: number;
    nf_potassium: number;
    nf_p: number;
    full_nutrients: Nutrient[];
    full_nutrient_map: Map<number, number>;
    nix_brand_name: string | null;
    nix_brand_id: string | null;
    nix_item_name: string | null;
    nix_item_id: string | null;
    upc: string | null;
    consumed_at: string;
    metadata: Metadata;
    source: number;
    ndb_no: number;
    tags: Tags;
    alt_measures: AltMeasure[];
    lat: number | null;
    lng: number | null;
    meal_type: number;
    photo: Photo;
    sub_recipe: string | null;
    class_code: string | null;
    brick_code: string | null;
    tag_id: number | null;
}

interface NutritionErrorObject {
    errorString: string;
}

class NutritionResponseObject {
    foods: FoodItem[] = [];
    errors: NutritionErrorObject[] = [];

    public getTotalNutritionData(decimal: number) : NutritionMacros {
        let ret = new NutritionMacros();
        this.foods.forEach(food => {
            ret.calories = roundToDecimal(ret.calories + food.nf_calories, 0)
            ret.cholesterol = roundToDecimal(ret.cholesterol + food.nf_cholesterol, decimal)
            ret.dietary_fiber = roundToDecimal(ret.dietary_fiber + food.nf_dietary_fiber, decimal)
            ret.phosphorus = roundToDecimal(ret.phosphorus + food.nf_p, decimal)
            ret.potassium = roundToDecimal(ret.potassium + food.nf_potassium, decimal)
            ret.protein = roundToDecimal(ret.protein + food.nf_protein, decimal)
            ret.saturated_fat = roundToDecimal(ret.saturated_fat + food.nf_saturated_fat, decimal)
            ret.sodium = roundToDecimal(ret.sodium + food.nf_sodium, decimal)
            ret.sugars = roundToDecimal(ret.sugars + food.nf_sugars, decimal)
            ret.total_carbohydrate = roundToDecimal(ret.total_carbohydrate + food.nf_total_carbohydrate, decimal)
            ret.total_fat = roundToDecimal(ret.total_fat + food.nf_total_fat, decimal)
            const nutArray: Nutrient[] = []
            food.full_nutrients.forEach(m => {
                const nut = ret.full_nutrients.find(s => s.attr_id == m.attr_id)
                const nutter: Nutrient = { attr_id: m.attr_id, value: 0 }
                // if its null, then this is (likely) the first value in the array
                if (nut) {
                    nutter.value = roundToDecimal(nut.value + m.value, decimal)
                } else {
                    nutter.value = roundToDecimal(m.value, decimal)
                }

                nutArray.push(nutter)
            })
            ret.full_nutrients = nutArray;
        });

        return ret;
    }      

}

class NutritionMacros {
    calories: number = 0;
    total_fat: number = 0;
    saturated_fat: number = 0;
    cholesterol: number = 0;
    sodium: number = 0;
    total_carbohydrate: number = 0;
    dietary_fiber: number = 0;
    sugars: number = 0;
    protein: number = 0;
    potassium: number = 0;
    phosphorus: number = 0;
    full_nutrients: Nutrient[] = [];
}

class RecipeResponseObject {
    recipeList: Recipe[] = []
}

interface Recipe {
    id: number;
    recipe_name: string;
    food_string: string;
    serving_size: number;
    nutrition_id: NutritionMacros;
}

interface NutritionixNutrient {
    id: number;
    macro_name: string;
    unit: string;
    daily_value: number | null;
}

// id values from https://docx.syndigo.com/developers/docs/list-of-all-nutrients-and-nutrient-ids-from-api
// daily values taken from https://www.fda.gov/food/nutrition-facts-label/how-understand-and-use-nutrition-facts-label
// when making updates, be sure to update nutrition.go/NutritionLabelContent
const NutritionLabelContent: NutritionixNutrient[] = [
    { id: 204, macro_name: "Total Fat", unit: "g", daily_value: 78 },
    { id: 606, macro_name: "Saturated Fat", unit: "g", daily_value: null },
    { id: 605, macro_name: "Trans Fat", unit: "g", daily_value: null },
    { id: 646, macro_name: "Polyunsaturated Fat", unit: "g", daily_value: null },
    { id: 645, macro_name: "Monounsaturated Fat", unit: "g", daily_value: null },
    { id: 203, macro_name: "Protein", unit: "g", daily_value: null },
    { id: 269, macro_name: "Sugar", unit: "g", daily_value: 50 },
    { id: 307, macro_name: "Sodium", unit: "mg", daily_value: 2300 },
    { id: 291, macro_name: "Dietary Fiber", unit: "g", daily_value: 28 },
    { id: 601, macro_name: "Cholesterol", unit: "mg", daily_value: 300 },
    { id: 306, macro_name: "Potassium", unit: "mg", daily_value: null },
    { id: 303, macro_name: "Iron", unit: "mg", daily_value: null },
    { id: 262, macro_name: "Caffine", unit: "mg", daily_value: null },
];

export { NutritionResponseObject, NutritionMacros, RecipeResponseObject, NutritionLabelContent }
export type { FoodItem, Recipe, NutritionixNutrient, NutritionErrorObject }

