<template>
<a-modal v-model:visible="visiblecheckpack" title="检查标识是否可用" title-align="start" :footer="false" :render-to-body="false" draggable width="530px" @before-close="handleClose">
    <div class="simplebox">
        <a-form ref="formRef" :model="formData" auto-label-width @submit="handleSubmit">
            <a-form-item field="type" label="标识类型" :rules="[{required:true,message:''}]" >
                <a-radio-group v-model="formData.type" direction="vertical">
                    <a-radio value="name">包名标识</a-radio>
                    <a-radio value="views">前端views下文件夹名</a-radio>
                    <a-radio value="components">前端components下组件文件夹名</a-radio>
                    <a-radio value="app">后端app下模块名</a-radio>
                    <a-radio value="business">后端app/business文件夹名</a-radio>
                    <a-radio value="utilsTool">后端utils/plugin下的插件文件夹名</a-radio>
                </a-radio-group>
            </a-form-item>
            <a-form-item field="name" label="检测标识" validate-trigger="input" :rules="[{required:true,message:'请填写检测标识名称'}]" >
                <a-input v-model="formData.name" placeholder="填写检测标识名称(只能以字母和下划线组成)" :max-length="50" allow-clear show-word-limit />
            </a-form-item>
            <a-form-item>
               <div class="submitbtn">
                    <a-space>
                        <a-button html-type="submit"  status="success">立即检测</a-button>
                        <a-button @click="handleSave" :loading="loading" type="primary" :disabled="!resdata">立即占用</a-button>
                    </a-space>
               </div>
            </a-form-item>
        </a-form>
        <a-alert type="warning">检测的目的是为了你开发的代码和其他开发代码包名不冲突，因为相同文件夹和文件名称在安装使用会被覆盖无法同时使用。
            检测可用后，如果您确定使用该“标识”请点击“立即占用”，防止您开发完后标识被别人占用。</a-alert>
    </div>
</a-modal>
</template>
<script lang="ts" setup>
import {ref,watch} from "vue";
import { checkPackName,savePackName } from '@/api/developer/codestore';
import { Message } from '@arco-design/web-vue';
import { FormInstance } from '@arco-design/web-vue/es/form';
import useLoading from '@/hooks/loading';
const emits = defineEmits(['update:modelValue'])
const props = defineProps({
  domian:{
    type: String,
    required: true
  },
  modelValue: {
    type: Boolean,
    required: true
  },
})
const visiblecheckpack=ref(false)
const formData=ref({
    type:"name",
    name:"",
})
watch(props, (nweProps)=>{
    visiblecheckpack.value=nweProps.modelValue
})
//关闭
const handleClose=()=>{
    emits('update:modelValue', false)
}
//提交检查
const formRef = ref<FormInstance>();
const resdata = ref(false);
const { loading, setLoading } = useLoading();
const handleSubmit=async()=>{
    try {
    const res = await formRef.value?.validate();
    if (!res) {
        Message.loading({content:"检查中",id:"post",duration:0})
        resdata.value=  await checkPackName(formData.value,props.domian);
        if(!resdata.value){
            Message.error({content:"标识已经被占用",id:"post",duration:2000}) 
        }else{
            Message.success({content:"标识可用",id:"post",duration:2000})
        }
    }
    } catch (error) {
        Message.loading({content:"检查中",id:"post",duration:1})
    }
}
//占用标识
const handleSave=async()=>{
    try {
        setLoading(true);
        Message.loading({content:"提交标识中",id:"save",duration:0})
        const res = await savePackName(formData.value,props.domian);
        if(res){
            resdata.value=false
            Message.success({content:"占用标识成功",id:"save",duration:2000})
        }else{
            Message.warning({content:"占用标识失败",id:"save",duration:2000})
        }
        setLoading(false);
    } catch (error) {
        setLoading(false);
        Message.loading({content:"提交标识中",id:"save",duration:1})
    }
}
</script>
<style scoped lang="less">
.simplebox{
    min-height: 400px;
}
.submitbtn{
    margin-top: 10px;
    width: 100%;
}
</style>