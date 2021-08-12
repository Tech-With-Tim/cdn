<script>
	export let routes;
</script>

<div class="container">
	
	<table class="title-container"><tr style="width: 100%">
		
		<td><img src="assets/favicon.png" alt="TWT Logo" class="logo"></td>
		
		<td style="width: 100%">
			<div class="heading">
				Tech With Tim <span class="subtitle">CDN Documentation</span>
			</div>
		</td>
	
	</tr></table>

	<div class="docs-container">
		{#each routes as route}
			<div style="margin-bottom: 1.5vw;">

				<table style="width: 100%;"><tr style="width: 100%;">
					<td>
						<a
							href={"#" + route.Name.replace(/\s/g, "-")}
							class="route-link"
						>
							<img
								src="assets/link.png"
								class="link"
								alt="link"
								id={route.Name.replace(/\s/g, "-") + "-img"}
							>
						</a>
					</td>

					<td style="width: 100%;">
						<div
							class="route-heading"
							id={route.Name.replace(/\s/g, "-")}
						>
							{ route.Name } <span>{ route.Route }</span>
						</div>
					</td>

				</tr></table>

				<div class="desc">
					<div
					style="padding-bottom: 10px; color: white; font-size: 16px"
					>
						{route.Description}
					</div>

					<span class="key">
						URL Parameters:
					</span>
					<span class="value">
						{route.URLParameters}
					</span>

					<br>

					<span class="key">
						Response:
					</span>
					<span class="value">
						{route.Response}
					</span>

					<br>

					<span class="key" style="margin-top: 30px;">
						Request Body:
					</span>

					<br>

					{#if route.RequestBody}
					<table class="req-body">
						<tr>
							<th style="text-align: left; min-width: 150px;">
								Field
							</th>
							<th style="text-align: left; min-width: 150px;">
								Type
							</th>
							<th style="text-align: left;">
								Description
							</th>
						</tr>
						{#each route.RequestBody as reqbody}
							<tr>
								<td>{reqbody.Name}</td>
								<td>{reqbody.Type}</td>
								<td>{reqbody.Desc}</td>
							</tr>
						{/each}
					</table>
					{/if}

				</div>
			</div>
		{/each}
	</div>

</div>

<style>

	.req-body {
		width: 100%;
		border-radius: 5px;
		padding: 10px;
		margin: 10px;
		margin-top: 10px;
		margin-left: 0px;
		background-color: rgba(0, 0, 0, 0.3);

		font-size: 16px;

		border-collapse: collapse;
	}

	.req-body tr:nth-child(even) {
		background-color: rgba(255, 255, 255, 0.1);
	}

	.req-body td, th {
		border: 2px solid rgba(0, 0, 0, 0.5);
		padding: 8px;
	}
	
	.route-link img {
		text-decoration: none;
		height: 20px;
		width: 20px;
	}

	.route-link img:hover {
		display: block;
		height: 20px;
		width: 20px;
	}

	.title-container {
		display: inline;
		width: 100%;
	}

	.container {
		margin-top: 3vw;
		align-items: left;
		padding-right: 10vw;
	}

	.route-heading {
		font-family: Nunito;
		font-size: 21px;
		font-weight: normal;
		margin-bottom: 10px;
		margin-right: 0px;

		width: 100%;

		color: white;
		border-bottom: 1px white solid;

		padding-top: 4px;
		padding-bottom: 4px;
	}

	span {
		color: #999999;
		font-size: 16px;
		font-family: "Secular One";
	}

	.key {
		color: white;
		font-size: 16px;
		font-weight: 600;
		font-family: Nunito;
		margin-bottom: 5px;
	}

	.value {
		color: #999;
		font-size: 16px;
		font-weight: normal;
		font-family: Nunito;
	}

	.desc {
		font-size: 17px;
		/*font-family: "Atkinson Hyperlegible";*/
		font-family: "Nunito";
		font-weight: 400;
		line-height: 1.4;
		color: white;

		padding-top: 5px;
		margin-bottom: 30px;

		margin-left: 25px;
	}

	.heading {
		border-bottom: 1px #AAAAAA solid;
		font-family: Nunito;
		font-size: 30px;
		color: white;
	}

	.subtitle {
		font-family: Nunito;
		font-size: 24px;
		color: #7D7D7D;
	}

	.logo {
		width: 3.4vw;
	}

	.docs-container {
		margin-top: 30px;
	}

</style>