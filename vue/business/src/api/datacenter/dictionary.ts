import { defHttp } from '@/utils/http';
import { FilterParam } from '/@/utils';
enum Api {
    getList = '/datacenter/dictionary/getList',
    getDicData= './common/basetool/getDicData',
    getTableDataForm= '/datacenter/dictionary/getTableDataForm',
    save = '/datacenter/dictionary/save',
    upStatus = '/datacenter/dictionary/upStatus',
    del = '/datacenter/dictionary/del',
}

//列表数据
export function getList(params: any) {
  params=FilterParam(params)
  return defHttp.get({ url: Api.getList, params:params }, { errorMessageMode: 'message' });
}
//获取字典数据列表
export function getDicData(params: object) {
  return defHttp.get({ url: Api.getDicData, params:params }, { errorMessageMode: 'message' });
}
//获取数据表数据-表单生成使用
export function getTableDataForm(params: object) {
  return defHttp.get({ url: Api.getTableDataForm, params:params }, { errorMessageMode: 'message' });
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
