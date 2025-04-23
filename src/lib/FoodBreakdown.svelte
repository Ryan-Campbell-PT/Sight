<script lang="ts">
    import { type FoodItem } from "./NutritionData";
    import { roundToDecimal } from "./util";

    export let item: FoodItem;
</script>
<tr class="food-breakdown {item.food_name}">
    <th scope="row"><img style="width: 50px; height: 50px;" alt={item.food_name} src={item.photo.thumb}/></th>
    <td>{item.food_name}</td>
    <td>{roundToDecimal(item.nf_calories, 0)}</td>
    <td>{roundToDecimal(item.nf_total_carbohydrate ?? "0", 0)}g</td>
    <td>{roundToDecimal(item.nf_protein ?? "0", 0)}g</td>
    <td>{roundToDecimal(item.nf_sodium ?? "0", 0)}mg</td>
    <td>{roundToDecimal(item.nf_sugars ?? "0", 0)}g</td>
    <td>{roundToDecimal(item.nf_total_fat ?? "0", 0)}g</td>
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