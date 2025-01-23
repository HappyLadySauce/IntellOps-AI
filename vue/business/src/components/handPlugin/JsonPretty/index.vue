<template>
  <div class="json_pretty_container">
    <VueJsonPretty path="res" :data="JSONObject" :show-length="true" />
    <icon-copy class="copy_icon" @click="onCopy(JSONObject,$t('sys.copy.success'))" />
  </div>
</template>

<script setup lang="ts">
  import { computed } from 'vue';
import VueJsonPretty from 'vue-json-pretty'
import 'vue-json-pretty/lib/styles.css'
import { CopyText } from '@/utils/tool'
defineOptions({ name: 'JsonPretty', inheritAttrs: false })
const props = defineProps<{
  json: string
}>()

const JSONObject = computed(() => JSON.parse(props?.json))

// 拷贝
const onCopy = (data: object,msg:string) => {
  CopyText(JSON.stringify(data),msg)
}
</script>

<style lang="less" scoped>
.json_pretty_container {
  width: 100%;
  height: 100%;
  overflow: auto;
  position: relative;
  .copy_icon {
    position: absolute;
    right: 10px;
    top: 10px;
    font-size: 18px;
    cursor: pointer;
  }
}
</style>
