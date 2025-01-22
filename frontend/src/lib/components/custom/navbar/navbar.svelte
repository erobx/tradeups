<script lang="ts">
	import Sun from "lucide-svelte/icons/sun";
	import Moon from "lucide-svelte/icons/moon";

	import { toggleMode } from "mode-watcher";
	import { Button } from "$lib/components/ui/button/index.js";
</script>

<header
	class="border-border/40 bg-background/95 sticky top-0 z-50 w-full border-b backdrop-blur"
>
	<div class="container flex h-14 max-w-screen-2xl items-center">
		<div class="mr-4 hidden md:flex">
			<nav class="flex items-center gap-6 text-sm">
				<a
					class="hover:text-foreground/80 transition-colors text-foreground/60"
					href="/">Home</a
				>
				<a
					class="hover:text-foreground/80 transition-colors text-foreground/60"
					href="/trades">Trades</a
				>
			</nav>
		</div>
		<div
			class="flex flex-1 items-center justify-between space-x-2 md:justify-end"
		>
			<nav class="flex items-center">
				<Button
					onclick={toggleMode}
					variant="outline"
					size="icon"
				>
					<Sun
						class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
					/>
					<Moon
						class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
					/>
					<span class="sr-only">Toggle theme</span
					>
				</Button>
			</nav>
		</div>
	</div>
</header>
