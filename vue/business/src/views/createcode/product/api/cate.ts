import { defHttp } from '@/utils/http';
enum Api {
    getTree = '/createcode/productcate/getTree',
    save = '/createcode/productcate/save',
    upStatus = '/createcode/productcate/upStatus',
    del = '/createcode/productcate/del',
}

//列表数据
export function getTree(params: object) {
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
export function delCate(params: object) {
    return defHttp.delete({ url: Api.del, params:params}, { errorMessageMode: 'message' });
}
/**数据类型 */
export interface DataItem {
    id:number,
    name: string;
}
