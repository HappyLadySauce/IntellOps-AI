<template>
  <a-config-provider :locale="locale">
    <router-view v-if="isRouter"/>
    <global-setting />
  </a-config-provider>
</template>

<script lang="ts" setup>
  import { computed ,nextTick, provide, ref} from 'vue';
  import enUS from '@arco-design/web-vue/es/locale/lang/en-us';
  import zhCN from '@arco-design/web-vue/es/locale/lang/zh-cn';
  import GlobalSetting from '@/components/global-setting/index.vue';
  import useLocale from '@/hooks/locale';

  const { currentLocale } = useLocale();
  const locale = computed(() => {
    switch (currentLocale.value) {
      case 'zh-CN':
        return zhCN;
      case 'en-US':
        return enUS;
      default:
        return enUS;
    }
  });
  document.title=window?.globalConfig.AppTitle
  //刷新
  const isRouter = ref(true)
  const reload = () => {
    isRouter.value = false 
    nextTick(() => {isRouter.value = true})
  }
  provide("reload", reload)
</script>
<style lang="less">
::-webkit-scrollbar {
  width: 10px;
  height: 8px;
}
::-webkit-scrollbar-track {
  background-color: var(--color-neutral-3);
}
::-webkit-scrollbar-thumb {
  background-color: var(--color-neutral-4);
  box-shadow: inset 0 0 6px var(--color-neutral-5);
  border-radius: 6px;
}
::-webkit-scrollbar-thumb:hover {
  background-color: var(--color-neutral-5);
}
</style>
