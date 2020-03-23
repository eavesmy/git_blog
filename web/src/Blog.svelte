<script>
	import { onMount, afterUpdate } from 'svelte';
	import { PATH_BLOG } from './lib/const.js';
	import { navigate } from "svelte-routing";
	import Get from './lib/net.js';
	import marked from 'marked';

	export let title;
	let content = "";
	
	onMount(async function(){
		let res = await Get(PATH_BLOG + "/" + title);
		if(res.status !== 200) {
			navigate("/");
			return 
		}
		res = await res.text();
		content = marked(res);
	})

</script>
<div id="content" class="card-content">
	<div class="container">
		{@html content}
	</div>
</div>
<style>

	@media (min-width: 800px) {
		.container {
			margin: 10px 20vw;
		}
	}
</style>
