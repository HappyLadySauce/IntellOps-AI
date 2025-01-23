<template>
    <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" :loading="loading" width="600px" :minHeight="modelHeight" :title="getTitle" @ok="handleSubmit">
      <div class="addFormbox"  >
        <div class="tabs-content">
          <a-form ref="formRef" :model="formData" auto-label-width>
            <div class="content_box">
                <!--基础信息-->
                  <div class="besecontent" >
                    <a-row :gutter="16">
                      <a-col :span="22">
                        <a-form-item field="name" label="名称" validate-trigger="input" :rules="[{required:true,message:'请填写名称'}]" >
                          <a-input v-model="formData.name" placeholder="请填名称" :max-length="50" allow-clear show-word-limit />
                        </a-form-item>
                      </a-col>
                      <a-col :span="22">
                        <a-form-item field="pid" label="上级" validate-trigger="input" :rules="[{required:true,message:'请填写名称'}]" >
                            <a-tree-select placeholder="选择上级" :data="treeList"  v-model="formData.pid" allow-search
                            :fieldNames="{
                                key: 'id',
                                title: 'name',
                                children: 'children'
                                }"/>
                        </a-form-item>
                      </a-col>
                      <a-col :span="11">
                        <a-form-item field="weigh" label="排序" validate-trigger="input" style="margin-bottom:15px;">
                          <a-input-number  v-model="formData.weigh" placeholder="请填排序" />
                        </a-form-item>
                      </a-col>
                      <a-col :span="12" >
                        <a-form-item field="status" label="状态" style="margin-bottom:5px;">
                            <a-radio-group v-model="formData.status" :options="OYoptions" />
                        </a-form-item>
                        </a-col>
                      <a-col :span="22">
                        <a-form-item field="remark" label="备注" style="margin-bottom:15px;">
                          <a-textarea v-model="formData.remark" placeholder="请填备注"  :max-length="200" allow-clear show-word-limit :auto-size="{minRows:3,maxRows:5}"/>
                        </a-form-item>
                      </a-col>
                    </a-row>
                  </div>
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
    import {save, getTree} from './api/cate';
    import { Message } from '@arco-design/web-vue';
    export default defineComponent({
      name: 'CateForm',
      components: { BasicModal },
      emits: ['success'],
      setup(_, { emit }) {
        const isUpdate = ref(false);
        const modelHeight= ref(350);
        const treeList= ref<any []>([]);
        //表单
        const { loading, setLoading } = useLoading();
        const formRef = ref<FormInstance>();
        //表单字段
        const basedata={
            id:0,
            name: '',
            pid: 0,
            status: 0,
            weigh:1,
            remark: "",
          }
        const formData = ref(basedata)
        //编辑器
        const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
            formRef.value?.resetFields()
            setLoading(true);
            setModalProps({ confirmLoading: false });
            const mdata = await getTree({});
            const parntList_df : any=[{id: 0,name: "一级数据"}];
            if(mdata){
                treeList.value=parntList_df.concat(mdata)
            }else{
                treeList.value=parntList_df
            }
            isUpdate.value = !!data?.isUpdate;
            if (unref(isUpdate)) {
              if(data.record.children){
                delete data.record.children
              }
              formData.value=cloneDeep(data.record)
            }else{
              formData.value=cloneDeep(basedata)
              formData.value.pid=data.record.pid
            }
            setLoading(false);
        });
        const getTitle = computed(() => (!unref(isUpdate) ? '新增数据' : '编辑数据'));
       //点击保存数据
       const handleSubmit = async () => {
        try {
            const res = await formRef.value?.validate();
            if (!res) {
              setLoading(true);
              Message.loading({content:"提交中",id:"submit",duration:0})
              let savedata=cloneDeep(unref(formData))
              await save(savedata);
              Message.success({content:"提交成功",id:"submit",duration:2000})
              closeModal()
              emit('success');
              setLoading(false);
            }
          } catch (error) {
            setLoading(false);
            Message.loading({content:"提交中",id:"submit",duration:1})
          }
        };
        return { 
          registerModal, 
          getTitle, 
          handleSubmit,
          formRef,
          loading,
          formData,
          OYoptions:[
            { label: '正常', value: 0 },
            { label: '禁用', value: 1 },
            ],
          modelHeight,treeList,
        };
      },
    });
  </script>
  <style lang="less" scoped>
    @import '@/assets/style/formlayer.less';
    .tabs-content{
      padding: 0px 25px;
    }
  </style>
  
  