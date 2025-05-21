<!-- this will handle the collection of Recipes, 
allow the display of Active and Inactive Recipes -->
<script lang="ts">
    import type { Recipe } from "$lib/NutritionData";

    let {
        listOfRecipes,
        isEditable,
        showInactive,
        showActive,
    }: {
        listOfRecipes: Recipe[];
        isEditable: boolean;
        showInactive: boolean;
        showActive: boolean;
    } = $props();

    let activeRecipes: Recipe[] = $state(listOfRecipes.filter((m) => m.active));
    let inactiveRecipes: Recipe[] = $state(
        listOfRecipes.filter((m) => !m.active),
    );
</script>

<div class="recipe-display">
    <!-- {#each listOfRecipes as recipe}
        <CustomRecipe {isEditable} {recipe} />
    {/each} -->
</div>
