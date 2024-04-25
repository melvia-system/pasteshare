<script lang="ts" setup>
const meta = {
  title: '📚 PASTESHARE - Create your paste and share it',
  description: 'Create your paste and share it with your friends, family, or anyone else.'
}
useSeoMeta({
  title: meta.title,
  description: meta.description,
  ogTitle: meta.title,
  ogDescription: meta.description,
  author: 'viandwi24',
  keywords: 'paste, share, code, text, viandwi24',
  ogType: 'website',
})
</script>

<template>
  <div class="w-full h-screen bg-gray-50 dark:bg-gray-950 flex flex-col">
    <div class="h-[74px] border-b-2 border-gray-500/50 flex items-center">
      <div class="w-full mx-auto max-w-screen-lg px-4 flex items-center justify-between">
        <ULink to="/" class="text-lg font-bold">
          📚 PASTESHARE
        </ULink>
        <div class="flex divide-x-2 divide-gray-500/50 gap-2">
          <UButton size="xs" variant="soft" to="/">New Paste</UButton>
          <div class="pl-2">
            <ButtonColorModeToggle />
          </div>
        </div>
      </div>
    </div>
    <div class="flex-1 flex">
      <NuxtPage />
    </div>
    <div class="h-[68px] border-t-2 border-gray-500/50 flex items-center">
      <div class="w-full mx-auto max-w-screen-lg px-4">
        <div>
          &copy; {{ new Date().getFullYear() }}
          by <UButton variant="link" href="https://viandwi24.com" target="_blank" label="viandwi24" class="px-0" />
        </div>
      </div>
    </div>
  </div>
</template>
