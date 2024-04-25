<script lang="ts" setup>
const { public: { API_URL } } = useRuntimeConfig()
</script>

<template>
  <div class="container max-w-screen-sm mx-auto py-6 flex-1 flex flex-col justify-center items-center">
    <div class="flex-1 flex justify-center items-center w-full">
      <ClientOnly>
        <CodeTest mode="create" />
      </ClientOnly>
    </div>
  </div>
</template>
