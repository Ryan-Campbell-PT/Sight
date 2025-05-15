<script lang="ts">
    import { Input, Label, Button } from "@sveltestrap/sveltestrap";
    import type {
        Recipe,
        NaturalLanguageResponseObject,
        RecipeResponseObject,
        FoodItem,
    } from "../../lib/NutritionData";
    import NutritionDisplay from "../../lib/NutritionDisplay.svelte";
    import NutritionLabel from "../../lib/NutritionLabel.svelte";
    import { onMount } from "svelte";
    import CustomRecipe from "$lib/CustomRecipe.svelte";
    import { error } from "@sveltejs/kit";
    import FoodListTextBox from "$lib/FoodListTextBox.svelte";

    let foodListString_calculate = $state("");
    let foodListString_recipe = $state("");
    let numServings_recipe = $state(1);
    let nutritionInfoIsVisible = $state(true);
    // let nutritionResponse = $state({
    //     nutritionResponseObject: NutritionResponseObject,
    //     // with the display variable being created, you may be able to get rid of one of the isVisible variables
    //     display: false,
    // });
    let userRecipeList = $state([]);
    let showNutritionBreakdown = $state(true);

    let setNutritionDisplayVisible = (isVisible: boolean) => {
        nutritionInfoIsVisible = isVisible;
        // nutritionResponse.display = isVisible;
        // nutritionInfoIsVisible = isVisible;
    };

    let testNutritionInformationFromChildComponent =
        $state<NaturalLanguageResponseObject>({
            errors: [],
            foods: [],
            totalNutritionInformation: {} as FoodItem,
        });

    let get_recipes = async () => {
        await fetch("http://localhost:8080/getRecipes", {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        })
            .then((res) => res.json())
            .then((data) => {
                var jsonData: RecipeResponseObject = JSON.parse(data);
                Object.assign(userRecipeList, jsonData);
            })
            .catch((err) => {
                throw new Error(err);
            });
    };
    // onMount(() => {get_recipes()})
</script>

<title>Nutrition Page</title>
<!-- <h2>Nutrition Page</h2> -->
<div class="container">
    <div class="row">
        <!-- this div will contain the food string and date picker -->
        <div id="first-column-half" class="col-md-6 container">
            <FoodListTextBox
                displayCalendar={true}
                primaryButtonText={"Visualize"}
                bind:nutritionResponse={
                    testNutritionInformationFromChildComponent
                }
            />
            <input
                type="button"
                onclick={() => {
                    console.log(
                        $state.snapshot(
                            testNutritionInformationFromChildComponent,
                        ),
                    );
                }}
            />
            {#if userRecipeList.length > 0}
                <div id="recipe-list">
                    <h3>Recipe List</h3>
                    <ul id="recipes">
                        {#each userRecipeList as recipe}
                            <CustomRecipe {recipe} />
                        {/each}
                    </ul>
                </div>
            {/if}
        </div>
        <div id="second-column-half" class="col-md-6 container">
            {#if nutritionInfoIsVisible && testNutritionInformationFromChildComponent}
                <NutritionDisplay
                    nutritionResponse={testNutritionInformationFromChildComponent}
                    nutritionLabelIsVisible={nutritionInfoIsVisible}
                    nutritionBreakdownIsVisible={showNutritionBreakdown}
                />
            {/if}

            <!--
            {#if nutritionResponse.display}
                <Label for="displayNutritionBreakdown">
                    Display Nutrition Breakdown?
                    <Input
                        id="displayNutritionBreakdown"
                        type="checkbox"
                        bind:checked={nutritionInfoIsVisible}
                    />
                </Label>
                {#if nutritionInfoIsVisible}
                    <NutritionDisplay
                        nutritionResponse={nutritionResponse.nutritionResponseObject}
                        nutritionLabelIsVisible={nutritionInfoIsVisible}
                        nutritionBreakdownIsVisible={showNutritionBreakdown}
                    />
                {/if}
            {/if}
            -->
        </div>
    </div>
</div>
