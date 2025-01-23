<template>
   <a-card class="general-card contentcard" v-if="ConfData">
        <template #title>
         <div class="flex">
          <div style="flex:1;"> 
            {{ ConfData?.title }}
            <a-tag v-if="ConfData?.pluginident" >
              标识： {{ConfData?.pluginident}}
            </a-tag>
            <span class="tig" >配置文件位置：resource/config/{{ConfData?.name}}.yaml</span>
            <a-tooltip content="配置字段命名说明：_txt结尾是多行文本，_read结尾是只读，_switch结尾开关,#注释加 &des 是字段说明。" position="top" mini>
              <icon-question-circle-fill  class="configdes"/>
            </a-tooltip>
          </div>
          <a-tooltip content="如果有相同配置标识文件，当点击启用时会禁用其他相同标识文件" position="tr" mini>
              <a-switch v-model="ConfData.status" @change="handleUpConfigStatus">
                <template #checked>
                  启用
                </template>
                <template #unchecked>
                  禁用
                </template>
              </a-switch>
          </a-tooltip>  
         </div>
        </template>
        <a-form :model="form"  auto-label-width>
          <a-row :gutter="20">
            <template v-for="item in ConfData.data">
              <a-col :span="12">
                  <a-form-item
                  :label="FontNameAndDes(item.keyname,0)"
                  :field="item.keyfield"
                  :extra="FontNameAndDes(item.keyname,1)"
                  >
                    <a-textarea v-model="item.keyvalue" :placeholder="`填写${FontNameAndDes(item.keyname,0)}`" v-if="item.keyfield.indexOf('_txt')>-1"/>
                    <a-tooltip :content="`“${FontNameAndDes(item.keyname,0)}”是只读不可编辑字段`" position="top" mini  v-else-if="item.keyfield.indexOf('_read')>-1">
                      <a-input v-model="item.keyvalue" :placeholder="`填写${FontNameAndDes(item.keyname,0)}`" readonly />
                    </a-tooltip>
                    <a-switch v-model="item.keyvalue" :checked-value="1" :unchecked-value="0" checked-text="是" unchecked-text="否" v-else-if="item.keyfield.indexOf('_switch')>-1"/>
                    <a-input v-model="item.keyvalue" :placeholder="`填写${FontNameAndDes(item.keyname,0)}`" allow-clear v-else/>
                  </a-form-item>
              </a-col>
            </template>
            <a-col :span="24">
                <a-form-item >
                  <div class="frombtn">
                      <a-button type="primary" html-type="submit" style="width: 120px;" @click="submitConfig">保存</a-button>
                  </div>
                </a-form-item>
            </a-col>
          </a-row>
        </a-form>
    </a-card>
  </template>
  
  <script lang="ts" setup>
    import { ref,PropType } from 'vue';
    //api
    import { menuItem,saveCodeStoreConfig,upConfigStatus} from '@/api/datacenter/configuration';
    import { Message } from '@arco-design/web-vue';
    const emits = defineEmits(['ok'])
    const props = defineProps({
      ConfData: {
        type: Object as PropType<menuItem>,
      },
    })
    //保存配置数据
    const form = ref({});
    const submitConfig=async()=>{
      if(props.ConfData){
        try {
          Message.loading({content:"保存中",id:"updata",duration:0})
          await saveCodeStoreConfig(props.ConfData);
          Message.success({content:"保存成功",id:"updata",duration:2000})
        } catch (error) {
          Message.error({content:"",id:"updata",duration:1})
        }
      }else{
        Message.warning({content:"配置数据不存在",id:"updata",duration:2000})
      }
    }
    //切换使用状态
    const handleUpConfigStatus=async(value:any)=>{
      if(props.ConfData){
        try {
          Message.loading({content:"更新状态中",id:"updata",duration:0})
          await upConfigStatus(props.ConfData);
          Message.success({content:"更新状态成功",id:"updata",duration:2000})
          emits('ok', true)
        } catch (error) {
          Message.loading({content:"更新状态中",id:"updata",duration:1})
        }
      }else{
        Message.warning({content:"配置数据不存在",id:"updata",duration:2000})
      }
    }
    //获取名称和描述
    const FontNameAndDes=(str:string,index:number):string=>{
      if(str){
        const str_arr=str.split("&des")
        if(index==0){
          return str_arr[index]
        }else{
          if(index==1&&str_arr.length==2)
            return str_arr[index]
          }
      }
      return ""
    }
  </script>
  
  <style scoped lang="less">
  .contentcard{
        overflow: hidden;
    }
    :deep(.general-card > .arco-card-header){
      padding: 10px 16px;
    }
    .iconbtn{
      user-select: none;
      cursor: pointer;
      opacity: .8;
      &:hover{
        opacity: 1;
      }
    }
    .frombtn{
      width: 100%;
      text-align: center;
    }
    .tig{
      font-size: 12px;
      color: var(--color-neutral-4);
      padding-left: 5px;
    }
    .configdes{
      color: var(--color-neutral-4);
      margin-left: 10px;
    }
  </style>
  