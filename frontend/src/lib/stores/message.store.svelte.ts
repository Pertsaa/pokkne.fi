type Message = {
	id: string;
	loading: boolean;
	sender: 'You' | 'Pokkne';
	message: string;
};

export class MessageStore {
	_messages = $state<Message[]>([]);

	get messages() {
		return this._messages;
	}

	async loadMessages() {
		try {
			const stored = localStorage.getItem('messages');
			if (stored) {
				this._messages = JSON.parse(stored);
			}
		} catch (error) {
			console.error(error);
		}
	}

	async createMessage(message: Message) {
		this._messages.push(message);
		localStorage.setItem('messages', JSON.stringify(this._messages.slice(0, 100)));
	}

	async replaceMessage(message: Message) {
		this._messages = this._messages.map((m) => (m.id === message.id ? message : m));
		localStorage.setItem('messages', JSON.stringify(this._messages.slice(0, 100)));
	}

	async clearMessages() {
		this._messages = [];
		localStorage.setItem('messages', JSON.stringify([]));
	}

	async sendMessage(message: string, mode?: string) {
		this.createMessage({
			id: crypto.randomUUID(),
			sender: 'You',
			message,
			loading: false
		});
		const responseMessageId = crypto.randomUUID();
		try {
			this.createMessage({
				id: responseMessageId,
				sender: 'Pokkne',
				message: '...',
				loading: true
			});
			const res = await fetch('https://api.pokkne.fi/v1/chat', {
				method: 'POST',
				body: JSON.stringify({ prompt: message, mode })
			});
			if (!res.ok) {
				throw Error('api error');
			}
			const data = await res.json();
			this.replaceMessage({
				id: responseMessageId,
				sender: 'Pokkne',
				message: data.message,
				loading: false
			});
		} catch (error) {
			console.error(error);
			this.replaceMessage({
				id: responseMessageId,
				sender: 'Pokkne',
				message: ':(',
				loading: false
			});
		}
	}
}
