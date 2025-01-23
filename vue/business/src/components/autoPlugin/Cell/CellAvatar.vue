<template>
  <a-space fill :size="spacesize">
    <a-avatar :size="size" shape="circle" :imageUrl="GetFullPath(props.avatar)"></a-avatar>
    <a-link v-if="props.isLink" @click="emit('click')">{{ props.name }}</a-link>
    <span v-else>{{ props.name }}</span>
  </a-space>
</template>

<script lang="ts" setup>
import { GetFullPath } from "@/utils/tool";
const props = withDefaults(defineProps<Props>(), {
  avatar: '',
  name: '',
  spacesize:"small",
  size:24,
  isLink: false // 是否可以点击
})

const emit = defineEmits<{
  (e: 'click'): void
}>()

interface Props {
  avatar: string
  name: string
  spacesize: any
  size: number
  isLink?: boolean
}
</script>

