<template>
  <a-card :title="$t('userSetting.title.basicInfo')" bordered class="gradient-card">
    <div class="body">
      <section>
        <a-avatar :size="100" @click="handleSelectImg">
          <template #trigger-icon><icon-camera /></template>
          <img v-if="userStore.avatar" :src="userStore.avatar" alt="avatar" />
        </a-avatar>
        <div class="name">
          <span style="margin-right: 10px">{{ userStore.nickname }}</span>
          <icon-edit :size="16" class="btn" @click="onUpdate" />
        </div>
        <div class="id">
          <Icon icon="svgfont-id" :size="16" />
          <span>{{ userStore.id }}</span>
        </div>
      </section>
      <footer>
        <a-descriptions :column="4" size="large">
          <a-descriptions-item :span="4">
            <template #label> <icon-idcard /><span style="margin-left: 5px">{{$t('userSetting.label.username')}}</span></template>
            {{ userdata.username }}
          </a-descriptions-item>
          <a-descriptions-item :span="4">
            <template #label> <icon-user /><span style="margin-left: 5px">{{$t('userSetting.label.name')}}</span></template>
            {{ userStore.name }}
            <!-- <icon-man v-if="userInfo.gender === 1" style="color: #19bbf1" />
            <icon-woman v-else-if="userInfo.gender === 2" style="color: #fa7fa9" /> -->
          </a-descriptions-item>
          <a-descriptions-item :span="4">
            <template #label> <icon-phone /><span style="margin-left: 5px">{{$t('userSetting.label.mobile')}}</span></template>
            {{ userStore.mobile }}
          </a-descriptions-item>
          <a-descriptions-item :span="4">
            <template #label> <icon-email /><span style="margin-left: 5px">{{$t('userSetting.label.email')}}</span></template>
            {{ userStore.email }}
          </a-descriptions-item>
          <a-descriptions-item :span="4">
            <template #label> <icon-mind-mapping /><span style="margin-left: 5px">{{$t('userSetting.label.department')}}</span></template>
            {{ userdata.deptname }}
          </a-descriptions-item>
          <a-descriptions-item :span="4">
            <template #label> <icon-user-group /><span style="margin-left: 5px">{{$t('userSetting.label.role')}}</span></template>
            {{ userdata.roles }}
          </a-descriptions-item>
        </a-descriptions>
      </footer>
    </div>
    <div class="footer">{{ $t('userSetting.label.registered') }} {{ userStore.createtime }}</div>
  </a-card>
  <BasicInfoUpdateModal :userdata="userdata" ref="BasicInfoUpdateModalRef" />
   <!--附件-->
   <FileManage @register="registerModal" @success="selectImg" />
</template>

<script setup lang="ts">
import { ref,PropType } from 'vue';
import { Message } from '@arco-design/web-vue'
import BasicInfoUpdateModal from './BasicInfoUpdateModal.vue'
import { useUserStore } from '@/store'
import { useModal } from '/@/components/Modal';
import FileManage from '@/components/attachment/FileManage.vue';
import { saveInfo,OptionUser } from './api'
const userStore = useUserStore()
import { GetFullPath } from "@/utils/tool";
const [registerModal, { openModal }] = useModal();
 defineProps({
  userdata: {
    type: Object as PropType<OptionUser>,
    default() {
      return {
      username:"",
      deptname:"",
      roles:"",
    }
    },
  },
})
// 选择头像
const handleSelectImg = async () => {
    openModal(true, {
      filetype:"image",
      getnumber: "one",//one 单张
      openfrom: "manage",
    });
  }
//选择附件返回
const selectImg=async(item:any)=>{
        userStore.avatar=GetFullPath(item.url)
        Message.loading({content:"更新头像",id:"delaction",duration:0})
        const upres = await saveInfo({avatar:item.url})
        if(upres){
          Message.success({content:"更新成功",id:"delaction",duration:2000})
        }else{
          Message.warning({content:"未更新",id:"delaction",duration:2000})
        }
  }

const BasicInfoUpdateModalRef = ref<InstanceType<typeof BasicInfoUpdateModal>>()
// 修改基本信息
const onUpdate =  () => {
  BasicInfoUpdateModalRef.value?.onUpdateUserinfo()
}
</script>

<style scoped lang="less">
:deep(.arco-avatar-trigger-icon-button) {
  width: 32px;
  height: 32px;
  line-height: 32px;
  background-color: #e8f3ff;
  .arco-icon-camera {
    margin-top: 8px;
    color: rgb(var(--arcoblue-6));
    font-size: 14px;
  }
}

.body {
  display: flex;
  flex-direction: column;
  padding: 28px 10px 20px 10px;
  .btn {
    cursor: pointer;
  }

  & > section {
    flex: 1 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 32px 0 50px;
    .name {
      font-size: 20px;
      margin: 20px 0;
    }
    .id {
      color: var(--color-text-3);
      span {
        font-size: 16px;
        font-weight: 400;
        line-height: 20px;
        padding: 0 6px;
      }
    }
  }

  & > footer .footer_item {
    display: flex;
    justify-content: space-between;
    margin-top: 10px;
    font-size: 12px;
  }
}

.footer {
  margin: 0 -16px;
  padding-top: 16px;
  font-size: 12px;
  text-align: center;
  border-top: 1px solid var(--color-border-2);
}

</style>
