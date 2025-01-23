<template>
   <a-avatar :size="32" :imageUrl="GetFullPath(iamges[0])" shape="square" @click="visible=true">
       <template #trigger-icon>
        {{ iamges.length }}
      </template>
   </a-avatar>
   <a-image-preview-group v-model:visible="visible" infinite :srcList="iamges"/>
</template>

<script lang="ts" setup>
import { ref,computed } from 'vue';
import { GetFullPath } from "@/utils/tool";
interface Props {
  data: string
}
const props =withDefaults(defineProps<Props>(), {
  data: () => ""
})
const visible = ref(false)
const iamges = computed(() => {
  if(props.data){
  const imgs=  props.data.split(",")
  let resion3 = imgs.map(item => GetFullPath(item))
  return resion3
  }else{
    return []
  }
});
</script>

