<script lang="ts">
    let { showModal = $bindable(false) }: { showModal: boolean } = $props();

    let dialog = $state() as HTMLDialogElement;

    $effect(() => {
        if (showModal) dialog.showModal();
    });
</script>

<div>
    {#if showModal}
        <dialog class="loader" bind:this={dialog}></dialog>
    {/if}
</div>

<style>
    .loader {
        border: 12px solid #f3f3f3;
        border-top: 12px solid tan;
        border-bottom: 12px solid tan;
        border-radius: 50%;
        width: 80px;
        height: 80px;
        animation: spin 2s linear infinite;
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }
</style>
