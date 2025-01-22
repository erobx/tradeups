<script lang="ts">
	import { ModeWatcher } from "mode-watcher";
	import { Navbar } from "$lib/components/custom/navbar/index.js";
	import '../app.css';
	let { children } = $props();
</script>

<ModeWatcher />
<Navbar />
{@render children()}
