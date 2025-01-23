<template>
  <a-drawer :visible="modelValue" title="日志详情" :width="700" :footer="false"  @cancel="closeDrawer">
    <a-descriptions title="基本信息" :column="2" size="large" class="general-description">
      <a-descriptions-item label="日志 ID">{{ dataDetail?.id }}</a-descriptions-item>
      <a-descriptions-item label="操作人">{{ dataDetail?.username }}</a-descriptions-item>
      <a-descriptions-item label="操作时间">{{ dataDetail?.createtime }}</a-descriptions-item>
      <a-descriptions-item label="操作内容">{{ dataDetail?.des }}</a-descriptions-item>
      <a-descriptions-item label="操作 IP"><a-typography-paragraph :copyable="!!dataDetail?.ip">{{ dataDetail?.ip }}</a-typography-paragraph></a-descriptions-item>
      <a-descriptions-item label="操作地点">{{ dataDetail?.address.replace("|0|","|") }}</a-descriptions-item>
      <a-descriptions-item label="状 态">
        <a-tag v-if="dataDetail?.status === 200" color="green">成功</a-tag>
        <a-tag v-else color="red">失败</a-tag>
      </a-descriptions-item>
      <a-descriptions-item label="耗 时">
        <a-tag v-if="dataDetail?.latency > 500" color="red"> {{ dataDetail?.latency }}ms </a-tag>
        <a-tag v-else-if="dataDetail?.latency > 200" color="orange"> {{ dataDetail?.latency }}ms </a-tag>
        <a-tag v-else color="green">{{ dataDetail?.latency }} ms</a-tag>
      </a-descriptions-item>
      <a-descriptions-item label="请求 URI" :span="2">
        <a-typography-paragraph :copyable="!!dataDetail?.url" :copy-text="dataDetail?.url"><span v-if="dataDetail?.request_method" :style="{fontSize: '12px',color:FontColor(dataDetail.request_method)}">{{ dataDetail.request_method }}：</span>{{ dataDetail?.url }}</a-typography-paragraph>
      </a-descriptions-item>
    </a-descriptions>
    <a-descriptions
      title="请求信息"
      :column="2"
      size="large"
      class="general-description http"
      style="margin-top: 10px; position: relative"
    >
      <a-descriptions-item :span="2">
        <a-tabs type="card">
          <a-tab-pane key="1" title="请求体">
            <JsonPretty v-if="dataDetail?.req_body" :json="dataDetail?.req_body" />
            <span v-else>无</span>
          </a-tab-pane>
          <a-tab-pane key="2" title="请求头">
            <JsonPretty v-if="dataDetail?.req_headers" :json="dataDetail?.req_headers" />
            <span v-else>无</span>
          </a-tab-pane>
        
        </a-tabs>
      </a-descriptions-item>
    </a-descriptions>
    <a-descriptions
      title="响应信息"
      :column="2"
      size="large"
      class="general-description http"
      style="margin-top: 10px; position: relative"
    >
      <a-descriptions-item :span="2">
        <a-tabs type="card">
          <a-tab-pane key="1" title="响应体">
            <JsonPretty v-if="dataDetail?.resp_body" :json="dataDetail?.resp_body" />
            <span v-else>无</span>
          </a-tab-pane>
          <a-tab-pane key="2" title="响应头">
            <JsonPretty v-if="dataDetail?.resp_headers" :json="dataDetail?.resp_headers" />
            <span v-else>无</span>
          </a-tab-pane>
        </a-tabs>
      </a-descriptions-item>
    </a-descriptions>
  </a-drawer>
</template>

<script lang="ts" setup>
 import { ref,watch } from 'vue';
 import { getOperationDetail } from '../api'
 import JsonPretty from '@/components/handPlugin/JsonPretty/index.vue'
 const emits = defineEmits(['update:modelValue','ok'])
  let props = defineProps({
    record: { type: Object, default: () => null },
    modelValue: { type: Boolean, default: () => false },
  });

const dataDetail = ref<any>()

// 查询详情
const getDataDetail = async () => {
  dataDetail.value = await getOperationDetail({id:props.record.id})
}
//取消
const closeDrawer = () => {
   emits('update:modelValue',false)
};
//返回请求颜色
const FontColor = (text:string) => {
  if(text=="GET"){
    return "rgb(var(--arcoblue-5))"
  }else if(text=="POST"||text=="PUT"){
    return "rgb(var(--green-6))"
  }else if(text=="DELETE"){
    return "rgb(var(--orange-6))"
  }else{
    return "var(--color-neutral-8)"
  }
};
watch(
    () => props.record,
    (recordVal) => {
      if(recordVal)
       getDataDetail()
    },
    { immediate: true, deep: true}//如果我们的需求是需要在初始化绑定数据的时候就执行，就可以将watch函数的immediate属性值设为true,反之设为false
  )
</script>

<style lang="less" scoped>
.http :deep(.arco-descriptions-item-label-block) {
  padding-right: 0;
}
 :deep(.arco-descriptions-item-labe) {
  line-height: unset;
}

:deep(.arco-tabs-content) {
  padding-top: 5px;
  padding-left: 15px;
}
</style>
