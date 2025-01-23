<template>
<div class="custom-full-table">
  <div class="table-toolbar flex flex-between">
    <div class="left">
      <a-space>
        <a-input v-model="formModel.user" placeholder="请输入登录用户" :style="{width:'180px'}" allow-clear>
          <template #prefix><icon-search /></template>
        </a-input>
        <a-input v-model="formModel.ip" placeholder="请输入登录 IP 或地点" :style="{width:'190px'}" allow-clear>
          <template #prefix><icon-search /></template>
        </a-input>
        <a-range-picker v-model="formModel.createtime" format="YYYY-MM-DD HH:mm" :show-time="true" :shortcuts="shortcuts" shortcuts-position="left" :style="{width:'310px'}" @change="fetchData"/>
        <a-select v-model="formModel.status"  :options="statusOptions" placeholder="状态" :style="{width:'120px'}"  @change="fetchData"/>
        <a-button type="primary" @click="fetchData">
          <template #icon>
            <icon-search />
          </template>
          查询
        </a-button>
        <a-button @click="reset">
          {{ $t('searchTable.form.reset') }}
        </a-button>
      </a-space>
    </div>
    <div class="right">
      <tabletool :showbtn="['refresh','settingcolumn','fullscreen']" :columns="columns" :disabledColumnKeys="['id']"
      @settingcolumn="(data)=>cloneColumns=data"
        @refresh="fetchData"  @fullscreen="handleFullscreen">
        <a-tooltip content="删除一个月之前的登录日志" >
            <a-button v-permission="'delLastLogin'" @click="handleDelLastLogin" class="gf_hover_btn-border">
            <template #default>删除日志</template>
          </a-button>
        </a-tooltip>
      </tabletool>
    </div>
  </div>
  <div class="table-container">
  <a-table
      row-key="id"
      :loading="loading"
      :scrollbar="true"
      :pagination="pagination"
      :columns="(cloneColumns as TableColumnData[])"
      :data="renderData"
      :scroll="{ x: '100%', y: '100%', minWidth: 1000 }"
      ref="artable"
      @page-change="handlePaageChange" 
      @page-size-change="handlePaageSizeChange" 
    >
      <template #user="{record}">
        <a-popover >
          <template #content>
            <div style="min-width: 180px;">
              <a-descriptions :column="1">
                <a-descriptions-item label="用户ID">{{ record.uid}}</a-descriptions-item>
                <a-descriptions-item label="用户名">{{ record.user.name}}</a-descriptions-item>
                <a-descriptions-item label="昵  称">{{ record.user.nickname}}</a-descriptions-item>
                <a-descriptions-item label="账  号"><a-typography-paragraph copyable>{{ record.user.username}}</a-typography-paragraph></a-descriptions-item>
              </a-descriptions>
            </div>
          </template>
          <CellAvatar v-if="record.user" :spacesize="1" :size="20"  :avatar="record.user.avatar" :name="record.user.name" is-link   />
        </a-popover>
      </template>
      <template #status="{record,column}">
        <a-tag v-if="record[column.dataIndex] === 0" color="green">
          <GfDot type="success" style="width: 5px; height: 5px" />
          <span style="margin-left: 5px">成功</span>
        </a-tag>
        <a-tooltip v-else :content="record.error_msg">
          <a-tag color="red" style="cursor: pointer">
            <GfDot type="danger" style="width: 5px; height: 5px" />
            <span style="margin-left: 5px">失败</span>
          </a-tag>
        </a-tooltip>
      </template>
      <template #createtime="{ record }">
        {{dayjs(record.createtime).format("YYYY-MM-DD HH:mm")}}
      </template>
      <template #user_agent="{record,column}">
        <a-popover >
          <a-link status="success">查看</a-link>
          <template #content>
            <div style="width: 260px;">
              {{ record[column.dataIndex] }}
            </div>
          </template>
        </a-popover>
      </template>
    </a-table>
  </div>
</div>
</template>

<script setup lang="ts">
  import { ref, reactive,onMounted,watch} from 'vue';
  import  { Message,Modal } from '@arco-design/web-vue'
  import { shortcuts, dayjs} from '@/utils/dayjs';
import useLoading from '@/hooks/loading';
import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
import { Pagination } from '@/types/global';
import { columns,statusOptions } from './data';
import { getLogin ,delLastLogin} from '../api'
import {tabletool} from '/@/components/tabletool';
import cloneDeep from 'lodash/cloneDeep';
defineProps({
  modelValue: {
    type: Boolean,
  },
})
const emits = defineEmits(['update:modelValue'])
const { loading, setLoading } = useLoading(true);
const renderData = ref([]);
const cloneColumns = ref<TableColumnData[]>([]);
  //分页
  const basePagination: Pagination = {
    current: 1,
    pageSize: 10,
  };
  const pagination = reactive({
    ...basePagination,
    showTotal:true,
    showPageSize:true,
  });
  //查询字段
  const generateFormModel = () => {
    return {
      user: '',
      ip: '',
      createtime: [],
      status: '',
    };
  };
  const formModel = ref(generateFormModel());
  const fetchData = async () => {
    setLoading(true);
    try {
      const data= await getLogin({page:pagination.current,pageSize:pagination.pageSize,...formModel.value});
      renderData.value = data.items;
      pagination.current = data.page;
      pagination.total = data.total;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };
  //分页
  const handlePaageChange = (page:any) => {
    pagination.current=page
    fetchData();
  }
  //分页总数
  const handlePaageSizeChange = (pageSize:any) => {
    pagination.pageSize=pageSize
    fetchData();
  }
  const reset = () => {
    formModel.value = generateFormModel();
    fetchData();
  };
  const handleFullscreen = (data:boolean) => {
    emits('update:modelValue', data)
  };
  //删除一个月之前日志数据
  const handleDelLastLogin = () => {
    Modal.warning({
      title: '是否确认删除一个月前的日志？',
      content: '删除后将无法恢复，敏感数据需要谨慎操作哦！',
      hideCancel:false,
      titleAlign:"start",
      onOk:async()=>{
        try {
          Message.loading({content:"删除中",id:"del",duration:0})
          await delLastLogin();
          fetchData();
          Message.success({content:"删除成功",id:"del",duration:2000})
        }catch (error) {
          Message.loading({content:"删除中",id:"del",duration:1})
        } 
      }
    });
  };
  //初始化
  onMounted(()=>{
    fetchData();
  })
  watch(
    () => columns.value,
    (val) => {
      cloneColumns.value = cloneDeep(val);
    },
    { deep: true, immediate: true }
  );
</script>

<style lang="less" scoped>
@import '@/assets/style/fulltable.less';
</style>
  