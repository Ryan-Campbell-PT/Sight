<script lang="ts">
    import { onMount } from "svelte";
    import EditRecipeModal from "$lib/EditRecipeModal.svelte";
    import { Button } from "@sveltestrap/sveltestrap";
    import EditRecipeBox from "$lib/EditRecipeBox.svelte";
    import { get_active_recipes, save_recipe } from "$lib/service/HttpService";

    let activeRecipeList: Recipe[] = $state([] as Recipe[]);

    onMount(get_active_recipes);

    // modal
    let newRecipeModalIsOpen: boolean = $state(false);

    let setRecipeModalOpen = (isOpen: boolean) => {
        newRecipeModalIsOpen = isOpen;
    };

    let onSave = (r: Recipe) => {
        save_recipe(r);
        setRecipeModalOpen(false);
    };
</script>

<div>
    <div id="activeRecipes" class="d-flex flex-wrap gap-4">
        {#each activeRecipeList as r}
            <EditRecipeBox recipe={r} />
        {/each}
    </div>
    <div>
        <Button onclick={() => setRecipeModalOpen(true)}>New Recipe</Button>

        <EditRecipeModal
            isOpen={newRecipeModalIsOpen}
            {onSave}
            onCancel={() => setRecipeModalOpen(false)}
            recipe={{
                id: -1,
                food_string: "",
                recipe_name: "",
                serving_size: 0,
                active: false,
                nutrition_id: -1,
            } as Recipe}
        />
    </div>
</div>
