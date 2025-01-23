<template>
    <BasicModal v-bind="$attrs" @register="registerModal" :footer="false" :isPadding="false" width="1000px" @height-change="onHeightChange" :minHeight="modelHeight" :title="getTitle" @before-close="handleClose">
      <div class="addFormbox" :style="{'min-height':`${windHeight}px`}">
        <div class="RichText" >
          <div ref="containerview"></div>
        </div>
      </div>
    </BasicModal>
  </template>
  <script lang="ts">
    import { defineComponent, ref} from 'vue';
    import { BasicModal, useModalInner } from '/@/components/Modal';
    import Engine,{ EngineInterface }from '@/components/gfeditor/emain';
    import { cards, plugins, pluginConfig } from "@/components/Editor/config";
    export default defineComponent({
      name: 'ShowPrivateContent',
      components: { BasicModal },
      emits: ['success'],
      setup(_, { emit }) {
        const loading=ref(true);
        const visibleimage=ref(false);
        const winHeight=document.documentElement.clientHeight-50
        const modelHeight= ref(winHeight);
        const windHeight= ref(winHeight);
        const getTitle = ref("查看插件详情")
        // 编辑器容器  编辑器引擎
        const containerview = ref<HTMLElement | null>(null);
        // 编辑器引擎
        const engine = ref<EngineInterface | null>(null);
        const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          getTitle.value=data.record.title
          if (containerview.value) {
              //实例化引擎
              engine.value  = new Engine(containerview.value, {
                // 启用的插件
                plugins,
                // 启用的卡片
                cards,
                readonly:true,
                autoPrepend:false,
                lazyRender:false,
                // 所有的卡片配置
                config: pluginConfig,
              });
            }
            if(engine.value){
              engine.value.setHtml(data.record.content, () => {
                loading.value = false;
              });
            }
        });
         //监听高度
         const onHeightChange=(val:any)=>{
           windHeight.value=val
        }
        //关闭界面销毁编辑器
         const handleClose=()=>{
          if (engine.value) engine.value.destroy();
        }
        return { 
          registerModal, 
          getTitle, 
          modelHeight,
          onHeightChange,windHeight,
          visibleimage,
          containerview,
          handleClose,
        };
      },
    });
  </script>
  <style lang="less" scoped>
    @import '@/assets/style/formlayer.less';
    .RichText{
      width: 100%;
      padding: 10px 25px;
    }
  </style>
  
  