import { navigate } from "svelte-routing";

export function Post(url,data) {
	return fetch(url,{
		method: "POST",
		headers: {
			"Content-Type": "application/json",
			"x-access-token": "eavesmy"
		},
		body: JSON.stringify(data),
	}).then(res => {
		if(res.status !== 200) {
			navigate("/login",{replace:true});
		}
		return res.json();
	});
}
