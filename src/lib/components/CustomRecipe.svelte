<!-- i would ideally like to be able to create new recipes through this component
pass in an empty Recipe, and have a prop that says "isNew" or something
and then when a button press, post it to the db and remove the object from the page -->
<script lang="ts">
    import type { Recipe } from "$lib/NutritionData";

    let { isEditable, recipe }: { isEditable: boolean; recipe: Recipe } =
        $props();

    let isChecked: boolean = $state(false);
    let setIsChecked = () => {
        isChecked = !isChecked;
    };
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
    class="border border-3 border-dark my-4"
    style="width: 100px; height: 100px;"
    onkeydown={setIsChecked}
    onclick={setIsChecked}
>
    <div class="row">
        <img
            alt={recipe.food_string}
            src="data:image/svg+xml;base64,PHN2ZyBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGZpbGwtcnVsZT0iZXZlbm9kZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIgc3Ryb2tlLW1pdGVybGltaXQ9IjIiIHZpZXdCb3g9IjAgMCAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cGF0aCBkPSJtOCAxNmgtNXY0YzAgLjYyMS41MiAxIDEgMWg0em02LjYgNXYtNWgtNS4ydjV6bTYuNC01aC01djVoNGMuNDc4IDAgMS0uMzc5IDEtMXptMC0xLjR2LTUuMmgtNXY1LjJ6bS0xOC01LjJ2NS4yaDV2LTUuMnptMTEuNiAwaC01LjJ2NS4yaDUuMnptMS40LTYuNHY1aDV2LTRjMC0uNDc4LS4zNzktMS0xLTF6bS04IDV2LTVoLTRjLS42MiAwLTEgLjUxOS0xIDF2NHptNi42LTVoLTUuMnY1aDUuMnoiIGZpbGwtcnVsZT0ibm9uemVybyIvPjwvc3ZnPg=="
        />
    </div>
    <div class="row">
        <div class="col-1">
            <input type="checkbox" checked={isChecked} />
        </div>
        <div class="col-5">
            {recipe.recipe_name}
        </div>
    </div>
    <div class="">
        {recipe.food_string}
    </div>
</div>

<!-- old code  -->
<!--
<script lang="ts">
    import { Input, Label, Button } from "@sveltestrap/sveltestrap";
</script>


<div class="d-flex justify-content-between">
    <div>
        <div class="my-2">
            <Label for="new-recipe-name">Recipe Name</Label>
            <Input id="new-recipe-name" type="text" bind:value={recipeName} />
        </div>
        <div class="d-flex my-2">
            <Label for="new-recipe-serving-input">Servings</Label>
            <Input
                id="new-recipe-serving-input"
                class="mx-2"
                style="width: 20%;"
                type="number"
                placeholder="Servings"
                defaultValue="1"
                min="1"
                bind:value={numServings}
            />
        </div>
        <Label for="new-recipe-input">
            <Input
                id="new-recipe-input"
                type="textarea"
                placeholder="List of foods, seperated by a comma"
                bind:value={foodListString}
            />
        </Label>
    </div>
    <div>
        <Button onclick={post_recipe}>Save Recipe</Button>
    </div>
</div>

-->
