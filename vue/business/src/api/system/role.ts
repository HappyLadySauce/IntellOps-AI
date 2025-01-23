import { defHttp } from '@/utils/http';
import { FilterParam } from '/@/utils';
//类型
export interface LoginData {
  username: string;
  password: string;
}
enum Api {
    getList = '/system/role/getList',
    getParent = '/system/role/getParent',
    getMenuList = '/system/role/getMenuList',
    save = '/system/role/save',
    upStatus = '/system/role/upStatus',
    del = '/system/role/del',
}

//列表数据
export function getList(params: any) {
  params=FilterParam(params)
  return defHttp.get({ url: Api.getList, params:params }, { errorMessageMode: 'message' });
}
//选择数据
export function getParent(params: any) {
  return defHttp.get({ url: Api.getParent, params:params }, { errorMessageMode: 'message' });
}
//获取菜单数据
export function getMenuList(params: object) {
  return defHttp.get({ url: Api.getMenuList, params:params}, { errorMessageMode: 'message' });
}
//提交数据
export function save(params: object) {
    return defHttp.post({ url: Api.save, params:params}, { errorMessageMode: 'message' });
}
//更新状态
export function upStatus(params: object) {
    return defHttp.post({ url: Api.upStatus, params:params}, { errorMessageMode: 'message' });
}
//删除数据
export function del(params: object) {
    return defHttp.delete({ url: Api.del, params:params}, { errorMessageMode: 'message' });
}
/**数据类型 */
export interface RuleItem {
    id:number,
    pid:number,
    locale: string;
    title: string;
  }