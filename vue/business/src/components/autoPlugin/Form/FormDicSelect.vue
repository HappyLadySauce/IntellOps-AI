
<template>
<a-select :model-value="modelValue" :multiple="multiple" :options="SelectList" :placeholder="placeholder" allow-clear allow-search 
     :loading="SelectLoading" @change="onChange"/>
</template>

<script lang="ts" setup>
import { ref,onMounted,watch } from 'vue'
import { getDicData } from '@/api/datacenter/dictionary';
const props = defineProps({
  placeholder: String, 
  modelValue: {
    type: String,
    required: true
  },
  multiple : {//是否多选
    type: Boolean,
    default: false
  },
  dicgroupid: {
    type: String,
    required: true
  },
})
const emits = defineEmits(['update:modelValue'])
//组件挂载完成后执行的函数
onMounted(()=>{
  if(props.dicgroupid){
      handleCommonSelect(props.dicgroupid)
    }
})
watch(
  () => props.dicgroupid,
  (nweProps) => {
    if(nweProps&&props.dicgroupid!=nweProps){
      handleCommonSelect(nweProps)
    }
  },
  { immediate: true, deep: true }
)
const SelectList = ref<any>([]);
const SelectLoading=ref(true)
const handleCommonSelect=async(dicgroupid:string)=>{
    SelectLoading.value=true
    const mlist= await getDicData({group_id:dicgroupid})
      if(mlist){
        SelectList.value=mlist
      }else{
        SelectList.value=[]
      }
    SelectLoading.value=false
}
//修改值
const onChange = (value:any) => {
    emits('update:modelValue', value)
}
</script>
