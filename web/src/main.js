import App from './App.svelte';

import { Title,Slogan } from './lib/const.js';

const app = new App({
	target: document.body,
	props: {
		Title, Slogan
	}
});

export default app;
