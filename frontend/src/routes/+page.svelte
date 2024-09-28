<script lang="ts">
	import { browser } from '$app/environment';
	import ChatBubble from '$lib/components/ChatBubble.svelte';
	import Icon from '$lib/components/Icon.svelte';
	import LoadingDots from '$lib/components/LoadingDots.svelte';
	import Textarea from '$lib/components/Textarea.svelte';
	import ToggleButton from '$lib/components/ToggleButton.svelte';
	import { MessageStore } from '$lib/stores/message.store.svelte';

	const messageStore = $derived.by(() => new MessageStore());

	let mode = $state('standard');
	let form: HTMLFormElement | null = $state(null);
	let chat: HTMLDivElement | null = $state(null);

	$effect(() => {
		if (browser) {
			messageStore.loadMessages();
		}
	});

	$effect(() => {
		if (messageStore.messages.length > 0) {
			chat?.scrollTo(0, chat?.scrollHeight ?? 0);
		}
	});

	const handleSubmit = async (prompt: string) => {
		messageStore.sendMessage(prompt, mode);
		form?.reset();
	};

	const handleModeToggle = () => {
		if (mode === 'standard') {
			mode = 'randomized';
		} else {
			mode = 'standard';
		}
	};
</script>

<div class="relative mx-auto flex h-full flex-col overflow-hidden pb-8 sm:max-w-[600px]">
	<div class="flex items-center justify-between px-8">
		<h1 class="py-6 text-2xl font-semibold">Pokkne Chat</h1>
		<div class="flex items-center gap-4">
			<ToggleButton checked={mode === 'randomized'} onclick={handleModeToggle} />
			<button
				type="submit"
				onclick={() => messageStore.clearMessages()}
				class="inline-flex items-center rounded-md bg-neutral-50 p-1 text-sm font-semibold text-neutral-900 shadow-sm hover:bg-neutral-200 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-neutral-600"
			>
				<Icon icon="trash" class="size-6 text-gray-800" />
			</button>
		</div>
	</div>

	<div bind:this={chat} class="flex h-full max-w-full flex-col gap-4 overflow-y-auto px-8 py-6">
		{#each messageStore.messages as message, i (i)}
			<ChatBubble align={message.sender === 'You' ? 'right' : 'left'}>
				{message.sender}: {#if message.loading}<LoadingDots />{:else}{message.message}{/if}
			</ChatBubble>
		{/each}
	</div>

	<div class="px-8">
		<Textarea bind:form onSubmit={handleSubmit} />
	</div>
</div>
