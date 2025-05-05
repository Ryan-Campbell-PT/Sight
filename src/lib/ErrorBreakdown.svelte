<script lang="ts">
    import { type NutritionErrorObject } from "./NutritionData";

    export let error: NutritionErrorObject;
    export let colSpan: number;
</script>

<tr class="error-breakdown ">
    <td colspan={colSpan} class="bg-danger text-nowrap">
        There was an error with the typed string <b>{error.foodString}</b>
    </td>
</tr>