<script>
	import { onMount } from 'svelte';
	import { fly } from 'svelte/transition';
	import { Router, Link, Route } from "svelte-routing";
	import { navigate } from "svelte-routing";
	import { PATH_LIST } from './lib/const.js';
	import moment from 'moment';

	import Get from './lib/net.js';
	
	let Top = [];
	let List = [];
	
	let list = [];
	let page = 1;
	let count = 10;

	onMount(async function(){
		let res = await Get(PATH_LIST);
		res = await res.json();

		for(let item of res) {

			item.LastModify = moment(item.LastModify).format("YYYY-MM-DD hh:mm")
			item.ShowName = item.Filename.replace("top_","");

			if(item.Filename.indexOf("top_") > -1) {
				item.Top = true;
				Top.push(item);
			} else {
				List.push(item);
			}
		}
		
		Top.sort((a,b) => {
			return a.LastModify < b.LastModify;
		});

		List.sort((a,b)=>{
			return a.LastModify < b.LastModify;
		});

		for (let item of Top) {
			List.unshift(item);
		}

		List = List;
		list = List;

		 // GetPage();
	});

	function GetPage(){
		page = 1;
		let index = count * (page - 1);
		let end = index + count;

		list = List.slice(index,end)
	}

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
		navigate(`/blog/${id}`);
	}

</script>

<div class="card-content">
	{#each list as item}
		<div class="container card" style="border-left:6px solid {randomColor()};" transition:fly>
			<a on:click={ToBlog} class="card-content columns" data-id={item.Filename}>
				<div class="column">
					{#if item.Top}
					<span class="tag is-success">Mark</span>
					{/if}
					<h3 class="subtitle is-4">
						{item.ShowName.replace(".md","")}
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
