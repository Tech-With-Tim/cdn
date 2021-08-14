<script>
	import RouteList from "./RouteList.svelte";
	import Docs from "./Docs.svelte";

	let routes = []

	fetch("/docs/docs.json").then(
		(resp) => resp.json()
	).then(
		(data) => {
			routes = data
		}
	)

</script>

<main>
	<table style="height: 100%;">
		<tr style="height: 100%">
			<td class="search">
				<RouteList routes={routes} />
			</td>

			<td class="docs-container">
				<Docs routes={routes}/>
			</td>
		</tr>
	</table>
</main>

<style>
	.search {
		width: 24vw !important;
		min-width: 24vw;
		margin-right: 40px;
	}

	.docs-container {
		width: 100%;
		vertical-align: top;
		height: max-content;
	}

	table, tr, td, main {
		height: 100% !important;
		margin: -1px;
		padding: 0;
	}
</style>
