<!-- nutrition display will contain the Label and all other associated information like the breakdown or edit functionality -->
<script lang="ts">
    import NutritionLabel from "./NutritionLabel.svelte";
    import type { NaturalLanguageResponseObject } from "$lib/NutritionData";
    import FoodBreakdown from "./FoodBreakdown.svelte";
    import ErrorBreakdown from "./ErrorBreakdown.svelte";

    let {
        nutritionResponse,
        nutritionLabelIsVisible = true,
        nutritionBreakdownIsVisible = true,
    }: {
        nutritionResponse: NaturalLanguageResponseObject;
        nutritionLabelIsVisible: boolean;
        nutritionBreakdownIsVisible: boolean;
    } = $props();

    let columns = [
        "Image",
        "Food",
        "Calories",
        "Carbs",
        "Protein",
        "Sodium",
        "Sugar",
        "Fat",
    ];
</script>

<div>
    {#if nutritionLabelIsVisible}
        <div class="nutritionLabel">
            <NutritionLabel
                totalNutritionInfo={nutritionResponse.total_nutrition_information}
                isVisible={nutritionLabelIsVisible}
            />
        </div>
    {/if}
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
                {#each nutritionResponse.foods as food}
                    <FoodBreakdown item={food} />
                {/each}
                {#each nutritionResponse.errors as error}
                    <ErrorBreakdown {error} colSpan={columns.length} />
                {/each}
            </tbody>
        </table>
    {/if}
</div>
