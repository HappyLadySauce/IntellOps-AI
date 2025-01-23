import { defHttp } from '@/utils/http';
import { UploadFileParams } from '/#/axios';
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
const DOMAIN =import.meta.env.VITE_APP_ENV=="production"? window?.globalConfig.Root_url: window?.globalConfig.Root_url_dev;
export function userUploadApi(
  params: UploadFileParams,
  onUploadProgress?: (progressEvent: any) => void
) {
  return defHttp.uploadFile<UploadItem>({ url:`${DOMAIN}/datacenter/upfile/upload`,onUploadProgress},params);
}