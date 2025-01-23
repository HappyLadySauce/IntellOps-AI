<template>
  <div class="container" >
    <page-card breadcrumb>
      <div class="flexbox">
        <div class="menu">
          <div class="tablebox">
            <a-table :data="nenudata" row-key="id" @row-click="handleClickMenu" :selected-keys="selectedKeys" :scroll="{ y: '100%'}" :pagination="false" :bordered="{wrapper: true, cell: true}" >
              <template #columns>
                <a-table-column :width="50" data-index="id">
                  <template #title><icon-loop :spin="loading" @click="handleCodestoreConfig" size="16" class="btncs"/></template>
                </a-table-column>
                <a-table-column  title="配置名称" data-index="title">
                  <template #cell="{ record }">
                      <div class="titlebox">
                      <div class="text">{{ record.title }}</div>
                      <div class="icon"><icon-right /></div>
                      </div>
                  </template>
                </a-table-column>
              </template>
            </a-table>
          </div>
        </div>
        <div class="content">
          <email-config v-if="selectedKeys[0]==1"></email-config>
          <app-config v-else-if="selectedKeys[0]==2"></app-config>
          <codestore-config v-else :ConfData="Menuitem" @ok="handleCodestoreConfig"></codestore-config>
        </div>
      </div>
  </page-card>
  </div>
</template>

<script lang="ts" setup>
  import {ref, onMounted } from 'vue';
  //api
  import { getCodestoreConfig,menuItem} from '@/api/datacenter/configuration';
  import EmailConfig from './components/EmailConfig.vue';
  import AppConfig from './components/AppConfig.vue';
  import CodestoreConfig from './components/CodestoreConfig.vue';
  import { cloneDeep } from 'lodash-es';
  import { useRoute } from 'vue-router'
  import { inArray } from '@/utils/is';
  const def_data=ref<menuItem []>([
    {
      id:1,
      title:"配置邮箱",
      name:"",
      pluginident:"",
      status:true,
      data:[],
    }
  ])
  onMounted(async()=>{
    const route = await useRoute();
    if(inArray(route.meta.btnroles,"syscnf")||inArray(route.meta.btnroles,"*")){
      def_data.value.push({
        id:2,
        title:"系统配置",
        name:"",
        pluginident:"",
        status:true,
        data:[],
      })
    }
    handleCodestoreConfig()
  })
  const nenudata=ref<menuItem []>([])

  //选择菜单数据
  const Menuitem=ref<menuItem>()
  const selectedKeys=ref<number []>([1])
  const handleClickMenu=(TableData:any)=>{
    selectedKeys.value=[TableData.id]
    Menuitem.value=TableData
  }
  //组件挂载完成后执行的函数

  //获取代码仓安装配置
  const loading=ref(false)
  const handleCodestoreConfig=async()=>{
     nenudata.value=cloneDeep(def_data.value)
     loading.value=true
     const codestorelist=await getCodestoreConfig({})
     for (var index in codestorelist) {
      const item=codestorelist[index]
      item["id"]=nenudata.value.length+1
      nenudata.value.push(item)
     }
     setTimeout(()=>{
      loading.value=false
     },800)
  }
</script>

<script lang="ts">
  export default {
    name: 'configuration',
  };
</script>

<style scoped lang="less">
  .flexbox{
      display: flex;
      justify-content: space-between; // 替代原先的space-between布局方式
      height: 100%;
      .menu{
        height: 100%;
        width: 230px;
        margin-right: 10px;
        background-color: var(--color-neutral-1);
      .tablebox{
        height: 100%;
          //分类标题背景色
          :deep(.arco-table-th){
            background-color: var(--color-neutral-3);
          }
          :deep(.arco-table-tr-checked .arco-table-td){
            color: rgb(var(--primary-6));
          }
        .titlebox{
          cursor: pointer;
          display: flex;
          .text{
            flex: 1;
          }
        }
      }
      }
      .content{
        flex: 1;
        min-width: 400px;
        background-color: var(--color-neutral-1);
        padding: 10px;
        height: 100%;
        overflow-y: auto;
        overflow-x: hidden;
      }
  }
  :deep(.arco-btn-size-medium){
    padding: 0 10px;
  }
  :deep(.arco-table-th) {
    &:last-child {
      .arco-table-th-item-title {
        margin-left: 16px;
      }
    }
  }
</style>
