import { defHttp } from '@/utils/http';
import { FilterParam } from '/@/utils';
//类型
export interface LoginData {
  username: string;
  password: string;
}
enum Api {
    getList = '/system/dept/getList',
    getParent = '/system/dept/getParent',
    save = '/system/dept/save',
    upStatus = '/system/dept/upStatus',
    del = '/system/dept/del',
}

//列表数据
export function getList(params: any) {
  params=FilterParam(params)
  return defHttp.get({ url: Api.getList, params:params }, { errorMessageMode: 'message' });
}
//选择数据
export function getParent() {
  return defHttp.get({ url: Api.getParent }, { errorMessageMode: 'message' });
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