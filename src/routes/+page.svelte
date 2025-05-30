<script lang="ts">
    import Recipe from "$lib/components/Recipe.svelte";
    import RecipeModal from "$lib/components/RecipeModal.svelte";
    import type { CustomRecipe } from "$lib/NutritionData";

    let recipe: CustomRecipe = {
        active: true,
        food_string: "1 banana",
        id: 1,
        nutrition_id: 1,
        recipe_name: "Banana cream pie",
        serving_size: 1,
    };

    let isOpen = $state(false);
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the
    documentation
</p>

<button
    type="button"
    class="btn btn-primary"
    onclick={() => {
        isOpen = !isOpen;
    }}
>
    Launch demo modal
</button>

<RecipeModal {recipe} isVisible={isOpen} />
