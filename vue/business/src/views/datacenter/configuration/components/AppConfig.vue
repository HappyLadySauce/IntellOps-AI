<template>
   <a-card class="general-card contentcard" v-if="formData">
        <template #title> 应用(系统)主要配置(对应resource/config.yaml) 是整个项目共用配置</template>
        <a-row :gutter="80">
          <a-col :span="24">
              <a-form-item
              label="前端代码开发根目录位置"
              field="vueobjroot"
              extra="配置代码生成时-前端代码根目录位置(开发必改)"
              >
              <a-input v-model="formData.vueobjroot" placeholder="填写附件请求路径前缀" allow-clear/>
              </a-form-item>
          </a-col>
          <a-col :span="24">
              <a-form-item
              label="附件访问域名、前缀路径"
              field="mainurl"
              extra="这里是配置系统内存储地址，示例地址：https://youdomain.cn/common/uploadfile/getfile?url=，如果安装腾讯云cos、阿里云oss、七牛等后会优先切换到这些三方存储。"
              >
              <a-input v-model="formData.mainurl" placeholder="填写附件请求路径前缀" allow-clear/>
              </a-form-item>
          </a-col>
          <a-col :span="24">
              <a-form-item >
              <div class="frombtn">
                  <a-button type="primary" html-type="submit" style="width: 120px;" @click="submitAttachmentConfig">保存</a-button>
              </div>
              </a-form-item>
          </a-col>
        </a-row>
    </a-card>
  </template>
  
  <script lang="ts" setup>
    import {  ref, onMounted } from 'vue';
    //api
    import { saveConfig,getConfig} from '@/api/datacenter/common_config';
    import { Message } from '@arco-design/web-vue';
    //数据配置
    const formData=ref({
      vueobjroot:"",
      mainurl:"",
    })
    //保存邮箱
    const submitAttachmentConfig=async()=>{
      try{
        Message.loading({content:"保存中",id:"updata",duration:0})
        await saveConfig(formData.value);
        Message.success({content:"保存成功",id:"updata",duration:2000})
      } catch (error) {
        Message.error({content:"",id:"updata",duration:1})
      }
    }
    //组件挂载完成后执行的函数
    onMounted(()=>{
      InitData()
    })
    //加载数据
    const InitData=async()=>{
      const emaildata = await getConfig({});
      formData.value=Object.assign({},formData.value,emaildata)
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
  </style>
  