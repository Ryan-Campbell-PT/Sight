<!-- nutrition display will contain the Label and all other associated information like the breakdown or edit functionality -->
<script lang="ts">
    import NutritionLabel from "./NutritionLabel.svelte";
    import type { NaturalLanguageResponseObject } from "$lib/NutritionData";
    import {
        NutritionLabelContent,
        MacroNutrientStrings,
    } from "$lib/NutritionData";
    import FoodBreakdown from "./FoodBreakdown.svelte";
    import ErrorBreakdown from "./ErrorBreakdown.svelte";
    import { Accordion } from "@sveltestrap/sveltestrap";
    import Page from "../../routes/+page.svelte";
    import { render } from "svelte/server";

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
        MacroNutrientStrings.Calorie,
        MacroNutrientStrings.TotalCarbohydrate,
        MacroNutrientStrings.Protein,
        MacroNutrientStrings.Sodium,
        MacroNutrientStrings.Sugar,
        MacroNutrientStrings.TotalFat,
    ];
    let getMacroUnit = (macro: string): string => {
        const ret = NutritionLabelContent.find((m) => m.macro_name === macro);
        if (!ret) return "";
        return ret.unit;
    };

    const columnTitle = (col: string): string => {
        const macroUnit = getMacroUnit(col);
        const macroUnitString = macroUnit ? ` (${macroUnit})` : "";
        return `<th scope="col">${col}${macroUnitString}</th>`;
    };
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
                        {@html columnTitle(col)}
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
