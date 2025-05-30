<script lang="ts">
    let {
        showError = $bindable(false),
        alertText = "",
    }: { showError: boolean; alertText: string } = $props();
</script>

<div>
    {#if showError}
        <div class="alert">
            <!-- &times; creates an X -->
            <button class="closebtn" onclick={() => (showError = false)}
                >&times;
            </button>
            {alertText}
        </div>
    {/if}
</div>

<style>
    /* The alert message box */
    .alert {
        margin: 20px;
        background-color: #f44336; /* Red */
        color: white;
        margin-bottom: 15px;
    }

    /* The close button */
    .closebtn {
        margin-left: 15px;
        color: white;
        font-weight: bold;
        float: right;
        font-size: 22px;
        line-height: 20px;
        cursor: pointer;
        transition: 0.3s;
    }

    /* When moving the mouse over the close button */
    .closebtn:hover {
        color: black;
    }
</style>
