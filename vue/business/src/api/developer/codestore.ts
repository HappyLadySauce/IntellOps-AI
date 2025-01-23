import { defHttp } from '@/utils/http';
import { FilterParam } from '/@/utils';
enum Api {
    getList = '/developer/codestore/codeList',
    getCodeCate= '/developer/codestore/getCodeCate',
    save = '/developer/codestore/save',
    upPrivateHouse = '/developer/codestore/upPrivateHouse',
    asyncVersion = '/developer/codestore/asyncVersion',
    upBaseCode = '/developer/codestore/upBaseCode',
    upPackToService = '/developer/codestore/upPackToService',
    checkPackName = '/developer/codestore/checkPackName',
    savePackName = '/developer/codestore/savePackName',
    addOrder = '/developer/codestore/addOrder',
    checkIsPay = '/developer/codestore/checkIsPay',
}

//列表数据
export function getList(params: any,baseurl:string) {
  params=FilterParam(params)
  return defHttp.post({ url: Api.getList+"?baseurl="+baseurl, params:params }, { errorMessageMode: 'message' });
}
//获取分类
export function getCodeCate(params: object) {
    return defHttp.get({ url: Api.getCodeCate, params:params }, { errorMessageMode: 'message' });
}
//提交数据
export function save(params: object) {
    return defHttp.post({ url: Api.save, params:params}, { errorMessageMode: 'message' });
}
//检查代码是否有新的代码更新
export function asyncVersion(params: object,baseurl:string) {
    return defHttp.post({ url: Api.asyncVersion+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}
//更新底座版本
export function upBaseCode(params: object,baseurl:string) {
    return defHttp.post({ url: Api.upBaseCode+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}
//更新私有仓地址
export function upPrivateHouse(params: object) {
    return defHttp.post({ url: Api.upPrivateHouse, params:params}, { errorMessageMode: 'message' });
}

//把打包数据提交到服务器
export function upPackToService(params: object,baseurl:string) {
    return defHttp.post({ url: Api.upPackToService+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}
//检测包名标识是否可用
export function checkPackName(params: object,baseurl:string) {
    return defHttp.post({ url: Api.checkPackName+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}
//占用包名标识
export function savePackName(params: object,baseurl:string) {
    return defHttp.post({ url: Api.savePackName+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}
// 购买插件
export function addOrder(params: object,baseurl:string) {
    return defHttp.post({ url: Api.addOrder+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}
// 检查是否支付
export function checkIsPay(params: object,baseurl:string) {
    return defHttp.post({ url: Api.checkIsPay+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}

/**数据类型 */
export interface DataItem {
    id:number,
    name: string;
    remark: string;
}
