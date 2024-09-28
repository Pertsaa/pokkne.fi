<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		align: 'left' | 'right';
		children?: Snippet;
	};

	let { children, align = 'right' }: Props = $props();
</script>

<div
	class="w-fit max-w-[90%] break-all rounded-md bg-neutral-800 px-4 py-2 sm:max-w-[70%] {align ===
	'right'
		? 'ml-auto'
		: 'mr-auto'}"
>
	{#if children}
		{@render children()}
	{/if}
</div>
