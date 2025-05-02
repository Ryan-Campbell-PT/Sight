<!-- the nutrition label is the html that will represent the standard Label you see regularly irl -->
<script lang="ts">
    import { NutritionMacros, NutritionLabelContent } from "./NutritionData";
    import { bootstrap } from "./bootstrapClasses";

    export let totalNutritionInfo: NutritionMacros;
    export let isVisible: boolean;

    const calculatePercentage = (currentValue: number, dailyValue: number) => {
        return Math.round(currentValue / dailyValue * 100)
    }

    const createNutritionRow = (name: string, isExtendedMacro: boolean): string => {
        const macroValue = NutritionLabelContent.find(m => m.macro_name == name)
        if(macroValue == null) return ""
        const totalNutritionMacro = totalNutritionInfo.full_nutrients.find(m => m.attr_id === macroValue.id)
        if(totalNutritionMacro == null) return ""
        return (
            `
                <div class="${isExtendedMacro ? bootstrap.nutritionLabel.extendedMacroRow : bootstrap.nutritionLabel.macroRow}">
                    <div class="${bootstrap.nutritionLabel.macroNameValue}">
                        ${isExtendedMacro ? `<span>${name}</span>` : `<b>${name}</b>`}
                        <span>${totalNutritionMacro.value ?? 0}${macroValue.unit}</span>
                    </div>
                    ${macroValue.daily_value ? `<b class="${bootstrap.nutritionLabel.percentage}">${calculatePercentage(totalNutritionMacro.value, macroValue.daily_value)}%</b>` : ``}
                </div>
            `
        )
    }
</script>

{#if isVisible}
    <div class="border border-3 border-dark" style="width: 300px;">
        <div class="m-2">
            <h3>Nutrition Facts</h3>
            <span class="fs-6">Amount Per Serving</span>
            <br/>
            <div class="fs-4 d-flex justify-content-between">
                <span>Calories</span>
                <span>{totalNutritionInfo.calories}</span>
            </div>
        </div>
        <hr class="mx-1"/>
        <div class="m-3">
            <span class="">% Daily Value*</span>
            {#each ["Total Fat"] as m}
                {@html createNutritionRow(m, false)}
            {/each}
            <div>
                {#each ["Saturated Fat", "Trans Fat", "Polyunsaturated Fat", "Monounsaturated Fat"] as m}
                    {@html createNutritionRow(m, true)}
                {/each}
            </div>
            {#each ["Cholesterol", "Sodium", "Total Carbohydrate"] as m}
                {@html createNutritionRow(m, false)}
            {/each}
            <div>
                {#each ["Dietary Fiber", "Sugar"] as m}
                    {@html createNutritionRow(m, true)}
                {/each}
            </div>
            {#each ["Protein"] as m}
                {@html createNutritionRow(m, false)}
            {/each}
        </div>
    </div>
{/if}