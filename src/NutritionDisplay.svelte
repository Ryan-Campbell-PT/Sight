<!-- nutrition display will contain the Label and all other associated information like the breakdown or edit functionality -->
<script lang="ts">
    import NutritionLabel from "./NutritionLabel.svelte";
    import { NutritionResponseObject } from "./NutritionData";
    import NutritionBreakdown from "./NutritionBreakdown.svelte";

    export let nutritionResponse: NutritionResponseObject;
    export let isVisible: boolean;
</script>

{#if isVisible}
    <div>
        <div class="nutritionLabel">
            <NutritionLabel totalNutritionInfo={nutritionResponse.getTotalNutritionData()} isVisible={isVisible}/>
        </div>
        <div class="totalNutritionBreakdown">
            {#each nutritionResponse.foods as food}
                <NutritionBreakdown item={food}/>
                <br/>
            {/each}
        </div>
    </div>
{/if}
