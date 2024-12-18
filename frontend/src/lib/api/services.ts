

const endpoint = "http://127.0.0.1:8080";

const api = {
	loadTest: endpoint + "/api/test",
}

export const loadTest = async (): Promise<any> => {
	try {
		return await (await fetch(api.loadTest)).json();
	} catch {
		return null
	}
}
