<template>
  <a-modal  v-model:visible="visibleModal" width="600px" title-align="start" :title="getTitle" @ok="handleSubmit">
    <a-form ref="formRef" :model="formData" auto-label-width style="padding:10px 20px;">
      <a-form-item field="name" label="名称" validate-trigger="input" :rules="[{required:true,message:'请填写名称'}]" >
        <a-input v-model="formData.name" placeholder="请填写名称" />
      </a-form-item>
      <a-form-item field="pid" label="上级" validate-trigger="input" :rules="[{required:true,message:'请填写名称'}]" >
          <a-tree-select placeholder="选择上级" :data="treeList"  v-model="formData.pid" allow-search
          :fieldNames="{
              key: 'id',
              title: 'name',
              children: 'children'
              }"/>
      </a-form-item>
      <a-form-item label="排序" field="weigh" style="margin-bottom:15px;">
        <a-input-number v-model="formData.weigh" placeholder="请填排序" />
      </a-form-item>
      <a-form-item field="des" label="备注" validate-trigger="input" style="margin-bottom:15px;">
        <a-textarea v-model="formData.des" placeholder="请填写备注" allow-clear/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref} from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import { cloneDeep } from 'lodash-es';
  //api
  import { save,getTree } from './api';
  import { Message } from '@arco-design/web-vue';
  export default defineComponent({
    name: 'AddForm',
    components: {  },
    emits: ['success'],
    setup(_, { emit }) {
      const visibleModal = ref(false);
      const isUpdate = ref(false);
      const parntList = ref([]);
      const treeList = ref([]);
      //表单
      const formRef = ref<FormInstance>();
      //表单字段
      const basedata={
            id:0,
            name: '',
            weigh: 1,
            des:"",
        }
      const formData = ref(basedata)
      const m_component=ref("")
      const ShowModal=async(data:any)=>{
        visibleModal.value=true
        isUpdate.value = !!data?.isUpdate;
        const mdata = await getTree({});
          const parntList_df : any=[{id: 0,name: "一级数据"}];
          if(mdata){
              treeList.value=parntList_df.concat(mdata)
          }else{
              treeList.value=parntList_df
          }
          if (unref(isUpdate)) {
            m_component.value=data.record.component
            formData.value=cloneDeep(data.record)
          }else{
            formData.value=cloneDeep(basedata)
          }
      }
      const getTitle = computed(() => (!unref(isUpdate) ? '新增' : '编辑'));
     //点击确认
     const { loading, setLoading } = useLoading();
     const handleSubmit = async () => {
      try {
          const res = await formRef.value?.validate();
          if (!res) {
            setLoading(true);
            Message.loading({content:"更新中",id:"upStatus",duration:0})
            await save(unref(formData));
            Message.success({content:"更新成功",id:"upStatus",duration:2000})
            emit('success');
            setLoading(false);
            visibleModal.value=false
          }
        } catch (error) {
          setLoading(false);
          Message.loading({content:"更新中",id:"upStatus",duration:1})
        }
      };
      return { 
        ShowModal, 
        getTitle, 
        handleSubmit,
        formRef,
        loading,
        formData,
        parntList,
        visibleModal,
        treeList,
      };
    },
  });
</script>
