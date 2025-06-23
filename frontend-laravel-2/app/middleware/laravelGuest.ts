import { useAuth } from "~/compossables/useAuth";

export default defineNuxtRouteMiddleware(() => {
  const { user } = useAuth();

  if (user.value) {
    return navigateTo("/dashboard");
  }
});
