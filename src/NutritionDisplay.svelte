<!-- nutrition display will contain the Label and all other associated information like the breakdown or edit functionality -->
<script lang="ts">
    let { nutritionInfo } = $props()
</script>

{#if nutritionInfo != null}
    <div>
        I am displaying information
    </div>
{/if}
