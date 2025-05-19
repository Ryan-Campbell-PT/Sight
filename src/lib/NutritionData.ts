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

interface NaturalLanguageResponseObject {
    foods: FoodItem[];
    totalNutritionInformation: FoodItem;
    errors: NutritionErrorObject[];

    /*
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
    */

}

interface RecipeResponseObject {
    recipeList: Recipe[];
}

interface Recipe {
    id: number;
    recipe_name: string;
    food_string: string;
    serving_size: number;
    nutrition_id: number;
    active: boolean;
}

export enum MacroNutrientStrings {
    Calorie = "Calories",
    TotalCarbohydrate = "Total Carbohydrate",
    TotalFat = "Total Fat",
    SaturatedFat = "Saturated Fat",
    TransFat = "Trans Fat",
    PolyunsaturatedFat = "Polyunsaturated Fat",
    MonounsaturatedFat = "Monounsaturated Fat",
    Protein = "Protein",
    Sugar = "Sugar",
    Sodium = "Sodium",
    DietaryFiber = "Dietary Fiber",
    Cholesterol = "Cholesterol",
    Potassium = "Potassium",
    Iron = "Iron",
    Caffeine = "Caffeine",
    Phosphorus = "Phosphorus",
}

export enum MacroNutrientIds {
    Calorie = 208,
    TotalCarbohydrate = 205,
    TotalFat = 204,
    SaturatedFat = 606,
    TransFat = 605,
    PolyunsaturatedFat = 646,
    MonounsaturatedFat = 645,
    Protein = 203,
    Sugar = 269,
    Sodium = 307,
    DietaryFiber = 291,
    Cholesterol = 601,
    Potassium = 306,
    Iron = 303,
    Caffeine = 262,
    Phosphorus = 305,
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
export const NutritionLabelContent: NutritionixNutrient[] = [
    { id: MacroNutrientIds.Calorie, macro_name: MacroNutrientStrings.Calorie, unit: "kcal", daily_value: 2000 },
    { id: MacroNutrientIds.TotalCarbohydrate, macro_name: MacroNutrientStrings.TotalCarbohydrate, unit: "g", daily_value: null },
    { id: MacroNutrientIds.TotalFat, macro_name: MacroNutrientStrings.TotalFat, unit: "g", daily_value: 78 },
    { id: MacroNutrientIds.SaturatedFat, macro_name: MacroNutrientStrings.SaturatedFat, unit: "g", daily_value: null },
    { id: MacroNutrientIds.TransFat, macro_name: MacroNutrientStrings.TransFat, unit: "g", daily_value: null },
    { id: MacroNutrientIds.PolyunsaturatedFat, macro_name: MacroNutrientStrings.PolyunsaturatedFat, unit: "g", daily_value: null },
    { id: MacroNutrientIds.MonounsaturatedFat, macro_name: MacroNutrientStrings.MonounsaturatedFat, unit: "g", daily_value: null },
    { id: MacroNutrientIds.Protein, macro_name: MacroNutrientStrings.Protein, unit: "g", daily_value: null },
    { id: MacroNutrientIds.Sugar, macro_name: MacroNutrientStrings.Sugar, unit: "g", daily_value: 50 },
    { id: MacroNutrientIds.Sodium, macro_name: MacroNutrientStrings.Sodium, unit: "mg", daily_value: 2300 },
    { id: MacroNutrientIds.DietaryFiber, macro_name: MacroNutrientStrings.DietaryFiber, unit: "g", daily_value: 28 },
    { id: MacroNutrientIds.Cholesterol, macro_name: MacroNutrientStrings.Cholesterol, unit: "mg", daily_value: 300 },
    { id: MacroNutrientIds.Potassium, macro_name: MacroNutrientStrings.Potassium, unit: "mg", daily_value: null },
    { id: MacroNutrientIds.Iron, macro_name: MacroNutrientStrings.Iron, unit: "mg", daily_value: null },
    { id: MacroNutrientIds.Caffeine, macro_name: MacroNutrientStrings.Caffeine, unit: "mg", daily_value: null },
    { id: MacroNutrientIds.Phosphorus, macro_name: MacroNutrientStrings.Phosphorus, unit: "mg", daily_value: null },
];

export function getNutrientValueFromString(str: MacroNutrientStrings, nutritionMap: Map<number, number>): number {
    const macro = NutritionLabelContent.find(m => m.macro_name === str.toString())
    if (!macro) return 0
    return getNutrientValueFromId(macro.id, nutritionMap)
}

export function getNutrientValueFromId(id: number, nutritionMap: Map<number, number>): number {
    if (nutritionMap)
        return nutritionMap.get(id) ?? 0
    return 0
}

// the json returned from golang does not correctly include Map objects, so it must be manually converted
export function foodItem_MapCorrection(food: FoodItem): Map<number, number> {
    let ret = new Map<number, number>();

    food.full_nutrients.forEach(m => {
        ret.set(m.attr_id, m.value)
    })

    return ret
}

export function naturalLanguageResponseObject_MapCorrection(response: NaturalLanguageResponseObject): NaturalLanguageResponseObject {

    response.foods.map(
        (m) => (m.full_nutrient_map = foodItem_MapCorrection(m)),
    );
    response.totalNutritionInformation.full_nutrient_map =
        foodItem_MapCorrection(
            response.totalNutritionInformation,
        );

    return response
}
export type { FoodItem, Recipe, NutritionixNutrient, NaturalLanguageResponseObject, NutritionErrorObject, RecipeResponseObject }

