<script>
    export let message = "Something went wrong!";
    export let type = "danger"; // "success", "info", "warning", "danger"
    export let duration = 3000; // milliseconds before fade out

    let visible = true;

    // Automatically hide after a few seconds
    const hideAfterDelay = () => {
        setTimeout(() => {
            visible = false;
        }, duration);
    };

    hideAfterDelay();
</script>

{#if visible}
    <div class="alert alert-{type} show fade" role="alert" id="alertBox">
        {message}
    </div>
{/if}

<style>
    #alertBox {
        position: fixed;
        top: 20px;
        right: 20px;
        min-width: 250px;
        z-index: 1050;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
        opacity: 1;
        transition: opacity 0.5s ease-in-out;
    }

    .fade {
        opacity: 1;
        transition: opacity 0.5s ease-in-out;
    }

    :global(.fade:not(.show)) {
        opacity: 0;
    }
</style>
