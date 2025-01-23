import { defHttp } from '@/utils/http';
import md5 from 'md5'
//类型
enum Api {
  //用户
   Login = '/developer/codestore/login',
   freeLogin = '/developer/codestore/freeLogin',
   registerUser = '/developer/codestore/registerUser ',
   getCode = '/developer/codestore/loginCode',
}
//登录
export function login(params: any,baseurl:string) {
  params=Object.assign({},params,{password:md5(params.password)})//加密推送
  return defHttp.post({ url: Api.Login+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}
//免密登录
export function freeLogin(params: any,baseurl:string) {
  return defHttp.post({ url: Api.freeLogin+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}
//注册
export function registerUser(params: any,baseurl:string) {
  params=Object.assign({},params,{password:md5(params.password)})//加密推送
  return defHttp.post({ url: Api.registerUser+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}
//获取验证码
export function getCode(params: object,baseurl:string) {
  return defHttp.post({ url: Api.getCode+"?baseurl="+baseurl, params:params }, { errorMessageMode: 'message' });
}
