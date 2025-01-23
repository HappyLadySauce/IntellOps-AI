<template>
  <a-card :title="$t('userSetting.title.security')" bordered class="gradient-card">
    <div v-for="item in modeList" :key="item.title">
      <div class="item">
        <div class="icon-wrapper"><Icon :icon="item.icon" :size="26" color="rgb(var(--primary-5))"/></div>
        <div class="info">
          <div class="info-top">
            <span class="label">{{ item.title }}</span>
            <span class="bind">
              <icon-check-circle-fill v-if="item.status" :size="14" class="success" />
              <icon-exclamation-circle-fill v-else :size="14" class="warning" />
              <span style="font-size: 12px" :class="item.status ? 'success' : 'warning'">{{ item.statusString }}</span>
            </span>
          </div>
          <div class="info-desc">
            <span class="value">{{ item.value }}</span>
            {{ item.subtitle }}
          </div>
        </div>
        <div class="btn-wrapper">
          <a-button
            v-if="item.jumpMode === 'modal'"
            class="btn"
            :type="item.status ? 'secondary' : 'primary'"
            @click="onUpdate(item.type)"
          >
            {{ ['password'].includes(item.type) || item.status ? '修改' : '绑定' }}
          </a-button>
        </div>
      </div>
    </div>
  </a-card>
  <VerifyModel ref="verifyModelRef" />
</template>

<script lang="ts" setup>
import { ref,PropType} from 'vue';
import type { ModeItem } from './type'
import VerifyModel from './components/VerifyModel.vue'
import { OptionUser  } from './api'
import { useUserStore } from '@/store'
const userStore = useUserStore()
 defineProps({
  userdata: {
    type: Object as PropType<OptionUser>,
    default() {
      return {
      username:"",
      deptname:"",
      roles:"",
      pwd_reset_time:"",
    }
    },
  },
})
const modeList = ref<ModeItem[]>([])
modeList.value = [
    {
      title: '安全手机',
      icon: 'icon-phone',
      value: `${`${userStore.$state.mobile} ` || '手机号'}`,
      subtitle: `可用于身份验证、密码找回、通知接收`,
      type: 'mobile',
      jumpMode: 'modal',
      status: !!userStore.$state.mobile,
      statusString: userStore.$state.mobile ? '已绑定' : '未绑定'
    },
    {
      title: '安全邮箱',
      icon: 'icon-email',
      value: `${`${userStore.$state.email} ` || '邮箱'}`,
      subtitle: `可用于身份验证、密码找回、通知接收`,
      type: 'email',
      jumpMode: 'modal',
      status: !!userStore.$state.email,
      statusString: userStore.$state.email ? '已绑定' : '未绑定'
    },
    {
      title: '登录密码',
      icon: 'icon-lock',
      subtitle: userStore.$state.pwd_reset_time ? `为了您的账号安全，建议定期修改密码` : '请设置密码，可通过账号+密码登录',
      type: 'password',
      jumpMode: 'modal',
      status: !!userStore.$state.pwd_reset_time,
      statusString: userStore.$state.pwd_reset_time ? '已设置' : '未设置'
    }
  ]

const verifyModelRef = ref<InstanceType<typeof VerifyModel>>()
// 修改
const onUpdate = (type: string) => {
  verifyModelRef.value?.open(type)
}

</script>

<style lang="less" scoped></style>
