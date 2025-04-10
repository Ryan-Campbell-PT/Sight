<script lang="ts">
    import { type FoodItem } from "./NutritionData";
    export let item: FoodItem;
</script>
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