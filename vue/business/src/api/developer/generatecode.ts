import { defHttp } from '@/utils/http';
import { FilterParam } from '/@/utils';
enum Api {
    getList = '/developer/generatecode/getList',
    upCodeTable = '/developer/generatecode/upCodeTable',
    getContent= '/developer/generatecode/getContent',
    getMenuParent= '/developer/generatecode/getMenuParent',
    save = '/developer/generatecode/save',
    upStatus = '/developer/generatecode/upStatus',
    del = '/developer/generatecode/del',
    delTplCode = '/developer/generatecode/delTplCode',
    uninstallcode = '/developer/generatecode/uninstallcode',
    getdbfield = '/developer/generatecode/getDbfield',
    get_tablelist = '/developer/generatecode/getTablelist',
    getDicTable = '/developer/generatecode/getDicTable',
    checkedHaseTable = '/developer/generatecode/checkedHaseTable',
    getGoVueDir = '/developer/generatecode/getGoVueDir',
    markeTplCode = '/developer/generatecode/markeTplCode',
}

//列表数据
export function getList(params: any) {
  params=FilterParam(params)
  return defHttp.get({ url: Api.getList, params:params }, { errorMessageMode: 'message' });
}
//获取内容
export function getContent(params: object) {
  return defHttp.get({ url: Api.getContent, params:params }, { errorMessageMode: 'message' });
}
//更新生成代码数据表
export function upCodeTable(params: object) {
    return defHttp.post({ url: Api.upCodeTable, params:params}, { errorMessageMode: 'message' });
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
//删除模板代码
export function delTplCode(params: object) {
    return defHttp.delete({ url: Api.delTplCode, params:params}, { errorMessageMode: 'message' });
}
//卸载
export function uninstallcode(params: object) {
    return defHttp.post({ url: Api.uninstallcode, params:params}, { errorMessageMode: 'message' });
}
//获取数据表字段
export function getdbfield(params: object) {
    return defHttp.get({ url: Api.getdbfield, params:params }, { errorMessageMode: 'message' });
  }
//获取数据库列表
export function get_tablelist(params: object) {
    return defHttp.get({ url: Api.get_tablelist, params:params }, { errorMessageMode: 'message' });
  }
//获取字典数据库
export function getDicTable(params: object) {
  return defHttp.get({ url: Api.getDicTable, params:params }, { errorMessageMode: 'message' });
}
//获取菜单数据
export function getMenuParent(params: object) {
  return defHttp.get({ url: Api.getMenuParent, params:params }, { errorMessageMode: 'message' });
}
//检查生成列表是否存在要添加的数据表
export function checkedHaseTable(params: object) {
  return defHttp.post({ url: Api.checkedHaseTable, params:params }, { errorMessageMode: 'message' });
}

//获取前后端代码目录
export function getGoVueDir(params: object) {
  return defHttp.get({ url: Api.getGoVueDir, params:params }, { errorMessageMode: 'message' });
}
//生成代码模板-用于开发
export function markeTplCode(params: object) {
  return defHttp.post({ url: Api.markeTplCode, params:params }, { errorMessageMode: 'message' });
}
/**数据类型 */
export interface DataItem {
    id:number,
    name: string;
}
/**数据类型 */
export interface TreeItem {
    value:string,
    label: string;
}
/**生成代码数据 */
export interface CodedataItem {
  id:number,
  tablename:string,
  codelocation:string,
  comment:string,
  pid:number,
  rule_id:number,
  rule_name:string,
  icon:string,
  is_install:number,
  routepath:string,
  routename:string,
  component:string,
  api_path:string,
  api_filename:string,
  tpl_type:string,
  cate_tablename:string,
}
//生成字段列表
export interface CodeListItem {
  id:number,
  generatecode_id:number,
  islist: boolean,
  name:string,
  field:string,
  isorder: boolean,
  align:string,
  show_ui:string,
  width:number,
  isform: boolean,
  required: boolean,
  formtype:string,
  datatable:string,
  datatablename:string,
  issearch:boolean,
  searchway:string,
  searchtype:string,
  searchwidth:number,
  field_weigh:number,
  list_weigh:number,
  search_weigh:number,
  def_value:string,
}