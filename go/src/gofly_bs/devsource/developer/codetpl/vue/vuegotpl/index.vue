<template>
  <div class="container">
    <page-card breadcrumb>
        在这里编写你的代码吧
    </page-card>
  </div>
</template>

<script lang="ts" setup>
 import { ref, onMounted } from 'vue';
 import { getList,save,putStatus,del} from './api';
 onMounted(()=>{
  handleGetList()
 })
 const handleGetList=async()=>{
  const res = await getList({});
  console.log("获取数据请求测试",res)
 }
</script>
<script lang="ts">
  export default {
    name: '{vuetplname}', // If you want the include property of keep-alive to take effect, you must name the component
  };
</script>

<style lang="less" scoped>

</style>
