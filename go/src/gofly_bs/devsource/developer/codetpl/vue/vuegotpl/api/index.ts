import { defHttp } from '@/utils/http';
import { join } from 'lodash';
import { isArray } from '/@/utils/is';
enum Api {
    getList = '/modname/filename/getList',
    save = '/modname/filename/save',
    putStatus = '/modname/filename/putStatus',
    del = '/modname/filename/del',
}

//列表数据
export function getList(params: any) {
    for(let key in params){
        if(isArray(params[key])){
            params[key]=join(params[key])
        }
    }
  return defHttp.get({ url: Api.getList, params:params }, { errorMessageMode: 'none' });
}

//提交数据
export function save(params: object) {
    return defHttp.post({ url: Api.save, params:params}, { errorMessageMode: 'message' });
}
//更新状态
export function putStatus(params: object) {
    return defHttp.put({ url: Api.putStatus, params:params}, { errorMessageMode: 'message' });
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