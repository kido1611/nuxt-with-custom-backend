import type { UseFetchOptions } from "nuxt/app";

export function useApiLaravel<T>(
  url: string | (() => string),
  options?: UseFetchOptions<T>,
) {
  return useFetch(url, {
    ...options,
    $fetch: useNuxtApp().$apiLaravel as typeof $fetch,
  });
}
