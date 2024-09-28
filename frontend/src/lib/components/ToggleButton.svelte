<script lang="ts">
	import type { HTMLButtonAttributes } from 'svelte/elements';

	type Props = { checked?: boolean } & HTMLButtonAttributes;

	const { checked, ...rest }: Props = $props();

	const buttonStyles = $derived(checked ? 'bg-cyan-600' : 'bg-neutral-800');
	const spanStyles = $derived(checked ? 'translate-x-5' : 'translate-x-0');
</script>

<button
	type="button"
	role="switch"
	aria-checked={checked}
	{...rest}
	class="relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-neutral-600 {buttonStyles}"
>
	<span class="sr-only">Use setting</span>
	<span
		aria-hidden="true"
		class="pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out {spanStyles}"
	></span>
</button>
