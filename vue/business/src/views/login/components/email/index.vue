<template>
  <a-form
    ref="formRef"
    :model="form"
    :rules="rules"
    :label-col-style="{ display: 'none' }"
    :wrapper-col-style="{ flex: 1 }"
    size="large"
    @submit="handleLogin"
  >
    <a-form-item field="email" hide-label>
      <a-input v-model="form.email" :placeholder="$t('login.form.getpemail.placeholder')"  allow-clear  >
        <template #prefix>
          <icon-email />
        </template>
      </a-input>
    </a-form-item>
    <a-form-item field="captcha" hide-label>
      <a-input v-model="form.captcha" :placeholder="$t('login.form.verification.placeholder')" :max-length="6" allow-clear style="flex: 1 1" />
      <a-button
        class="captcha-btn"
        :loading="captchaLoading"
        :disabled="captchaDisable"
        size="large"
        @click="onCaptcha"
      >
        {{ captchaBtnName }}
      </a-button>
    </a-form-item>
    <a-form-item>
      <a-space direction="vertical" fill class="w-full">
        <a-button :disabled="form.captcha.length!=6" class="btn" type="primary" :loading="loading" html-type="submit" size="large" long>{{$t('login.form.login')}}</a-button>
      </a-space>
    </a-form-item>
  </a-form>
</template>

<script setup lang="ts">
import { ref,reactive,} from 'vue';
import { type FormInstance, Message } from '@arco-design/web-vue'
import { useUserStore } from '@/store'
import { useRouter } from 'vue-router';
import * as Regexp from '@/utils/regexp'
import { getCaptcha } from '@/api/user';
import { useI18n } from 'vue-i18n';
const { t } = useI18n();
const formRef = ref<FormInstance>()
const form = reactive({
  email: "",
  captcha: ""
})
const rules: FormInstance['rules'] = {
  email: [
    { required: true, message: '请输入邮箱' },
    { match: Regexp.Email, message: '请输入正确的邮箱' }
  ],
  captcha: [{ required: true, message: '请输入验证码' }]
}

const userStore = useUserStore()
const router = useRouter()
const loading = ref(false)
// 登录
const handleLogin = async () => {
  try {
    const isInvalid = await formRef.value?.validate()
    if (isInvalid) return
    loading.value = true
    await userStore.login({
      email: form.email,
      captcha: form.captcha,
    })
    const { redirect, ...othersQuery } = router.currentRoute.value.query;
    var toURl=(redirect as string)
    if(toURl=="notFound"){
        toURl="home"
    }
    router.replace({
      name: toURl || 'home',
      query: {
        ...othersQuery,
      },
    });
    Message.success({content:t('login.form.login.success'),id:"errmsg"})
  } catch (error) {
    form.captcha = ""
  } finally {
    loading.value = false
  }
}

const captchaTimer = ref()
const captchaTime = ref(60)
const captchaBtnName = ref('获取验证码')
const captchaDisable = ref(false)
// 重置验证码
const resetCaptcha = () => {
  window.clearInterval(captchaTimer.value)
  captchaTime.value = 60
  captchaBtnName.value = '获取验证码'
  captchaDisable.value = false
}

const captchaLoading = ref(false)
// 获取验证码
const onCaptcha = async () => {
  if (captchaLoading.value) return
  const isInvalid = await formRef.value?.validateField('email')
  if (isInvalid) return
  try {
    captchaLoading.value = true
    captchaBtnName.value = '发送中...'
    await getCaptcha({
      type:"email",
      email: form.email
    })
    captchaLoading.value = false
    captchaDisable.value = true
    captchaBtnName.value = `获取验证码(${(captchaTime.value -= 1)}s)`
    Message.success({content:'邮件发送成功',id:"email"})
    captchaTimer.value = window.setInterval(() => {
      captchaTime.value -= 1
      captchaBtnName.value = `获取验证码(${captchaTime.value}s)`
      if (captchaTime.value <= 0) {
        resetCaptcha()
      }
    }, 1000)
  } catch (error) {
    resetCaptcha()
  } finally {
    captchaLoading.value = false
  }
}
</script>

<style lang="less" scoped>
.arco-input-wrapper,
:deep(.arco-select-view-single) {
  height: 40px;
  border-radius: 4px;
  font-size: 13px;
}

.arco-input-wrapper.arco-input-error {
  background-color: rgb(var(--danger-1));
  border-color: rgb(var(--danger-3));
}
.arco-input-wrapper.arco-input-error:hover {
  background-color: rgb(var(--danger-1));
  border-color: rgb(var(--danger-6));
}

.arco-input-wrapper :deep(.arco-input) {
  font-size: 13px;
  color: var(--color-text-1);
}
.arco-input-wrapper:hover {
  border-color: rgb(var(--arcoblue-6));
}

.captcha-btn {
  height: 40px;
  margin-left: 12px;
  min-width: 98px;
  border-radius: 4px;
}

.btn {
  height: 40px;
}
</style>
