<!-- nutrition display will contain the Label and all other associated information like the breakdown or edit functionality -->
<script lang="ts">
    import NutritionLabel from "./NutritionLabel.svelte";
    import { NutritionResponseObject } from "./NutritionData";
    import FoodBreakdown from "./FoodBreakdown.svelte";

    export let nutritionResponse: NutritionResponseObject;
    export let nutritionLabelIsVisible: boolean;
    export let nutritionBreakdownIsVisible: boolean;
</script>

{#if nutritionLabelIsVisible || nutritionBreakdownIsVisible}
    <div>
        <div class="nutritionLabel">
            <NutritionLabel
                totalNutritionInfo={nutritionResponse.getTotalNutritionData(0)}
                isVisible={nutritionLabelIsVisible}
            />
        </div>
        {#if nutritionBreakdownIsVisible}
            <table class="total-food-nutrition-breakdown table">
                <thead>
                    <tr>
                        <th scope="col">Image</th>
                        <th scope="col">Food</th>
                        <th scope="col">Calories</th>
                        <th scope="col">Carbs</th>
                        <th scope="col">Protein</th>
                        <th scope="col">Sodium</th>
                        <th scope="col">Sugar</th>
                        <th scope="col">Fat</th>
                    </tr>
                </thead>
                <tbody>
                    {#each nutritionResponse.foods as food}
                        <FoodBreakdown
                            item={food}
                        />
                    {/each}
                </tbody>
            </table>
        {/if}
    </div>
{/if}
