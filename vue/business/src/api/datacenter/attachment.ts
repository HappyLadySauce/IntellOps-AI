import { defHttp } from '@/utils/http';
import { join } from 'lodash';
enum Api {
    getPictureCate = '/datacenter/attachment/getPictureCate',
    getMyFiles = '/datacenter/attachment/getMyFiles',
    getPicture = '/datacenter/attachment/getPicture',
    getContent= '/datacenter/attachment/getContent',
    saveFile = '/datacenter/attachment/save',
    upImgPid = '/datacenter/attachment/upImgPid',
    delFile = '/datacenter/attachment/del',
    delDir = '/datacenter/attachment/delDir',
}

//获取我的附件
export function getMyFiles(params: object) {
  return defHttp.get({ url: Api.getMyFiles, params:params }, { errorMessageMode: 'message' });
}
//获取图片库
export function getPictureCate(params: object) {
  return defHttp.get({ url: Api.getPictureCate, params:params }, { errorMessageMode: 'message' });
}
//获取图片库
export function get_picture(params: object) {
  return defHttp.get({ url: Api.getPicture, params:params }, { errorMessageMode: 'message' });
}
//获取内容
export function getContent(params: object) {
  return defHttp.get({ url: Api.getContent, params:params }, { errorMessageMode: 'message' });
}
//提交数据
export function saveFile(params: object) {
    return defHttp.post({ url: Api.saveFile, params:params}, { errorMessageMode: 'message' });
}
//更新文件夹
export function upImgPid(params: object) {
    return defHttp.post({ url: Api.upImgPid, params:params}, { errorMessageMode: 'message' });
}
//删除文件
export function delFile(params: object) {
    return defHttp.delete({ url: Api.delFile, params:params}, { errorMessageMode: 'message' });
}
//删除文件夹
export function delDir(params: object) {
    return defHttp.delete({ url: Api.delDir, params:params}, { errorMessageMode: 'message' });
}

/**数据类型 */
export interface DataItem {
    id:number,
    name: string;
}
//图片类型
export interface PictureItem {
    id:number;
    type: number;
    title: string;
    url: string;
    cover_url: string;
  }
  export interface CateItem {
    id:number;
    type: number;
    name: string;
  }
  export interface PictureItem {
    id:number;
    type: number;
    title: string;
    url: string;
  }
  export interface FileItem {
    id:number;
    pid: number;
    type: number;
    title: string;
    url: string;
    storage: string;
    mimetype: string;
    cover_url: string;
    is_common : number;
  }