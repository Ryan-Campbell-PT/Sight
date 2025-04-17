import { roundToDecimal } from "../src/util"

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
    
    class NutritionResponseObject {
        foods: FoodItem[] = [];

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
        //for now, the overhead of full_nutrients is too much, can be resolved later
        // full_nutrients: Nutrient[];
    }

    export { NutritionResponseObject, NutritionMacros }
    export type { FoodItem }