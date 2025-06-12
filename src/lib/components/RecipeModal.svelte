<script lang="ts">
    import type { CustomRecipe } from "$lib/NutritionData";
    import { Modal, Input, Label } from "@sveltestrap/sveltestrap";
    import RecipeDisplay from "./RecipeDisplay.svelte";

    let {
        recipe,
        isVisible = $bindable(),
        isEdit,
    }: { recipe: CustomRecipe; isVisible: boolean; isEdit: boolean } = $props();

    let recipeEdit: CustomRecipe = $state(recipe);

    let closeModal = () => {
        isVisible = false;
        recipeEdit = {} as CustomRecipe;
    };

    $effect(() => {
        console.log(recipeEdit.food_string);
    });
</script>

<!-- https://getbootstrap.com/docs/5.3/components/modal/ -->
<!-- https://svelte.dev/playground/27a9b36c6b2a48fb9c98fd9358a8861e?version=5.33.9 -->
<!--
    my idea for the design of this is to have about 2/3 of the screen occupied by the modal
    and the other 1/3 occupied by the nutrition label
    there should be a 'close' 'save' and a 'test' button to check that will run the food string against the api
    so you can confirm the food string matches what youd expect from it 
 -->
<Modal id="testModal" bind:isOpen={isVisible}>
    <div class="modal-header">
        {#if isEdit}
            <Input
                type="text"
                placeholder="Recipe Name"
                bind:value={recipeEdit.recipe_name}
            />
        {:else}
            <h5 class="modal-title">{recipeEdit.recipe_name}</h5>
        {/if}
        <button
            type="button"
            class="btn-close"
            aria-label="Close"
            onclick={closeModal}
        >
        </button>
    </div>
    <div class="modal-body">
        <h5>List of Foods</h5>
        <div class="my-2">
            {#if isEdit}
                <Input
                    type="textarea"
                    placeholder="List of Foods"
                    bind:value={recipeEdit.food_string}
                />
            {:else}
                {recipeEdit.food_string}
            {/if}
        </div>
        <div class="my-2">
            <div class="col-6">
                <Label for="servings">
                    Number of servings
                    <Input
                        type="number"
                        placeholder="Number of Servings"
                        min="1"
                        defaultValue="1"
                        bind:value={recipeEdit.serving_size}
                    />
                </Label>
            </div>
            <div class="col-6"></div>
        </div>
    </div>
    <div class="modal-footer">
        <button type="button" class="btn btn-secondary" onclick={closeModal}
            >Close</button
        >
        <button type="button" class="btn btn-primary">Save changes</button>
    </div>
</Modal>
