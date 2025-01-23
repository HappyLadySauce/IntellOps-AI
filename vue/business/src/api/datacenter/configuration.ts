import { defHttp } from '@/utils/http';
enum Api {
    getCodestoreConfig = '/datacenter/configuration/getCodestoreConfig',
    getEmail = '/datacenter/configuration/getEmail',
    saveEmail = '/datacenter/configuration/saveEmail',
    saveCodeStoreConfig = '/datacenter/configuration/saveCodeStoreConfig',
    upConfigStatus = '/datacenter/configuration/upConfigStatus',
}

//获取代码仓代码数据
export function getCodestoreConfig(params: object) {
  return defHttp.get({ url: Api.getCodestoreConfig, params:params }, { errorMessageMode: 'message' });
}
//列表数据
export function getEmail(params: object) {
  return defHttp.get({ url: Api.getEmail, params:params }, { errorMessageMode: 'message' });
}
//提交数据
export function saveEmail(params: object) {
    return defHttp.post({ url: Api.saveEmail, params:params}, { errorMessageMode: 'message' });
}
//修改代码仓配置
export function saveCodeStoreConfig(params: object) {
    return defHttp.post({ url: Api.saveCodeStoreConfig, params:params}, { errorMessageMode: 'message' });
}
//修改代码仓配置
export function upConfigStatus(params: object) {
    return defHttp.post({ url: Api.upConfigStatus, params:params}, { errorMessageMode: 'message' });
}
/**数据类型 */
export interface menuItem {
    id:number,
    title: string;
    status: Boolean;
    name: string;
    pluginident: string;
    data:DataObj[];
}

export interface DataObj {
    keyname: string;
    keyvalue: string;
    keyfield: string;
}
