<script>
	export let Title;
	export let Slogan;
	
	import { onMount } from 'svelte';
	import { Router, Link, Route, navigate } from "svelte-routing";
	import { beforeUpdate } from "svelte";
	import Home from './Home.svelte'
	import Blog from './Blog.svelte'
	import { PATH_INDEX } from './lib/const.js';
	import Get from './lib/net.js';

	let List = [];
	let url = ""


	onMount(async ()=>{
	
		var _hmt = _hmt || [];

		(function() {
		  var hm = document.createElement("script");
		  hm.src = "https://hm.baidu.com/hm.js?58398e0930b0f096f0003d8daa533dae";
		  hm.type = "text/javascript";
		  var s = document.getElementsByTagName("script")[0];
		  s.parentNode.insertBefore(hm, s);
		})();


		let blog = new URL(location.href).searchParams.get("blog");
		if (!!blog) {
			navigate(`/blog/${blog}`,{repalce: true});
			return
		}
		
		let res = await Get(PATH_INDEX);
		if(!res) return;

		res = await res.text();
		Slogan = res;
	});

</script>

<main>
	<section class="hero is-dark is-small">
		<div class="hero-body">
			<h1 class="title">
				{Title}
			</h1>
			<h3 class="subtitle">
				{Slogan}
			</h3>
		</div>
	</section>
	<Router url={url}>
		<Route path="/" component="{Home}"/>
		<Route path="/blog/:title" component="{Blog}"/>
	</Router>
	<div>
	<footer class="footer">
		<div class="container">
			<a href="https://github.com/eavesmy"> Github </a>
		</div>
	</footer>
</main>

<style>
	.footer {
		padding: 1.5rem;
	}
</style>
