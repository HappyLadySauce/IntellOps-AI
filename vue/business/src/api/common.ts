import { defHttp } from '@/utils/http';
import { UploadFileParams } from '/#/axios';
enum Api {
  getTables = './common/table/getTables',
  tableWeigh = './common/table/tableWeigh',
  checkHaseRule = './common/basetool/checkHaseRule',

  parseImport = './common/dataimport/parseImport',
  importData = './common/dataimport/importData',
}
//获取数据库表
export function getTables(params: object) {
  return defHttp.get({ url: Api.getTables, params:params }, { errorMessageMode: 'message'});
}
//更新排序
export function tableWeigh(params: object) {
  return defHttp.post({ url: Api.tableWeigh, params:params}, { errorMessageMode: 'message' });
}
//检查路由是否存在-路由不能重复
export function checkHaseRule(params: object) {
    return defHttp.post({ url: Api.checkHaseRule, params:params}, { errorMessageMode: 'message' });
}

//导入数据-解析数据
export function parseImport(params: object) {
    return defHttp.post({ url: Api.parseImport, params:params}, { errorMessageMode: 'message' });
}
//导入数据-执行导入
export function importData(params: object) {
  return defHttp.post({ url: Api.importData, params:params}, { errorMessageMode: 'message' });
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
//获取配置中的上传文件路径
const DOMAIN =import.meta.env.VITE_APP_ENV=="production"? window?.globalConfig.Upload_url: window?.globalConfig.Upload_url_dev;
export function userUploadApi(
  params: UploadFileParams,
  onUploadProgress?: (progressEvent: any) => void
) {
  return defHttp.uploadFile<UploadItem>({ url:`${DOMAIN}/upload/image`,onUploadProgress},params);
}
/**数据库表类型 */
export interface TableItem {
  id:number,
  name: string;
  title: string;
}