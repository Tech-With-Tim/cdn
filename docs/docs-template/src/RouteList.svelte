<script>

	export let routes;
	let search = "";

	let routeList = routes.slice(0, routes.length)

	$: {
		if (search == "") {
			routeList = routes
		} else {
			routeList = routes.filter(
				(route) => route.toLowerCase().includes(search.toLowerCase())
			)
		}
	}

	function ScrollTo(route) {
		route = route.replace(/\s/g, "-")
		console.log(route)
		let routeDiv = document.getElementById(route)
		routeDiv.scrollIntoView({behavior: 'smooth'})
	}

</script>

<div class="container">
	<div class="inner">
		<input bind:value={search} />
		{#if routeList.length != 0}
			{#each routeList as route}
				<div class="route" on:click={ScrollTo(route)}>
					{route}
				</div>
			{/each}
		{:else}
			<div class="not-found">No results found</div>
		{/if}
	</div>
</div>

<style>
	.container {
		display: block;
		height: 100%;
	}

	.inner {
		padding-top: 4vw;
		padding-left: 20px;
	}

	.route {
		font-family: Cabin;
		font-size: 16px;
		font-weight: 500;
		padding-top: 15px;
		padding-left: 1px;
		cursor: pointer;
	}

	input {
		border-radius: 5px;
		border: 2px black solid;
	}

	.not-found {
		padding-top: 15px;
		color: #888;
		font-family: Cabin;
	}

</style>