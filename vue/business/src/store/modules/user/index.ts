import { defineStore } from 'pinia';
import router from '@/router';
import type { RouteRecordRaw} from 'vue-router';
import {
  login as userLogin,
  logout as userLogout,
  getUserInfo,
  // LoginData,
} from '@/api/user';
import { setToken,getToken, clearToken } from '@/utils/auth';
import { removeRouteListener } from '@/utils/route-listener';
import { UserState} from './types';
import useAppStore from '../app';

const useUserStore = defineStore('user', {
  state: (): UserState => ({
    name: undefined,
    nickname: undefined,
    mobile: undefined,
    email: undefined,
    avatar: undefined,
    introduction: undefined,
    id: 0,
    business_id: 0,
    rooturl: '',
    localurl: '',
    city: '',
    company: '',
    sessionTimeout: false,//登录是否已过期
    btnroles:[],//页面对应按钮权限
    createtime:"",
    pwd_reset_time:"",
  }),

  getters: {
    userInfo(state: UserState): UserState {
      return { ...state };
    },
  },

  actions: {
    setTokenData(info: string | undefined) {
      setToken(info);
    },
    // Set user's information
    setInfo(partial: Partial<UserState>) {
      this.$patch(partial);
    },
    // Reset user's information
    resetInfo() {
      this.$reset();
    },

    // Get user's information/获取用户信息
    async info() {
      const res = await getUserInfo();
      this.setInfo(res);
    },
    //设置token
    async setTokenArr(token:any) {
      setToken(token);
    },
    // 登录/Login
    async login(loginForm: any) {
      try {
        const res = await userLogin(loginForm);
        setToken(res);
        const appStore = useAppStore();
        appStore.setDynamicAddedRoute(false);//重注册路由
        return this.afterLoginAction();
      } catch (err) {
        clearToken();
        throw err;
      }
    },
    //登录成功后加载菜单
    async afterLoginAction(){
      if (!getToken()) return false;
      this.info() 
      const appStore = useAppStore();
      await appStore.fetchServerMenuConfig({});
      //注册路由
      if (!appStore.getIsDynamicAddedRoute) {
        appStore.appAsyncRoute.forEach( route => {
          if(route.name&&!router.hasRoute(route.name))
            router.addRoute(route as unknown as RouteRecordRaw);
        })
        appStore.setDynamicAddedRoute(true);
      }
      return true;
    },
    logoutCallBack() {
      const appStore = useAppStore();
      this.resetInfo();
      clearToken();
      removeRouteListener();
      appStore.clearServerMenu();
    },
    setSessionTimeout(flag: boolean) {
      this.sessionTimeout = flag;
    },
    // 退出登录
    async logout(goLogin = false) {
      try {
        await userLogout();
      } finally {
        this.logoutCallBack();
      }
      setToken(undefined);
      goLogin && router.push('/login');
    },
    //清除登录信息
    clearloginfo() {
      clearToken();
      removeRouteListener();
      const currentRoute = router.currentRoute.value;
      router.push({
        name:"login",
        query: {
          ...router.currentRoute.value.query,
          redirect: currentRoute.name as string,
        },
      });
    },
  },
});

export default useUserStore;
