<!-- nutrition display will contain the Label and all other associated information like the breakdown or edit functionality -->

<script lang="ts">
    import NutritionLabel from "./NutritionLabel.svelte";
    import FoodBreakdown from "./FoodBreakdown.svelte";
    import ErrorBreakdown from "./ErrorBreakdown.svelte";
    import type { UserFoodQueryResponse } from "./models/ResponseModels";

    let {
        nutritionResponse,
        nutritionLabelIsVisible = true,
        nutritionBreakdownIsVisible = true,
    }: {
        nutritionResponse: UserFoodQueryResponse;
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
                totalNutritionInfo={nutritionResponse.total_nutrition_data}
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
                {#each nutritionResponse.food_list.food_list as food}
                    <FoodBreakdown item={food} />
                {/each}
                {#each nutritionResponse.error_list as error}
                    <ErrorBreakdown {error} colSpan={columns.length} />
                {/each}
            </tbody>
        </table>
    {/if}
</div>
