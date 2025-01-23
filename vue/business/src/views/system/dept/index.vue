<template>
  <div class="container">
    <page-card breadcrumb scrollPage :isFullscreen="isFullscreen">
      <div class="table-toolbar flex flex-between">
        <div class="left">
          <a-space>
            <a-input :style="{width:'200px'}"  v-model="formModel.name" placeholder="搜索部门名称" allow-clear />
            <a-range-picker v-model="formModel.createtime" :style="{width:'230px'}" @change="fetchData"/>
            <a-select v-model="formModel.status"  :options="statusOptions" placeholder="状态" :style="{width:'90px'}" @change="fetchData"/>
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
        <tabletool :showbtn="['create','refresh','selectdensity','fullscreen']"
           @create="createData" @refresh="fetchData" @selectdensity="(data)=>size=data" @fullscreen="(data)=>isFullscreen=data"></tabletool>
        </div>
        </div>
      <a-table
         row-key="id"
        :loading="loading"
        :pagination="false"
        :columns="(cloneColumns as TableColumnData[])"
        :data="renderData"
        :bordered="{wrapper:true,cell:true}"
        :size="size"
        :default-expand-all-rows="true"
        ref="artable"
        :scroll="{x:'100%'}"
        @change="handleChange" 
      >
        <template #title="{ record }">
          <span v-html="record.spacer" style="padding-right: 5px;color: var(--color-neutral-4);"></span>{{ record.name }}
        </template>
        <template #icon="{ record }">
          <Icon :icon="record.icon" :size="20"></Icon>
        </template>
        <template #createtime="{ record }">
          {{dayjs(record.createtime).format("YYYY-MM-DD")}}
        </template>
        <template #status="{ record }">
          <a-switch type="round" v-model="record.status" :checked-value="0" :unchecked-value="1" @change="handleStatus(record)">
              <template #checked>
                开
              </template>
              <template #unchecked>
                关
              </template>
            </a-switch>
        </template>
        <template #operations="{ record }">
          <Icon icon="svgfont-bianji1" class="iconbtn" @click="handleEdit(record)" :size="18" color="#0960bd"></Icon>
          <a-divider direction="vertical" />
          <a-popconfirm content="您确定要删除吗?" @ok="handleDel(record)" position="tr">
            <Icon icon="svgfont-icon7" class="iconbtn" :size="18" color="#ed6f6f"></Icon>
          </a-popconfirm>
        </template>
      </a-table>
    </page-card>
    <!--表单-->
    <AddForm  @register="registerModal" @success="handleData"/>
  </div>
</template>

<script lang="ts" setup>
  import {  ref, watch, nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  import { getList,upStatus,del} from '@/api/system/dept';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import dayjs from 'dayjs';
  //数据
  import { columns} from './data';
  //表单
  import AddForm from './AddForm.vue';
  import { useModal } from '/@/components/Modal';
  import {Icon} from '@/components/Icon';
  import { Message } from '@arco-design/web-vue';
  import {tabletool,SizeProps,statusOptions} from '/@/components/tabletool';
  const [registerModal, { openModal }] = useModal();
  type Column = TableColumnData & { checked?: true };
  const { loading, setLoading } = useLoading(true);
  const renderData = ref([]);
  const cloneColumns = ref<Column[]>([]);
  const isFullscreen = ref(false);
  const size = ref<SizeProps>('large');
    //查询字段
    const generateFormModel = () => {
    return {
      name: '',
      createtime: [],
      status: '',
    };
  };
  const formModel = ref(generateFormModel());
  const artable=ref()
  const fetchData = async () => {
    setLoading(true);
    try {
      const data= await getList({...formModel.value});
      renderData.value = data;
      nextTick(()=>{
        artable.value.expandAll()
      })
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const reset = () => {
    formModel.value = generateFormModel();
    fetchData();
  };
  fetchData();

  watch(
    () => columns.value,
    (val) => {
      cloneColumns.value = cloneDeep(val);
    },
    { deep: true, immediate: true }
  );
  //添加
  const createData=()=>{
    openModal(true, {
      isUpdate: false,
      record:null
    });
  }
  //编辑数据
  const handleEdit=async(record:any)=>{
    openModal(true, {
      isUpdate: true,
      record:record
    });
  }
  //更新数据
  const handleData=async()=>{
    fetchData();
  }
  //排序拖拽
  const handleChange = (_data:any) => {
    renderData.value = _data
  }
  //更新状态
  const handleStatus=async(record:any)=>{
    try {
      Message.loading({content:"更新状态中",id:"upStatus",duration:0})
      await upStatus({id:record.id,status:record.status});
      Message.success({content:"更新状态成功",id:"upStatus",duration:2000})
    }catch (error) {
      record.status=record.status==1?0:1
      Message.loading({content:"更新状态中",id:"upStatus",duration:1})
    } 
  }
  //删除数据
  const handleDel=async(record:any)=>{
    try {
      Message.loading({content:"删除中",id:"del",duration:0})
      await del({ids:[record.id]});
      fetchData();
      Message.success({content:"删除成功",id:"del",duration:2000})
    }catch (error) {
      Message.loading({content:"删除中",id:"del",duration:1})
    } 
  }
</script>

<script lang="ts">
  export default {
    name: 'dept',//页签是名字和路由name相同则缓存生效
  };
</script>

<style scoped lang="less">
  :deep(.arco-table-th) {
    &:last-child {
      .arco-table-th-item-title {
        margin-left: 16px;
      }
    }
  }
  .action-icon {
    margin-left: 12px;
    cursor: pointer;
  }
  .active {
    color: #0960bd;
    background-color: #e3f4fc;
  }
  .setting {
    display: flex;
    align-items: center;
    width: 200px;
    .title {
      margin-left: 12px;
      cursor: pointer;
    }
  }
  :deep(.general-card > .arco-card-header){
    padding: 10px 16px;
  }
  .iconbtn{
    user-select: none;
    cursor: pointer;
    opacity: .8;
    &:hover{
      opacity: 1;
    }
  }
</style>
