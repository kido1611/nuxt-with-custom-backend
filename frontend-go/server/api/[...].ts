import { joinURL } from "ufo";

/**
 * Taken from video Alexander Lichter
 * https://youtu.be/J4E5uYz5AY8?si=z4_1E5wx4GNScn1x&t=816
 */
export default defineEventHandler(async (event) => {
  const proxyUrl = useRuntimeConfig(event).apiUrl;

  const path = event.path.replace(/^\/api\//, "");

  const target = joinURL(proxyUrl, path);

  return proxyRequest(event, target);
});
