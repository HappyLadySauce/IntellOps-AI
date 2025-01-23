<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" :loading="loading" width="900px" @height-change="onHeightChange" :minHeight="modelHeight" :title="getTitle" @ok="handleSubmit">
    <div class="addFormbox" :style="{'min-height':`${windHeight}px`}">
      <div class="tabs-content" >
        <a-form ref="formRef" :model="formData" auto-label-width>
          <div class="content_box">
              <!--基础信息-->
              <a-scrollbar   style="overflow: auto;" :style="{height:`${windHeight}px`}">
                <div class="besecontent" >
                  <a-row :gutter="16">
<!--replaceTpl-->
                  </a-row>
                </div>
              </a-scrollbar>
          </div>
        </a-form>
      </div>
    </div>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref} from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import { cloneDeep } from 'lodash-es';
  //api
  import { save,getContent } from './api';
  import { Message } from '@arco-design/web-vue';
  export default defineComponent({
    name: 'AddBook',
    components: { BasicModal},
    emits: ['success'],
    setup(_, { emit }) {
      const isUpdate = ref(false);
      const modelHeight= ref(520);
      const windHeight= ref(520);
      //表单
      const formRef = ref<FormInstance>();
      //表单字段
      const basedata={
            id:0,
replaceField:null
        }
      const formData = ref(basedata)
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          formRef.value?.resetFields()
          setModalProps({ confirmLoading: false });
          isUpdate.value = !!data?.isUpdate;
          if (unref(isUpdate)) {
            formData.value=cloneDeep(data.record)
            setLoading(true);
            const mewdata = await getContent({id:data.record.id});
            formData.value=Object.assign({},formData.value,mewdata)
            setLoading(false);
          }else{
            formData.value=cloneDeep(basedata)
          }
      });
      const getTitle = computed(() => (!unref(isUpdate) ? '新增数据' : '编辑数据'));
     //提交数据
     const { loading, setLoading } = useLoading();
     const handleSubmit = async () => {
      try {
          const res = await formRef.value?.validate();
          if (!res) {
            setLoading(true);
            Message.loading({content:"提交中",id:"submit",duration:0})
            let savedata=cloneDeep(unref(formData))
            await save(savedata);
            Message.success({content:"提交成功",id:"submit",duration:1500})
            closeModal()
            emit('success');
            setLoading(false);
          }
        } catch (error) {
          setLoading(false);
          Message.loading({content:"提交中",id:"submit",duration:1})
        }
      };
       //监听高度
       const onHeightChange=(val:any)=>{
        windHeight.value=val
      }
      return { 
        registerModal, 
        getTitle, 
        handleSubmit,
        formRef,
        loading,
        formData,
        modelHeight,
        onHeightChange,windHeight,
      };
    },
  });
</script>
<style lang="less" scoped>
  @import '@/assets/style/formlayer.less';
  .besecontent{
    padding: 15px 30px 5px 30px!important;
  }
</style>

