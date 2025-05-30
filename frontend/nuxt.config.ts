// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: ["@nuxt/ui"],
  css: ["~/assets/css/main.css"],
  compatibilityDate: "2025-05-15",
  devtools: { enabled: true },
  future: {
    compatibilityVersion: 4,
  },
});
