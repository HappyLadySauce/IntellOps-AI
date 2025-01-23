<template>
  <div class="container" >
    <page-card breadcrumb fullTable :isFullscreen="isFullscreen">
       <template #searchLeft>
          <a-space>
            <a-input :style="{width:'200px'}"  v-model="formModel.name" placeholder="搜索用户名" allow-clear />
            <a-range-picker v-model="formModel.createtime" :shortcuts="shortcuts" shortcuts-position="left" :style="{width:'230px'}" @change="fetchData"/>
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
      </template>
      <template #searchRight>
        <tabletool :showbtn="['create','refresh','selectdensity','fullscreen']"
        @create="createData" @refresh="fetchData" @selectdensity="(data)=>size=data" @fullscreen="(data)=>isFullscreen=data"></tabletool>
      </template>
      <template #table>
        <a-table
          row-key="id"
          :loading="loading"
          :pagination="pagination"
          :columns="(cloneColumns as TableColumnData[])"
          :data="renderData"
          :bordered="{wrapper:true,cell:true}"
          :size="size"
          :default-expand-all-rows="true"
          ref="artable"
          :scroll="{x:'100%', y: '100%'}"
          @page-change="handlePaageChange" 
          @page-size-change="handlePaageSizeChange" 
        >
          <template #name="{ record }">
            <CellNodeAvatar :avatar="record.avatar" :size="20" :name="record.name" />
          </template>
          <template #createtime="{ record }">
            {{dayjs(record.createtime*1000).format("YYYY-MM-DD")}}
          </template>
          <template #status="{ record }">
            <cell-status v-if="record.id==userStore.business_id||record.id==userStore.id" :status="record.status"/>
            <a-switch v-else type="round" v-model="record.status" :checked-value="0" :unchecked-value="1" @change="handleStatus(record)">
              <template #checked>开</template>
              <template #unchecked>关</template>
            </a-switch>
          </template>
          <template #operations="{ record }">
            <template v-if="record.id==userStore.id">
              <span style="font-size: 12px" class="success">本账号</span>
            </template>
            <template v-else-if="record.id==userStore.business_id">
              <span style="font-size: 12px" class="success">主账号</span>
            </template>
            <template v-else>
              <Icon icon="svgfont-bianji1" class="iconbtn" @click="handleEdit(record)" :size="18" color="#0960bd"></Icon>
              <a-divider direction="vertical" />
              <a-popconfirm content="您确定要删除吗?" @ok="handleDel(record)" position="tr">
                <Icon icon="svgfont-icon7" class="iconbtn" :size="18" color="#ed6f6f"></Icon>
              </a-popconfirm>
            </template>
          </template>
        </a-table>
      </template>
    </page-card>
    <!--表单-->
    <AddForm @register="registerModal"  @success="fetchData"/>
  </div>
</template>

<script lang="ts" setup>
  import { ref, reactive, watch } from 'vue';
  import useLoading from '@/hooks/loading';
  import { getList,upStatus,del} from '@/api/system/account';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import { shortcuts, dayjs} from '@/utils/dayjs';
  //数据
  import { columns} from './data';
  //表单
  import { useModal } from '/@/components/Modal';
  import AddForm from './AddForm.vue';
  import {Icon} from '@/components/Icon';
  import { Message } from '@arco-design/web-vue';
  import { Pagination } from '@/types/global';
  import { useUserStore } from '@/store'
  import {tabletool,SizeProps,statusOptions} from '/@/components/tabletool';
  const userStore = useUserStore()
  const isFullscreen = ref<boolean>(false);
  const [registerModal, { openModal }] = useModal();
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
  type Column = TableColumnData & { checked?: true };
  const { loading, setLoading } = useLoading(true);
  const renderData = ref([]);
  const cloneColumns = ref<Column[]>([]);
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
  const fetchData = async () => {
    setLoading(true);
    try {
      const data= await getList({page:pagination.current,pageSize:pagination.pageSize,...formModel.value});
      renderData.value = data.items;
      pagination.current = data.page;
      pagination.total = data.total;
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
  //更新状态
  const handleStatus=async(record:any)=>{
    try {
        Message.loading({content:"更新状态中",id:"upStatus",duration:0})
        const res= await upStatus({id:record.id,status:record.status});
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
      const res= await del({ids:[record.id]});
      fetchData();
      Message.success({content:"删除成功",id:"del",duration:2000})
    }catch (error) {
      Message.loading({content:"删除中",id:"del",duration:2})
    } 
  }
</script>

<script lang="ts">
  export default {
    name: 'account',//页签是名字和路由name相同则缓存生效
  };
</script>

<style scoped lang="less">

</style>
