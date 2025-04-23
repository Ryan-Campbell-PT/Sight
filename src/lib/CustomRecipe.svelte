<script lang='ts'>
    import type { Recipe } from "./NutritionData";

    let { recipe }: { recipe: Recipe } = $props()

    
    let isChecked = $state(false)

    let check = () => {isChecked = !isChecked}
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="border border-3 border-dark my-4" style="width: 100px; height: 100px;" onkeydown={check} onclick={check}>
    <div class="row">
        <img alt={recipe.food_string} src="data:image/svg+xml;base64,PHN2ZyBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGZpbGwtcnVsZT0iZXZlbm9kZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIgc3Ryb2tlLW1pdGVybGltaXQ9IjIiIHZpZXdCb3g9IjAgMCAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cGF0aCBkPSJtOCAxNmgtNXY0YzAgLjYyMS41MiAxIDEgMWg0em02LjYgNXYtNWgtNS4ydjV6bTYuNC01aC01djVoNGMuNDc4IDAgMS0uMzc5IDEtMXptMC0xLjR2LTUuMmgtNXY1LjJ6bS0xOC01LjJ2NS4yaDV2LTUuMnptMTEuNiAwaC01LjJ2NS4yaDUuMnptMS40LTYuNHY1aDV2LTRjMC0uNDc4LS4zNzktMS0xLTF6bS04IDV2LTVoLTRjLS42MiAwLTEgLjUxOS0xIDF2NHptNi42LTVoLTUuMnY1aDUuMnoiIGZpbGwtcnVsZT0ibm9uemVybyIvPjwvc3ZnPg==">    </div>
    <div class="row">
        <div class="col-1">
            <input type="checkbox" checked={isChecked}/>
        </div>
        <div class="col-5">
            {recipe.recipe_name}
        </div>
    </div>
</div>