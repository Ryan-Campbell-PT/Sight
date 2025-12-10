
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
export function foodItem_MapCorrection(food: NixFoodItem): Map<number, number> {
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