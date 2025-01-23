import { defHttp } from '@/utils/http';
import { join } from 'lodash';
import { isArray } from '/@/utils/is';
enum Api {
    getLogin = '/system/log/getLogin',
    getOperation = '/system/log/getOperation',
    getOperationDetail = '/system/log/getOperationDetail',
    delLastLogin = '/system/log/delLastLogin',
    delLastOperation = '/system/log/delLastOperation',
}

//1.获取登录日志
export function getLogin(params: any) {
    for(let key in params){
        if(isArray(params[key])){
            params[key]=join(params[key])
        }
    }
  return defHttp.get({ url: Api.getLogin, params:params }, { errorMessageMode: 'message' });
}

//2.获取操作日志
export function getOperation(params: any) {
    for(let key in params){
        if(isArray(params[key])){
            params[key]=join(params[key])
        }
    }
  return defHttp.get({ url: Api.getOperation, params:params }, { errorMessageMode: 'message' });
}

//3.获取操作日志
export function getOperationDetail(params: any) {
  return defHttp.get({ url: Api.getOperationDetail, params:params }, { errorMessageMode: 'message' });
}

//4.删除登录日志
export function delLastLogin() {
    return defHttp.delete({ url: Api.delLastLogin}, { errorMessageMode: 'message' });
}
//5.删除操作日志
export function delLastOperation() {
    return defHttp.delete({ url: Api.delLastOperation}, { errorMessageMode: 'message' });
}
