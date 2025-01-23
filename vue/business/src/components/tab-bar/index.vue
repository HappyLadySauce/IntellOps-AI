<template>
  <div class="tab-bar-container">
      <div class="tab-bar-box">
        <a-tabs type="card-gutter" editable hide-content size="medium" :active-key="route.name" @tab-click="handleTabclick" @delete="handleDel">
          <a-tab-pane v-for="(tag, index) in tagList" :key="tag.name" :closable="index>0">
            <template #title>    
              <tab-item  :index="index" :item-data="tag"/>
              <svg class="tabs_chrome_left" height="7" width="7"><path data-v-bae60330="" d="M 0 7 A 7 7 0 0 0 7 0 L 7 7 Z"></path></svg>
              <svg class="tabs_chrome_right" height="7" width="7"><path data-v-bae60330="" d="M 0 0 A 7 7 0 0 0 7 7 L 0 7 Z"></path></svg>
            </template>
          </a-tab-pane>
          <template #extra>
            <a-dropdown trigger="hover">
              <icon-drag-dot-vertical />
              <template #content>
                <a-doption @click="tagClose">
                  <template #icon><icon-close /></template>
                  <template #default>关闭当前</template>
                </a-doption>
                <a-doption @click="tagCloseOthers">
                  <template #icon><icon-eraser /></template>
                  <template #default>关闭其他</template>
                </a-doption>
                <a-doption @click="tagCloseAll">
                  <template #icon><icon-minus /></template>
                  <template #default>关闭全部</template>
                </a-doption>
              </template>
            </a-dropdown>
          </template>
        </a-tabs>
      </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed, watch, onUnmounted } from 'vue';
  import type { RouteLocationNormalized } from 'vue-router';
  import type { TagProps } from '@/store/modules/tab-bar/types';
  import { useRouter, useRoute } from 'vue-router';
  import { DEFAULT_ROUTE_NAME } from '@/router/constants';
  import {listenerRouteChange,removeRouteListener} from '@/utils/route-listener';
  import { useAppStore, useTabBarStore } from '@/store';

  import tabItem from './tab-item.vue';
  const route = useRoute();
  const router = useRouter();
  const appStore = useAppStore();
  const tabBarStore = useTabBarStore();
  const affixRef = ref();
  const tagList = computed(() => {
    return tabBarStore.getTabList.filter((item, index, self) => index === self.findIndex(t => t.name === item.name));
  });
  const offsetTop = computed(() => {
    return appStore.navbar ? 50 : 0;
  });
  watch(
    () => appStore.navbar,
    () => {
      affixRef.value.updatePosition();
    }
  );
  listenerRouteChange((route: RouteLocationNormalized) => {
    if (
      !route.meta.noAffix &&
      !tagList.value.some((tag) => tag.fullPath === route.fullPath)
    ) {
      tabBarStore.updateTabList(route);
    }
  }, true);
  //关闭点击删除的标签
  const handleDel=(pathname: any)=>{
   tagList.value.forEach((tag:TagProps,idx:number)=>{
      if(tag.name==pathname&&DEFAULT_ROUTE_NAME!=tag.name){
          tabBarStore.deleteTag(idx, tag);
          if(tag.name==route.name){
            const latest = tagList.value[idx - 1]; // 获取队列的前一个tab
            router.push({ name: latest.name });
          }
      }
    })
  }
  //关闭当前
  const tagClose = () => {
    tagList.value.forEach((tag:TagProps,idx:number)=>{
      if(tag.fullPath==route.fullPath&&DEFAULT_ROUTE_NAME!=tag.name){
          tabBarStore.deleteTag(idx, tag);
          const latest = tagList.value[idx - 1]; // 获取队列的前一个tab
          router.push({ name: latest.name });
      }
    })
  };
  //关闭其他
  const tagCloseOthers = () => {
    tagList.value.forEach((tag:TagProps,index:number)=>{
      if(tag.fullPath==route.fullPath){
        const filterList = tagList.value.filter((el, idx) => {
          return idx === 0 || idx ===index;
        });
        tabBarStore.freshTabList(filterList);
        router.push({ name: tag.name });
      }
    })
  };
  //关闭全部
  const tagCloseAll = () => {
    tabBarStore.resetTabList();
      router.push({ name: DEFAULT_ROUTE_NAME });
  };
  //切换tab
  const handleTabclick = (pathname:any) => {
    const data=tagList.value.find((item)=>item.name==pathname)
    if(data)
    router.push({...data});
  };
  onUnmounted(() => {
    removeRouteListener();
  });
</script>

<style scoped lang="less">
:deep(.arco-tabs-nav-ink) {
   display: none;
}
:deep(.arco-tabs-nav::before) {
   display: none;
}
:deep(.arco-tabs-nav-tab) {
  padding-left: 3px;
  .arco-tabs-tab {
    position: relative;
    border:0;
    border-radius: 10px 10px 0 0;
    background-color: var(--color-bg-2);
    user-select: none;
    border-bottom-color: transparent !important;
    .arco-tabs-tab-title::before {
      right:0px;
      left:0px;
    }
    &:hover{
      .tabs_chrome_left, .tabs_chrome_right{
        display: unset;
      }
      background-color: var(--color-neutral-2);
      &:before, & + div:before {
        background-color:transparent !important;
      }
    }
    &:not(:first-of-type):before {
        position: absolute;
        top: 50%;
        left: 0px;
        display: block;
        width: 2px;
        height: 16px;
        background-color: rgb(var(--arcoblue-1));
        transform: translateY(-50%);
        content: "";
    }
  }
  //圆角
  .tabs_chrome_left,.tabs_chrome_right{
    display: none;
    position: absolute;
    fill: var(--color-fill-2);
  }
  .tabs_chrome_left{
      left: -7px;
      bottom: 0px;
    }
  .tabs_chrome_right{
    right: -7px;
    bottom: 0px;
  }
  .arco-icon-hover:hover::before{
    background-color: var(--color-fill-3) !important;
  }
}
//选中
:deep(.arco-tabs-nav-type-card-gutter){
  .arco-tabs-tab-active {
    background-color: var(--color-fill-2);
    padding-bottom: 4px !important;
    &:before {
      background-color: transparent !important;
    }
    .tabs_chrome_left, .tabs_chrome_right{
      display: unset;
    }
    .arco-tabs-tab-close-btn{
       color: var(--color-text-2);
    }
    &:hover{
      border-radius: 10px 10px 0 0;
    }
    & + div:before{
      background-color: transparent !important;
    }
  }
}
//导航框
.tab-bar-container {
  .tab-bar-box{
    padding: 5px 5px 0px 5px;
    margin-left: 1px;
    background-color: var(--color-bg-2);
    border-bottom: solid 1px  var(--color-fill-2);
  }
}
</style>
