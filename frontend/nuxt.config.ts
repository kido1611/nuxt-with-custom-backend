// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: ["@nuxt/ui", "nuxt-auth-sanctum"],
  css: ["~/assets/css/main.css"],
  compatibilityDate: "2025-05-15",
  devtools: { enabled: true },
  future: {
    compatibilityVersion: 4,
  },
  sanctum: {
    baseUrl: "http://localhost:8000",
    endpoints: {
      login: "/api/login",
      logout: "/api/logout",
    },
    redirect: {
      onLogin: "/dashboard",
    },
  },
});
