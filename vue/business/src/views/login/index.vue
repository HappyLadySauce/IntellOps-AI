<template>
  <div class="login pc">
    <h3 class="login-logo" v-if="!isDesktop">
      <img  src="/logo.png" alt="logo" />
      <span>{{ title }}</span>
    </h3>
    <a-row align="stretch" class="login-box">
      <a-col :xs="0" :sm="12" :md="13" class="bgimg">
        <h3 class="logotitle">
          <a-avatar-group>
            <a-avatar style="background-color: var(--color-neutral-2)" >
              <img  src="/logo.png" alt="logo" />
            </a-avatar>
          </a-avatar-group>
          <span>{{ title }}</span>
        </h3>
        <div class="login-left">
          <img class="login-left__img" src="@/assets/images/banner.png" alt="banner" />
        </div>
      </a-col>
      <a-col :xs="24" :sm="12" :md="11">
        <div class="login-right">
          <h3 v-if="isEmailLogin" class="login-right__title">邮箱登录</h3>
          <EmailLogin v-show="isEmailLogin" />
          <a-tabs v-show="!isEmailLogin" class="login-right__form">
            <a-tab-pane key="1" :title="$t('login.form.tabacount')">
              <AccountLogin />
            </a-tab-pane>
            <a-tab-pane key="2" :title="$t('login.form.tabmobile')">
              <PhoneLogin />
            </a-tab-pane>
          </a-tabs>
          <div class="login-right__oauth">
            <a-divider orientation="center">{{ $t('login.form.other') }}</a-divider>
            <div class="list">
              <div v-if="isEmailLogin" class="mode item" @click="toggleLoginMode"><icon-user />{{ $t('login.form.tabacountmobile') }}</div>
              <div v-else class="mode item" @click="toggleLoginMode"><icon-email />{{ $t('login.form.tabemail') }}</div>
              <a class="item" title="使用 飞书 账号登录" @click="onOauth('lark')">
                <Icon icon="icon-lark-color" :size="24" />
              </a>
              <a class="item" title="使用 微信 账号登录" @click="onOauth('wechat')">
                <Icon icon="icon-wechat" :size="24" color="#28c445" />
              </a>
            </div>
          </div>
        </div>
      </a-col>
    </a-row>
    <div v-if="isDesktop" class="loginfooter">
      <Footer></Footer>
    </div>
    <GfThemeBtn class="theme-btn" />
    <GfLocaleBtn class="locale-btn" />
    <Background />
  </div>
</template>

<script setup lang="ts">
import { ref} from 'vue';
import Background from './components/background/index.vue'
import AccountLogin from './components/account/index.vue'
import PhoneLogin from './components/phone/index.vue'
import EmailLogin from './components/email/index.vue'
import GfThemeBtn from '@/components/handPlugin/Gf/GfThemeBtn/index.vue'
import GfLocaleBtn from '@/components/handPlugin/Gf/GfLocaleBtn/index.vue'
import { useDevice } from '@/hooks/useDevice'
import {Icon} from '@/components/Icon';
import Footer from '@/components/footer/index.vue';
import { Message } from '@arco-design/web-vue'
const { isDesktop } = useDevice()
const title = window?.globalConfig.AppTitle
const isEmailLogin = ref(false)
// 切换登录模式
const toggleLoginMode = () => {
  isEmailLogin.value = !isEmailLogin.value
}

// 第三方登录授权
const onOauth = async (source: string) => {
  Message.warning({content:'请根据需求接第三方登录授权',id:"socialAuth"})
  // const { data } = await socialAuth(source)
  // window.location.href = data.authorizeUrl
}
</script>

<style lang="less" scoped>
@media screen and (max-width: 570px) {
  .pc {
    background-color: white !important;
  }
  .login {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: start;
    align-items: center;
    background-color: var(--color-bg-5);
    color: #121314;
    &-logo {
      width: 100%;
      height: 104px;
      font-weight: 700;
      font-size: 20px;
      line-height: 32px;
      display: flex;
      padding: 0 20px;
      align-items: center;
      justify-content: start;
      background-size: 100% 100%;
      box-sizing: border-box;
      img {
        width: 34px;
        height: 34px;
        margin-right: 8px;
      }
    }
    &-box {
      width: 100%;
      display: flex;
      z-index: 999;
    }
  }
  .login-right {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    padding: 30px 30px 0;
    box-sizing: border-box;
    &__title {
      color: var(--color-text-1);
      font-weight: 500;
      font-size: 20px;
      line-height: 32px;
      margin-bottom: 20px;
    }
    &__form {
      :deep(.arco-tabs-nav-tab) {
        display: flex;
        justify-content: start;
        align-items: center;
      }
      :deep(.arco-tabs-tab) {
        color: var(--color-text-2);
        margin: 0 20px 0 0;
      }
      :deep(.arco-tabs-tab-title) {
        font-size: 16px;
        font-weight: 500;
        line-height: 22px;
      }
      :deep(.arco-tabs-content) {
        margin-top: 10px;
      }
      :deep(.arco-tabs-tab-active),
      :deep(.arco-tabs-tab-title:hover) {
        color: rgb(var(--arcoblue-6));
      }
      :deep(.arco-tabs-nav::before) {
        display: none;
      }
      :deep(.arco-tabs-tab-title:before) {
        display: none;
      }
    }
    &__oauth {
      width: 100%;
      position: fixed;
      bottom: 0;
      left: 0;
      padding-bottom: 20px;
      :deep(.arco-divider-text) {
        color: var(--color-text-4);
        font-size: 12px;
        font-weight: 400;
        line-height: 20px;
      }
      .list {
        align-items: center;
        display: flex;
        justify-content: center;
        width: 100%;
        .item {
          margin-right: 15px;
        }
        .mode {
          color: var(--color-text-2);
          font-size: 12px;
          font-weight: 400;
          line-height: 20px;
          padding: 6px 10px;
          align-items: center;
          border: 1px solid var(--color-border-3);
          border-radius: 32px;
          box-sizing: border-box;
          display: flex;
          height: 32px;
          justify-content: center;
          cursor: pointer;
          user-select: none;
          .icon {
            width: 21px;
            height: 20px;
          }
        }
        .mode svg {
          font-size: 16px;
          margin-right: 10px;
        }
        .mode:hover {
          background: rgba(var(--primary-6), 0.05);
          border: 1px solid rgb(var(--primary-3));
          color: rgb(var(--arcoblue-6));
        }
      }
    }
  }
  .locale-btn {
    position: fixed;
    top: 20px;
    right: 30px;
    z-index: 9999;
  }
  .theme-btn {
    position: fixed;
    top: 20px;
    right: 70px;
    z-index: 9999;
  }

  .loginfooter {
    align-items: center;
    box-sizing: border-box;
    position: absolute;
    bottom: 10px;
    z-index: 999;
    .footer{
      background-color:transparent;
    }
  }
}
@media screen and (min-width: 571px) {
  .login {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-color: var(--color-bg-5);
    &-logo {
      position: fixed;
      top: 20px;
      left: 30px;
      z-index: 9999;
      color: var(--color-text-1);
      font-weight: 500;
      font-size: 20px;
      line-height: 32px;
      margin-bottom: 20px;
      display: flex;
      justify-content: center;
      align-items: center;
      img {
        width: 34px;
        height: 34px;
        margin-right: 8px;
      }
    }
    &-box {
      width: 86%;
      max-width: 850px;
      height: 490px;
      display: flex;
      z-index: 999;
      box-shadow: 0 2px 4px 2px rgba(0, 0, 0, 0.08);
      border-radius: 5px;
      overflow: hidden;
    }
  }

  .login-left {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
    overflow: hidden;
    background: linear-gradient(60deg, rgb(var(--primary-6)), rgb(var(--primary-3)));
    &__img {
      width: 100%;
      position: absolute;
      bottom: 0;
      right: 0;
      top: 50%;
      left: 50%;
      transform: translateX(-50%) translateY(-50%);
      transition: all 0.3s;
      object-fit: cover;
    }
  }

  .login-right {
    width: 100%;
    height: 100%;
    background: var(--color-bg-1);
    display: flex;
    flex-direction: column;
    padding: 30px 30px 0;
    box-sizing: border-box;
    &__title {
      color: var(--color-text-1);
      font-weight: 500;
      font-size: 20px;
      line-height: 32px;
      margin-bottom: 20px;
    }
    &__form {
      :deep(.arco-tabs-nav-tab) {
        display: flex;
        justify-content: center;
        align-items: center;
      }
      :deep(.arco-tabs-tab) {
        color: var(--color-text-2);
      }
      :deep(.arco-tabs-tab-title) {
        font-size: 16px;
        font-weight: 500;
        line-height: 22px;
      }
      :deep(.arco-tabs-content) {
        margin-top: 10px;
      }
      :deep(.arco-tabs-tab-active),
      :deep(.arco-tabs-tab-title:hover) {
        color: rgb(var(--arcoblue-6));
      }
      :deep(.arco-tabs-nav::before) {
        display: none;
      }
      :deep(.arco-tabs-tab-title:before) {
        display: none;
      }
    }
    &__oauth {
      margin-top: auto;
      margin-bottom: 20px;
      :deep(.arco-divider-text) {
        color: var(--color-text-4);
        font-size: 12px;
        font-weight: 400;
        line-height: 20px;
      }
      .list {
        align-items: center;
        display: flex;
        justify-content: center;
        width: 100%;
        .item {
          margin-right: 15px;
        }
        .mode {
          color: var(--color-text-2);
          font-size: 12px;
          font-weight: 400;
          line-height: 20px;
          padding: 6px 10px;
          align-items: center;
          border: 1px solid var(--color-border-3);
          border-radius: 32px;
          box-sizing: border-box;
          display: flex;
          height: 32px;
          justify-content: center;
          cursor: pointer;
          user-select: none;
          .icon {
            width: 21px;
            height: 20px;
          }
        }
        .mode svg {
          font-size: 16px;
          margin-right: 10px;
        }
        .mode:hover {
          background: rgba(var(--primary-6), 0.05);
          border: 1px solid rgb(var(--primary-3));
          color: rgb(var(--arcoblue-6));
        }
      }
    }
  }
  .locale-btn {
    position: fixed;
    top: 20px;
    right: 30px;
    z-index: 9999;
  }
  .theme-btn {
    position: fixed;
    top: 20px;
    right: 70px;
    z-index: 9999;
  }

  .loginfooter {
    align-items: center;
    box-sizing: border-box;
    position: absolute;
    bottom: 10px;
    z-index: 999;
    .footer{
      background-color:transparent;
    }
  }
}
.bgimg{
  position: relative;
  .logotitle{
    z-index: 999;
    position: absolute;
    top: 20px;
    left: 0px;
    width: 100%;
    font-weight: 700;
    font-size: 21px;
    display: flex;
    color: #ffffff;
    text-align: center;
    padding-left: 20px;
    span{
      padding-left: 6px;
      height: 36px;
      line-height: 36px;
    }
  }
}
</style>
