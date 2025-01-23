export const WHITE_LIST = [
  { name: 'notFound', children: [] },
  { name: 'login', children: [] },
];

export const NOT_FOUND = {
  name: 'notFound',
};

export const REDIRECT_ROUTE_NAME = 'Redirect';

export const DEFAULT_ROUTE_NAME = window?.globalConfig.RouterHome;

export const DEFAULT_ROUTE = {
  title: 'menu.dashboard.workplace',
  name: DEFAULT_ROUTE_NAME,
  fullPath: '/'+window?.globalConfig.RouterHome,
};
//动态路由
export const REDIRECT_NAME = 'Redirect';
export const PARENT_LAYOUT_NAME = 'ParentLayout';
export const PAGE_NOT_FOUND_NAME = 'PageNotFound';
export const EXCEPTION_COMPONENT = () => import('@/views/systool/404/index.vue');
/**
 * @description: default layout
 */
export const LAYOUT = () => import('@/layout/default-layout.vue');
/**
 * @description: parent-layout
 */
export const getParentLayout = (_name?: string) => {
  return () =>
    new Promise((resolve) => {
      resolve({
        name: PARENT_LAYOUT_NAME,
      });
    });
};
