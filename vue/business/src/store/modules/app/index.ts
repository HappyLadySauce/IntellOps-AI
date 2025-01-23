import { defineStore } from 'pinia';
import type { RouteRecordNormalized } from 'vue-router';
import { generate, getRgbStr } from '@arco-design/color'
import defaultSettings from '@/config/settings.json';
import { getMenuList } from '@/api/user';
import { AppState } from './types';
import { cloneDeep } from 'lodash-es';
import {Notification} from '@arco-design/web-vue';
import type { AppRouteRecordRaw } from '/@/router/types';
import { transformObjToRoute } from '/@/router/helper/routeHelper';
//获取本地保存配置
const settingsval=localStorage.getItem("settingsval");
const useAppStore = defineStore('app', {
  state: (): AppState => settingsval?({...JSON.parse(settingsval)}):({ ...defaultSettings }),
  getters: {
    appCurrentSetting(state: AppState): AppState {
      return { ...state };
    },
    appDevice(state: AppState) {
      return state.device;
    },
    appAsyncMenus(state: AppState): RouteRecordNormalized[] {
      return state.serverMenu as unknown as RouteRecordNormalized[];
    },
    appAsyncRoute(state: AppState): RouteRecordNormalized[] {
      return state.serverRoute as unknown as RouteRecordNormalized[];
    },
    //获取路由状态
    getIsDynamicAddedRoute():boolean{
      return this.isDynamicAddedRoute;
    },
    //获取主题
    getTheme():string{
      return this.theme;
    },
  },

  actions: {
    //设置动态路由加载状态
    setDynamicAddedRoute(added: boolean) {
      this.isDynamicAddedRoute = added;
    },
    // Update app settings
    updateSettings(partial: Partial<AppState>) {
      // @ts-ignore-next-line
      this.$patch(partial);
    },

    // Change theme color
    toggleTheme(dark: boolean) {
      if (dark) {
        this.theme = 'dark';
        document.body.setAttribute('arco-theme', 'dark');
      } else {
        this.theme = 'light';
        document.body.removeAttribute('arco-theme');
      }
    },
    toggleDevice(device: string) {
      this.device = device;
    },
    toggleMenu(value: boolean) {
      this.hideMenu = value;
    },
    async fetchServerMenuConfig(rouid:any) {//加载服务端菜单
      try {
        const menudata = await getMenuList(rouid);
        if(menudata){
          this.serverMenu = menudata;
          // 动态引入组件
          let routeList: AppRouteRecordRaw[] = cloneDeep(menudata);
          var routeLists = transformObjToRoute(routeList);
          this.serverRoute =routeLists 
        }else{//提示账号没有菜单权限
          Notification.warning({
            title: '温馨提示：',
            content: '您登录的账号没有菜单权限，请联系管理检查后再登录！',
            position: 'bottomRight'
          })
        }
      } catch (error) {
      }
    },
    clearServerMenu() {
      this.serverMenu = [];
    },
    // 设置主题色
   setThemeColor (color: string)  {
      if (!color) return
      this.themeColor = color
      const list = generate(this.themeColor, { list: true, dark: this.theme === 'dark' })
      list.forEach((color: string, index: number) => {
        const rgbStr = getRgbStr(color)
        document.body.style.setProperty(`--primary-${index + 1}`, rgbStr)
        document.body.style.setProperty(`--arcoblue-${index + 1}`, rgbStr)
      })
    }
  },
});

export default useAppStore;
