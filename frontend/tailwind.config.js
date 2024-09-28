/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				'neutral-925': '#111111'
			}
		}
	},
	plugins: [require('@tailwindcss/forms')]
};
