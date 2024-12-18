// @ts-nocheck
import type { PageLoad } from "./$types";
import { loadTest } from "$lib/api/services";

export const load = async () => {
	return await loadTest()
};
;null as any as PageLoad;