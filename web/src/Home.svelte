<script>
	import { onMount } from 'svelte';
	import { Router, Link, Route } from "svelte-routing";
	import { navigate } from "svelte-routing";
	import { PATH_LIST } from './lib/const.js';
	import moment from 'moment';

	import Get from './lib/net.js';

	let List = [];

	onMount(async function(){
		let res = await Get(PATH_LIST);
		res = await res.json();

		for(let item of res) {
			item.LastModify = moment(item.LastModify).format("YYYY-MM-DD hh:mm")
		}
		List = res;
	});

	function randomColor(){
		let c = (~~(Math.random()*(1<<24))).toString(16);
		if(c.length < 6) {
			c += 3
		}
		return  "#" + c;
	}

	function ToBlog(){
		const $this = this;
		let id = $this.getAttribute("data-id");
		navigate(`/blog/${id}`,{replace:true});
	}

</script>

<div class="card-content">
	{#each List as item}
		<div class="container card" style="border-left:6px solid {randomColor()};">
			<a on:click={ToBlog} class="card-content columns" data-id={item.Filename}>
				<div class="column">
					<h3 class="subtitle is-4">
						{item.Filename.replace(".md","")}
					</h3>
				</div>
				<div class="column is-6 subtitle">
					{item.LastModify}
				</div>
			</a>
		</div>
	{/each}
</div>

<style>

	a:hover {
		text-decoration: none;
	}

	.container {
		margin: 10px 5px;
		padding: 10px 50px;
		padding-left: 5px;
	}

	@media (min-width: 800px) {
		.container {
			margin: 10px 20vw;
		}
	}
</style>
