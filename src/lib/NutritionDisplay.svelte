<!-- nutrition display will contain the Label and all other associated information like the breakdown or edit functionality -->
<script lang="ts">
    import NutritionLabel from "./NutritionLabel.svelte";
    import { NutritionResponseObject } from "./NutritionData";
    import FoodBreakdown from "./FoodBreakdown.svelte";
    import ErrorBreakdown from "./ErrorBreakdown.svelte";

    export let nutritionResponse: NutritionResponseObject;
    export let nutritionLabelIsVisible: boolean;
    export let nutritionBreakdownIsVisible: boolean;

    let columns = ['Image', 'Food', 'Calories', 'Carbs', 'Protein', 'Sodium', 'Sugar', 'Fat']
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
                        {#each columns as col}
                            <th scope="col">{col}</th>
                        {/each}
                    </tr>
                </thead>
                <tbody>
                    {#each nutritionResponse.foodInfo.foods as food}
                        <FoodBreakdown
                            item={food}
                        />
                    {/each}
                    {#each nutritionResponse.errors as error}
                        <ErrorBreakdown
                            error={error}
                            colSpan={columns.length}
                        />
                    {/each}
                </tbody>
            </table>
        {/if}
    </div>
{/if}
