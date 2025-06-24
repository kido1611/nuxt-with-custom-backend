import type { FetchContext } from "ofetch";
import { useAuth } from "~/compossables/useAuth";
import type { User } from "~/types";

export default defineNuxtPlugin(async () => {
  const { user } = useAuth();

  // taken from: https://github.com/manchenkoff/nuxt-auth-sanctum/blob/cd3f86384b62f1f2f07ef7236f10f0c16c6220b3/src/runtime/httpFactory.ts#L35
  const determineCredentialsMode = () => {
    // Fix for Cloudflare workers - https://github.com/cloudflare/workers-sdk/issues/2514
    return "credentials" in Request.prototype ? "include" : undefined;
  };

  // create custom $fetch
  const apiLaravel = $fetch.create({
    baseURL: "/api/laravel",
    credentials: determineCredentialsMode(),
    onRequest: async (context: FetchContext): Promise<void> => {
      if (import.meta.server) {
        // when data is fetched on SSR, append cookie (taken from client) and
        // Referrer/Origin to headers (Required when using sanctum SPA mode. Written in docs).
        const clientCookies = useRequestHeaders(["cookie"]);

        const origin = useRequestURL().origin;

        context.options.headers.append("Referrer", origin);
        context.options.headers.append("Origin", origin);
        if (clientCookies && clientCookies.cookie) {
          context.options.headers.append("Cookie", clientCookies.cookie);
        }
      }

      // TODO: nice to have to transform when using FormData with method PUT/PATCH
      // when using laravel, need to convert it into POST and add body `_method`
      // https://github.com/qirolab/nuxt-sanctum-authentication/blob/14f9d571f195d472231c135e1491698b4c00b4e8/src/runtime/services/createFetchService.ts#L135

      const method = context.options.method ?? "GET";

      // Add csrf-token when doing a request with method POST/PUT/PATCH/DELETE
      // if cookie already have XSRF-TOKEN, just take it and append to header
      // if cookie is missing, call /sanctum/csrf-cookie and take the XSRF-TOKEN cookie
      if (!["POST", "PUT", "PATCH", "DELETE"].includes(method)) {
        return;
      }

      let csrfCookie = useCookie("XSRF-TOKEN", { readonly: true });

      if (!csrfCookie.value) {
        await $fetch("/sanctum/csrf-cookie", {
          baseURL: "/api/laravel",
          method: "GET",
          credentials: "include",
        });

        csrfCookie = useCookie("XSRF-TOKEN", { readonly: true });
      }

      context.options.headers.append("X-XSRF-TOKEN", csrfCookie.value ?? "");
    },
    onResponseError: async (context: FetchContext): Promise<void> => {
      if (context.response?.status === 401) {
        user.value = null;
      }
      // TODO: handle clean user when status unauthorized (401)
      // TODO: handle when token missmatch (419)

      console.log("on response error");
    },
  });

  // check is currently logged in
  if (!user.value) {
    try {
      const response = await apiLaravel<User>("/api/user", {
        method: "GET",
        headers: {
          Accept: "application/json",
        },
      });

      user.value = response;
    } catch (error) {}
  }

  return {
    provide: {
      apiLaravel,
    },
  };
});
