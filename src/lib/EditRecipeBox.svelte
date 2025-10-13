<script lang="ts">
    import { Button } from "@sveltestrap/sveltestrap";
    import EditRecipeModal from "./EditRecipeModal.svelte";
    import type { PostRecipeRequest } from "./models/RequestModels";
    import type { SaveRecipeResponse } from "./models/ResponseModels";

    // import { Pencil } from 'lucide-svelte'; // nice lightweight icon set

    let modalIsOpen = $state(false);

    let {
        recipe = undefined,
        onClick = () => {},
    }: {
        recipe: Recipe | undefined;
        onClick: () => void;
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
        if (!r) return;

        const request: PostRecipeRequest = {
            recipe_name: r.recipe_name,
            recipe_servings: r.serving_size,
            user_food_query: r.food_string,
            recipe_id: r.id,
        };

        const res = await fetch("http://localhost:8080/post_recipe", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(request),
        });

        if (res.ok) {
            const response: SaveRecipeResponse = await res.json();
            if (response && response.success) {
                // display some success alert
                console.log("Save Recipe success");
            } else {
                //display some error
                console.log("Save recipe errror");
            }
        }
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
    isOpen={modalIsOpen}
    onCancel={() => setModalIsOpen(false)}
    onSave={saveRecipe}
    bind:recipe
/>

<style>
    .truncate {
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
</style>
