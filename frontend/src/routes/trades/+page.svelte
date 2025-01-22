<script lang="ts">
	import { type PageData } from "./$types";

	let { data }: { data: PageData } = $props();
	let test = $state(data.test);
</script>

<div class="container flex-col">
	<div>Active Tradeups</div>
	<div>
		<div>{test}</div>
	</div>
</div>
