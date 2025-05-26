export interface Nutrient {
    attr_id: number;
    value: number;
}

export interface AltMeasure {
    serving_weight: number;
    measure: string;
    seq: number | null;
    qty: number;
}

export interface Photo {
    thumb: string;
    highres: string;
    is_user_uploaded: boolean;
}

// aligns with nutrition_objects.go/CustomFoodItem
export interface FoodItem {
    food_name: string;
    serving_qty: number;
    serving_unit: string;
    full_nutrient_map: Map<number, number>;
    photo: Photo;
}

export interface NutritionErrorObject {
    errorString: string;
}

// aligns with NutritionInfoResponse in nutrition_objects.go
export interface NaturalLanguageResponseObject {
    foods: FoodItem[];
    total_nutrition_information: FoodItem;
    errors: NutritionErrorObject[];
}

export interface RecipeResponseObject {
    recipeList: CustomRecipe[];
}

export interface CustomRecipe {
    id: number;
    recipe_name: string;
    food_string: string;
    serving_size: number;
    nutrition_id: number;
    active: boolean;
}

export interface NutritionixNutrient {
    id: number;
    macro_name: string;
    unit: string;
    daily_value: number | null;
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