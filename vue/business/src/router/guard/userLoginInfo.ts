import type { Router, LocationQueryRaw } from 'vue-router';
import NProgress from 'nprogress'; // progress bar
import { useUserStore } from '@/store';
import { isLogin } from '@/utils/auth';

export default function setupUserLoginInfoGuard(router: Router) {
  router.beforeEach(async (to, from, next) => {
    NProgress.start();
    const userStore = useUserStore();
    if (isLogin()) {
      if (userStore.id) {
        next();
      } else {
        try {
          await userStore.info();
          next();
        } catch (error) {
          await userStore.logout();
          next({
            name: 'login',
            query: {
              redirect: to.name,
              ...to.query,
            } as LocationQueryRaw,
          });
        }
      }
    }else {
      if (to.name === 'login') {
        next();
        return;
      }else if(to.query.rouid){
        next();
        return;
      }
      //判断是否需要-用外部独立页面，如数据大屏
      let redirectval= to.name
      if(redirectval=="notFound"&&to.path){
        redirectval=to.path.split("/")[1]
      }
      next({
        name: 'login',
        query: {
          redirect: redirectval,
          ...to.query,
        } as LocationQueryRaw,
      });
    }
  });
}
