<script lang="ts" setup>
import hljs, { type LanguageFn } from 'highlight.js';
// @ts-ignore
import CodeEditor from "simple-code-editor"
// import 'simple-code-editor/themes/themes.css';

import hljsjavascript from 'highlight.js/lib/languages/javascript'
import hljsphp from 'highlight.js/lib/languages/php'
import hljsplaintext from 'highlight.js/lib/languages/plaintext'
import hljstypescript from 'highlight.js/lib/languages/typescript'
import hljspython from 'highlight.js/lib/languages/python'
import hljsgo from 'highlight.js/lib/languages/go'

const codesLangs: [LanguageFn, string][] = [
  [hljsjavascript, 'javascript'],
  [hljsphp, 'php'],
  [hljsplaintext, 'plaintext'],
  [hljstypescript, 'typescript'],
  [hljspython, 'python'],
  [hljsgo, 'go'],
]

const props = defineProps({
  mode: {
    type: String as PropType<"create" | "view">,
    default: "view",
  },
  uuid: {
    type: String,
    required: false,
  },
  title: {
    type: String,
    required: false,
  },
  content: {
    type: String,
    required: false,
  },
})

const { public: { API_URL } } = useRuntimeConfig()

const code = ref({
  lang: 'plaintext',
})
const inputs = ref({
  title: props.title || '',
  content: props.content || 'console.log("Hello, World!")',
})
const langs = ref([
  ['plaintext', 'Plain Text'],
  ['cpp', 'C++'],
  ['python', 'Python'],
  ['php', 'PHP'],
  ['javascript', 'JS'],
  ['typescript', 'TypeScript'],
  ['html', 'HTML'],
  ['go', 'Go'],
])
// langs, but code.lang first
const langsSorted = computed(() => {
  const lgs = [...langs.value]
  const idx = lgs.findIndex((l) => l[0] === code.value.lang)
  if (idx !== -1) {
    lgs.splice(idx, 1)
    lgs.unshift(langs.value[idx])
  }
  console.log(lgs)
  return lgs
})
// const options = ref<EditorConfiguration>({
//   mode: "php",
//   theme: "dracula",
//   readOnly: props.mode === 'view',
//   lineNumbers: true,
//   lineWiseCopyCut: true,
//   lineWrapping: true,
//   // height: '450px'
// })

onMounted(async () => {
  codesLangs.forEach((lang) => {
    hljs.registerLanguage(lang[1], lang[0])
  })
})

const save = async () => {
  console.log(inputs.value)
  const res = await $fetch(`${API_URL}/doc`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(inputs.value),
  })
  const data = (res as any).Data as {
    Id: string
  }
  navigateTo(`${data.Id}`)
}

const copy = async () => {
  await navigator.clipboard.writeText(`${location.origin}/${props.uuid}`)
  alert('Copied to clipboard!')
}

const getLanguage = (lang: string) => {
  console.log(lang)
}
</script>

<template>
  <div class="flex-1 flex flex-col items-center justify-center">
    <div class="relative max-w-[350px] sm:max-w-[600px] md:max-w-[700px] lg:max-w-[900px] xl:max-w-[1000px] flex flex-col">
      <div class="rounded-t-lg overflow-hidden px-4 py-2 bg-gray-800 flex gap-4 items-center">
        <div class="flex gap-2">
          <div class="h-3 w-3 bg-red-500 rounded-full"></div>
          <div class="h-3 w-3 bg-yellow-500 rounded-full"></div>
          <div class="h-3 w-3 bg-green-500 rounded-full"></div>
        </div>
        <div>
          <input
            type="text"
            class="bg-gray-800 text-white border-none focus:outline-none"
            placeholder="give this paste a title..."
            autofocus
            v-model="inputs.title"
            :disabled="mode === 'view'"
            :readonly="mode === 'view'"
          />
        </div>
      </div>
      <div class="relative bg-gray-800">
        <CodeEditor
          v-model="inputs.content"
          :languages="langsSorted"
          :read-only="mode === 'view'"
          line-nums
          @lang="getLanguage"
          width="100%"
        />
      </div>
    </div>
    <div v-if="mode === 'create'" class="mt-4 flex justify-center">
      <UButton label="Create Paste" variant="soft" @click="save" icon="i-heroicons-pencil" />
    </div>
    <div v-else class="mt-4 flex justify-center">
      <UButton label="Share Paste" variant="soft" icon="i-heroicons-link" @click="copy" />
    </div>
  </div>
</template>