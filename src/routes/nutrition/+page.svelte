<script lang="ts">
    import { Input, Label, Button } from "@sveltestrap/sveltestrap";
    import type {
        Recipe,
        NaturalLanguageResponseObject,
        RecipeResponseObject,
        FoodItem,
    } from "../../lib/NutritionData";
    import NutritionDisplay from "$lib/components/NutritionDisplay.svelte";
    import NutritionLabel from "$lib/components/NutritionLabel.svelte";
    import { onMount } from "svelte";
    import CustomRecipe from "$lib/components/CustomRecipe.svelte";
    import { error } from "@sveltejs/kit";
    import FoodListTextBox from "$lib/components/NaturalLanguageTextBox.svelte";

    let nutritionDisplayIsVisible = $state(false);
    // let nutritionResponse = $state({
    //     nutritionResponseObject: NutritionResponseObject,
    //     // with the display variable being created, you may be able to get rid of one of the isVisible variables
    //     display: false,
    // });
    let userRecipeList = $state([]);
    let showNutritionBreakdown = $state(false);

    let setNutritionDisplayVisible = (isVisible: boolean) => {
        nutritionDisplayIsVisible = isVisible;
        showNutritionBreakdown = isVisible;
    };

    let testNutritionInformationFromChildComponent =
        $state<NaturalLanguageResponseObject>({
            errors: [],
            foods: [],
            total_nutrition_information: {} as FoodItem,
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
                fetchFailCallback={() => setNutritionDisplayVisible(false)}
                fetchSuccessCallback={() => setNutritionDisplayVisible(true)}
            />
        </div>
        <div id="second-column-half" class="col-md-6 container">
            <NutritionDisplay
                nutritionResponse={testNutritionInformationFromChildComponent}
                nutritionLabelIsVisible={nutritionDisplayIsVisible}
                nutritionBreakdownIsVisible={showNutritionBreakdown}
            />

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
