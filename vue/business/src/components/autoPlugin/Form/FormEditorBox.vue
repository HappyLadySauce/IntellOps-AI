<template>
    <div class="hcontent" v-show="activeKey==subnum" :style="{height:`${winHeights}px`}">
      <Editor :minHeight="winHeights" :valueTyle="valueTyle" :placeholder="placeholder" @doclick="handleFocus" ref="editorRef" @updata="handleEditUpdta"/>
    </div>
</template>

<script lang="ts" setup>
import { ref,onUnmounted,watch } from 'vue'
import Editor from "@/components/Editor/Main.vue"; 
const props = defineProps({
  activeKey: Number,
  subnum: Number,
  winHeights: Number,
  placeholder: String,
  modelValue: {
    type: String,
    required: true
  },
  valueTyle: {
    type: String,
     default: "getHtml"
  },
})
const emits = defineEmits(['update:modelValue'])
const editorRef = ref();
const lockdit= ref(true);
const handleFocus=()=>{
  lockdit.value=false
}
onUnmounted(() => {
  lockdit.value=true
});
//编辑器
watch(
  () => props.modelValue,
  (nweProps) => {
    if(editorRef.value&&lockdit.value){
    editorRef.value.setVal(nweProps)
  }
  },
  { immediate: true, deep: true }
)
//编辑器返回数据
const handleEditUpdta=(val:string)=>{
  emits('update:modelValue', val)
}
</script>
<style lang="less" scoped>
</style>
