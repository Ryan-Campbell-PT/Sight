<script lang="ts">
    import { Button } from "@sveltestrap/sveltestrap";
    import EditRecipeModal from "./EditRecipeModal.svelte";
    import { save_recipe } from "./service/HttpService";

    let modalIsOpen = $state(false);

    let {
        recipe = undefined,
    }: {
        recipe: Recipe | undefined;
    } = $props();

    const color = "#4a90e2";

    // Create initials from the name
    const initials = recipe?.recipe_name
        .split(" ")
        .filter(Boolean)
        .map((word) => word[0].toUpperCase())
        .join("");

    let setModalIsOpen = (isOpen: boolean) => {
        modalIsOpen = isOpen;
    };

    let saveRecipe = async (r: Recipe) => {
        save_recipe(r);
        setModalIsOpen(false);
    };
</script>

<div class="card d-inline-block m-2" style="width: 12rem; height: 12rem;">
    <!-- Top portion -->
    <div
        class="d-flex align-items-center justify-content-center text-white font-weight-bold"
        style="background-color: {color}; height: 70%;"
    >
        <span style="font-size: 2rem;">{initials}</span>
    </div>

    <!-- Bottom portion -->
    <div
        class="card-body position-relative p-2 d-flex align-items-center justify-content-center text-center"
    >
        <span class="text-truncate w-100">{recipe?.recipe_name}</span>

        <Button
            color="light"
            size="sm"
            class="position-absolute"
            style="bottom: 0.5rem; right: 0.5rem; padding: 0.25rem 0.5rem;"
            on:click={() => setModalIsOpen(true)}
        >
            ✎
        </Button>
    </div>
</div>

<EditRecipeModal
    {recipe}
    isOpen={modalIsOpen}
    onCancel={() => setModalIsOpen(false)}
    onSave={saveRecipe}
/>
