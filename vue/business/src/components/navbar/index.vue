<template>
  <div class="navbar" :style="{ backgroundImage:(theme === 'dark'||!appStore.navBg)?'':'url(https://cms.goflys.cn/navbarbg.jpg) ', backgroundSize:'100% 60px'}">
    <div class="left-side">
      <a-space>
        <img
          alt="logo"
          src="/logo.png"
          style="height: 33px;overflow: hidden;"
        />
        <div class="sitetitle" v-if="appStore.device === 'desktop'">
          {{AppTitle}}
        </div>
        <icon-menu-fold
          v-if="!topMenu && appStore.device === 'mobile'"
          style="font-size: 22px; cursor: pointer"
          @click="toggleDrawerMenu"
        />
      </a-space>
    </div>
    <div class="center-side">
      <Menu v-if="topMenu" />
    </div>
    <ul class="right-side">
      <li>
        <a-tooltip :content="$t('settings.language')">
          <a-button
            class="nav-btn"
            type="outline"
            :shape="'circle'"
            @click="setDropDownVisible"
          >
            <template #icon>
              <icon-language />
            </template>
          </a-button>
        </a-tooltip>
        <a-dropdown trigger="click" @select="changeLocale as any">
          <div ref="triggerBtn" class="trigger-btn"></div>
          <template #content>
            <a-doption
              v-for="item in locales"
              :key="item.value"
              :value="item.value"
            >
              <template #icon>
                <icon-check v-show="item.value === currentLocale" />
              </template>
              {{ item.label }}
            </a-doption>
          </template>
        </a-dropdown>
      </li>
      <li>
        <a-tooltip
          :content="
            theme === 'light'
              ? $t('settings.navbar.theme.toDark')
              : $t('settings.navbar.theme.toLight')
          "
        >
          <a-button
            class="nav-btn"
            type="outline"
            :shape="'circle'"
            @click="handleToggleTheme"
          >
            <template #icon>
              <icon-moon-fill v-if="theme === 'dark'" />
              <icon-sun-fill v-else />
            </template>
          </a-button>
        </a-tooltip>
      </li>
      <li>
        <a-tooltip :content="$t('settings.navbar.alerts')">
          <div class="message-box-trigger">
            <a-badge :count="9" dot>
              <a-button
                class="nav-btn"
                type="outline"
                :shape="'circle'"
                @click="setPopoverVisible"
              >
                <icon-notification />
              </a-button>
            </a-badge>
          </div>
        </a-tooltip>
        <a-popover
          trigger="click"
          :arrow-style="{ display: 'none' }"
          :content-style="{ padding: 0, width: '420px' }"
          content-class="message-popover"
        >
          <div ref="refBtn" class="ref-btn"></div>
          <template #content>
             <message-box />
          </template>
        </a-popover>
      </li>
      <li>
        <a-tooltip
          :content="
            isFullscreen
              ? $t('settings.navbar.screen.toExit')
              : $t('settings.navbar.screen.toFull')
          "
        >
          <a-button
            class="nav-btn"
            type="outline"
            :shape="'circle'"
            @click="toggleFullScreen"
          >
            <template #icon>
              <icon-fullscreen-exit v-if="isFullscreen" />
              <icon-fullscreen v-else />
            </template>
          </a-button>
        </a-tooltip>
      </li>
      <li>
        <a-tooltip :content="$t('settings.title')">
          <a-button
            class="nav-btn"
            type="outline"
            :shape="'circle'"
            @click="setVisible"
          >
            <template #icon>
              <icon-settings />
            </template>
          </a-button>
        </a-tooltip>
      </li>
      <li>
        <!-- 管理员账户 -->
        <a-dropdown trigger="hover">
          <a-row align="center" :wrap="false" class="user">
            <!-- 管理员头像 -->
            <a-avatar :size="32" :imageUrl="GetFullPath(userStore.avatar)"></a-avatar>
            <span class="username">{{ userStore.name }}</span>
            <icon-down />
          </a-row>
          <template #content>
            <a-doption @click="$router.push({ path: '/usersetting' })">
              <template #icon>
                <icon-user />
              </template>
              <template #default>{{ $t('messageBox.userCenter') }}</template>
            </a-doption>
            <a-divider :margin="1" />
            <a-doption @click="handleLogout">
              <template #icon>
                <icon-export />
              </template>
              <template #default> {{ $t('messageBox.logout') }}</template>
            </a-doption>
          </template>
        </a-dropdown>
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, inject } from 'vue';
  import { useDark, useToggle, useFullscreen } from '@vueuse/core';
  import { useAppStore, useUserStore } from '@/store';
  import { LOCALE_OPTIONS } from '@/locale';
  import useLocale from '@/hooks/locale';
  import useUser from '@/hooks/user';
  import Menu from '@/components/menu/index.vue';
  import MessageBox from '../message-box/index.vue';
  import { GetFullPath } from "@/utils/tool";
  const appStore = useAppStore();
  const userStore = useUserStore();
  const { logout } = useUser();
  const { changeLocale, currentLocale } = useLocale();
  const { isFullscreen, toggle: toggleFullScreen } = useFullscreen();
  const locales = [...LOCALE_OPTIONS];
  const theme = computed(() => {
    return appStore.theme;
  });
  const topMenu = computed(() => appStore.topMenu && appStore.menu);
  const isDark = useDark({
    selector: 'body',
    attribute: 'arco-theme',
    valueDark: 'dark',
    valueLight: 'light',
    storageKey: 'arco-theme',
    onChanged(dark: boolean) {
      // overridden default behavior
      appStore.toggleTheme(dark);
    },
  });
  const toggleTheme = useToggle(isDark);
  const handleToggleTheme = () => {
    toggleTheme();
  };
  const setVisible = () => {
    appStore.updateSettings({ globalSettings: true });
  };
  const refBtn = ref();
  const triggerBtn = ref();
  const setPopoverVisible = () => {
    const event = new MouseEvent('click', {
      view: window,
      bubbles: true,
      cancelable: true,
    });
    refBtn.value.dispatchEvent(event);
  };
  const handleLogout = () => {
    logout();
  };
  const setDropDownVisible = () => {
    const event = new MouseEvent('click', {
      view: window,
      bubbles: true,
      cancelable: true,
    });
    triggerBtn.value.dispatchEvent(event);
  };
  const toggleDrawerMenu = inject('toggleDrawerMenu') as () => void;
  //获取网站配置-应用名称
  const AppTitle = window?.globalConfig.AppTitle
</script>

<style scoped lang="less">
  .navbar {
    display: flex;
    justify-content: space-between;
    height: 100%;
    background-color: var(--color-bg-2);
    border-bottom: 1px solid var(--color-border);
  }

  .left-side {
    display: flex;
    align-items: center;
    padding-left: 20px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    .sitetitle{
         font-size: 18px;
         overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }
  }

  .center-side {
    flex: 1;
    :deep(.arco-menu-light){
      background-color:transparent;
      .arco-menu-pop-header,.arco-menu-item{
        background-color:transparent;
      }
      .arco-menu-pop-header,.arco-menu-item{
        &:hover{
          color: rgb(var(--primary-6));
          .arco-icon{
            color: rgb(var(--primary-6));
          }
        }
        &::after{
          bottom: -8px;
        }
      }
      .arco-menu-selected{
        &:hover{
          background-color:transparent;
        }
      }
      .arco-menu-selected-label{
        bottom: -8px;
      }
    }
    :deep(.arco-menu-inner){
      padding: 0px 20px;
      height: 49px;
      .arco-menu-overflow-wrap{
        height: 28px;
      }
    }
  }
  .right-side {
    display: flex;
    list-style: none;
    :deep(.locale-select) {
      border-radius: 20px;
    }
    li {
      display: flex;
      align-items: center;
      padding: 0 10px;
    }
    a {
      color: var(--color-text-1);
      text-decoration: none;
    }
    .nav-btn {
      border-color: rgb(var(--gray-2));
      color: rgb(var(--gray-8));
      font-size: 16px;
    }
    .trigger-btn,
    .ref-btn {
      position: absolute;
      bottom: 14px;
    }
    .trigger-btn {
      margin-left: 14px;
    }
  }
  .arco-dropdown-open .arco-icon-down {
    transform: rotate(180deg);
  }

.user {
  cursor: pointer;
  color: var(--color-text-1);
  .username {
    margin-left: 8px;
    white-space: nowrap;
  }
  .arco-icon-down {
    transition: all 0.3s;
    margin-left: 2px;
  }
}
</style>