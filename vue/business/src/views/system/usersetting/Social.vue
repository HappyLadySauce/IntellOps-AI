<template>
  <a-card title="第三方账号" bordered class="gradient-card">
    <div v-for="item in modeList" :key="item.title">
      <div class="item">
        <div class="icon-wrapper"><Icon :icon="item.icon" :size="26" /></div>
        <div class="info">
          <div class="info-top">
            <span class="label">{{ item.title }}</span>
            <span class="bind">
              <icon-check-circle-fill v-if="item.status" :size="14" class="success" />
              <icon-exclamation-circle-fill v-else :size="14" class="warning" />
              <span style="font-size: 12px" :class="item.status ? 'success' : 'warning'">{{
                item.status ? '已绑定' : '未绑定'
              }}</span>
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
            {{ item.status ? '修改' : '绑定' }}
          </a-button>
          <a-button
            v-else-if="item.jumpMode === 'link'"
            class="btn"
            :type="item.status ? 'secondary' : 'primary'"
            @click="onBinding(item.type, item.status)"
          >
            {{ item.status ? '解绑' : '绑定' }}
          </a-button>
        </div>
      </div>
    </div>
  </a-card>
  <VerifyModel ref="verifyModelRef" />
</template>

<script setup lang="ts">
import { ref} from 'vue';
import type { ModeItem } from './type'
import VerifyModel from './components/VerifyModel.vue'

const socialList = ref<any>([])
const modeList = ref<ModeItem[]>([])
modeList.value = [
  {
    title: '绑定 Gitee',
    icon: 'icon-google-circle-fill',
    subtitle: `${socialList.value.includes('gitee') ? '' : '绑定后，'}可通过 Gitee 进行登录`,
    jumpMode: 'link',
    type: 'gitee',
    status: socialList.value.includes('gitee')
  },
  {
    title: '绑定 GitHub',
    icon: 'icon-github',
    subtitle: `${socialList.value.includes('gitee') ? '' : '绑定后，'}可通过 GitHub 进行登录`,
    type: 'github',
    jumpMode: 'link',
    status: socialList.value.includes('github')
  }
] as any

// 绑定
const onBinding = (type: string, status: boolean) => {
  if (!status) {
    window.open("https://gitee.com/login?redirect_to_url=https%3A%2F%2Fgitee.com%2Foauth%2Fauthorize%3Fresponse_type%3Dcode%26client_id%3D39efc49924f2505b1033372b82a46b2a21f55ff544c9d17005bffce50a649ad0%26redirect_uri%3Dhttps://api.goflys.cn/home%26state%3D93f8790827104c1098280da11b3b2ad4%26scope%3Duser_info", '_self')
  } else {
   
  }
}

const verifyModelRef = ref<InstanceType<typeof VerifyModel>>()
// 修改
const onUpdate = (type: string) => {
  verifyModelRef.value?.open(type)
}

</script>

<style lang="less" scoped></style>
