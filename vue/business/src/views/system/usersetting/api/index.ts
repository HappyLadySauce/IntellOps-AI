import { defHttp } from '@/utils/http';
import md5 from 'md5'
//类型
export interface LoginData {
  email: string;
  mobile: string;
  username: string;
  password: string;
  captcha: number|undefined;
  codeid: string;
} 
enum Api {
  saveInfo = '/user/setting/saveInfo',
  getUserinfo = '/user/setting/getUserinfo',
}
//更新头像
export function saveInfo(params: any) {
  if(params.oldpassword){
    params=Object.assign({},params,{oldpassword:md5(params.oldpassword)})//加密推送
  }
  if (params.type&&params.type === 'password') {
    if(params.newpassword){
      params=Object.assign({},params,{newpassword:md5(params.newpassword)})//加密推送
    }
    delete params["repassword"]
    delete params["email"]
    delete params["mobile"]
  }else if (params.type&&params.type === 'mobile') {
    delete params["email"]
  }else if (params.type&&params.type === 'email') {
    delete params["mobile"]
  }
  return defHttp.post({ url: Api.saveInfo, params:params}, { errorMessageMode: 'message' });
}

//获取用户信息
export function getUserinfo() {
  return defHttp.get({ url: Api.getUserinfo }, { errorMessageMode: 'message' });
}


/** 用户类型 */
export interface OptionUser {
  username:string
  deptname:string
  roles:string
}
