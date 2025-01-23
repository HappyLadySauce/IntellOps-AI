
<template>
  <a-select :model-value="modelValue" :multiple="multiple" :placeholder="placeholder" allow-clear allow-search 
       :loading="SelectLoading" @change="onChange">
        <template v-for="item in SelectList">
          <a-option :value="item?.value" v-if="item?.label_txt" :label="item?.label"><span v-html="item?.label_txt"></span></a-option>
          <a-option :value="item?.value" v-else>{{ item?.label }}</a-option>
        </template>
    </a-select>
  </template>
  
  <script lang="ts" setup>
  import { ref,onMounted,watch } from 'vue'
  import { getTableDataForm } from '@/api/datacenter/dictionary';
  const props = defineProps({
    placeholder: String, 
    modelValue: {
      type: Number,
      required: true
    },
    tablename: {
      type: String,
      required: true
    },
    multiple : {//是否多选
      type: Boolean,
      default: false
    },
    showfield: {
      type: String,
      required: true
    },
    from: {
      type: String,
      default: "form"//默认是表单，search是搜索输入框
    },
    custom: {
      type: String,
      default: ""//查找条件
    },
  })
  const emits = defineEmits(['update:modelValue'])
  //组件挂载完成后执行的函数
  onMounted(()=>{
    if(props.tablename){
        handleCommonSelect(props.tablename)
      }
  })
  watch(
      () => props.tablename,
      (nweProps) => {
        if(nweProps&&props.tablename!=nweProps){
          handleCommonSelect(nweProps)
        }
      },
      { immediate: true, deep: true }
  )
  const SelectList = ref<any>([]);
  const SelectLoading=ref(true)
  const handleCommonSelect=async(tablename:string)=>{
      SelectLoading.value=true
      const mlist= await getTableDataForm({tablename:tablename,showfield:props.showfield,custom:props.custom})
      const parntList_df : any=[{value: 0,label:FontDefaultTitle()}];
      if(mlist){
        SelectList.value=parntList_df.concat(mlist)
      }else{
        SelectList.value=parntList_df
      }
      SelectLoading.value=false
  }
  //修改值
  const onChange = (value:any) => {
      emits('update:modelValue', value)
  }
  //过滤默认值提示
  const FontDefaultTitle = ():string=> {
       if(props.from=="form"){
          return "未选"
       }else if(props.from=="search"){
          return "搜索全部"
       }else{
        return "默认值"
       }
  }
  </script>
  