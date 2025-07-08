<script lang="ts">
    import { Input, Label, Button } from "@sveltestrap/sveltestrap";
    import type {
        CustomRecipe,
        NaturalLanguageResponseObject,
        FoodItem,
    } from "../../lib/NutritionData";
    import NutritionDisplay from "$lib/components/NutritionDisplay.svelte";
    import NaturalLanguageTextBox from "$lib/components/NaturalLanguageTextBox.svelte";
    import Loading from "$lib/components/LoadingModal.svelte";
    import AlertBox from "$lib/components/AlertBox.svelte";
    import RecipeModal from "$lib/components/RecipeModal.svelte";

    let isLoading = $state(false);
    let nutritionDisplayIsVisible = $state(false);
    let displayError = $state(false);
    let newRecipeIsVisible = $state(false);

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

    // onMount(() => {get_recipes()})
</script>

<title>Nutrition Page</title>
<!-- <h2>Nutrition Page</h2> -->
<div class="container">
    <Loading showModal={isLoading} />
    <AlertBox
        bind:showError={displayError}
        alertText="There was an error getting nutrition information"
    />

    <div class="row">
        <!-- this div will contain the food string and date picker -->
        <div id="first-column-half" class="col-md-6 container">
            <NaturalLanguageTextBox
                displayCalendar={true}
                bind:isLoading
                primaryButtonText={"Visualize"}
                bind:nutritionResponse={
                    testNutritionInformationFromChildComponent
                }
                fetchFailCallback={() => {
                    displayError = true;
                    setNutritionDisplayVisible(false);
                }}
                fetchSuccessCallback={() => {
                    displayError = true;
                    setNutritionDisplayVisible(true);
                }}
            />

            <Button onclick={() => (newRecipeIsVisible = true)}>
                Create New Recipe
            </Button>
            <RecipeModal
                recipeProp={{} as CustomRecipe}
                bind:isVisible={newRecipeIsVisible}
            />
        </div>
        <div id="second-column-half" class="col-md-6 container">
            <NutritionDisplay
                nutritionResponse={testNutritionInformationFromChildComponent}
                nutritionLabelIsVisible={nutritionDisplayIsVisible}
                nutritionBreakdownIsVisible={showNutritionBreakdown}
            />
        </div>
    </div>
</div>
