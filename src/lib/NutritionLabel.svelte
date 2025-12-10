<!-- the nutrition label is the html that will represent the standard Label you see regularly irl -->
<script lang="ts">
    import { FoodItem } from "./models/FoodQueryModels";
    import { roundToDecimal } from "../lib/util";
    import { bootstrap } from "./bootstrapClasses";
    import {
        NutritionValues,
        getNutritionId,
        NutritionMap,
    } from "$lib/models/NutritionConstants";
    import { getNutrientValueFromId } from "$lib/service/NutritionService";

    let {
        totalNutritionInfo,
        isVisible,
    }: { totalNutritionInfo: FoodItem; isVisible: boolean } = $props();

    const calculateDailyValuePercentage = (
        currentValue: number,
        dailyValue: number,
    ) => {
        return Math.round((currentValue / dailyValue) * 100);
    };

    // TODO could change isIndented to indentLevel as a number to provide more indentation options
    const createNutritionRow = (
        macro: NutritionValues,
        isIndented: boolean,
    ): string => {
        if (!totalNutritionInfo || !totalNutritionInfo.full_nutrient_dict)
            return "";
        const macroInfo = NutritionMap[macro];

        if (!macroInfo) return "";
        const macroValue = roundToDecimal(
            getNutrientValueFromId(
                macroInfo.id,
                totalNutritionInfo.full_nutrient_dict,
            ),
            0,
        );

        return `
                <div class="${isIndented ? bootstrap.nutritionLabel.extendedMacroRow : bootstrap.nutritionLabel.macroRow}">
                    <div class="${bootstrap.nutritionLabel.macroNameValue}">
                        ${isIndented ? `<span>${macroInfo.displayName}</span>` : `<b>${macroInfo.displayName}</b>`}
                        <span>${macroValue ?? 0}${macroInfo.unit}</span>
                    </div>
                    ${macroInfo.dailyValue ? `<b class="${bootstrap.nutritionLabel.percentage}">${calculateDailyValuePercentage(macroValue, macroInfo.dailyValue)}%</b>` : ``}
                </div>
            `;
    };
</script>

{#if isVisible}
    <div class="border border-3 border-dark" style="width: 300px;">
        <div class="m-2">
            <h3>Nutrition Facts</h3>
            <span class="fs-6">Amount Per Serving</span>
            <br />
            <div class="fs-4 d-flex justify-content-between">
                <span>Calories</span>
                <span>
                    {roundToDecimal(
                        getNutrientValueFromId(
                            getNutritionId(NutritionValues.Calories),
                            totalNutritionInfo.full_nutrient_dict,
                        ),
                        0,
                    )}
                </span>
            </div>
        </div>
        <hr class="mx-1" />
        <div class="m-3">
            <span class="">% Daily Value*</span>
            {#each [NutritionValues.TotalFat] as m}
                {@html createNutritionRow(m, false)}
            {/each}
            <div>
                {#each [NutritionValues.SaturatedFat, NutritionValues.TransFat, NutritionValues.PolyFat, NutritionValues.MonoFat] as m}
                    {@html createNutritionRow(m, true)}
                {/each}
            </div>
            {#each [NutritionValues.Cholesterol, NutritionValues.Sodium, NutritionValues.TotalCarbohydrate] as m}
                {@html createNutritionRow(m, false)}
            {/each}
            <div>
                {#each [NutritionValues.Fiber, NutritionValues.Sugar] as m}
                    {@html createNutritionRow(m, true)}
                {/each}
            </div>
            {#each [NutritionValues.Protein] as m}
                {@html createNutritionRow(m, false)}
            {/each}
        </div>
    </div>
{/if}
