<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" :loading="loading" width="1000px" @height-change="onHeightChange" :minHeight="modelHeight" :title="getTitle" @ok="handleSubmit">
    <div class="addFormbox" :style="{'min-height':`${windHeight}px`}">
      <div class="tabs-header" v-if="isEditor">
        <div class="tabs-nav-wrap">
            <div class="tap_item" v-for="iten in tapList" :class="{item_active:activeKey==iten.id}" @click="()=>{activeKey=iten.id}">
                <div class="label">{{iten.name}}</div>
            </div>
        </div>
        <div class="tabs-bar" :style="{top: `${(activeKey-1)*64}px`,height: `64px`}"></div>
      </div>
      <div class="tabs-content" :class="{addpadding:!isEditor}">
        <a-form ref="formRef" :model="formData" auto-label-width>
          <div class="content_box">
              <!--基础信息-->
              <a-scrollbar v-show="activeKey==1" style="overflow: auto;" :style="{height:`${windHeight}px`}">
                <div class="besecontent" >
                  <a-row :gutter="16">
<!--replaceTpl-->  
                  </a-row>
                </div>
              </a-scrollbar>
              <!--编辑器-->
<!--replaceEditHtml-->               
          </div>
        </a-form>
      </div>
    </div>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref} from 'vue';
  import { BasicModal, useModalInner} from '/@/components/Modal';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import { cloneDeep } from 'lodash-es';
  //api
  import { save,getContent } from './api';
  import { Message} from '@arco-design/web-vue';
  export default defineComponent({
    name: 'AddForm',
    components: { BasicModal},
    emits: ['success'],
    setup(_, { emit }) {
      const visibleimage=ref(false);
      //判断是否存在编辑器
      const isEditor=ref(true);
      const isUpdate = ref(false);
      const activeKey= ref(1);
      const modelHeight= ref(620);
      const windHeight= ref(620);
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
          activeKey.value=1
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
        tapList:[
          {id:1,name:"基础内容"},{replacEdit:null}
        ],
        activeKey,
        modelHeight,
        onHeightChange,windHeight,
        isEditor,
        visibleimage,
      };
    },
  });
</script>
<style lang="less" scoped>
  @import '@/assets/style/formlayer.less';
</style>

