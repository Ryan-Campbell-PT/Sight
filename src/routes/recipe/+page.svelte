<!-- this page will show all active recipes,
let you add recipes,
make modifications to recipes -->
<script lang="ts">
    import { Input, Label, Button } from "@sveltestrap/sveltestrap";

    // aligns with PostRecipe_RequestBody
    class RecipeRequestObject {
        recipeName: string = "";
        foodListString: string = "";
        numServings: number = 1;
    }

    var recipeName = $state("");
    var foodListString = $state("");
    var numServings = $state(1);

    let post_recipe = async () => {
        const body: RecipeRequestObject = {
            recipeName: recipeName,
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
            }
        } catch (error) {
            console.log(error);
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
        <div class="d-flex justify-content-between">
            <div>
                <div class="my-2">
                    <Label for="new-recipe-name">Recipe Name</Label>
                    <Input
                        id="new-recipe-name"
                        type="text"
                        bind:value={recipeName}
                    />
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
    </div>
</div>
