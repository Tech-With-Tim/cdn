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
								height="20px"
								class="link"
								alt="link"
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
					{#each route.Description.split("<br>") as descriptionLine}
						<div>
							<span class="key">
								{descriptionLine.split(":")[0]}
							</span>
							
							{#if descriptionLine.split(":")[1]}
								<span class="value">
									:{descriptionLine.split(":")[1]}
								</span>
							{/if}
						</div>
					{/each}
				</div>
			</div>
		{/each}
	</div>

</div>

<style>
	
	.route-link {
		text-decoration: none;
		height: 25px;
		width: 25px;
	}

	.route-link:hover {
		display: block;
		height: 25px;
		width: 25px;
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
		font-weight: normal;
		font-family: "Secular One";
	}

	.key {
		color: white;
		font-size: 16px;
		font-weight: normal;
		font-family: Nunito;
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