<template>
    <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" :loading="loading" width="800px" @height-change="onHeightChange" :minHeight="modelHeight" :title="getTitle" @ok="handleSubmit">
      <div class="addFormbox" :style="{'min-height':`${windHeight}px`}">
        <div class="tabs-content">
          <a-form ref="formRef" :model="formData" auto-label-width>
            <div class="content_box">
                <!--基础信息-->
                <a-scrollbar style="overflow: auto;" :style="{height:`${windHeight}px`}">
                  <div class="besecontent" >
                    <a-row :gutter="16">
                      <a-col :span="24">
                        <a-form-item field="title" label="名称" validate-trigger="input" :rules="[{required:true,message:'请填写名称'}]" >
                          <a-input v-model="formData.title" placeholder="请填名称" :max-length="50" allow-clear show-word-limit />
                        </a-form-item>
                      </a-col>
                      <a-col :span="12" >
                        <a-form-item field="db_way" label="数据存储表" tooltip="公共表是common_dictionary_data，单独建表是你自己创建的数据表(根据业务需求分表存储)">
                            <a-radio-group v-model="formData.db_way" >
                              <a-radio value="sys">公共表</a-radio>
                              <a-radio value="alone">单独建表</a-radio>
                            </a-radio-group>
                        </a-form-item>
                      </a-col>
                      <a-col :span="12" >
                        <a-form-item field="status" label="状态" >
                            <a-radio-group v-model="formData.status" :options="OYoptions" />
                        </a-form-item>
                      </a-col>
                      <a-col :span="24" v-if="formData.db_way=='alone'">
                        <a-form-item field="tablename" label="数据表名称" validate-trigger="input" :rules="[{required:true,message:'请填写名称'}]" >
                          <a-input v-model="formData.tablename" placeholder="这里是填写您自己新建的数据表，表结构参考common_dictionary_data表设计" :max-length="50" allow-clear show-word-limit />
                        </a-form-item>
                      </a-col>
                      <a-col :span="24">
                        <a-form-item field="weigh" label="排序" validate-trigger="input" style="margin-bottom:15px;">
                          <a-input-number :min="0" v-model="formData.weigh" placeholder="请填排序" />
                        </a-form-item>
                      </a-col>
                      <a-col :span="24">
                        <a-form-item field="remark" label="备注" style="margin-bottom:15px;">
                          <a-textarea v-model="formData.remark" placeholder="请填备注"  :max-length="200" allow-clear show-word-limit :auto-size="{minRows:3,maxRows:5}"/>
                        </a-form-item>
                      </a-col>
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
    import { save } from '@/api/datacenter/tabledata';
    import { Message } from '@arco-design/web-vue';
    export default defineComponent({
      name: 'AddMenu',
      components: { BasicModal },
      emits: ['success'],
      setup(_, { emit }) {
        const isUpdate = ref(false);
        const modelHeight= ref(350);
        const windHeight= ref(350);
        //表单
        const { loading, setLoading } = useLoading();
        const formRef = ref<FormInstance>();
        //表单字段
        const basedata={
              id:0,
              title: '',
              tablename: "",
              remark: "",
              status: 0,
              db_way:"sys",
              weigh:0,
          }
        const formData = ref(cloneDeep(basedata))
        //编辑器
        const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
            formRef.value?.resetFields()
            setLoading(true);
            setModalProps({ confirmLoading: false });
            isUpdate.value = !!data?.isUpdate;
            if (unref(isUpdate)) {
              formData.value=cloneDeep(data.record)
            }else{
              formData.value=cloneDeep(basedata)
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
              Message.loading({content:"提交中",id:"save",duration:0})
              let savedata=cloneDeep(unref(formData))
              await save(savedata);
              Message.success({content:"提交成功",id:"save",duration:2000})
              closeModal()
              emit('success');
              setLoading(false);
            }
          } catch (error) {
            setLoading(false);
            Message.loading({content:"提交中",id:"save",duration:1})
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
          OYoptions:[
            { label: '正常', value: 0 },
            { label: '禁用', value: 1 },
            ],
          modelHeight,
          onHeightChange,windHeight,
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
  
  