<script lang="ts">
    import type { CustomRecipe } from "$lib/NutritionData";
    import { Modal } from "@sveltestrap/sveltestrap";

    let { recipe, isVisible }: { recipe: CustomRecipe; isVisible: boolean } =
        $props();

    let recipeEdit: CustomRecipe = recipe;

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
        <h5 class="modal-title">{recipeEdit.recipe_name}</h5>
        <button
            type="button"
            class="btn-close"
            aria-label="Close"
            onclick={() => (isVisible = false)}
        >
        </button>
    </div>
    <div class="modal-body">
        <h5>Food List</h5>
        <div>
            {recipeEdit.food_string}
        </div>
    </div>
    <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal"
            >Close</button
        >
        <button type="button" class="btn btn-primary">Save changes</button>
    </div>
</Modal>
