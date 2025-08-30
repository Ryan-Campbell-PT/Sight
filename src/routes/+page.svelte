<script lang="ts">
    import { onMount } from "svelte";

    let ding = async () => {
        const body = {
            userFoodQuery: "1 banana, 2 cookies",
        };
        var res = await fetch("http://localhost:8080/userFoodQuery", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body),
        });
        if (res.ok) {
            console.log(await res.json());
        }
    };

    let dong = async () => {
        var res = await fetch("http://localhost:8080/coolStuff", {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        });
        if (res.ok) {
            console.log(await res.json());
        }
    };

    onMount(() => dong());
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the
    documentation
</p>
