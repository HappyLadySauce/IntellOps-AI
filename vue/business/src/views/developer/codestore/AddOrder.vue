<template>
<a-modal v-model:visible="visible" unmountOnClose :footer="false" :closable="false" :width="500" :modal-style="{minHeight: '550px'}" >
   <div class="login-box">
        <div class="login-header">
            <div class="left"></div>
            <div class="right" @click="visible=false">
                <icon-close :size="23" class="closebtn"/>
            </div>
        </div>
        <div class="login-content"  >
            <div class="formdata">
                <div class="datainfo">
                    <div class="filed-item">
                        <div class="label">名称：</div>
                        <div class="text">{{ orderdata.title }}</div>
                    </div>
                    <div class="filed-item">
                        <div class="label">支付金额：</div>
                        <div class="text" style="color:red;font-size: 16px;">{{ orderdata.price }}<span class="unit">元</span></div>
                    </div>
                    <div class="filed-item">
                        <div class="label">收款账号：</div>
                        <div class="text" style="color:#165dff;">{{ orderdata.paycode.pay_type=="self"?"插件开发者自己收款":"GoFly平台收款" }}</div>
                    </div>
                    <div class="texttig"> 
                       选择支付方式，输入插件金额付款完成，点击“立即提交订单”，社区接收都订单后核对支付金额处理订单，处理后即可点击“安装”使用插件。平台7x24小时都可以处理订单，提交后会在10秒~1分钟处理
                    </div>
                </div>
                <div class="formbox">
                    <div class="paybox" v-if="orderdata.paycode.pay_type=='self'">
                        <div class="paymind">
                            <a-image
                                width="220"
                                :preview="false"
                                :src="orderdata.paycode.alipay"
                            />
                        </div>
                        <div class="paymind">
                            <a-image
                                width="220"
                                :preview="false"
                                :src="orderdata.paycode.wxapy"
                            />
                        </div>
                    </div>
                    <div class="paybox" v-else>
                        <div class="paymind">
                            <a-image
                                width="220"
                                :preview="false"
                                src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/staticfile/alapy2.png"
                            />
                        </div>
                        <div class="paymind">
                            <a-image
                                width="220"
                                :preview="false"
                                src="https://api.goflys.cn/common/uploadfile/get_image?url=resource/staticfile/wxpay.jpg"
                            />
                        </div>
                    </div>
                </div>
            </div>
            <div class="footer">
                <div >推荐使用支付宝、微信转账有限制，可以先付款再提交订单、也可先提交订单再付款、我们根据付款备注处理订单<span style="color:red;">（本次是{{ orderdata.paycode.pay_type=="self"?"插件作者自己收款":"GoFly平台收款" }}）</span></div>
                <div class="loginbtn">
                    <a-button type="primary" long shape="round" size="large" @click="onLogin">立即提交订单</a-button>
                </div>
               <div class="agree">
                  购买即代表同意<a-link href="https://goflys.cn/ardetail?id=12"  target="_blank">gofly快速开发框架插件协议</a-link>
               </div>
            </div>
        </div>
   </div>
</a-modal>
</template>
<script lang="ts">
import { defineComponent,ref} from "vue";
import { Modal,Input,Button,Image,Link,Message} from '@arco-design/web-vue';
//api
import { addOrder } from '@/api/developer/codestore';
export default defineComponent({
name: "AddOrder",
components: {
    AModal: Modal,
    AButton: Button,
    AInput: Input,
    ALink: Link,
    AImage: Image,
},
emits: ['ok'],
setup(_, { emit }) {
    //数据存储
    const visible=ref(false)
    const orderdata=ref({
        codestore_id:0,
        code_token:"",
        type:"",
        title:"",
        price:"",
        name:"",
        paycode:{alipay:"",wxapy:"",pay_type:""},
    })
    //展示
    const baseurldata=ref("")
    const showModel=(data:any,baseurl:any)=>{
        visible.value=true
        baseurldata.value=baseurl
        orderdata.value=Object.assign({},orderdata.value,data)
    }
    //提交订单
    const onLogin=async()=>{
        try {
            Message.loading({content:"提交中",id:"vcode",duration:0})
            const resultdata = await addOrder(orderdata.value,baseurldata.value);
            if(resultdata){
            Message.success({content:"提交成功",id:"vcode",duration:2000})
            emit('ok');
            visible.value=false
            }else{
                Message.clear()
            }
        } catch (error) {
            Message.loading({content:"提交中",id:"vcode",duration:1})
        }
    }
    return {
        visible,orderdata,
        showModel,
        //登录
        onLogin,
    }
  }
})
</script>
<style lang="less" scoped>
   .login-box{
    .login-header{
        display: flex;
        margin-top: -10px;
        .left{
            flex:1;
            padding-left: 10px;
            font-weight: 700;
        }
        .right{
            user-select: none;
            .closebtn{
                padding: 3px;
                border-radius: 100%;
                display: flex;
                justify-content: center;
                align-items: center;
                cursor: pointer;
                color:var(--color-neutral-6);
            }
            &:hover{
                .closebtn{
                    background-color: var(--color-neutral-3);
                    padding: 3px;
                    border-radius: 100%;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    color:var(--color-neutral-8);
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
            .datainfo{
                margin-top: 2px;
                margin-bottom: 18px;
                padding-left: 10px;
                .filed-item{
                    display: flex;
                    align-items: center;
                    .label{
                        font-weight: 600;
                    }
                    .text{
                        color: var(--color-neutral-8);
                        .unit{
                            font-size: 12px;
                            padding-left: 3px;
                            color: rgb(var(--red-5));
                        }
                    }
                }
                .texttig{
                    color: var(--color-neutral-6);
                    margin-top: 3px;
                    font-size: 12px;
                }
            }
        }
        .formbox{
            .passwod-login{
                text-align: center;
            }
            .paybox{
                width: 100%;
                display: flex;
                .paymind{
                    text-align: center;
                    &:last-child{
                        margin-left: 20px;
                    }
                }
            }
        }
        .footer{
            width:100%;
            .loginbtn{
                padding: 10px 25px;
                margin-top: 10px;
            }
            .agree{
                text-align: center;
            }
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
    }
   }
</style>
    
    