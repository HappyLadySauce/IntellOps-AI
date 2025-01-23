<template>
  <a-input v-model="fileval" :placeholder="placeholder" @input="onChange" @clear="onChange('')" allow-clear/>
  <a-button type="primary" @click="UpFile">
    <template #default>选择附件</template>
  </a-button>
  <FileManage @register="registerFileModal" @success="selectImg"/>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import FileManage from '@/components/attachment/FileManage.vue';
import { useModal} from '/@/components/Modal';
const [registerFileModal, { openModal:openFileModal }] = useModal();
const props = defineProps({
  placeholder: String,
  modelValue: {
    type: String,
    required: true
  },
})
const fileval = computed(() => props.modelValue)
const emits = defineEmits(['update:modelValue'])
//上传
const UpFile=()=>{
  openFileModal(true, {
    filetype:"file",
    getnumber: "one",//one 单张 more多张
    openfrom: "use",//manage管理 use选择使用
  });
}
//选择附件返回
const selectImg=(item:any)=>{
   if(item.type=="one"){
      emits('update:modelValue', item.url)
    }
}
//修改值
const onChange = (value:any) => {
  emits('update:modelValue', value)
}
</script>
<style lang="less" scoped>
 .upfilezip{
    display: flex;
    .upbtn{
      padding-right: 10px;
    }
    .showfile{
      flex: 1;
      height: 32px;
      line-height: 32px;
      a{
        text-decoration: none;
      }
    }
  }
</style>
