import type { PageLoad } from "./$types";
import { loadTest } from "$lib/api/services";

export const load: PageLoad = async () => {
	return await loadTest()
};
