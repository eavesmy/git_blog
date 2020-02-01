<script>
	  import { onMount } from 'svelte';
	  import { Router, Link, Route } from "svelte-routing";
	  import Home from './Home.svelte';
	  import Login from './Login.svelte';

	  let url = "";
	
	  onMount(async function(){
		console.log(123)
	  });

</script>

<Router url="{url}">
	<Route path="/login" component="{Login}"/>
	<Route path="/" component="{Home}"/>
	<Route path="/">
		<Home />
	</Route>
</Router>

<style></style>
