<!-- the nutrition label is the html that will represent the standard Label you see regularly irl -->

<script lang="ts">
    import { NutritionMacros } from "./NutritionData";
    import { bootstrap } from "./bootstrapClasses";

    export let totalNutritionInfo: NutritionMacros;
    export let isVisible: boolean;

    // values taken from https://www.fda.gov/food/nutrition-facts-label/how-understand-and-use-nutrition-facts-label
    const dailyValueTotals = {
        fat: 20,
        sodium: 2300,
        fiber: 28,
        sugar: 50,
        cholesterol: 300,
        carbs: 275
    }

    const calculatePercentage = (currentValue: number, dailyValue: number) => {
        return Math.round(currentValue / dailyValue * 100)
    }

</script>

{#if isVisible}
    <div class="border border-3 border-dark" style="width: 300px;">
        <div class="m-2">
            <h3>Nutrition Facts</h3>
            <!-- might be fs-1 -->
            <span class="fs-6">Amount Per Serving</span>
            <br/>
            <div class="fs-4 d-flex justify-content-between">
                <span>Calories</span>
                <span>{totalNutritionInfo.calories}</span>
            </div>
        </div>
        <hr class="mx-1"/>
        <div class="m-2">
            <span class="">% Daily Value*</span>
            <div id="fats">
                <div class={bootstrap.nutritionLabel.macro}>
                    <b class="me-1">Total Fat</b>
                    <span class="me-auto">{totalNutritionInfo.total_fat ?? 0}g</span>
                    <b class="">{calculatePercentage(totalNutritionInfo.total_fat, dailyValueTotals.fat)}%</b>
                </div>
                <div class="ms-1 fs-6">
                    <div class={bootstrap.nutritionLabel.extendedMacro}>
                        <span>Saturated Fat</span>
                        <span>{totalNutritionInfo.saturated_fat ?? 0}g</span>
                    </div>
                    <div class={bootstrap.nutritionLabel.extendedMacro}>
                        <span>Trans Fat</span>
                        <span>0g</span>
                    </div>
                    <div class={bootstrap.nutritionLabel.extendedMacro}>
                        <span>Polyunsatured Fat</span>
                        <span>0.7g</span>
                    </div>
                    <div class={bootstrap.nutritionLabel.extendedMacro}>
                        <span>Monunsatured Fat</span>
                        <span>0.7g</span>
                    </div>
                </div>
            </div>
            <div class={bootstrap.nutritionLabel.macro}>
                <b class="me-1">Cholesterol</b>
                <span class="me-auto">{totalNutritionInfo.cholesterol ?? 0}mg</span>
                <b class="">{calculatePercentage(totalNutritionInfo.cholesterol, dailyValueTotals.cholesterol)}%</b>
            </div>
            <div class={bootstrap.nutritionLabel.macro}>
                <b class="me-1">Sodium</b>
                <span class="me-auto">{totalNutritionInfo.sodium}mg</span>
                <b class="">{calculatePercentage(totalNutritionInfo.sodium, dailyValueTotals.sodium)}%</b>
            </div>
            <div>
                <div class={bootstrap.nutritionLabel.macro}>
                    <b class="me-1">Total Carbohydrates</b>
                    <span class="me-auto">{totalNutritionInfo.total_carbohydrate ?? 0}mg</span>
                    <b class="">{calculatePercentage(totalNutritionInfo.total_carbohydrate, dailyValueTotals.carbs)}%</b>
                </div>
                <div class="ms-2 fs-6">
                    <div class={bootstrap.nutritionLabel.extendedMacro}>
                        <span class="me-1">Dietary Fiber</span>
                        <span class="me-auto">{totalNutritionInfo.dietary_fiber ?? 0}g</span>
                        <b>34%</b>
                    </div>
                    <div class={bootstrap.nutritionLabel.extendedMacro}>
                        <span>Sugar</span>
                        <span>{totalNutritionInfo.sugars ?? 0}g</span>
                        <b class="">{calculatePercentage(totalNutritionInfo.sugars, dailyValueTotals.sugar)}%</b>
                    </div>
                </div>
            </div>
            <div class={bootstrap.nutritionLabel.macro}>
                <b class="me-1">Protein</b>
                <span class="me-auto">{totalNutritionInfo.protein ?? 0}g</span>
            </div>

        </div>
    </div>
{/if}