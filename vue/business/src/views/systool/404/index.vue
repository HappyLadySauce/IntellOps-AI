<template>
  <div class="container">
    <Breadcrumb :items="['menu.exception', 'menu.exception.404']" />
    <page-card >
    <div class="content">
      <a-result
        class="result"
        status="404"
        :subtitle="$t('exception.result.404.description')"
      >
      </a-result>
      <div class="operation-row">
        <a-button key="again" style="margin-right: 16px" @click="reflast">
          {{ $t('exception.result.404.retry') }}
        </a-button>
        <a-button key="back" type="primary" @click="onback">
          {{ $t('exception.result.404.back') }}
        </a-button>
      </div>
    </div>
    </page-card>
  </div>
</template>

<script lang="ts" setup>
import { inject } from 'vue'
import { useRouter } from 'vue-router';
const router = useRouter();
const reload = inject("reload")
//刷新
const reflast=()=>{
  if (typeof reload == "function") reload();
}
//返回
const onback=()=>{
  router.go(-1)
}
</script>

<script lang="ts">
  export default {
    name: '404',
  };
</script>

<style scoped lang="less">
    .content {
      position: relative;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      height: 100%;
      text-align: center;
      background-color: var(--color-bg-1);
      border-radius: 4px;
    }
</style>
