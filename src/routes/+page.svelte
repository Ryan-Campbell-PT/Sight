<script lang="ts">
    import { Input, Label, Button } from "@sveltestrap/sveltestrap";
    import { onMount } from "svelte";
    import type {
        NaturalLanguageRequest,
        SaveRecipeRequest,
    } from "../lib/models/RequestModels";

    // state
    let userFoodQuery = $state("");
    let nameOfRecipe = $state("");
    let numberOfRecipeServings = $state("");

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
        const request: SaveRecipeRequest = {
            recipe_id: undefined,
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
            id="userFoodQuery"
            type="textarea"
            placeholder="List of foods or recipes, seperated by comma"
            bind:value={userFoodQuery}
        />

        <Button onclick={() => post_saveRecipe()} />
    </div>
</div>
