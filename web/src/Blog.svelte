<script>
	import { onMount, afterUpdate } from 'svelte';
	import { PATH_BLOG } from './lib/const.js';
	import Get from './lib/net.js';
	import marked from 'marked';

	export let title;
	let content = "";
	
	onMount(async function(){
		let res = await Get(PATH_BLOG + "/" + title);
		res = await res.text();
		content = marked(res);
		let area = document.querySelector("#content")
		area.setAttribute("contenteditable","false")
	})

</script>
<div id="content" class="card-content" contenteditable="true" bind:innerHTML={content}/>
<style>
</style>
