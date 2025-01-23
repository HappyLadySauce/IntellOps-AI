<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :loading="loading" helpMessage="编辑和修改部门" width="900px" :minHeight="500" :title="getTitle" @ok="handleSubmit">
    <a-form ref="formRef" :model="formData" auto-label-width style="padding:10px 20px;">
      <a-form-item field="name" label="部门名称" validate-trigger="input" :rules="[{required:true,message:'请填写部门名称'}]" style="margin-bottom:15px;">
        <a-input v-model="formData.name" placeholder="请填写部门名称" />
      </a-form-item>
      <a-form-item label="上级部门" field="pid" style="margin-bottom:15px;">
        <a-tree-select placeholder="选择上级部门" :data="parntList" 
        :fieldNames="{
            key: 'id',
            title: 'name',
            children: 'children'
          }"
        v-model="formData.pid">
        </a-tree-select>
      </a-form-item>
      <a-form-item label="排序" field="weigh" style="margin-bottom:15px;">
        <a-input-number v-model="formData.weigh" placeholder="请填排序" />
      </a-form-item>
      <a-form-item field="remark" label="备注" validate-trigger="input" style="margin-bottom:15px;">
        <a-textarea v-model="formData.remark" placeholder="请填写备注" allow-clear/>
      </a-form-item>
    </a-form>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref} from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import { useI18n } from 'vue-i18n';
  import { cloneDeep } from 'lodash-es';
  //api
  import { getParent, save } from '@/api/system/dept';
  import { IconPicker ,Icon} from '@/components/Icon';
  import { Message } from '@arco-design/web-vue';
  export default defineComponent({
    name: 'AddForm',
    components: { BasicModal,IconPicker,Icon },
    emits: ['success'],
    setup(_, { emit }) {
      const { t } = useI18n();
      const isUpdate = ref(false);
      const parntList = ref([]);
      //表单
      const formRef = ref<FormInstance>();
      //表单字段
      const basedata={
            id:0,
            name: '',
            pid: 0,
            weigh: 1,
            remark:"",
        }
      const formData = ref(basedata)
      const m_component=ref("")
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          formRef.value?.resetFields()
          setModalProps({ confirmLoading: false });
          const resultdata = await getParent();
          const parntList_df : any=[{id: 0,name: "一级部门",pid: 0}];
          if(resultdata){
            parntList.value=parntList_df.concat(resultdata)
          }else{
            parntList.value=parntList_df
          }
          isUpdate.value = !!data?.isUpdate;
          if (unref(isUpdate)) {
            m_component.value=data.record.component
            formData.value=cloneDeep(data.record)
          }else{
            formData.value= cloneDeep(basedata)
          }
      });
      const getTitle = computed(() => (!unref(isUpdate) ? '新增部门' : '编辑部门'));
     //点击确认
     const { loading, setLoading } = useLoading();
     const handleSubmit = async () => {
      try {
          const res = await formRef.value?.validate();
          if (!res) {
            setLoading(true);
            Message.loading({content:"提交中",id:"save",duration:0})
            await save(unref(formData));
            Message.success({content:"提交成功",id:"save",duration:2000})
            closeModal()
            emit('success');
            setLoading(false);
          }
        } catch (error) {
          setLoading(false);
          Message.loading({content:"提交中",id:"save",duration:2})
        }
      };
      return { 
        registerModal, 
        getTitle, 
        handleSubmit,
        formRef,
        loading,
        formData,
        parntList,
        t,
        OYoptions:[
          { label: '否', value: 0 },
          { label: '是', value: 1 },
        ],
      };
    },
  });
</script>
