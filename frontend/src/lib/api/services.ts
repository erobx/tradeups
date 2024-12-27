

const endpoint = "http://localhost:8080";

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
