// @ts-ignore
import { InstallCodemirro } from "codemirror-editor-vue3"

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(InstallCodemirro)
})