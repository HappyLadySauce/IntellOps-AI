<template>
  <div class="container">
    <a-row wrap :gutter="16" align="stretch">
      <a-col :xs="24" :sm="24" :md="10" :lg="10" :xl="7" :xxl="7">
        <LeftBox :userdata="newuserdata"/>
      </a-col>
      <a-col :xs="24" :sm="24" :md="14" :lg="14" :xl="17" :xxl="17">
        <div>
          <PasswordPolicy :userdata="newuserdata"/>
        </div>
        <div style="margin-top: 16px">
          <RightBox />
        </div>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
import { ref,onMounted } from 'vue';
import LeftBox from './BasicInfo.vue'
import RightBox from './Social.vue'
import PasswordPolicy from './Security.vue'
import { getUserinfo,OptionUser} from './api'
const newuserdata=ref<OptionUser>({
  username:"",
  deptname:"",
  roles:"",
})
onMounted(async() => {
  newuserdata.value = await getUserinfo()
})
</script>

<script lang="ts">
  export default {
    name: 'usersetting',//页签是名字和路由name相同则缓存生效
  };
</script>