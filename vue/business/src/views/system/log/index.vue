<template>
  <div class="container" >
      <div class="flex-page" :class="{ 'gf-table--fullscreen': isFullscreen}" :style="isFullscreen?'height:100%;':''">
        <a-row justify="space-between" align="center" class="header">
          <a-space wrap>
            <slot name="custom-title">
              <div class="title">系统日志</div>
            </slot>
          </a-space>
        </a-row>
        <a-tabs v-model:active-key="activeKey" type="card-gutter" size="large">
          <a-tab-pane key="1" >
            <template #title>
              <icon-find-replace /> 操作日志
            </template>
          </a-tab-pane>
          <a-tab-pane key="2">
            <template #title>
              <icon-lock /> 登录日志
            </template>
          </a-tab-pane>
        </a-tabs>
        <keep-alive>
          <component :is="PaneMap[activeKey]" v-model="isFullscreen"/>
        </keep-alive>
      </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import OperationLog from './operation/index.vue'
import LoginLog from './login/index.vue'
import { useRoute,useRouter} from 'vue-router'
const route = useRoute(),router=useRouter()
const isFullscreen = ref(false);

const PaneMap: any= {
1: OperationLog,
2: LoginLog
}

const activeKey = ref('1')
watch(
() => route.query,
() => {
  if (route.query.tabKey) {
    activeKey.value = String(route.query.tabKey)
  }
},
{ immediate: true }
)
//更新路由参数
const changeRouter = (key: string | number) => {
activeKey.value = key as string
router.replace({ path: route.path, query: { tabKey: key } })
}
</script>
<script lang="ts">
export default {
  name: 'log',//页签是名字和路由name相同则缓存生效
};
</script>
<style lang="less" scoped>
:deep(.arco-card-body){
padding: 0px;
}
:deep(.arco-tabs .arco-tabs-nav-type-card-gutter .arco-tabs-tab-active) {
box-shadow: inset 0 2px 0 rgb(var(--primary-6)), inset -1px 0 0 var(--color-border-2),
  inset 1px 0 0 var(--color-border-2);
   position: relative;
}

:deep(.arco-tabs-nav-type-card-gutter .arco-tabs-tab) {
border-radius: var(--border-radius-medium) var(--border-radius-medium) 0 0;
}

:deep(.arco-tabs-type-card-gutter > .arco-tabs-content) {
border: none;
}

:deep(.arco-tabs-nav::before) {
left: -20px;
right: -20px;
}

:deep(.arco-tabs) {
overflow: visible;
}

:deep(.arco-tabs-nav) {
overflow: visible;
}
.container{
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  position: relative;
  padding:14px;
}
// 表格页面
.flex-page {
  height: 100%;
  overflow: hidden;
  background: var(--color-bg-1);
  padding: 16px;
  padding-bottom: 0;
  border-radius: 4px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  position: relative;
  .header {
    padding: 0 0 10px;
    .title {
      color: var(--color-text-1);
      font-size: 18px;
      font-weight: 500;
      line-height: 1.5;
    }
  }
  &:after {
    content: "";
    height: 13px;
    font-size: 12px;
    color: var(--color-text-3);
    box-sizing: border-box;
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>
