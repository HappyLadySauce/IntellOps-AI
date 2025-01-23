<template>
  <a-image :height="props.height" :width="props.width" :src="GetFullPath(props.src)"/>
</template>

<script lang="ts" setup>
import { GetFullPath } from "@/utils/tool";
const props = withDefaults(defineProps<Props>(), {
  src: '',
  height: '32',
  width: '32',
})
interface Props {
  src: string
  height: string
  width: string
}
</script>

