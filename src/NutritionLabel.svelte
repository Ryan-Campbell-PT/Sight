<!-- the nutrition label is the html that will represent the standard Label you see regularly irl -->

<script lang="ts">
    import { NutritionMacros } from "./NutritionData";

    export let totalNutritionInfo: NutritionMacros;
    export let isVisible: boolean;
</script>

{#if isVisible}
    <div>
            <span> Calories: {totalNutritionInfo.calories} </span>
            <br/>
            <span> Protein: {totalNutritionInfo.protein} </span>
            <br/>
            <span> Carbs: {totalNutritionInfo.total_carbohydrate} </span>
            <br/>
            <span> Fiber: {totalNutritionInfo.dietary_fiber} </span>

    </div>
{/if}

