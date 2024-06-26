// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxt/ui'],

  runtimeConfig: {
    public: {
      API_URL: "https://pasteshare-api.viandwi24.com"
    }
  }
})
