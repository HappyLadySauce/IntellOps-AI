import { defHttp } from '@/utils/http';
enum Api {
    getList = '/datacenter/tabledata/getList',
    save = '/datacenter/tabledata/save',
    del = '/datacenter/tabledata/del',
}

//列表数据
export function getmenuList(params: object) {
  return defHttp.get({ url: Api.getList, params:params }, { errorMessageMode: 'message' });
}
//提交数据
export function save(params: object) {
    return defHttp.post({ url: Api.save, params:params}, { errorMessageMode: 'message' });
}
//删除数据
export function delMenuList(params: object) {
    return defHttp.delete({ url: Api.del, params:params}, { errorMessageMode: 'message' });
}
/**数据类型 */
export interface menuItem {
    id:number,
    business_id:number,
    title: string;
    remark: string;
    tablename: string;
    data_from: string;
    weigh: number;
    db_way: string;
    status: number;
}
