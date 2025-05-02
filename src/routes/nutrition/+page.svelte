<script lang="ts">
    import { Input, Label, Button } from "@sveltestrap/sveltestrap";
        import type { Recipe } from "../../lib/NutritionData";
    import NutritionDisplay from  "../../lib/NutritionDisplay.svelte"
    import NutritionLabel from "../../lib/NutritionLabel.svelte";
    import { NutritionResponseObject, RecipeResponseObject } from "../../lib/NutritionData"
    import { onMount } from "svelte";
    import CustomRecipe from "$lib/CustomRecipe.svelte";
    
    function formatDateToYYYYMMDD(date: Date): string {
        const year = date.getFullYear();
        const month = (date.getMonth() + 1).toString().padStart(2, '0');
        const day = date.getDate().toString().padStart(2, '0');
        return `${year}-${month}-${day}`;
    }

    let foodListString_calculate = $state("")
    let foodListString_recipe = $state("")
    let numServings_recipe = $state(1)
    let currentSelectedDate = $state(formatDateToYYYYMMDD(new Date()))
    let nutritionInfoIsVisible = $state(false)
    let nutritionResponse = $state(
        {
            nutritionResponseObject: new NutritionResponseObject(),
            // with the display variable being created, you may be able to get rid of one of the isVisible variables
            display: false
        }
    )
    let test: Recipe[] = $state([])
    let recipeResponse = $state(
        {
            recipeResponseObject: new RecipeResponseObject(),
            display: false
        }
    )
    let showAddNewRecipe = $state(true)
    let showNutritionBreakdown = $state(true)

    let post_foodList = async (saveToDb = false) => {
        const body = {
            foodListString: foodListString_calculate,
            date: currentSelectedDate,
            saveToDb: saveToDb
        }
        try {
            const res = await fetch("http://localhost:8080/postFoodList", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body),
            })

            if(!res.ok) {
                nutritionInfoIsVisible = false
                nutritionResponse.display = false
                throw new Error("Failed to fetch")
            }
            else {
                Object.assign(nutritionResponse.nutritionResponseObject, JSON.parse(await res.json()))
                console.log(nutritionResponse.nutritionResponseObject)
                nutritionResponse.display = true
                nutritionInfoIsVisible = true
            }
        } catch(error) {
            console.error(error)
        }
    }

    let post_recipe = async () => {
        const body = {
            foodListString: foodListString_recipe,
            numServings: numServings_recipe,
        }

        try {
            const res = await fetch("http://localhost:8080/postRecipe", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body),
            })

            if(!res.ok) {
                //TODO display some error
            }
            else {
                //nothing to save, just show something on page to confirm it saved correctly
                //TODO save recipe, add it to recipe list on page
                //OR
                //refresh page to fill information, which should reach out to db to get recipe information
            }
        }
        catch(error) {
                console.log(error)
        }
    }

    let get_recipes = async () => {
        try {
            const res = await fetch("http://localhost:8080/getRecipes", {
                method: "GET",
                headers: { "Content-Type": "application/json" },
            })
            if(!res.ok) {
                throw new Error("error getting recipes")
            }

            Object.assign(test, JSON.parse(await res.json()))
            console.log(test)
            recipeResponse.display = true

        } catch(error) {
            console.log(error)
        }
    }


    onMount(() => {get_recipes()})
</script>


<title>Nutrition Page</title>
<!-- <h2>Nutrition Page</h2> -->
<div class="container">
    <div class="row">
        <!-- this div will contain the food string and date picker -->
        <div id="first-column-half" class="col-md-6 container">
            <div id="food-list" class="my-2">
                <!-- this text string assortment can be turned into a reusable class -->
                <Label for="FoodListString">
                    <p style="font-size: small"> Enter a query like: <b>1 banana, .5 cup of white rice, 1 pound ground beef</b> to get the nutrition information </p>
                    <Input id="FoodListString" type="textarea" placeholder="List of foods, seperated by a comma" bind:value={foodListString_calculate}/>
                </Label>

                <div class="d-flex justify-content-between">
                    <Label for="DatePicker" class="inline">
                        <Input type="date" bind:value={currentSelectedDate} onchange={() => console.log(currentSelectedDate)}/>
                    </Label>
                    <Button onclick={() => post_foodList()}>Visualize</Button>
                    <Button onclick={() => post_foodList(true)}>Visualize and Save</Button>
                </div>
            </div>
            <div id="recipe-options">

            </div>

            <Label for="show-add-new-recipe">
                Add new recipe?
                <Input id="show-add-new-recipe" type="checkbox" bind:checked={showAddNewRecipe}/>
            </Label>

            {#if showAddNewRecipe}
                <div id="add-new-recipe" class="my-2">
                    <p style="font-size: small;">Enter in the recipe ingredients below, along with the serving size.<br/>
                        The recipe will be saved for use later, where you can specify the servings
                    </p>
                    <div class="d-flex justify-content-between">
                        <div>
                            <div class="d-flex my-2">
                                <Label for="new-recipe-serving-input">
                                    Servings
                                </Label>
                                <Input id="new-recipe-serving-input" class="mx-2" style="width: 20%;" type="number" placeholder="Servings" defaultValue=1 min=1 bind:value={numServings_recipe}/>
                            </div>
                            <Label for="new-recipe-input">
                                <Input id="new-recipe-input" type="textarea" placeholder="List of foods, seperated by a comma" bind:value={foodListString_recipe}/>
                            </Label>
                        </div>
                        <div>
                            <Button onclick={post_recipe}>Save Recipe</Button>
                        </div>
                    </div>
                </div>
            {/if}
            <div id="recipe-list">
                <h3> Recipe List </h3>
                    <ul id="recipes">
                        {#each test as recipe}
                            <CustomRecipe recipe={recipe}/>
                        {/each}
                    </ul>
            </div>

        </div>
        <div id="second-column-half" class="col-md-6 container">
            <!-- this div will display the nutrition label
            and maybe the breakdown, depending on space
            breakdown may be better suited in general middle of page -->
            {#if nutritionResponse.display}
                <Label for="displayNutritionBreakdown">
                    Display Nutrition Breakdown?
                    <Input id="displayNutritionBreakdown" type="checkbox" bind:checked={nutritionInfoIsVisible}/>
                </Label>
                {#if nutritionInfoIsVisible}
                    <NutritionDisplay
                        nutritionResponse={nutritionResponse.nutritionResponseObject}
                        nutritionLabelIsVisible={nutritionInfoIsVisible}
                        nutritionBreakdownIsVisible={showNutritionBreakdown}
                    />
                {/if}
            {/if}
        </div>
    </div>
</div>