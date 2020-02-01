import { Host } from './const.js';

export default function Get(path){
	return fetch(Host + path);
}
