<template>
 <span class="celllink"  v-if="data">
  <span class="link" :title="data">{{ data }}</span>
   <a-link @click="handleCopy"><icon-copy /></a-link>
 </span>
</template>

<script lang="ts" setup>
import { Message } from '@arco-design/web-vue';
import { useClipboard } from '@vueuse/core';
const { copy } = useClipboard();
const props = withDefaults(defineProps<Props>(), {
  data: ""
})

interface Props {
  data: string
}
const handleCopy=async()=>{
  await copy(props.data);
  Message.success({content:"复制成功",id:"copysuccess"});
}
</script>
<style scoped lang="less">
  .celllink {
    display:flex;
    width: 100%;
    .link{
      display:block;
      overflow:hidden;
      text-overflow:ellipsis;
      white-space:nowrap;
      color: var(--color-neutral-6);
    }
  }
</style>
