<template>
  <Modal v-bind="getBindValue"  @cancel="handleCancel" :unmount-on-close="true" :maskClosable="getProps.maskClosable" :fullscreen="fullscreenVal" :body-style="{padding:0}" modal-class="self-modal">
    <template #title v-if="!$slots.title">
      <ModalHeader
        :helpMessage="getProps.helpMessage"
        :title="getMergeProps.title"
        :canFullscreen="getProps.canFullscreen"
        :fullscreenVal="fullscreenVal"
        @fullscreen="handleFullscreens"
        @fullexit="handleFullscreens"
        @dblclick="handleTitleDbClick"
      />
    </template>
    <!--滚动区-->
    <ModalWrapper
      :useWrapper="getProps.useWrapper"
      :footerOffset="wrapperFooterOffset"
      :fullScreen="fullscreenVal"
       ref="modalWrapperRef"
      :loading="getProps.loading"
      :loading-tip="getProps.loadingTip"
      :isPadding="getProps.isPadding"
      :minHeight="getProps.minHeight"
      :height="getWrapperHeight"
      :visible="visibleRef"
      :modalFooterHeight="footer !== undefined && !footer ? 0 : undefined"
      v-bind="omit(getProps.wrapperProps, 'visible', 'height', 'modalFooterHeight')"
      @ext-height="handleExtHeight"
      @height-change="handleHeightChange"
    >
      <slot></slot>
    </ModalWrapper>
    <template #footer v-if="!$slots.footer">
      <ModalFooter v-bind="getBindValue"  :loading="getProps.loading" @ok="handleOk" @cancel="handleCancel">
        <template #[item]="data" v-for="item in Object.keys($slots)">
          <slot :name="item" v-bind="data || {}"></slot>
        </template>
      </ModalFooter>
    </template>
  </Modal>
</template>
<script lang="ts">
  import type { ModalProps, ModalMethods } from './typing';
  import {
    defineComponent,
    computed,
    ref,
    watch,
    unref,
    watchEffect,
    getCurrentInstance,
    nextTick,
  } from 'vue';
  import Modal from './components/Modal';
  import ModalWrapper from './components/ModalWrapper.vue';
  import ModalFooter from './components/ModalFooter.vue';
  import ModalHeader from './components/ModalHeader.vue';
  import { isFunction } from '/@/utils/is';
  import { deepMerge } from '/@/utils';
  import { basicProps } from './props';
  import { omit } from 'lodash-es';
  declare type Recordable<T = any> = Record<string, T>;
  export default defineComponent({
    name: 'BasicModal',
    components: { Modal, ModalFooter, ModalHeader ,ModalWrapper},
    inheritAttrs: false,
    props: basicProps,
    emits: ['visible-change', 'height-change', 'cancel', 'ok', 'register', 'update:visible'],
    setup(props, { emit, attrs }) {
      const visibleRef = ref(false);
      const propsRef = ref<Partial<ModalProps> | null>(null);
      const modalWrapperRef = ref<any>(null);
      const fullscreenVal =ref(false);

      // modal   Bottom and top height
      const extHeightRef = ref(0);
      const modalMethods: ModalMethods = {
        setModalProps,
        emitVisible: undefined,
        redoModalHeight: () => {
          nextTick(() => {
            if (unref(modalWrapperRef)) {
              (unref(modalWrapperRef) as any).setModalHeight();
            }
          });
        },
      };

      const instance = getCurrentInstance();
      if (instance) {
        emit('register', modalMethods, instance.uid);
      }

      // Custom title component: get title
      const getMergeProps = computed((): Recordable => {
        return {
          ...props,
          ...(unref(propsRef) as any),
        };
      });


      // modal component does not need title and origin buttons
      const getProps = computed((): Recordable => {
        const opt = {
          ...unref(getMergeProps),
          visible: unref(visibleRef),
          okButtonProps: undefined,
          cancelButtonProps: undefined,
          title: undefined,
        };
        return {
          ...opt,
        };
      });

      const getBindValue = computed((): Recordable => {
        const attr = {
          ...attrs,
          ...unref(getMergeProps),
          visible: unref(visibleRef),
        };
        if (unref(fullscreenVal)) {
          return omit(attr, ['height', 'title']);
        }
        return omit(attr, 'title');
      });

      const getWrapperHeight = computed(() => {
        if (unref(fullscreenVal)) return undefined;
        return unref(getProps).height;
      });

      watchEffect(() => {
        visibleRef.value = !!props.visible;
        fullscreenVal.value = !!props.defaultFullscreen;
      });

      watch(
        () => unref(visibleRef),
        (v) => {
          emit('visible-change', v);
          emit('update:visible', v);
          instance && modalMethods.emitVisible?.(v, instance.uid);
          nextTick(() => {
            if (props.scrollTop && v && unref(modalWrapperRef)) {
              (unref(modalWrapperRef) as any).scrollTop();
            }
          });
        },
        {
          immediate: false,
        },
      );

      // 取消事件
      async function handleCancel(e: Event) {
        e?.stopPropagation();
        // 过滤自定义关闭按钮的空白区域
        if (e&&(e.target as HTMLElement)?.classList?.contains('-close--custom')) return;
        if (props.closeFunc && isFunction(props.closeFunc)) {
          const isClose: boolean = await props.closeFunc();
          visibleRef.value = !isClose;
          return;
        }
        visibleRef.value = false;
        fullscreenVal.value = false;
        emit('cancel', e);
      }

      /**
       * @description: 设置modal参数
       */
      function setModalProps(props: Partial<ModalProps>): void {
        // Keep the last setModalProps
        propsRef.value = deepMerge(unref(propsRef) || ({} as any), props);
        if (Reflect.has(props, 'visible')) {
          visibleRef.value = !!props.visible;
        }
        if (Reflect.has(props, 'defaultFullscreen')) {
          fullscreenVal.value = !!props.defaultFullscreen;
        }
      }
      function handleOk(e: Event) {
        fullscreenVal.value = false;
        emit('ok', e);
      }

      function handleHeightChange(height: string) {
        emit('height-change', height);
      }

      function handleExtHeight(height: number) {
        extHeightRef.value = height;
      }

      function handleTitleDbClick(isFull:boolean) {
        fullscreenVal.value=isFull
      }
     //全屏
     function handleFullscreens(isFull:boolean){
        fullscreenVal.value=isFull
     }
      return {
        handleCancel,
        getBindValue,
        getProps,
        handleFullscreens,
        fullscreenVal,
        getMergeProps,
        handleOk,
        visibleRef,
        omit,
        modalWrapperRef,
        handleExtHeight,
        handleHeightChange,
        handleTitleDbClick,
        getWrapperHeight,
      };
    },
  });
</script>
<style lang="less" >
.self-modal{
  .arco-modal-footer{
    padding: 10px 20px;
  }
}
</style>