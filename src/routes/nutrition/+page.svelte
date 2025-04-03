<script>
    let foodListString = $state("")
    let currentSelectedDate = $state(Date.now)
    let nutritionInfoIsVisible = $state(false)
</script>
<script>
    import { onMount } from "svelte";
    let data = null

    let get = async () => {
        try {
            const res = await fetch("http://localhost:8080/nutritionix", {
                method: "GET",
                headers: { "Content-Type": "application/json" },
            })

            if(!res.ok) throw new Error("Failed to fetch")
            data = await res.json()
        } catch(error) {
            console.error(error)
        }
    }

    let post = async () => {
        try {
            const res = await fetch("http://localhost:8080/postFoodList", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: foodListString,
            })

            if(!res.ok) throw new Error("Failed to fetch")
            data = await res.json()
        } catch(error) {
            console.error(error)
        }
    }

    onMount(post)
</script>

<pre>{JSON.stringify(data, null, 2)}</pre>

<head>
    <title>Nutrition Page</title>
</head>
<body>
    <h2>Nutrition Page</h2>
    <div style="display: flex;">
        <div id="foodStringAndDatePicker" style="display: flex;">
            <!-- this div will contain the food string and ate picker -->
            <div>
                <!-- this text string assortment can be turned into a reusable class -->
                <label for="FoodListString">List of foods</label>
                <input id="FoodListString" type="text" placeholder="List of foods, seperated by a comma" bind:value={foodListString}/>
            </div>
            <div>
                <label for="DatePicker"></label>
                <input type="date" bind:value={currentSelectedDate}/>
            </div>
            <button>Visualize</button>
        </div>
        <div>
            <!-- this div will contain established recipes and images associated with them.
            The actual display of these should be a class -->
            <div style="display: flex;">
                <div id="establishedRecipeClass_recipe1">
                    <img src="" alt="" width="75px" height="75px"/> <!--turn these width/height values into a variable-->
                    <span>recipe1</span>
                </div>
                <div id="establishedRecipeClass_recipe2">
                    <img src="" alt="" width="75px" height="75px"/>
                    <span>recipe2</span>
                </div>
            </div>
            <div id="addNewRecipe">
                <!-- this div will be used to add a new recipe -->
                <!-- clicking this button will set state property `addNewRecipe` to true and display the div below -->
                <label for="isNewRecipeCheckbox">Add New Recipe?</label>
                <input id="isNewRecipeCheckbox" type="checkbox"/>
                <div id="isNewRecipe?">
                    <label for="newRecipeName"></label>
                    <input id="newRecipeName" type="text" placeholder="Name of new recipe">

                    <label for="newRecipeTextString"></label>
                    <input id="newRecipeTextString" type="text" placeholder="New Recipe">

                    <button>Save Recipe</button>
                </div>
            </div>
        </div>
        <div>
            <!-- this div will display the nutrition label
            and maybe the breakdown, depending on space
            breakdown may be better suited in general middle of page -->
            
        </div>
    </div>
</body>