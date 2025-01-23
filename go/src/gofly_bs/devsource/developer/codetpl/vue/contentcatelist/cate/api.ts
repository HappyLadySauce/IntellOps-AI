import { defHttp } from '@/utils/http';
import { FilterParam } from '/@/utils';
enum Api {
    getTree = '/modname/filename/getTree',
    save = '/modname/filename/save',
    upStatus = '/modname/filename/upStatus',
    del = '/modname/filename/del',
}

//列表数据
export function getTree(params: object) {
  params=FilterParam(params)
  return defHttp.get({ url: Api.getTree, params:params }, { errorMessageMode: 'message' });
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
export interface DataItem {
    id:number,
    name: string;
}