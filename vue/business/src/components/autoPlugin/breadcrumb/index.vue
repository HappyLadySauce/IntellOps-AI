<template>
  <a-breadcrumb class="container-breadcrumb" v-if="breadcrumb">
    <a-breadcrumb-item >
      <a-link  @click="handleNaction('/home')"><icon-home /></a-link>
    </a-breadcrumb-item>
    <template v-if="items&&items.length>0">
      <a-breadcrumb-item v-for="item in items" :key="item">
        {{ $t(item) }}
      </a-breadcrumb-item>
    </template>
    <template v-else>
      <a-breadcrumb-item v-for="item in routeList">
        <a-link v-if="item.redirect&&item.redirect!=route.path" @click="handleNaction(item.redirect)">
          {{ item.meta.locale?$t(item.meta.locale):item.meta.title }}
        </a-link>
        <template v-else>{{ item.meta.locale?$t(item.meta.locale):item.meta.title }}</template>
      </a-breadcrumb-item>
    </template>
    <slot></slot>
  </a-breadcrumb>
</template>

<script lang="ts" setup>
  import { PropType,computed} from 'vue';
  import { useAppStore } from '@/store';
  import { useRoute,useRouter} from 'vue-router'
  const route = useRoute(),router=useRouter()
  const routeList = computed(() => route.matched.filter((item) => item.path!=""));
  const appStore = useAppStore();
  const breadcrumb = computed(() => appStore.breadcrumb);
  defineProps({
    items: {
      type: Array as PropType<string[]>,
      default() {
        return [];
      },
    },
  });
  //跳转路由
  const handleNaction=(path:any)=>{
    router.push({path:path})
  }
</script>

<style scoped lang="less">
  .container-breadcrumb {
    padding-bottom: 14px;
    :deep(.arco-breadcrumb-item) {
      color: rgb(var(--gray-6));
      &:last-child {
        color: rgb(var(--gray-8));
      }
    }
  }
</style>
