<script>
	export let Title;
	export let Slogan;
	
	import { Router, Link, Route } from "svelte-routing";
	import Home from './Home.svelte'
	import Blog from './Blog.svelte'

	let List = [];
	let url = ""

</script>


<main class="card">
	<div class="card-header">
		<section class="hero">
			<div class="hero-body">
				<h1 class="title">
					{Title}
				</h1>
				<h3 class="subtitle">
					{Slogan}
				</h3>
			</div>
		</section>
	</div>

	<Router url={url}>
		<Route path="/" component="{Home}"/>
		<Route path="#:title" component="{Blog}"/>
	</Router>
</main>

<style>
</style>
