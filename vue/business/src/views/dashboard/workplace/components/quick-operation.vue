<template>
  <a-card
    class="general-card"
    :title="$t('workplace.quick.operation')"
    :header-style="{ paddingBottom: '0' }"
    :body-style="{ padding: '24px 20px 0 20px' }"
    :loading="loading"
  >
    <template #extra>
      <a-link @click="isedit=!isedit">{{isedit?"完成":"管理" }}</a-link>
    </template>
    <a-row :gutter="8">
      <a-col v-for="link in links" :key="link.name" :span="8" class="wrapper" @click="openAndEdit(link)">
        <a-popconfirm content="您确定要删除吗?" @ok="handleDel(link)" position="tr">
          <div class="del" @click.stop="()=>{}" v-if="isedit&&link.is_common==0"><icon-minus-circle-fill style="color: red;" size="15"/></div>
          </a-popconfirm>
        <div class="icon">
          <Icon :icon="link.icon" ></Icon>
        </div>
        <a-typography-paragraph class="text">
          {{ link.name }}
        </a-typography-paragraph>
      </a-col>
      <a-col v-if="isedit" :span="8" class="wrapper" @click="addEditQuick(null)">
        <div class="icon">
          <component is="icon-plus" />
        </div>
        <a-typography-paragraph class="text">
           添加
        </a-typography-paragraph>
      </a-col>
    </a-row>
    <a-divider class="split-line" style="margin: 0" />
     <QuickForm @register="registerModal" @success="fetchData"/>
  </a-card>
</template>

<script lang="ts" setup>
import { ref,onMounted } from 'vue';
import useLoading from '@/hooks/loading';
import { Message } from '@arco-design/web-vue';
import { useModal } from '/@/components/Modal';
import QuickForm from './QuickForm.vue';
import {Icon} from '@/components/Icon';
import { useRouter } from 'vue-router'
const  router = useRouter();
//api 
import { getQuick,delQuick,QuickItem} from '@/api/dashboard/workplace';
const [registerModal, { openModal }] = useModal();
const { loading, setLoading } = useLoading(true);
  onMounted(()=>{
    setTimeout(()=>{
      fetchData();
    },800)
  })
  const isedit=ref(false)
  const links=ref<QuickItem []>([])
  const fetchData = async () => {
    setLoading(true);
    try {
      links.value= await getQuick({});
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };
  //删除
  const handleDel=async(item:any)=>{
       Message.loading({content:"删除中",id:"upStatus",duration:0})
       const res= await delQuick({id:item.id});
       if(res){
        fetchData();
         Message.success({content:"删除成功",id:"upStatus",duration:2000})
       }
  }
  //打开或者编辑
  const openAndEdit=(item:any)=>{
   if(isedit.value){
    addEditQuick(item)
   }else{//跳转
      if(item.type==1){//外部
        window.open(item.path_url, '_blank');
      }else{//内部链接
        router.push({name:item.path_url})
      }
   }
  }
  //编辑会计按钮
  const addEditQuick=(item:any)=>{
      if(item==null){//新增
        openModal(true, {
          isUpdate: false,
          record:null
        });
      }else{//编辑
        openModal(true, {
          isUpdate: true,
          record:item
        });
      }
  }
</script>

<style scoped lang="less">
.wrapper{
  position: relative;
  .del{
    position: absolute;
    right: 15px;
    top:-10px
  }
}
</style>
