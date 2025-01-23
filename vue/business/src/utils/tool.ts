import { useUserStore } from '@/store'
import { useClipboard } from '@vueuse/core';
import { isAllUrl } from '@/utils/is';
import { Message } from '@arco-design/web-vue';
//获取随机数
export function randomRange(min:number, max:number){
    var returnStr = "",
    range = (max ? Math.round(Math.random() * (max-min)) + min : min),
    arr = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'];
    for(var i=0; i<range; i++){
    var index = Math.round(Math.random() * (arr.length-1));
    returnStr += arr[index];
    }
    return returnStr;
}
//获取附件完整访问路径
export function GetFullPath(url:any):string{
    if(!url){
        return ""
    }else if(isAllUrl(url)){
        return url
    }else{
        const userStore = useUserStore()
         if(url.startsWith("resource/uploads/")){//本地访问域名
            return userStore.localurl+url
         }else{//获取设置访问域名
            return userStore.rooturl+url
         }
    }
}
//复制
var copyObj :any
export function CopyText(text:any,msg:string){
    if(!copyObj){
        const { copy } = useClipboard();
        copyObj=copy
    }
    copyObj(text);
    Message.success({content:msg,id:"copy"});
}
