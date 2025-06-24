import { useAuth } from "~/compossables/useAuth";
import type { User } from "~/types";
import { createFetch } from "~/utils/fetch";

export default defineNuxtPlugin(async () => {
  const { user } = useAuth();

  const apiFetch = createFetch({
    clearUser: () => {
      user.value = null;
    },
  });

  // check is currently logged in
  if (!user.value) {
    try {
      const response = await apiFetch<User>("/api/api/user", {
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
      apiFetch,
    },
  };
});
