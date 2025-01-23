<template>
<a-modal v-model:visible="visible" unmountOnClose :footer="false" :closable="false" :width="410" :modal-style="{minHeight: '480px'}" >
   <div class="login-box">
        <div class="login-header">
            <div class="left"></div>
            <div class="right" @click="visible=false">
                <icon-close class="fonticon" size="20"/>
            </div>
        </div>
        <div class="login-content" v-if="dotype=='login'">
            <div class="formdata">
                <div class="tabs">
                    <div class="item"><span :class="{active:logintype=='psd'}" @click="changtype('psd')">密码登录</span></div>
                    <div class="item"><span :class="{active:logintype=='free'}" @click="changtype('free')">免密登录</span></div>
                    <div class="item"><span :class="{active:logintype=='wx'}" @click="changtype('wx')">微信登录</span></div>
                </div>
                <div class="formbox">
                    <div class="passwod-login" v-if="logintype=='psd'">
                        <div class="inputbox">
                            <a-input :style="{width:'320px'}" v-model="logindata.username" placeholder="手机号/邮箱/用户名" size="large" allow-clear />
                        </div>
                        <div class="inputbox">
                            <a-input-password :style="{width:'320px'}" v-model="logindata.password"  placeholder="密码"  size="large" allow-clear/>
                        </div>
                        <div class="loginbtn">
                            <a-button type="primary" long shape="round" size="large" @click="onLogin">登录</a-button>
                        </div>
                    </div>
                    <div class="free-login" v-else-if="logintype=='free'">
                        <a-form  size="large" :model="freeformdata" auto-label-width :style="{width:'320px'}" @submit="handleFreeSubmit">
                                <a-form-item field="email" label="邮箱"
                                >
                                   <a-input v-model="freeformdata.email" placeholder="填写您的邮箱" />
                                </a-form-item>
                                <a-form-item field="verifycode" label="验证码"
                                >
                                    <a-input v-model="freeformdata.verifycode" placeholder="邮箱验证码" />
                                    <a-button @click="handleFreeSend()" :style="{marginLeft:'10px'}" status="success"> {{ codeNum == 60 ? "发送" : `(${codeNum})后重发` }}</a-button>
                                </a-form-item>
                                <div class="submintbtn">
                                    <a-button type="primary" html-type="submit" long shape="round" size="large">立即登录</a-button>
                                </div>
                        </a-form>
                    </div>
                    <div class="wx-login" v-else-if="logintype=='wx'">
                        <div  style="text-align: center;">微信登录暂未开放请使用其他登录</div>
                    </div>
                </div>
            </div>
            <div class="options">
                <div class="register" @click="dotype='register'" v-if="logintype=='psd'||logintype=='free'">
                    注册账号
                </div>
            </div>
            <!-- <div class="footer">
                登录即代表同意 服务条款 和 隐私协议
            </div> -->
        </div>
        <div class="registerbox" v-else-if="dotype=='register'">
            <div class="regform">
                <a-form ref="registerRef" size="large" :model="formdata" auto-label-width :style="{width:'100%'}" @submit="handleSubmit">
                    <a-form-item field="email" label="邮箱"
                    :rules="[{required:true,message:'请输入您的邮箱'}]"
                    >
                      <a-input v-model="formdata.email" placeholder="填写您的邮箱" />
                    </a-form-item>
                    <a-form-item field="verifycode" label="验证码"
                    :rules="[{required:true,message:'请输入您的邮箱验证码'}]"
                    >
                      <a-input v-model="formdata.verifycode" placeholder="邮箱验证码" />
                      <a-button @click="handleSend()" :style="{marginLeft:'10px'}" status="success"> {{ codeNum == 60 ? "发送" : `(${codeNum})后重发` }}</a-button>
                    </a-form-item>
                    <a-form-item field="mobile" label="手机号码">
                      <a-input v-model="formdata.mobile" placeholder="填写您的手机号码(可用于登录)" />
                    </a-form-item>
                    <a-form-item field="password" label="密码" :rules="[{required:true,message:'请输入登录密码'}]">
                        <a-input-password  v-model="formdata.password"  placeholder="登录密码"  allow-clear/>
                    </a-form-item>
                    <a-form-item field="confirmpwd" label="确认密码" :rules="rulessecond">
                        <a-input-password  v-model="formdata.confirmpwd"  placeholder="登录密码"  allow-clear/>
                    </a-form-item>
                    <div class="submintbtn">
                        <a-button type="primary" html-type="submit" long shape="round" size="large">立即注册</a-button>
                    </div>
                </a-form>
                <div class="options">
                    <div class="register" @click="dotype='login'" >
                        有账号 去登录
                    </div>
                </div>
            </div>
        </div>
   </div>
</a-modal>
</template>
<script lang="ts">
import { defineComponent,nextTick,ref} from "vue";
import { Modal,Input,Button ,Form,Message} from '@arco-design/web-vue';
//api
import {getCode,registerUser,login,freeLogin} from  '@/api/developer/codestoreLogin';
import { setCstoreToken } from '@/utils/auth';
import { useUserStore } from '@/store';
export default defineComponent({
name: "Login",
components: {
    AModal: Modal,
    AButton: Button,
    AInput: Input,
    AInputPassword: Input.Password,
    AForm: Form,
    AFormItem: Form.Item,
},
emits: ['ok'],
setup(_, { emit }) {
    //数据存储
    const userStore = useUserStore();
    const dotype=ref("login")
    const visible=ref(false)
    const baseurl=ref("")
    //展示
    const showLogin=(baseurldata:string)=>{
        baseurl.value=baseurldata
        visible.value=true
        dotype.value="login"
        logindata.value={username:"",password:"",}
        formdata.value={
            email:"",
            password:"",
            mobile:"",
            confirmpwd:"",
            verifycode:"",
        }
    }
    const logintype=ref('psd')
    const changtype=(val:string)=>[
        logintype.value=val
    ]
    /*****登录 */
    const logindata=ref({
        username:"",
        password:"",
    })
    //提交登录
    const onLogin=async()=>{
        if(!logindata.value.username){
            Message.error({content:"请填写登录账号",id:"vcode",duration:2000})
        }else if(!logindata.value.password){
            Message.error({content:"请填写登录密码",id:"vcode",duration:2000})
        }else{
            try {
                Message.loading({content:"登录中",id:"vcode",duration:0})
                const resultdata = await login(logindata.value,baseurl.value);
                if(resultdata){
                    Message.success({content:"登录成功",id:"vcode",duration:2000})
                    setCstoreToken(resultdata)
                    emit('ok',resultdata);
                    visible.value=false
                }else{
                    Message.error({content:"登录失败",id:"vcode",duration:1})
                }
            } catch (error) {
                Message.loading({content:"登录中",id:"vcode",duration:1})
            }
        }
    }
    /*********** 注册 */
    const formdata=ref({
        email:"",
        password:"",
        mobile:"",
        confirmpwd:"",
        verifycode:"",
    })
    //注册账号
    const handleSubmit=async(data:any)=>{
        if(!data.errors){
            Message.loading({content:"注册中",id:"delaction",duration:0})
            const resultdata = await registerUser({email:formdata.value.email,password:formdata.value.password,
                mobile:formdata.value.mobile,code:formdata.value.verifycode},baseurl.value);
            if(resultdata){
              Message.success({content:"注册成功",id:"delaction",duration:2000})
              logindata.value.username=formdata.value.email
              dotype.value="login"
            }else{
                Message.loading({content:"注册中",id:"delaction",duration:1})
            }
        }
    }
   
    // 定时器id
    let clearId=ref();
    // 倒计时时间
    const codeNum = ref(60);
    // 是否发送了验证码 防止连点
    const isClickSend = ref(false)
     //发送短信
    const handleSend=async()=>{
        if(!formdata.value.email){
            Message.error({content:"请填写邮箱账号",id:"vcode",duration:2000})
        }else{
            if (isClickSend.value || codeNum.value != 60) return;
            isClickSend.value = true;
            Message.loading({content:"发送中，请稍后！",id:"vcode",duration:0})
            const data= await getCode({email:formdata.value.email},baseurl.value);
           if(!data){
                clearId.value = setInterval(() => {
                    codeNum.value--;
                    if (codeNum.value == 0) {
                    clearInterval(clearId.value);
                    codeNum.value = 60;
                    isClickSend.value = false;
                    }
                }, 1000);
             Message.success({content:"发送成功！",id:"vcode",duration:2000})
           }else{
              Message.error({content:"发送失败",id:"vcode",duration:2000})
           }
        }
    }
  //确认密码
  const rulessecond = [{
      validator: (value:any, cb:any) => {
        return new Promise(resolve => {
          if(value==undefined){
            cb('请输入确认密码')
          }else{
            if(value!=formdata.value.password){
              cb('两次密码不一致')
            }
          }
          resolve(null)
        })
      }
    }];
    //3快速登录
    const freeformdata=ref({
        email:"",
        verifycode:"",
    })
    //3.1登录
    const handleFreeSubmit=async()=>{
        if(!freeformdata.value.email){
            Message.error({content:"请填写邮箱账号",id:"vcode",duration:2000})
        }else if(!freeformdata.value.verifycode){
            Message.error({content:"请填写验证码",id:"vcode",duration:2000})
        }else{
            Message.loading({content:"登录中",id:"vcode",duration:0})
            const resultdata = await freeLogin({...freeformdata.value},baseurl.value);
            if(resultdata){
                    Message.success({content:"登录成功",id:"vcode",duration:2000})
                    setCstoreToken(resultdata)
                    emit('ok',resultdata);
                    visible.value=false
            }else{
                Message.loading({content:"登录中",id:"vcode",duration:1})
            }
        }
    }
    //3.2发验证码
    const handleFreeSend=async()=>{
        if(!freeformdata.value.email){
            Message.error({content:"请填写邮箱账号",id:"vcode",duration:2000})
        }else{
            try {
                if (isClickSend.value || codeNum.value != 60) return;
                isClickSend.value = true;
                Message.loading({content:"发送中，请稍后！",id:"vcode",duration:0})
                const data= await getCode({email:freeformdata.value.email},baseurl.value);
                if(!data){
                    clearId.value = setInterval(() => {
                        codeNum.value--;
                        if (codeNum.value == 0) {
                        clearInterval(clearId.value);
                        codeNum.value = 60;
                        isClickSend.value = false;
                        }
                    }, 1000);
                    Message.success({content:"发送成功！",id:"vcode",duration:2000})
                }else{
                    isClickSend.value = false;
                    Message.error({content:"发送失败",id:"vcode",duration:2000})
                }
         } catch (err) {
            Message.loading({content:"发送中，请稍后！",id:"vcode",duration:1})
            isClickSend.value = false;
         }
        }
    }
    return {
        visible,
        showLogin,
        changtype,logintype,
        dotype,
        //登录
        onLogin,logindata,
        formdata,handleSubmit,handleSend,codeNum,rulessecond,
        handleFreeSubmit,freeformdata,handleFreeSend,
    }
  }
})
</script>
<style lang="less" scoped>
   .login-box{
    height: 432px;
    .login-header{
        display: flex;
        .left{
            flex:1;
        }
        .right{
            user-select: none;
            .fonticon{
                    padding: 2px;
                    border-radius: 100%;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    cursor: pointer;
                }
            &:hover{
                .fonticon{
                    background-color: var(--color-neutral-3);
                    padding: 2px;
                    border-radius: 100%;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                }
            }
        }
    }
    .login-content{
        height: calc(100% - 22px);
        margin: 0 auto;
        display: flex;
        flex-wrap: wrap;
        align-content: space-between;
        .formdata{
            width:100%;
            .tabs{
                margin-top: 26px;
                display: flex;
                justify-content:space-between;
                padding-bottom: 8px;
                margin-bottom: 32px;
                .item{
                    flex:1;
                    text-align: center;
                    span{
                        cursor: pointer;
                        font-size: 16px;
                        font-weight: 600;
                        color: var(--color-neutral-6);
                        user-select: none;
                    }
                    .active{
                        color: var(--color-neutral-10);
                    }
                }
            }
        }
        .formbox{
            .passwod-login{
                text-align: center;
                .inputbox{
                    margin-bottom: 20px;
                }
                .loginbtn{
                    padding: 10px 25px;
                    margin-top: 25px;
                }
            }
            .free-login{
                display: flex;
                justify-content: center;
                .submintbtn{
                  padding: 10px 0px;
                }
            }
        }
        .footer{
            text-align: center;
            width:100%;
        }
    }
    //注册
    .registerbox{
        .regform{
            margin-top: 26px;
            .submintbtn{
                padding: 10px 25px;
            }
        }
        .options{
            margin-top: 15px;
        }
    }
   }
.options{
    width:100%;
    .register{
        width:100%;
        height: 24px;
        font-size: 13px;
        font-weight: 400;
        color: #999aaa;
        line-height: 24px;
        margin-bottom: 8px;
        text-align: center;
        cursor: pointer;
        user-select: none;
    }
}
</style>
    
    