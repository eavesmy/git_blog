<script>
	export let Title;
	export let Slogan;
	
	import { beforeUpdate } from "svelte";
	import { Router, Link, Route } from "svelte-routing";
	import { navigate } from "svelte-routing";
	import Home from './Home.svelte'
	import Blog from './Blog.svelte'

	let List = [];
	let url = ""

	beforeUpdate(async function(){
		var urlParams = new URLSearchParams(location.search);
		let title = urlParams.get("blog");
		if(!title) return
		
		navigate(`/blog/${title}`,{replace:true});
	})

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
		<Route path="/blog/:title" component="{Blog}"/>
	</Router>
</main>

<style>
</style>
