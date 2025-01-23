<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :loading="loading" :footer="false" width="600px" :minHeight="390" :title="getTitle" >
   <div class="modalbox">
    <a-form ref="formRef" :model="formData" auto-label-width style="padding:10px 20px;"  @submit="handleSubmit">
      <a-form-item field="name" label="名称" validate-trigger="input" :rules="[{required:true,message:'请填写名称'}]" >
        <a-input v-model="formData.name" placeholder="请填写名称" />
      </a-form-item>
      <a-form-item field="path_url" label="跳转路径" validate-trigger="input" :rules="[{required:true,message:'请填写跳转路径'}]" >
        <a-input v-model="formData.path_url" placeholder="请填跳转路径" />
      </a-form-item>
      <a-form-item field="icon" label="选择图标" :rules="[{required:true,message:'请填写选择图标'}]" >
          <a-input-search v-model="formData.icon" placeholder="选择图标/填写"  search-button>
            <template  v-if="formData.icon" #prefix>
              <Icon :icon="formData.icon" ></Icon>
            </template>
              <template #button-icon>
                <a-popover position="br" trigger="click">
                  <icon-apps :size="23"/>
                  <template #content>
                      <IconPicker @change="(icon)=>{formData.icon=icon}"></IconPicker>
                  </template>
                </a-popover>
              </template>
            </a-input-search>
      </a-form-item>
      <a-form-item label="类型" field="type" style="margin-bottom:10px;">
        <a-radio-group v-model="formData.type" >
          <a-radio :value="0">跳转系统</a-radio>
          <a-radio :value="1">跳转外部</a-radio>
        </a-radio-group>
      </a-form-item>
      <a-form-item label="排序" field="weigh" style="margin-bottom:15px;">
        <a-input-number v-model="formData.weigh" placeholder="请填排序" />
      </a-form-item>
      <a-form-item >
        <a-button type="primary" html-type="submit" :loading="loading" style="width: 150px;margin-left: 80px;margin-top: 15px;">提交</a-button>
      </a-form-item>
    </a-form>
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
  import {  saveQuick } from '@/api/dashboard/workplace';
  import { IconPicker ,Icon} from '@/components/Icon';
  import { Message } from '@arco-design/web-vue';
  export default defineComponent({
    name: 'QuickForm',
    components: { BasicModal,IconPicker,Icon },
    emits: ['success'],
    setup(_, { emit }) {
      const isUpdate = ref(false);
      //表单
      const formRef = ref<FormInstance>();
      //表单字段
      const basedata={
            id:0,
            type:0,
            name: "",
            path_url: "",
            icon: "",
            weigh:0,
        }
      const formData = ref(basedata)
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          formRef.value?.resetFields()
          setModalProps({ confirmLoading: false });
          isUpdate.value = !!data?.isUpdate;
          if (unref(isUpdate)) {
            formData.value=cloneDeep(data.record)
          }else{
            formData.value=cloneDeep(basedata)
          }
      });
      const getTitle = computed(() => (!unref(isUpdate) ? '新增快捷按键' : '编辑快捷按键'));
     //回复留言
     const { loading, setLoading } = useLoading();
     const handleSubmit = async () => {
      try {
          const res = await formRef.value?.validate();
          if (!res) {
            setLoading(true);
            Message.loading({content:"提交中",id:"upStatus",duration:0})
            await saveQuick(formData.value);
            Message.success({content:"提交成功",id:"upStatus",duration:2000})
            closeModal()
            emit('success');
            setLoading(false);
          }
        } catch (error) {
          setLoading(false);
          Message.clear("top")
        }
      };
      return { 
        registerModal, 
        getTitle, 
        handleSubmit,
        formRef,
        loading,
        formData, 
      };
    },
  });
</script>
<style lang="less" scoped>
.modalbox{
  margin-bottom: 30px;
}
</style>