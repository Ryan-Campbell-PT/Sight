<script lang="ts">
    let {
        error = {} as AnalysisErrorObject,
        colSpan = 0,
    }: {
        error: AnalysisErrorObject;
        colSpan: number;
    } = $props();
</script>

<tr class="error-breakdown">
    <td colspan={colSpan} class="bg-danger text-nowrap">
        There was an error with the typed: <b>{error.original_string}</b>
    </td>
</tr>
