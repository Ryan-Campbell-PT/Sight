import { NutritionLabelContent, type FoodItem, type MacroNutrientStrings } from "./NutritionData"

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
export function createNutrientMap(food: FoodItem): Map<number, number> {
    const ret = new Map<number, number>();

    // because TS thinks full_nutrient_map is a real map
    // even though its actually json poorly parsed into a map
    // you have to tell TS to first ignore the typing it associates with it (as unknown)
    // then you tell it its actually formatted as json (Record<string, number>) that looks like below
    // {
    //     "203": 24.5,
    //     "204": 8.9
    // }
    // all this to bypass the issue that is JSON cant send over true maps
    const rawObj = food.full_nutrient_map as unknown as Record<string, number>;

    for (const key in rawObj) {
        if (rawObj.hasOwnProperty(key)) {
            const numKey = Number(key);
            ret.set(numKey, rawObj[key]);
        }
    }

    return ret
}
