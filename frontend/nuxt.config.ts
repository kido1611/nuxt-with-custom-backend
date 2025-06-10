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
    baseUrl: "/api/laravel",
    endpoints: {
      login: "/api/auth/login",
      logout: "/api/auth/logout",
    },
    redirect: {
      onLogin: "/dashboard",
      onGuestOnly: "/dashboard",
    },
  },
  routeRules: {
    "/api/laravel/**": {
      proxy: "http://localhost:8000/**",
    },
  },
});
