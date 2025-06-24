import type { User } from "~/types";

export function useAuth() {
  const user = useState<User | null>("user-auth", () => null);

  const apiFetch = createFetch({
    clearUser: () => {
      user.value = null;
    },
  });

  const fetchUser = async () => {
    try {
      const response = await apiFetch<User>("/api/api/user", {
        method: "GET",
        headers: {
          Accept: "application/json",
        },
      });

      user.value = response;
    } catch (error) {
      user.value = null;
    }
  };

  const login = async (body: Record<string, any>) => {
    if (user.value) {
      throw new Error("Already authenticated.");
    }

    await apiFetch("/api/api/auth/login", {
      method: "POST",
      body,
    });

    await fetchUser();
  };

  const logout = async () => {
    if (!user.value) {
      return;
    }

    await apiFetch("/api/api/auth/logout", {
      method: "DELETE",
    });

    user.value = null;
  };

  return {
    user,
    fetchUser,
    login,
    logout,
  };
}
