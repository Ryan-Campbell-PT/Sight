<!-- this page will show all active recipes,
let you add recipes,
make modifications to recipes -->
<!--
<script lang="ts">
    import CustomRecipe from "$lib/CustomRecipe.svelte";
    import type {
        NaturalLanguageResponseObject,
        Recipe,
    } from "$lib/NutritionData";
    import { Input, Label, Button } from "@sveltestrap/sveltestrap";

    // aligns with PostRecipe_RequestBody
    class PostRecipeRequestObject {
        recipeName: string = "";
        alternativeRecipeNames: string[] = [];
        foodListString: string = "";
        numServings: number = 1;
    }

    var recipeName = $state("");
    var altNames = $state([]);
    var foodListString = $state("");
    var numServings = $state(1);
    var naturalLanguageResponse = $state() as NaturalLanguageResponseObject;

    let refreshRecipeList = async () => {
        getAllRecipes();
    };

    let post_recipe = async () => {
        const body: PostRecipeRequestObject = {
            recipeName: recipeName,
            alternativeRecipeNames: altNames,
            foodListString: foodListString,
            numServings: numServings,
        };

        try {
            const res = await fetch("http://localhost:8080/postRecipe", {
                method: "GET",
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

    let getAllRecipes = async () => {
        try {
            const res = await fetch("http://localhost:8080/getAllRecipes", {
                method: "GET",
                headers: { "Content-Type": "application/json" },
            });

            if (!res.ok) {
                throw new Error();
            }
        } catch (e) {
            console.log("Error fetching all recipes: " + e);
            throw new Error();
        }
    };
</script>

<div>
    <div id="add-new-recipe" class="my-2">
        <p style="font-size: small;">
            Enter in the recipe ingredients below, along with the serving size.<br
            />
            The recipe will be saved for use later, where you can specify the servings
        </p>
    </div>
    <div>
        <CustomRecipe
            isEditable={false}
            recipe={{
                active: true,
                food_string: "1 banana",
                id: 1,
                nutrition_id: 1,
                recipe_name: "Ding",
                serving_size: 1,
            } as Recipe}
        />
    </div>
</div>
-->
