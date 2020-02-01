import App from './App.svelte';
import {
	Title	
} from './lib/const.js';

const app = new App({
	target: document.body,
	props: {
		Title
	}
});

export default app;
