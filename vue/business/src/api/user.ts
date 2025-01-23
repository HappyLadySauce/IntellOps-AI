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
  //用户
   Login = '/user/login ',
   Logout = '/user/logout',
   resetPassword = '/user/resetPassword',
   GetUserInfo = '/user/getUserinfo',
   GetMenu = '/user/account/getMenu',
   getCaptcha = './common/basetool/getCaptcha',
}
//登录
export function login(params: LoginData) {
  if(params.password){
    params=Object.assign({},params,{password:md5(params.password)})//加密推送
  }
  return defHttp.post({ url: Api.Login, params:params}, { errorMessageMode: 'message' });
}
//重置密码
export function resetPassword(params: LoginData) {
  params=Object.assign({},params,{password:md5(params.password)})//加密推送
  return defHttp.post({ url: Api.resetPassword, params:params}, { errorMessageMode: 'message' });
}
//退出登录
export function logout() {
  return defHttp.post({ url: Api.Logout}, { errorMessageMode: 'message' });
}
//获取用信息
export function getUserInfo() {
  return defHttp.get({ url: Api.GetUserInfo }, { errorMessageMode: 'message' });
}

//获取后台菜单
export function getMenuList(params:object) {
  return defHttp.get({ url: Api.GetMenu, params:params }, { errorMessageMode: 'notification' });
}

//获取邮箱验证码、手机验证码、人机验证图片
export function getCaptcha(params:object) {
  return defHttp.get({ url: Api.getCaptcha, params:params }, { errorMessageMode: 'message' });
}
