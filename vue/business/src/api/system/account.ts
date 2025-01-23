import { defHttp } from '@/utils/http';
import md5 from 'md5'
import { FilterParam } from '/@/utils';
//类型
export interface LoginData {
  username: string;
  password: string;
}
enum Api {
    getList = '/system/account/getList',
    getRole = '/system/account/getRole',
    save = '/system/account/save',
    Isaccountexist = '/system/account/isaccountexist',
    upStatus = '/system/account/upStatus',
    del = '/system/account/del',
}

//数据列表
export function getList(params: any) {
  params=FilterParam(params)
  return defHttp.get({ url: Api.getList, params:params }, { errorMessageMode: 'message' });
}
//选择角色
export function getRole() {
  return defHttp.get({ url: Api.getRole }, { errorMessageMode: 'message' });
}
//提交菜单
export function save(params: any) {
  if(params.password){
    params=Object.assign({},params,{password:md5(params.password)})//加密推送
  }
  return defHttp.post({ url: Api.save, params:params}, { errorMessageMode: 'message' });
}
//判断账号是否已经存在
export function isAccountexist(params: object) {
  return defHttp.post({ url: Api.Isaccountexist, params:params}, { errorMessageMode: 'message',isTransformResponse:false});
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
export interface DataItem {
    id:number,
    pid:number,
    locale: string;
    title: string;
  }