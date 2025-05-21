<!-- the nutrition label is the html that will represent the standard Label you see regularly irl -->
<script lang="ts">
    import type { FoodItem } from "$lib/NutritionData";
    import {
        NutritionLabelContent,
        MacroNutrientIds,
        MacroNutrientStrings,
    } from "$lib/NutritionData";
    import { getNutrientValueFromId } from "$lib/NutritionFunctions";
    import { roundToDecimal } from "$lib/util";
    import { bootstrap } from "$lib/bootstrapClasses";

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

    const createNutritionRow = (name: string, isIndented: boolean): string => {
        if (!totalNutritionInfo || !totalNutritionInfo.full_nutrients)
            return "";
        const macroInfo = NutritionLabelContent.find(
            (m) => m.macro_name == name,
        );
        if (!macroInfo) return "";
        const macroValue = roundToDecimal(
            getNutrientValueFromId(
                macroInfo.id,
                totalNutritionInfo.full_nutrient_map,
            ),
            0,
        );
        return `
                <div class="${isIndented ? bootstrap.nutritionLabel.extendedMacroRow : bootstrap.nutritionLabel.macroRow}">
                    <div class="${bootstrap.nutritionLabel.macroNameValue}">
                        ${isIndented ? `<span>${name}</span>` : `<b>${name}</b>`}
                        <span>${macroValue ?? 0}${macroInfo.unit}</span>
                    </div>
                    ${macroInfo.daily_value ? `<b class="${bootstrap.nutritionLabel.percentage}">${calculateDailyValuePercentage(macroValue, macroInfo.daily_value)}%</b>` : ``}
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
                <!-- TODO im not totally sure i like this, may need to be reworked with the similar functionality to whats used in createNutritionRow -->
                <span>
                    {roundToDecimal(
                        getNutrientValueFromId(
                            MacroNutrientIds.Calorie,
                            totalNutritionInfo.full_nutrient_map,
                        ),
                        0,
                    )}
                </span>
            </div>
        </div>
        <hr class="mx-1" />
        <div class="m-3">
            <span class="">% Daily Value*</span>
            {#each [MacroNutrientStrings.TotalFat] as m}
                {@html createNutritionRow(m, false)}
            {/each}
            <div>
                {#each [MacroNutrientStrings.SaturatedFat, MacroNutrientStrings.TransFat, MacroNutrientStrings.PolyunsaturatedFat, MacroNutrientStrings.MonounsaturatedFat] as m}
                    {@html createNutritionRow(m, true)}
                {/each}
            </div>
            {#each [MacroNutrientStrings.Cholesterol, MacroNutrientStrings.Sodium, MacroNutrientStrings.Cholesterol] as m}
                {@html createNutritionRow(m, false)}
            {/each}
            <div>
                {#each [MacroNutrientStrings.DietaryFiber, MacroNutrientStrings.Sugar] as m}
                    {@html createNutritionRow(m, true)}
                {/each}
            </div>
            {#each [MacroNutrientStrings.Protein] as m}
                {@html createNutritionRow(m, false)}
            {/each}
        </div>
        <!-- <div>
            secondary component
        </div>
        <div>
            literally everything
        </div> -->
    </div>
{/if}
