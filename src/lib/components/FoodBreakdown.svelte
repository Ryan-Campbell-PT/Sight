<script lang="ts">
    import {
        type FoodItem,
        MacroNutrientIds,
        MacroNutrientStrings,
    } from "$lib/NutritionData";
    import { getNutrientValueFromId } from "$lib/NutritionFunctions";
    import { roundToDecimal } from "$lib/util";

    let { item }: { item: FoodItem } = $props();
</script>

<tr class="food-breakdown {item.food_name}">
    <th scope="row">
        <img
            style="width: 50px; height: 50px;"
            alt={item.food_name}
            src={item.photo.thumb}
        />
    </th>
    <td>{item.food_name}</td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                MacroNutrientIds.Calorie,
                item.full_nutrient_map,
            ),
            0,
        )}
    </td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                MacroNutrientIds.TotalCarbohydrate,
                item.full_nutrient_map,
            ),
            0,
        )}
    </td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                MacroNutrientIds.Protein,
                item.full_nutrient_map,
            ),
            0,
        )}
    </td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                MacroNutrientIds.Sodium,
                item.full_nutrient_map,
            ),
            0,
        )}
    </td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                MacroNutrientIds.Sugar,
                item.full_nutrient_map,
            ),
            0,
        )}
    </td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                MacroNutrientIds.TotalFat,
                item.full_nutrient_map,
            ),
            0,
        )}
    </td>
</tr>
