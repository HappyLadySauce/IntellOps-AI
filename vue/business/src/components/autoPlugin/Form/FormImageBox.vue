<template>
  <div class="upimagebox" >
    <div class="imagebtn">
      <div class="upload-show-picture" v-if="modelValue">
        <a-image :src="GetFullPath(modelValue)" height="90" :preview-visible="visibleimage"
          @preview-visible-change="() => { visibleimage = false }" />
        <div class="upload-show-picture-mask">
          <a-space>
            <icon-eye @click="() => visibleimage = true" class="opbtn" />
            <IconEdit @click="UpImage" class="opbtn" />
            <icon-delete @click="DelImage" class="opbtn"/>
          </a-space>
        </div>
      </div>
      <div class="upload-picture-card" v-else @click="UpImage">
        <div class="upload-picture-card-text">
          <IconPlus />
          <div style="margin-top: 10px; font-weight: 600">上传图片</div>
        </div>
      </div>
    </div>
  </div>
  <!--附件-->
  <FileManage @register="registerFileModal" @success="selectImg"/>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import FileManage from '@/components/attachment/FileManage.vue';
import { useModal} from '/@/components/Modal';
import { GetFullPath } from "@/utils/tool";
const props = defineProps({
  placeholder: String,
  modelValue: {
    type: String,
    required: true
  },
})
const emits = defineEmits(['update:modelValue'])
const [registerFileModal, { openModal:openFileModal }] = useModal();
const visibleimage = ref(false)
//组件挂载完成后执行的函数
onMounted(() => {
})
//上传图片
const UpImage=()=>{
  openFileModal(true, {
    filetype:"image",
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
//删除
const DelImage=()=>{
  emits('update:modelValue', "")
}
</script>
<style lang="less" scoped>
 //上传图片
 .upimagebox{
    display: flex;
    .imagebtn{
      position: relative;
      width: 145px;
      height: 90px;
      background-color: var(--color-neutral-1);
      border-radius: 4px;
      overflow: hidden;
      flex-shrink: 0;
      //预览
      .upload-show-picture{
        position: relative;
        box-sizing: border-box;
        width: 100%;
        height: 100%;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
        img{
          height: 100%;
        }
        &:hover{
          .upload-show-picture-mask{
             opacity: 1;
          }
        }
        .upload-show-picture-mask{
            position: absolute;
            top: 0;
            right: 0;
            bottom: 0;
            left: 0;
            color: var(--color-white);
            font-size: 16px;
            line-height: 90px;
            text-align: center;
            background: rgba(0, 0, 0, 0.5);
            cursor: pointer;
            opacity: 0;
            transition: opacity 0.1s cubic-bezier(0, 0, 1, 1);
        }
      }
      .upload-picture-card{
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        .upload-picture-card-text{
           text-align: center;
           color:  var(--color-neutral-5);
        }
      }
    }
  }
</style>
