import { defHttp } from '@/utils/http';
import { UploadFileParams } from '/#/axios';
enum Api {
    downCode = '/developer/codestoreoption/downCode',
    installCode = '/developer/codestoreoption/installCode',
    uninstallCode = '/developer/codestoreoption/uninstallCode',
    packCode = '/developer/codestoreoption/packCode',
    getInstallPack = '/developer/codestoreoption/getInstallPack',
    getPackdirs = '/developer/codestoreoption/getPackdirs',
    getMenutreeData = '/developer/codestoreoption/getMenutreeData',
    menuTreeToJson = '/developer/codestoreoption/menuTreeToJson',
}
//上传文件
export interface UploadItem {
    id: number;
    name: string;
    response: string;
    status: string;
    time: number;
    uid: string;
    url: string;
  }
//上传服务器域名
const DOMAIN= import.meta.env.VITE_APP_ENV=="production"? window?.globalConfig.Root_url: window?.globalConfig.Root_url_dev;
//下载安装代码包
export function downCode(params: object,baseurl:string) {
    return defHttp.post({ url: Api.downCode+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}
//安装
export function installCode(params: object) {
    return defHttp.post({ url: Api.installCode, params:params}, { errorMessageMode: 'message' });
}
//卸载
export function uninstallCode(params: object) {
    return defHttp.post({ url: Api.uninstallCode, params:params}, { errorMessageMode: 'message' });
}
//打包
export function packCode(params: object) {
    return defHttp.post({ url: Api.packCode, params:params}, { errorMessageMode: 'message' });
}

//获取已经安装的包
export function getInstallPack(params: object) {
    return defHttp.get({ url: Api.getInstallPack, params:params}, { errorMessageMode: 'message' });
}

//获取打包文件目录
export function getPackdirs(params: object) {
    return defHttp.get({ url: Api.getPackdirs, params:params}, { errorMessageMode: 'message' });
}
//获取菜单数据
export function getMenutreeData(params: object) {
    return defHttp.get({ url: Api.getMenutreeData, params:params}, { errorMessageMode: 'message' });
}
//菜单id转JSON数据
export function menuTreeToJson(params: object) {
    return defHttp.post({ url: Api.menuTreeToJson, params:params}, { errorMessageMode: 'message' });
}
//安装本地代码包
export function installLocalCode(  params: UploadFileParams, onUploadProgress?: (progressEvent: any) => void) {
    return defHttp.uploadFile({url:`${DOMAIN}/developer/codestoreoption/installLocalCode`,onUploadProgress},params);
}
//上传文件到公共仓
export function userUploadFile(
  params: UploadFileParams,
  onUploadProgress?: (progressEvent: any) => void
) {
  return defHttp.uploadFile<UploadItem>({ url:`${DOMAIN}/developer/codestoreoption/upfile`,onUploadProgress},params);
}
