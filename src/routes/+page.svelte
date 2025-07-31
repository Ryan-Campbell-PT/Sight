<script lang="ts">
    import { onMount } from "svelte";

    let ding = async () => {
        var res = await fetch("http://localhost:8080/", {});
        if (res.ok) {
            console.log(await res.json());
        }
    };

    onMount(() => ding());
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the
    documentation
</p>
