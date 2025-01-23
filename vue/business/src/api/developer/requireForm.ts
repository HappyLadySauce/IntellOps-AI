import { defHttp } from '@/utils/http';
enum Api {
    requirement = '/developer/codestore/requirement',
}

//发布需求数据
export function requirement(params: object,baseurl:string) {
    return defHttp.post({ url: Api.requirement+"?baseurl="+baseurl, params:params}, { errorMessageMode: 'message' });
}

