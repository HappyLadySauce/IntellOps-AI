<template>
  <slot name="breadcrumb">
    <Breadcrumb v-if="breadcrumb"/>
  </slot>
  <a-card v-bind="attrs" class="general-card " :class="{ 'gf-table--fullscreen': isFullscreen,'flex-page': fullTable}" 
  :title="title" v-resize="onResize" :style="`${scrollPage&&!isFullscreen?'min-height':'height'}:${isFullscreen?'100%':'calc(100% - '+(isbreadcrumb&&pbreadcrumb?appStore.breadcrumbHeight:0)+'px)'}`">
  <div v-if="fullTable" class="custom-full-table">
    <div class="table-toolbar flex flex-between" v-if="slotsearchLeft">
      <div class="left"><slot name="searchLeft"></slot></div>
      <div class="right"><slot name="searchRight"></slot></div>
    </div> 
    <slot name="searchBar"></slot>
    <div class="table-container" v-if="slottable">
      <slot name="table"></slot>
    </div>
  </div> 
  <slot v-else></slot>
  </a-card>
</template>

<script lang="ts" setup>
  import {computed,nextTick,onMounted,ref,useAttrs,useSlots} from 'vue';
  import { useAppStore } from '@/store';
  const appStore = useAppStore();
  const attrs = useAttrs()
  const isbreadcrumb = computed(() => appStore.breadcrumb);
  const props = defineProps({
    isFullscreen: {
      type:Boolean,
      default() {
        return false;
      },
    },
    breadcrumb: {//是否显示Breadcrumb面包导航，设置为true
      type:Boolean,
      default() {
        return false;
      },
    },
    scrollPage: {//card可以超出可以滚动，并白色背景铺满页面
      type:Boolean,
      default() {
        return false;
      },
    },
    fullTable: {//表格充满框
      type:Boolean,
      default() {
        return false;
      },
    },
    title:{
      type:String,
      default() {
        return "";
      },
    },
  });
  const emits = defineEmits(['onheight'])
  const onResize = (dom: any) => {
    emits('onheight', dom.height)
  }
  const pbreadcrumb = ref(false) // 可以直接获取，不需要使用 toRef
  onMounted(()=>{
    nextTick(()=>{
      if(document.getElementsByClassName("arco-breadcrumb").length>0){
        pbreadcrumb.value=true
      }else{
        pbreadcrumb.value=false
      }
    })
  })
  //判断<slot name="search_left"/>是否有传值
  const slotsearchLeft = !!useSlots().searchLeft;
  const slottable = !!useSlots().table;
</script>

<style scoped lang="less">
 .general-card {
  :deep(.arco-card-header) {
    padding: 10px 16px;
  }
  :deep(.arco-card-body) {
      padding: 12px;
    }
  }
  //表格铺满卡片
  .flex-page{
    flex: 1;
    overflow: hidden;
    padding-bottom: 0;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    position: relative;
    .custom-full-table{
      flex: 1;
      display: flex;
      flex-direction: column;
      overflow: hidden;
      height: 100%;
      background: var(--color-bg-1);
      position: relative;
      box-sizing: border-box;
      .table-container{
        max-height: 100%;
        overflow: hidden;
        flex: 1;
      }
    }
    &:after {
      content: "";
      height:0px;
      font-size: 12px;
      color: var(--color-text-3);
      box-sizing: border-box;
      display: flex;
      justify-content: center;
      align-items: center;
    }
  }
  :deep(.arco-card-body){
    height: 100%;
  }
</style>
<style lang="less">
//全局table铺满卡片样式
.flex-page{
   // 控制table高度占满
   .arco-table-border .arco-table-container {
      height: 100%;
    }
     .arco-table-body  {
      height: 100%;
    }
     .arco-table-content .arco-scrollbar-type-embed:last-child {
      height: 100%;
    }
    //如果为空时，将表格铺满
     .arco-table-element:has(tbody .arco-table-tr-empty) {
      height: 100%;
    }
    // 控制表格最后一行的下边框显示
     .arco-table-border .arco-table-scroll-y .arco-table-body .arco-table-tr:last-of-type .arco-table-td,
      .arco-table-border .arco-table-scroll-y tfoot .arco-table-tr:last-of-type .arco-table-td  {
      border-bottom: 1px solid var(--color-neutral-3);
    }
    //表格边框
    .table-container{
      .arco-table-col-fixed-right-first{
        border-right:0px !important;
      }
      .arco-table-container{
        border-right: 1px solid var(--color-neutral-3);
      }
    }
  }
</style>
