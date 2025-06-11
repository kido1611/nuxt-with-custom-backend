import type { User } from "~/types";

export function useAuth() {
  const user = useState<User | null>("user-auth", () => null);

  const apiLaravelFetch = useNuxtApp().$apiLaravel;

  const fetchUser = async () => {
    try {
      const response = await apiLaravelFetch<User>("/api/user", {
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

    await apiLaravelFetch("/api/auth/login", {
      method: "POST",
      body,
    });

    await fetchUser();
  };

  const logout = async () => {
    if (!user.value) {
      return;
    }

    await apiLaravelFetch("/api/auth/logout", {
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
