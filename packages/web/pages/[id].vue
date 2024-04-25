<script lang="ts" setup>
const { public: { API_URL } } = useRuntimeConfig()

const $route = useRoute()

const { data } = await useFetch<{
  Data: {
    Title: string
    Content: string
  }
}>(`${API_URL}/doc/${$route.params.id as string}`)

const inputs = ref({
  title: data.value?.Data.Title || '',
  content: data.value?.Data.Content || '',
})

const meta = {
  title: `${inputs.value.title} - 📚 PASTESHARE`,
  description: 'Create your paste and share it with your friends, family, or anyone else.'
}
useSeoMeta({
  title: meta.title,
  description: meta.description,
  ogTitle: meta.title,
  ogDescription: meta.description,
})

// const fetch = async () => {
//   try {
//     const res = await $fetch(`${API_URL}/doc/${$route.params.id as string}`)
//     const data = (res as any).Data as {
//       Title: string
//       Content: string
//     }
//     inputs.value = {
//       title: data.Title,
//       content: data.Content,
//     }
//     console.log(data)
//   } catch (error) {
//     console.error(error)
//     navigateTo('/')
//   }
// }
</script>

<template>
  <div class="container max-w-screen-sm mx-auto py-6 flex-1 flex flex-col justify-center items-center">
    <div class="flex-1 flex justify-center items-center w-full">
      <ClientOnly>
        <CodeTest mode="view" :uuid="($route.params.id as string)" :title="inputs.title" :content="inputs.content" />
      </ClientOnly>
    </div>
  </div>
</template>
