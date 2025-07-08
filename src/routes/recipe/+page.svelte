<!-- this page will show all active recipes,
let you add recipes,
make modifications to recipes -->
<script lang="ts">
    import LoadingModal from "$lib/components/LoadingModal.svelte";
    import NaturalLanguageTextBox from "$lib/components/NaturalLanguageTextBox.svelte";
    import Recipe from "$lib/components/Recipe.svelte";
    import type {
        NaturalLanguageResponseObject,
        GetUserRecipesResponseObject,
        CustomRecipe,
    } from "$lib/NutritionData";
    import { json } from "@sveltejs/kit";
    import { Input, Label, Button } from "@sveltestrap/sveltestrap";
    import { onMount } from "svelte";

    var isLoading = $state(false);
    var recipeName = $state("");
    var altNames = $state([]);
    var foodListString = $state("");
    var numServings = $state(1);
    var userRecipeResponseObject = $state({}) as GetUserRecipesResponseObject;
    var inactive_recipes = $state([] as CustomRecipe[]);
    var active_recipes = $state([] as CustomRecipe[]);

    let resetPage = () => {
        if (userRecipeResponseObject.recipeList) {
            Object.assign(
                inactive_recipes,
                userRecipeResponseObject.recipeList.filter((m) => !m.active),
            );
            Object.assign(
                active_recipes,
                userRecipeResponseObject.recipeList.filter((m) => m.active),
            );
        }
    };

    let refreshRecipeList = async () => {
        getAllRecipes();
        resetPage();
    };
    /*
    let saveRecipe = async () => {
        const body: PostRecipeRequestObject = {
            recipeName: recipeName,
            alternativeRecipeNames: altNames,
            foodListString: foodListString,
            numServings: numServings,
        };

        try {
            const res = await fetch("http://localhost:8080/postRecipe", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body),
            });

            if (!res.ok) {
                //TODO display some error
            } else {
                //nothing to save, just show something on page to confirm it saved correctly
                //TODO save recipe, add it to recipe list on page
                //OR
                //refresh page to fill information, which should reach out to db to get recipe information
                refreshRecipeList();
            }
        } catch (error) {
            console.log(error);
        }
    };
*/

    let getAllRecipes = async () => {
        try {
            const res = await fetch("http://localhost:8080/getUserRecipes", {
                method: "GET",
                headers: { "Content-Type": "application/json" },
            });

            if (!res.ok) {
                throw new Error();
            }

            Object.assign(
                userRecipeResponseObject,
                JSON.parse(await res.json()),
            );
            resetPage();
        } catch (e) {
            console.log("Error fetching all recipes: " + e);
            throw new Error();
        }
    };

    let display_error = () => {
        // TODO
        // probs want a boolean parameter
    };

    $effect(() => {
        refreshRecipeList();
    });

    const recipeSections = [
        { title: "Active Recipes", recipes: active_recipes },
        { title: "Inactive Recipes", recipes: inactive_recipes },
    ];
</script>

<div class="container">
    <LoadingModal showModal={isLoading} />
    <div>
        {#each recipeSections as section}
            <div class="m-3">
                <h3>{section.title}</h3>
                <div class="row d-flex my-2">
                    {#each section.recipes as r}
                        <Recipe recipe={r} isEditable={false} />
                    {/each}
                </div>
            </div>
        {/each}
    </div>
</div>
