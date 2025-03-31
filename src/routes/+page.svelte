<script>
    import { onMount } from "svelte";

    let data = null;

    onMount(async () => {
        try {
            const res = await fetch("http://localhost:8080/", {
                method: "GET",
                headers: { "Content-Type": "application/json" },
            });

            if (!res.ok) throw new Error("Failed to fetch");

            data = await res.json();
        } catch (error) {
            console.error(error);
        }
    });
</script>

<pre>{JSON.stringify(data, null, 2)}</pre>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
