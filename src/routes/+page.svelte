<script lang="ts">
    import { getAllContexts, onMount } from "svelte";
    import type {
        NaturalLanguageRequest,
        PostRecipeRequest,
    } from "../lib/models/RequestModels";
    import type { GetActiveRecipes } from "../lib/models/ResponseModels";
    import { json } from "@sveltejs/kit";
    import EditRecipeModal from "$lib/EditRecipeModal.svelte";
    import {
        Modal,
        ModalHeader,
        ModalBody,
        ModalFooter,
        Button,
        Form,
        FormGroup,
        Label,
        Input,
    } from "@sveltestrap/sveltestrap";
    import EditRecipeBox from "$lib/EditRecipeBox.svelte";
    import { bootstrap } from "$lib/bootstrapClasses";

    // state
    let selectedRecipeId: number = $state(-1);
    let userFoodQuery: string = $state("");
    let nameOfRecipe: string = $state("");
    let numberOfRecipeServings: string = $state("");

    let activeRecipeList: Recipe[] = $state([] as Recipe[]);

    let modalIsOpen: boolean = $state(false);

    // props
    let {}: {} = $props();

    /*
    let post_UserFoodQuery = async () => {
        // basic endpoint, just returns response from NIX
        const request: NaturalLanguageRequest = {
            user_food_query: userFoodQuery,
        };

        var res = await fetch("http://localhost:8080/NaturalLanguageRequest", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(request),
        });
        if (res.ok) {
            const response = (await res.json()) as NaturalLanguageResponse;
            // const data = JSON.parse(response)
        }
    };
*/
    let post_saveRecipe = async () => {
        /*
        const request: PostRecipeRequest = {
            recipe_id: selectedRecipeId,
            recipe_name: nameOfRecipe,
            recipe_servings: numberOfRecipeServings,
            user_food_query: userFoodQuery,
        };

        const res = await fetch("http://localhost:8080/post_recipe", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(request),
        });

        if (res.ok) {
            console.log("Successful recipe save");
        }
        */
    };

    let get_activeRecipes = async () => {
        const res = await fetch("http://localhost:8080/get_active_recipes", {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        });

        if (res.ok) {
            const response = (await res.json()) as GetActiveRecipes;
            activeRecipeList = response.recipe_list;
            console.log(activeRecipeList);
        }
    };

    onMount(() => get_activeRecipes());

    // modal
    let setModalIsOpen = (isOpen: boolean) => {
        modalIsOpen = isOpen;
    };
</script>

<div>
    <div id="recipe">
        <Input
            id="recipeName"
            type="text"
            placeholder="Recipe Name"
            bind:value={nameOfRecipe}
        />

        <Input
            id="numServings"
            type="number"
            placeholder="Number of Servings"
            bind:value={numberOfRecipeServings}
        />

        <Input
            id="userFoodQuery"
            type="textarea"
            placeholder="List of foods or recipes, seperated by comma"
            bind:value={userFoodQuery}
        />

        <Button onclick={() => post_saveRecipe()}>Save</Button>
    </div>
    <div id="activeRecipes" class="d-flex flex-wrap gap-4">
        {#each activeRecipeList as r}
            <EditRecipeBox onClick={() => {}} recipe={r} />
        {/each}
    </div>
    <div>
        <EditRecipeBox
            onClick={() => {}}
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
