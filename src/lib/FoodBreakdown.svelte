<script lang="ts">
    import type { FoodItem } from "./models/FoodQueryModels";
    import {
        getNutritionId,
        NutritionValues,
    } from "./models/NutritionConstants";
    import { getNutrientValueFromId } from "./service/NutritionService";
    import { roundToDecimal } from "./util";

    let { item }: { item: FoodItem } = $props();
</script>

<tr class="food-breakdown {item.food_name}">
    <th scope="row">
        <img style="width: 50px; height: 50px;" alt={item.food_name} />
        <!-- src={item.photo.thumb}  -->
    </th>
    <td>{item.food_name}</td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                getNutritionId(NutritionValues.Calories),
                item.full_nutrient_dict,
            ),
            0,
        )}
    </td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                getNutritionId(NutritionValues.TotalCarbohydrate),
                item.full_nutrient_dict,
            ),
            0,
        )}
    </td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                getNutritionId(NutritionValues.Protein),
                item.full_nutrient_dict,
            ),
            0,
        )}
    </td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                getNutritionId(NutritionValues.Sodium),
                item.full_nutrient_dict,
            ),
            0,
        )}
    </td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                getNutritionId(NutritionValues.Sugar),
                item.full_nutrient_dict,
            ),
            0,
        )}
    </td>
    <td>
        {roundToDecimal(
            getNutrientValueFromId(
                getNutritionId(NutritionValues.TotalFat),
                item.full_nutrient_dict,
            ),
            0,
        )}
    </td>
</tr>
