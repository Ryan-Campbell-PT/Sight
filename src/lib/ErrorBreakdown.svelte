<script lang="ts">
    export let error: NutritionErrorObject;
    export let colSpan: number;
</script>

<tr class="error-breakdown">
    <td colspan={colSpan} class="bg-danger text-nowrap">
        <!-- There was an error with the typed: <b>{error.errorString}</b> -->
    </td>
</tr>
