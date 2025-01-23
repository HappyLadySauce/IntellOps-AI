<template>
  <a-tag v-if="isHex(color)" :size="size" :style="{backgroundColor:hexToRgba(color),color:color}">{{ value }}</a-tag>
  <a-tag v-else :size="size" :color="color">{{ value }}</a-tag>
</template>

<script lang="ts" setup>
//升级了tag标签
const props = defineProps({
  color: {
    type: String,
    default: ""
  },
  size: {
    type: String as any,
    default: "medium"
  },
  value: {
    type: String,
    required: true
  }
})
//判断是hex
const isHex=(color:string):boolean=>{
   if(color.indexOf("#")>-1){
    return true
   }
   return false
}
// 颜色格式 hex 转 rgba
const hexToRgba = (bgColor:string) => {
    let color = bgColor.slice(1);   // 去掉'#'号
    let rgba = [
        parseInt('0x'+color.slice(0, 2)),
        parseInt('0x'+color.slice(2, 4)),
        parseInt('0x'+color.slice(4, 6)),
        0.1
    ];
    return 'rgba(' + rgba.toString() + ')';
};
</script>

