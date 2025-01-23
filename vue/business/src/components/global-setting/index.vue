<template>
  <div v-if="!appStore.navbar" class="fixed-settings" @click="setVisible">
    <a-button type="primary">
      <template #icon>
        <icon-settings />
      </template>
    </a-button>
  </div>
  <a-drawer
    :width="300"
    unmount-on-close
    :visible="visible"
    :cancel-text="$t('settings.resetSetting')"
    :ok-text="$t('settings.saveSetting')"
    @cancel="cancel"
  >
    <template #title> {{ $t('settings.title') }} </template>
    <Block :options="contentOpts" />
    <a-divider orientation="center">{{ $t('settings.sysTheme') }}</a-divider>
    <a-row justify="center">
      <a-color-picker v-model="appStore.themeColor" format="hex" @change="changeColor" hideTrigger showPreset :preset-colors="defaultColorList"/>
    </a-row>
    <!-- <a-row justify="center" style="padding-top: 5px;user-select: none;">
      <a-link @click="handleCopySettings">{{ $t("settings.copySettings") }}</a-link>
    </a-row> -->
    <template #footer>
      <a-space>
        <a-link @click="handleCopySettings">{{ $t("settings.copySettings") }}</a-link>
        <a-button @click="handleReset">{{ $t('settings.resetSetting') }}</a-button>
        <a-button type="primary" @click="handleSaveSettings">{{ $t('settings.saveSetting') }}</a-button>
      </a-space>
    </template>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { computed,ref,onMounted, nextTick} from 'vue';
  import { Message } from '@arco-design/web-vue';
  import { useI18n } from 'vue-i18n';
  import { useClipboard } from '@vueuse/core';
  import { useAppStore } from '@/store';
  import Block from './block.vue';
  import defaultSettings from '@/config/settings.json';
  import { cloneDeep } from 'lodash-es';

  const emit = defineEmits(['cancel']);
  const appStore = useAppStore();
  const { t } = useI18n();
  const { copy } = useClipboard();
  const visible = computed(() => appStore.globalSettings);
  const contentOpts = computed(() => [
    { name: 'settings.navbar', key: 'navbar', defaultVal: appStore.navbar },
    {
      name: 'settings.menu',
      key: 'menu',
      defaultVal: appStore.menu,
    },
    {
      name: 'settings.topMenu',
      key: 'topMenu',
      defaultVal: appStore.topMenu,
    },
    { name: 'settings.footer', key: 'footer', defaultVal: appStore.footer },
    { name: 'settings.breadcrumb', key: 'breadcrumb', defaultVal: appStore.breadcrumb },
    { name: 'settings.tabBar', key: 'tabBar', defaultVal: appStore.tabBar },
    { name: 'settings.menuDark', key: 'menuDark', defaultVal: appStore.menuDark },
    { name: 'settings.navBg', key: 'navBg', defaultVal: appStore.navBg },
    { name: 'settings.menuAccordion', key: 'menuAccordion', defaultVal: appStore.menuAccordion },
    {
      name: 'settings.colorWeak',
      key: 'colorWeak',
      defaultVal: appStore.colorWeak,
    },
    {
      name: 'settings.menuWidth',
      key: 'menuWidth',
      defaultVal: appStore.menuWidth,
      type: 'number',
    },
  ]);
  //组件挂载完成后执行的函数
  onMounted(()=>{
    nextTick(()=>{
      if(defaultSettings.themeColor!=appStore.themeColor){
        changeColor(appStore.themeColor)
      }
    })
  })
  //关闭
  const cancel = () => {
    appStore.updateSettings({ globalSettings: false });
    emit('cancel');
  };
  //恢复设置
  const handleReset = () => {
    localStorage.setItem("settingsval", JSON.stringify(defaultSettings));
    localStorage.removeItem("settingsval")
    Message.success(t('settings.resetSetting.message'));
    location.reload()
  };
  //保存设置
  const defaultSetData=ref(cloneDeep(defaultSettings))
  const handleSaveSettings = async () => {
    defaultSetData.value=Object.assign(defaultSetData.value,{...appStore.$state,serverMenu:[],serverRoute:[],globalSettings:false,isDynamicAddedRoute:false})
    localStorage.setItem("settingsval", JSON.stringify(defaultSetData.value));
    Message.success(t('settings.saveSetting.message'));
  };
  //复制配置
  const handleCopySettings = async () => {
    defaultSetData.value=Object.assign(defaultSetData.value,{...appStore.$state,serverMenu:[],serverRoute:[],globalSettings:false,isDynamicAddedRoute:false})
    const text = JSON.stringify(defaultSetData.value, null, 2);
    await copy(text);
    Message.success({content:t('settings.copySettings.message'),id:"settings"});
  };
  const setVisible = () => {
    appStore.updateSettings({ globalSettings: true });
  };
  // 改变主题色
  const changeColor = (colorval:any ) => {
    if (!/^#[0-9A-Za-z]{6}/.test(colorval)) return
    appStore.setThemeColor(colorval)
  }
  const defaultColorList = [
  '#165DFF',
  '#409EFF',
  '#18A058',
  '#2d8cf0',
  '#007AFF',
  '#5ac8fa',
  '#5856D6',
  '#536dfe',
  '#9c27b0',
  '#AF52DE',
  '#0096c7',
  '#00C1D4',
  '#43a047',
  '#e53935',
  '#f4511e',
  '#6d4c41',
  '#C396ED',
  '#F7BA1E',
  '#FF7D00',
  '#F5319D',
  '#722ED1',
  '#12D2AC',
]
</script>

<style scoped lang="less">
  .fixed-settings {
    position: fixed;
    top: 280px;
    right: 0;
    svg {
      font-size: 18px;
      vertical-align: -4px;
    }
  }
</style>
