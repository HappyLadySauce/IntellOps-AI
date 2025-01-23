<template>
  <a-modal v-model:visible="visible" :title="$t('userSetting.title.EditbasicInfo')" :ok-loading="okloading" @before-ok="save" >
    <a-form :model="formData" ref="formRef" auto-label-width>
      <a-form-item field="nickname" :label="$t('userSetting.label.nickname')" :rules="[{required:true,message:$t('userSetting.placeholder.Please')+$t('userSetting.label.nickname')}]">
        <a-input v-model="formData.nickname" :placeholder="$t('userSetting.placeholder.Fill')+$t('userSetting.label.nickname')" allow-clear/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Message } from '@arco-design/web-vue'
import { useUserStore } from '@/store'
import { saveInfo } from './api'
import { useI18n } from 'vue-i18n';
const { t } = useI18n();
const userStore = useUserStore()
const formData = ref({
  nickname:userStore.$state.nickname,
})


const visible = ref(false)
// 修改
const onUpdateUserinfo = () => {
  visible.value = true
}

// 保存
const formRef = ref()
const okloading=ref(false)
const save = async () => {
  const isInvalid = await formRef.value?.validate();
  if (isInvalid) return false
  try {
    okloading.value=true
    await saveInfo({nickname:formData.value.nickname})
    userStore.nickname=formData.value.nickname
    Message.success(t('userSetting.saveSuccess'))
    okloading.value=false
    return true
  } catch (error) {
    return false
  }
}

defineExpose({ onUpdateUserinfo })
</script>
