<script lang="ts">
    import { onMount } from "svelte";
    import EditRecipeModal from "$lib/EditRecipeModal.svelte";
    import { Button, Input } from "@sveltestrap/sveltestrap";
    import EditRecipeBox from "$lib/EditRecipeBox.svelte";
    import {
        get_active_recipes,
        save_recipe,
        post_user_food_query,
    } from "$lib/service/HttpService";

    let activeRecipeList: Recipe[] = $state([] as Recipe[]);
    let userFoodQuery = $state("");

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

    let postUserQuery = async () => {
        const ding = await post_user_food_query(userFoodQuery);
        console.log(ding);
    };
</script>

<div>
    <div>
        <Input id="userfoodquery" type="textarea" bind:value={userFoodQuery} />
        <Button onclick={postUserQuery}>Post user query</Button>
    </div>
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
