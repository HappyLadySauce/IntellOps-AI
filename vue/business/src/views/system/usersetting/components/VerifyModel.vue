<template>
  <a-modal v-model:visible="visible" :title="title" :ok-loading="okloading" @before-ok="save" @close="reset">
    <a-form :model="formData" ref="formRef" auto-label-width size="large" >
        <a-form-item v-if="verifyType=='mobile'" field="mobile" :label="$t('userSetting.label.mobile')" :rules="[{required:true,message:$t('userSetting.placeholder.Please')+$t('userSetting.label.mobile')}, { match: Regexp.Phone, message: '请输入正确的手机号' }]">
          <a-input v-model="formData.mobile" :placeholder="$t('userSetting.placeholder.Fill')+$t('userSetting.label.mobile')" allow-clear/>
        </a-form-item>
        <a-form-item v-if="verifyType=='email'" field="email" :label="$t('userSetting.label.email')" :rules="[{required:true,message:$t('userSetting.placeholder.Please')+$t('userSetting.label.email')},{ match: Regexp.Email, message: '请输入正确的邮箱' }]">
          <a-input v-model="formData.email" :placeholder="$t('userSetting.placeholder.Fill')+$t('userSetting.label.email')" allow-clear/>
        </a-form-item>
        <a-form-item v-if="verifyType=='mobile'||verifyType=='email'" field="captcha" label="验证码" :rules="[{required:true,message:'请输入验证码'}]">
          <a-input v-model="formData.captcha" placeholder="输入验证码" :max-length="6" allow-clear style="flex: 1 1" />
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
        <a-form-item field="oldpassword" label="当前密码" :rules="[{required:true,message:'请输入当前密码'}]">
          <a-input-password v-model="formData.oldpassword" placeholder="输入当前密码" ></a-input-password>
        </a-form-item>
        <template v-if="verifyType=='password'">
          <a-form-item field="newpassword" label="新密码" :rules="[{required:true,message:'请输入新密码'}]">
            <a-input-password v-model="formData.newpassword" placeholder="输入新密码" ></a-input-password>
          </a-form-item>
          <a-form-item field="repassword" label="确认新密码" :rules="[{required:true,message:'请输入确认新密码'}]">
            <a-input-password v-model="formData.repassword" placeholder="输入确认新密码" ></a-input-password>
          </a-form-item>
        </template>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { ref,computed} from 'vue';
import { Message } from '@arco-design/web-vue'
import { getCaptcha } from '@/api/user';
import { saveInfo } from '../api';
import * as Regexp from '@/utils/regexp'
const okloading=ref(false)
const formRef = ref()
const verifyType = ref()
const title = computed(
  () => `修改${verifyType.value === 'mobile' ? '手机号' : verifyType.value === 'email' ? '邮箱' : '密码'}`
)
const formData = ref({
  type:"",
  mobile:"",
  email:"",
  captcha:"",
  oldpassword:"",
  newpassword:"",
  repassword:"",
})
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

// 重置
const reset = () => {
  formRef.value?.resetFields()
  resetCaptcha()
}

const captchaLoading = ref(false)
// 获取验证码
const onCaptcha = async () => {
  const isInvalid = await formRef.value?.validateField(verifyType.value === 'mobile' ? 'mobile' : 'email')
  if (isInvalid) return false
  // 发送验证码
  try {
    captchaLoading.value = true
    captchaBtnName.value = '发送中...'
    if (verifyType.value === 'phone') {
      await getCaptcha({
        type:"mobile",
        mobile: formData.value.mobile
      })
    } else if (verifyType.value === 'email') {
      await getCaptcha({
        type:"email",
        email: formData.value.email
      })
    }
    captchaLoading.value = false
    captchaDisable.value = true
    captchaBtnName.value = `获取验证码(${(captchaTime.value -= 1)}s)`
    Message.success('发送成功')
    // Message.success('仅提供效果演示，实际使用请查看代码取消相关注释')
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

// 保存
const save = async () => {
  const isInvalid = await formRef.value?.validate();
  if (isInvalid) return false
  try {
    okloading.value=true
    if (verifyType.value === 'password') {
      if (formData.value.newpassword !== formData.value.repassword) {
        Message.error('两次新密码不一致')
        okloading.value=false
        return false
      }
      if (formData.value.newpassword  === formData.value.oldpassword ) {
        Message.error('新密码与旧密码不能相同')
        okloading.value=false
        return false
      }
    }
    await saveInfo(formData.value)
    okloading.value=false
    Message.success('修改成功')
    // 修改成功后，重新获取用户信息
    return true
  } catch (error) {
    okloading.value=false
    return false
  }
}

const visible = ref(false)
// 打开弹框
const open = (type: string) => {
  verifyType.value = type
  formData.value.type=type
  visible.value = true
}

defineExpose({ open })
</script>

<style lang="less" scoped>
.captcha-btn {
  margin-left: 12px;
  min-width: 98px;
  border-radius: 4px;
}
</style>
