<script>
	export let Title;
	export let Slogan;
	
<<<<<<< HEAD
	import { onMount } from 'svelte';
	import { Router, Link, Route, navigate } from "svelte-routing";
=======
	import { beforeUpdate } from "svelte";
	import { Router, Link, Route } from "svelte-routing";
	import { navigate } from "svelte-routing";
>>>>>>> f977ad069801eeb6079049640523206667b7b6e9
	import Home from './Home.svelte'
	import Blog from './Blog.svelte'
	import { PATH_INDEX } from './lib/const.js';
	import Get from './lib/net.js';

	let List = [];
	let url = ""

<<<<<<< HEAD
	onMount(async ()=>{
		let blog = new URL(location.href).searchParams.get("blog");
		if (!!blog) {
			navigate(`/blog/${blog}`);
			return
		}
		
		let res = await Get(PATH_INDEX);
		if(!res) return;
=======
	beforeUpdate(async function(){
		var urlParams = new URLSearchParams(location.search);
		let title = urlParams.get("blog");
		if(!title) return
		
		navigate(`/blog/${title}`,{replace:true});
	})

</script>
>>>>>>> f977ad069801eeb6079049640523206667b7b6e9

		res = await res.text();
		Slogan = res;
	});

</script>

<<<<<<< HEAD
<main>
	<div class="card">
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
	<div>
	<footer class="footer">
		<div class="container">
			<a href="https://github.com/eavesmy"> Github </a>
			<script src="https://v1.cnzz.com/z_stat.php?id=1278631690&amp;show=pic" type="text/javascript"></script>
		</div>
	</footer>
=======
	<Router url={url}>
		<Route path="/" component="{Home}"/>
		<Route path="/blog/:title" component="{Blog}"/>
	</Router>
>>>>>>> f977ad069801eeb6079049640523206667b7b6e9
</main>

<style>
	.footer {
		padding: 1.5rem;
	}
</style>
