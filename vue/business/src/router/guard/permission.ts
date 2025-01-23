import type { Router, RouteRecordNormalized ,RouteRecordRaw} from 'vue-router';
import NProgress from 'nprogress'; // progress bar

import usePermission from '@/hooks/permission';
import { useUserStore, useAppStore } from '@/store';
import { WHITE_LIST, NOT_FOUND } from '../constants';
import { isLogin } from '@/utils/auth';

export default function setupPermissionGuard(router: Router) {
  router.beforeEach(async (to, from, next) => {
    const appStore = useAppStore();
    const userStore = useUserStore();
    const Permission = usePermission();
    const permissionsAllow = Permission.accessRouter(to);
    userStore.btnroles=to.meta.btnroles// 把按钮权限赋值给用户的按钮权限存储
    if (appStore.menuFromServer&&isLogin()) {
      // 针对来自服务端的菜单配置进行处理
      // Handle routing configuration from the server
      if (appStore.getIsDynamicAddedRoute) {//判断是否已经加载路由组件
        next();
        NProgress.done();
        return;
      }
      // 根据需要自行完善来源于服务端的菜单配置的permission逻辑
      // Refine the permission logic from the server's menu configuration as needed
      if (!appStore.appAsyncRoute.length) {
        await appStore.fetchServerMenuConfig({});
        //加载路由组件
        appStore.appAsyncRoute.forEach( route => {
          router.addRoute(route as unknown as RouteRecordRaw);
        })
        appStore.setDynamicAddedRoute(true);
      }
      const serverMenuConfig = [...appStore.appAsyncRoute, ...WHITE_LIST];
      let exist = false;
      while (serverMenuConfig.length && !exist) {
        const element = serverMenuConfig.shift();
        if (element?.name === to.name) exist = true;
        if (element?.children) {
          serverMenuConfig.push(
            ...(element.children as unknown as RouteRecordNormalized[])
          );
        }
      }
      if (exist && permissionsAllow) {
        // 动态添加路由后，此处应当重定向到fullPath，否则会加载404页面内容
        next({ path: to.fullPath, replace: true, query: to.query });
      } else next(NOT_FOUND);
    } else {
      if(to.query.rouid){//独立界面无需登录查看
        if (appStore.getIsDynamicAddedRoute) {//判断是否已经加载路由组件
          next();
          NProgress.done();
          return;
        }
        // 根据需要自行完善来源于服务端的菜单配置的permission逻辑
        if (!appStore.appAsyncRoute.length) {
          await appStore.fetchServerMenuConfig({rouid:to.query.rouid});
          //加载路由组件
          appStore.appAsyncRoute.forEach( route => {
            router.addRoute(route as unknown as RouteRecordRaw);
          })
          appStore.setDynamicAddedRoute(true);
        }
        const serverMenuConfig = [...appStore.appAsyncRoute, ...WHITE_LIST];
        let exist = false;
        while (serverMenuConfig.length && !exist) {
          const element = serverMenuConfig.shift();
          if (element?.name === to.name) exist = true;
          if (element?.children) {
            serverMenuConfig.push(
              ...(element.children as unknown as RouteRecordNormalized[])
            );
          }
        }
        if (exist && permissionsAllow) {
          next({ path: to.fullPath, replace: true, query: to.query });
        } else next(NOT_FOUND);
      }else{
        if (permissionsAllow) next();
        else {
          next(NOT_FOUND);
        }
      }
    }
    NProgress.done();
  });
}
