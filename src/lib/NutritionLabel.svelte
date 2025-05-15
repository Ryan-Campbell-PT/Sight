<!-- the nutrition label is the html that will represent the standard Label you see regularly irl -->
<script lang="ts">
    import type { FoodItem } from "./NutritionData";
    import {
        NutritionLabelContent,
        MacroNutrientIds,
        MacroNutrientStrings,
    } from "./NutritionData";
    import { bootstrap } from "./bootstrapClasses";
    import { getNutritionValueFromName } from "./util";

    let {
        totalNutritionInfo,
        isVisible,
    }: { totalNutritionInfo: FoodItem; isVisible: boolean } = $props();

    const calculatePercentage = (currentValue: number, dailyValue: number) => {
        return Math.round((currentValue / dailyValue) * 100);
    };

    const createNutritionRow = (
        name: string,
        isExtendedMacro: boolean,
    ): string => {
        if (!totalNutritionInfo || !totalNutritionInfo.full_nutrients)
            return "";
        const macroValue = NutritionLabelContent.find(
            (m) => m.macro_name == name,
        );
        if (macroValue == null) return "";
        const totalNutritionMacro = totalNutritionInfo.full_nutrients.find(
            (m) => m.attr_id === macroValue.id,
        );
        if (totalNutritionMacro == null) return "";
        return `
                <div class="${isExtendedMacro ? bootstrap.nutritionLabel.extendedMacroRow : bootstrap.nutritionLabel.macroRow}">
                    <div class="${bootstrap.nutritionLabel.macroNameValue}">
                        ${isExtendedMacro ? `<span>${name}</span>` : `<b>${name}</b>`}
                        <span>${totalNutritionMacro.value ?? 0}${macroValue.unit}</span>
                    </div>
                    ${macroValue.daily_value ? `<b class="${bootstrap.nutritionLabel.percentage}">${calculatePercentage(totalNutritionMacro.value, macroValue.daily_value)}%</b>` : ``}
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
                <span
                    >{getNutritionValueFromName(
                        MacroNutrientIds.Calorie,
                        totalNutritionInfo.full_nutrient_map,
                    )}</span
                >
            </div>
        </div>
        <hr class="mx-1" />
        <div class="m-3">
            <span class="">% Daily Value*</span>
            <!-- TODO make these variables, not hard strings -->
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
    </div>
{/if}
