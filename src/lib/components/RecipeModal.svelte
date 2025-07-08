<script lang="ts">
    import type { CustomRecipe } from "$lib/NutritionData";
    import { Modal, Input, Label } from "@sveltestrap/sveltestrap";
    import RecipeDisplay from "./RecipeDisplay.svelte";

    // aligns with SaveRecipeRequestBody
    interface SaveRecipeRequestObject {
        /*
        recipeName: string;
        alternativeRecipeNames: string[];
        foodListString: string;
        numServings: number;
        */
        recipe: CustomRecipe;
        isMacroInfo: boolean;
    }

    let {
        recipeProp,
        isVisible = $bindable(),
    }: { recipeProp: CustomRecipe; isVisible: boolean } = $props();

    let isMacroInfo = $state(false);

    let closeModal = () => {
        isVisible = false;
    };

    // function to check if the passed in recipe is real/populated, or a new recipe (empty)
    let isEmptyRecipe = () => {
        return !(
            recipeProp &&
            recipeProp.food_string &&
            recipeProp.food_string.length > 0 &&
            recipeProp.recipe_name &&
            recipeProp.recipe_name.length > 0 &&
            recipeProp.serving_size
        );
    };

    let saveRecipe = async () => {
        const body: SaveRecipeRequestObject = {
            /*
            alternativeRecipeNames: [],
            foodListString: recipeProp.food_string,
            numServings: recipeProp.serving_size,
            recipeName: recipeProp.recipe_name,
            */
            recipe: recipeProp,
            isMacroInfo: isMacroInfo,
        };
        const res = await fetch("http://localhost:8080/saveRecipe", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body),
        });

        if (!res) {
            // TODO display error
        } else {
            // TODO display success
            console.log("successful saving of recipe: ");
            closeModal();
            // TODO refresh recipe list
        }
    };

    // on state change (modal open) check if its an empty recipe. If so, create something to work with
    $effect(() => {
        if (isEmptyRecipe()) {
            recipeProp = {
                active: true,
                food_string: "",
                recipe_name: "New Recipe",
                serving_size: 1,
            } as CustomRecipe;
        }
        // console.log(recipeEdit.food_string);
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
<Modal bind:isOpen={isVisible}>
    <div class="modal-header">
        <h5 class="modal-title">
            <Input
                type="text"
                placeholder="Recipe Name"
                bind:value={recipeProp.recipe_name}
            />
        </h5>
        <button
            type="button"
            class="btn-close"
            aria-label="Close"
            onclick={closeModal}
        >
        </button>
    </div>
    <div class="modal-body">
        <div>
            <Input
                type={"checkbox"}
                bind:value={isMacroInfo}
                label="Manually input macro nutrients, instead of list of foods"
                onclick={() => (isMacroInfo = !isMacroInfo)}
            />
        </div>
        <div>
            {#if !isMacroInfo}
                <h5>List of Foods</h5>
                <div class="my-2">
                    <Input
                        type="textarea"
                        placeholder="List of Foods"
                        bind:value={recipeProp.food_string}
                    />
                </div>
            {:else}
                <h5>List of Macros</h5>
                <div class="my-2"></div>
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
                        bind:value={recipeProp.serving_size}
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
        <button type="button" class="btn btn-primary" onclick={saveRecipe}
            >Save changes</button
        >
    </div>
</Modal>
