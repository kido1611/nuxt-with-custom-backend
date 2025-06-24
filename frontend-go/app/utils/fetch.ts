import type { FetchContext, $Fetch } from "ofetch";

const CSRF_COOKIE = "XSRF-TOKEN";
const CSRF_HEADER = "X-XSRF-TOKEN";

const determineCredentialsMode = () => {
  // Fix for Cloudflare workers - https://github.com/cloudflare/workers-sdk/issues/2514
  return "credentials" in Request.prototype ? "include" : undefined;
};

export const createFetch = (options: { clearUser: () => void }): $Fetch => {
  return $fetch.create({
    credentials: determineCredentialsMode(),
    onRequest: async (context: FetchContext): Promise<void> => {
      if (import.meta.server) {
        // when data is fetched on SSR, append cookie (taken from client) and
        // Referrer/Origin to headers (Required when using sanctum SPA mode. Written in docs).
        const clientCookies = useRequestHeaders(["cookie"]);
        if (clientCookies && clientCookies.cookie) {
          context.options.headers.append("Cookie", clientCookies.cookie);
        }

        const origin = useRequestURL().origin;
        context.options.headers.append("Referrer", origin);
        context.options.headers.append("Origin", origin);
      }

      const method = context.options.method ?? "GET";

      // Add csrf-token when doing a request with method POST/PUT/PATCH/DELETE
      // if cookie already have XSRF-TOKEN, just take it and append to header
      // if cookie is missing, call /sanctum/csrf-cookie and take the XSRF-TOKEN cookie
      if (!["post", "put", "patch", "delete"].includes(method.toLowerCase())) {
        return;
      }

      const csrfToken = await initCsrfToken();
      if (csrfToken) {
        context.options.headers.append(CSRF_HEADER, csrfToken);
      }
    },
    onResponseError: async (context: FetchContext): Promise<void> => {
      if (context.response?.status === 401) {
        options.clearUser();
      }
    },
  }) as $Fetch;
};

const initCsrfToken = async (
  force = false,
): Promise<string | null | undefined> => {
  if (!force) {
    const csrfToken = useCookie(CSRF_COOKIE, {
      readonly: true,
    }).value;

    if (csrfToken) {
      return csrfToken;
    }
  }

  await $fetch("/api/sanctum/csrf-cookie", {
    method: "GET",
    credentials: "include",
  });

  return useCookie(CSRF_COOKIE, { readonly: true }).value;
};
