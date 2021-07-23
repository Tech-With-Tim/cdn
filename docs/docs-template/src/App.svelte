<script>
	import { fly } from "svelte/transition";
	import RouteList from "./RouteList.svelte";
	import Docs from "./Docs.svelte";
	import Description from "./Description.svelte";
</script>

<main>
	<div transition:fly={{ y: 50, duration: 1500 }}>
		<table>
			<tr>
				<td class="search">
					<RouteList routes={["Hello World", "Create Asset"]} />
				</td>

				<td style="width: 99%;">
					<Docs />
				</td>
				
				<td class="description">
					<Description />
				</td>
			</tr>
		</table>
	</div>
</main>

<style>
	/*.description {
		background-color: #273338;
	}*/

	.search {
		width: 30vw !important;
		min-width: 30vw;
	}

	table, tr, td, main, div {
		height: 80%;
	}
</style>
