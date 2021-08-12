<script>

	export let routes;
	let search = "";

	let routeList = routes.slice(0, routes.length)

	$: {
		if (search == "") {
			routeList = routes
		} else {
			routeList = routes.filter(
				(route) => route.Name.toLowerCase().includes(search.toLowerCase())
			)
		}
	}

	function ScrollTo(route) {
		route = route.replace(/\s/g, "-")
		document
			.getElementById(route)
			.scrollIntoView({behavior: 'smooth'})

		let obj = { Title: "Tech With Tim CDN", Url: "#" + route };
		history.pushState(obj, obj.Title, obj.Url);
	}

</script>

<div class="container">
	<div class="inner">
		<input bind:value={search} placeholder="Search" />
		{#if routeList.length != 0}
			{#each routeList as route}
				<table>
					<tr
						on:click={ScrollTo(route.Name)}
					>

						<td style="display: flex; justify-content: center;">
							<span class="route">
								{route.Name}
							</span>
						</td>
						<td style="width: 100%" />
						<td>
							<!-- svelte-ignore a11y-missing-attribute -->
							<img src="assets/arrow.png" width=20>
						</td>

					</tr>
				</table>
			{/each}
		{:else}
			<div class="not-found">No results found</div>
		{/if}
	</div>
</div>

<style>

	.container {
		background-color: #28282d;
		display: block;
		height: 100%;
		position: fixed;

		top: 0;
		left: 0;
		width: 23vw;
	}

	.inner {
		padding-top: 3vw;
		padding-left: 15px;
		padding-right: 15px;
		width: 20vw;
	}

	.route {
		font-family: "Atkinson Hyperlegible";
		font-size: 16px;
		font-weight: 500;
		padding-top: 7px;
		padding-bottom: 7px;
		padding-left: 1px;
		width: 15vw;

		transition: 100ms all;

		color: #AAA;
	}

	tr {
		cursor: pointer;
	}

	input {
		border-radius: 5px;
		/*border: 1px white solid;*/
		background-color: #38383d;
		color: white;
		width: 20vw;
	}

	input:focus {
		outline: none;
	}

	.not-found {
		padding-top: 15px;
		color: #888;
		font-family: Nunito;
		font-size: 16px;
		font-weight: 500;
	}

</style>