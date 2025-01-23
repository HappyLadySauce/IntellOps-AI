import { defHttp } from '@/utils/http';
enum Api {
    getConfig = '/datacenter/common_config/getConfig',
    saveConfig = '/datacenter/common_config/saveConfig',
}

//列表数据
export function getConfig(params: object) {
  return defHttp.get({ url: Api.getConfig, params:params }, { errorMessageMode: 'message' });
}
//提交数据
export function saveConfig(params: object) {
    return defHttp.post({ url: Api.saveConfig, params:params}, { errorMessageMode: 'message' });
}
/**数据类型 */
export   interface menuItem {
    id:number,
    title: string;
}
