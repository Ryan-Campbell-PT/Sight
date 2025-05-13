<script>
    let { children } = $props();
</script>

<nav>
    <ul class="nav-links">
        <li><a href="/">Home</a></li>
        <li><a href="/nutrition">Nutrition</a></li>
        <li><a href="/recipe">Recipe</a></li>
    </ul>
</nav>

{@render children()}

<style>
    .nav-links {
        list-style-type: none;
        margin: 0;
        padding: 0;
        text-align: center;
        background-color: tan;
    }

    .nav-links li {
        display: inline-block;
        margin-left: 20px;
    }

    .nav-links li a {
        color: black;
        text-decoration: none;
        padding: 10px 15px;
        display: inline-block;
        transition: background-color 0.3s;
    }

    .nav-links li a:hover {
        background-color: #555;
    }
</style>
