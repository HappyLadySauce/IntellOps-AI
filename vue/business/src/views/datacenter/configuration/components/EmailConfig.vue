<template>
   <a-card class="general-card contentcard" >
        <template #title>
        配置邮箱（每个saas账号单独配置）
        </template>
        <a-row :gutter="80">
        <a-col :span="12">
            <a-form-item
            label="发送者邮箱"
            field="sender_email"
            >
            <a-input v-model="formEmail.sender_email" placeholder="填写发送者邮箱" allow-clear/>
            </a-form-item>
        </a-col>
        <a-col :span="12">
            <a-form-item
            label="邮箱授权码"
            field="auth_code"
            >
            <a-input v-model="formEmail.auth_code" placeholder="填写邮箱授权码" allow-clear/>
            </a-form-item>
        </a-col>
        <a-col :span="12">
            <a-form-item
            label="邮件服务器"
            field="service_host"
            >
            <a-input v-model="formEmail.service_host" placeholder="填写邮件服务器" allow-clear/>
            </a-form-item>
        </a-col>
        <a-col :span="12">
            <a-form-item
            label="邮件服务器端口"
            field="service_port"
            >
            <a-input v-model="formEmail.service_port" placeholder="填写邮件服务器端口" allow-clear/>
            </a-form-item>
        </a-col>
        <a-divider orientation="center" style="width: 100%;">发送邮箱验证码</a-divider>
        <a-col :span="16">
            <a-form-item
            label="（验证码）邮件标题"
            field="mail_title"
            >
            <a-input v-model="formEmail.mail_title" placeholder="填写邮件标题" allow-clear/>
            </a-form-item>
        </a-col>
        <a-col :span="24">
            <a-form-item
            label="（验证码）邮件内容"
            field="mail_body"
            extra="{code} 会替换成动态验证码"
            >
            <a-textarea v-model="formEmail.mail_body" placeholder="填写邮件内容" allow-clear :auto-size="{minRows:3,maxRows:5}"/>
            </a-form-item>
        </a-col>
        <a-col :span="24">
            <a-form-item >
            <div class="frombtn">
                <a-button type="primary" html-type="submit" style="width: 120px;" @click="submitEmail">保存</a-button>
            </div>
            </a-form-item>
        </a-col>
        </a-row>
    </a-card>
  </template>
  
  <script lang="ts" setup>
    import {  ref, onMounted } from 'vue';
    //api
    import { saveEmail,getEmail} from '@/api/datacenter/configuration';
    import { Message } from '@arco-design/web-vue';
  
    //邮箱配置
    const formEmail=ref({
      sender_email:"",
      auth_code:"",
      mail_title:"",
      mail_body:"",
      service_host:"",
      service_port:"",
    })
    //保存邮箱
    const submitEmail=async()=>{
      Message.loading({content:"保存邮箱中",id:"upStatus",duration:0})
      await saveEmail(formEmail.value);
      Message.success({content:"保存邮箱成功",id:"upStatus",duration:2000})
    }
    //组件挂载完成后执行的函数
    onMounted(()=>{
      InitData()
    })
    //加载数据
    const InitData=async()=>{
      const emaildata = await getEmail({});
      formEmail.value=Object.assign({},formEmail.value,emaildata)
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
  