<script lang="ts">
	import Icon from './Icon.svelte';

	type Props = {
		form: HTMLFormElement | null;
		onSubmit: (prompt: string) => void;
	};

	const handleSubmit = (e: SubmitEvent) => {
		e.preventDefault();
		const data = new FormData(e.currentTarget as HTMLFormElement);
		const prompt = data.get('prompt') as string;
		onSubmit(prompt);
	};

	const handleEnter = (e: KeyboardEvent) => {
		if (e.key === 'Enter' && form) {
			e.preventDefault();
			const data = new FormData(form);
			const prompt = data.get('prompt') as string;
			onSubmit(prompt);
		}
	};

	let { form = $bindable(), onSubmit }: Props = $props();
</script>

<form bind:this={form} onsubmit={handleSubmit}>
	<div
		class="relative flex items-center justify-between overflow-hidden rounded-md py-1.5 ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-neutral-600"
	>
		<div class="w-full">
			<label for="comment" class="sr-only">Add your comment</label>
			<textarea
				rows="1"
				name="prompt"
				id="prompt"
				class="block w-full resize-none border-0 bg-transparent placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6"
				placeholder="Add your comment..."
				onkeydown={handleEnter}
			></textarea>
		</div>

		<button
			type="submit"
			class="mr-3 inline-flex items-center rounded-md bg-neutral-50 p-1 text-sm font-semibold text-neutral-900 shadow-sm hover:bg-neutral-200 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-neutral-600"
		>
			<Icon icon="paper" class="size-6 text-neutral-800" />
		</button>
	</div>
</form>
