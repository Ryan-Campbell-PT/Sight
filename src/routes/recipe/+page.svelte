<!-- this page will show all active recipes,
let you add recipes,
make modifications to recipes -->
<script lang="ts">
    import NaturalLanguageTextBox from "$lib/components/NaturalLanguageTextBox.svelte";
    import Recipe from "$lib/components/Recipe.svelte";
    import type {
        NaturalLanguageResponseObject,
        CustomRecipe,
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
    var all_recipes = $state([] as CustomRecipe[]);
    var inactive_recipes = all_recipes.filter((m) => !m.active);
    var active_recipes = all_recipes.filter((m) => m.active);

    let reset_page = () => {};

    let refresh_recipe_list = async () => {
        get_all_recipes();
        reset_page();
    };

    let save_recipe = async () => {
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
                refresh_recipe_list();
            }
        } catch (error) {
            console.log(error);
        }
    };

    let get_all_recipes = async () => {
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

    let display_error = () => {
        // TODO
        // probs want a boolean parameter
    };
</script>

<div>
    <div>
        <Label class="m-2" for="recipe-name">
            Recipe Name
            <Input id="recipe-name" class="my-1" bind:value={recipeName} />
        </Label>
    </div>
    <div>
        <Label class="m-2" for="recipe-servings">
            Number of Servings
            <Input
                id="recipe-servings"
                class="my-1"
                type="number"
                min="1"
                bind:value={numServings}
            />
        </Label>
    </div>
    <!-- may want to be additive inputs, where you have a button that creates or removes boxes for more/less names -->
    <!-- may also want alt-names to be a drop down or show/hide, since id imagine most recipes wont have alternative names -->
    <div>
        <Label class="m-2" for="recipe-alt-names">
            Alternative Recipe Names
            <Input
                id="recipe-alt-names"
                class="my-1"
                min="1"
                bind:value={altNames}
            />
        </Label>
    </div>
    <div>
        <NaturalLanguageTextBox
            displayCalendar={false}
            primaryButtonText={"Save Recipe"}
            fetchSuccessCallback={save_recipe}
            fetchFailCallback={display_error}
        />
    </div>
    <div>
        <Recipe
            isEditable={false}
            recipe={{
                active: true,
                food_string: "1 banana",
                id: 1,
                nutrition_id: 1,
                recipe_name: "Ding",
                serving_size: 1,
            } as CustomRecipe}
        />
    </div>
</div>
