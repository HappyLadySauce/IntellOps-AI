<template>
  <div v-if="iconList.length" class="IconPicker">
       <a-scrollbar  style="height:250px;overflow: auto;">
        <div class="title">arco design图标</div>
        <ul class="iconUl">
          <li
            v-for="icon in iconList"
            :key="icon"
            class="border border-primary"
            @click="handleClick(icon)"
            :title="icon"
          >
           <component :is="icon"></component>  
          </li>
        </ul>
        <div class="title">阿里iconfont图标</div>
        <ul class="iconUl">
          <li
            v-for="icon in alilist"
            :key="icon"
            class="border border-primary"
            @click="handleClick(icon)"
            :title="icon"
          >
            <SvgIcon :size="18" :icon="icon"  />
          </li>
        </ul>
      </a-scrollbar>
    </div>
</template>
<script lang="ts" setup>
  import iconsData from '../data/icons.data';
  import SvgIcon from './SvgIcon.vue';
  // 没有使用别名引入，是因为WebStorm当前版本还不能正确识别，会报unused警告
  const emit = defineEmits(['change','update:modelValue']);
  function getIcons() {
    const data = iconsData as any;
    const prefix: string = data?.prefix ?? '';
    let result: string[] = [];
    if (prefix) {
      result = (data?.icons ?? []).map((item:any) => `${prefix}-${item}`);
    } else if (Array.isArray(iconsData)) {
      result = iconsData as string[];
    }
    return result;
  }
  //阿里图标
  function getAlilist() {
    const data = iconsData as any;
    const aliprefix: string = data?.aliprefix ?? '';
    let result: string[] = [];
    if (aliprefix) {
      result = (data?.alifont ?? []).map((item:any) => `${aliprefix}-${item}`);
    } else if (Array.isArray(iconsData)) {
      result = iconsData as string[];
    }
    return result;
  }
  const iconList = getIcons();
  const alilist = getAlilist();
  //选择菜单
  const handleClick=(icons:any)=>{
    emit('change', icons)
    emit('update:modelValue', icons)
  }
</script>
<style lang="less" scoped>
.IconPicker{
  width: 250px;
  height: 250px;
  .title{
    font-size: 12px;
    padding: 5px 10px;
  }
  .iconUl{
    list-style: none;
    margin: 0;
    padding: 0;
    padding-left: 0.5rem;
    padding-right: 0.5rem;
    display: flex;
    flex-wrap: wrap;
    li{
      border-style: solid;
      border-width: 1px;
      cursor: pointer;
      display: flex;
      justify-content: center;
      align-items: center;
      margin-right: 0.25rem;
      margin-top: 0.25rem;
      padding: 0.5rem;
      width: 12.5%;
      border-color: #e5e7eb;
    }
  }
}
</style>
