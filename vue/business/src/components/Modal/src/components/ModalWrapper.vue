<template>
  <a-spin :loading="loading" :tip="loadingTip?loadingTip:'加载中'" style="width: 100%;display:block" >
    <a-scrollbar ref="wrapperRef" style="overflow: auto;" :style="spinStyle">
      <div ref="spinRef" :class="{modelbody:isPadding}" >
        <slot></slot>
      </div>
    </a-scrollbar>
  </a-spin>
</template>
<script lang="ts">
  import type { CSSProperties } from 'vue';
  import {
    defineComponent,
    computed,
    ref,
    watchEffect,
    unref,
    watch,
    onMounted,
    nextTick,
    onUnmounted,
  } from 'vue';
  import { createModalContext } from '../hooks/useModalContext';
  import { useMutationObserver } from '@vueuse/core';

  const props = {
    loading: { type: Boolean },
    useWrapper: { type: Boolean, default: true },
    isPadding: { type: Boolean, default: true },
    modalHeaderHeight: { type: Number, default: 48 },
    modalFooterHeight: { type: Number, default: 53 },
    minHeight: { type: Number, default: 200 },
    height: { type: Number },
    footerOffset: { type: Number, default: 0 },
    visible: { type: Boolean },
    fullScreen: { type: Boolean },
    loadingTip: { type: String },
  };

  export default defineComponent({
    name: 'ModalWrapper',
    components: {  },
    inheritAttrs: false,
    props,
    emits: ['height-change', 'ext-height'],
    setup(props, { emit }) {
      const wrapperRef = ref<ComponentRef>(null);
      const spinRef = ref<ElRef>(null);
      const realHeightRef = ref(0);
      const minRealHeightRef = ref(0);

      let realHeight = 0;

      let stopElResizeFn: Fn = () => {};


      useMutationObserver(
        spinRef,
        () => {
          setModalHeight();
        },
        {
          attributes: true,
          subtree: true,
        },
      );

      createModalContext({
        redoModalHeight: setModalHeight,
      });

      const spinStyle = computed((): CSSProperties => {
        return {
          minHeight: `${props.minHeight}px`,
          [props.fullScreen ? 'height' : 'maxHeight']: `${unref(realHeightRef)}px`,
        };
      });

      watchEffect(() => {
        props.useWrapper && setModalHeight();
      });

      watch(
        () => props.fullScreen,
        (v) => {
          setModalHeight();
          if (!v) {
            realHeightRef.value = minRealHeightRef.value;
          } else {
            minRealHeightRef.value = realHeightRef.value;
          }
        },
      );

      onMounted(() => {
        const { modalHeaderHeight, modalFooterHeight } = props;
        emit('ext-height', modalHeaderHeight + modalFooterHeight);
      });

      onUnmounted(() => {
        stopElResizeFn && stopElResizeFn();
      });

      async function scrollTop() {
        nextTick(() => {
          const wrapperRefDom = unref(wrapperRef);
          if (!wrapperRefDom) return;
          (wrapperRefDom as any)?.scrollTo?.(0);
        });
      }

      async function setModalHeight() {
        // 解决在弹窗关闭的时候监听还存在,导致再次打开弹窗没有高度
        // 加上这个,就必须在使用的时候传递父级的visible
        if (!props.visible) return;
        const wrapperRefDom = unref(wrapperRef);
        if (!wrapperRefDom) return;

        const bodyDom = wrapperRefDom.$el.parentElement;
        if (!bodyDom) return;
        bodyDom.style.padding = '0';
        await nextTick();
        try {
          const modalDom = bodyDom.parentElement && bodyDom.parentElement.parentElement;
          if (!modalDom) return;

          const modalRect = getComputedStyle(modalDom as Element).top;
          const modalTop = Number.parseInt(modalRect);
          let maxHeight =
            window.innerHeight -
            modalTop * 2 +
            (props.footerOffset! || 0) -
            props.modalFooterHeight -
            props.modalHeaderHeight;

          // 距离顶部过进会出现滚动条
          if (modalTop < 40) {
            maxHeight -= 26;
          }
          await nextTick();
          const spinEl = unref(spinRef);

          if (!spinEl) return;
          await nextTick();
          if (!realHeight) {
             realHeight = spinEl.scrollHeight;
          }
          if (props.fullScreen) {
            var minheigt=0
            if(props.modalFooterHeight==0){
              minheigt=6
            }
            realHeightRef.value =
              window.innerHeight - props.modalFooterHeight - props.modalHeaderHeight-minheigt;
          } else {
            realHeightRef.value = props.height
              ? props.height
              :props.minHeight;
          }
          emit('height-change', unref(realHeightRef));
        } catch (error) {
          console.error(error);
        }
      }

      return { wrapperRef, spinRef, spinStyle, scrollTop, setModalHeight };
    },
  });
</script>
<style scoped>
.modelbody{
  padding: 15px;
}
</style>
