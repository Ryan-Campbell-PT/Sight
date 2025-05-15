<script lang="ts">
    import {
        type FoodItem,
        MacroNutrientIds,
        MacroNutrientStrings,
    } from "./NutritionData";
    import { roundToDecimal, getNutritionValueFromName } from "./util";

    let { item }: { item: FoodItem } = $props();
</script>

<tr class="food-breakdown {item.food_name}">
    <th scope="row"
        ><img
            style="width: 50px; height: 50px;"
            alt={item.food_name}
            src={item.photo.thumb}
        /></th
    >
    <td>{item.food_name}</td>
    <td
        >{roundToDecimal(
            getNutritionValueFromName(
                MacroNutrientIds.Calorie,
                item.full_nutrient_map,
            ),
            0,
        )}
    </td>
    <td
        >{roundToDecimal(
            getNutritionValueFromName(
                MacroNutrientIds.TotalCarbohydrate,
                item.full_nutrient_map,
            ),
            0,
        )}</td
    >
    <td
        >{roundToDecimal(
            getNutritionValueFromName(
                MacroNutrientIds.Protein,
                item.full_nutrient_map,
            ),
            0,
        )}</td
    >
    <td
        >{roundToDecimal(
            getNutritionValueFromName(
                MacroNutrientIds.Sodium,
                item.full_nutrient_map,
            ),
            0,
        )}</td
    >
    <td
        >{roundToDecimal(
            getNutritionValueFromName(
                MacroNutrientIds.Sugar,
                item.full_nutrient_map,
            ),
            0,
        )}</td
    >
    <td
        >{roundToDecimal(
            getNutritionValueFromName(
                MacroNutrientIds.TotalFat,
                item.full_nutrient_map,
            ),
            0,
        )}</td
    >
</tr>
<!-- Just keeping this here in case a simple text display is needed
<div class="nutritionBreakdown">
    <div class="breakdown_foodname">
        <span>Food: {item.food_name}</span>
    </div>
    <div class="breakdown_calories">
        <span>Calories: {item.nf_calories}</span>
    </div>
    <div class="breakdown_protein">
        <span>Protein: {item.nf_protein}</span>
    </div>
    <div class="breakdown_totalfat">
        <span>Total Fat: {item.nf_total_fat}</span>
    </div>
    <div class="breakdown_sodium">
        <span>Sodium: {item.nf_sodium}</span>
    </div>
    <div class="breakdown_sugar">
        <span>Sugars: {item.nf_sugars}</span>
    </div>
</div>
-->
