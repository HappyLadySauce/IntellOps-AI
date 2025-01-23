<template>
  <div class="container" >
    <page-card breadcrumb :isFullscreen="isFullscreen">
      <div class="flexbox">
        <div class="menu flexcolumn">
          <div class="btn">
            <a-space>
              <a-button type="primary" @click="AddMenuData(1)" v-permission="'addcate'">
                <icon-plus /><span style="padding-left: 3px;">添加</span>
              </a-button>
              <a-button type="primary" status="warning" @click="AddMenuData(2)" v-permission="'addcate'">
                <icon-edit /><span style="padding-left: 3px;">修改</span>
              </a-button>
              <a-button type="primary" status="danger" @click="delMenuData" v-permission="'delcate'">
                <icon-delete /><span style="padding-left: 3px;">删除</span>
              </a-button>
            </a-space>
          </div>
          <div class="tablebox">
            <a-table  :data="nenudata" row-key="id" @row-click="handleClickMenu" :selected-keys="selectedKeys" :scroll="{ y: '100%'}" :pagination="false" :bordered="{wrapper: true, cell: true}" style="margin-top: 10px">
              <template #columns>
                    <a-table-column title="" :width="50" data-index="id"></a-table-column>
                    <a-table-column  title="字典名称" data-index="title">
                      <template #cell="{ record }">
                         <div class="titlebox">
                          <div class="text">{{ record.title }}</div>
                          <div class="icon"><icon-right /></div>
                         </div>
                      </template>
                    </a-table-column>
                  </template>
            </a-table>
          </div>
      </div>
      <div class="flex-page content">
        <div class="custom-full-table">
        <div class="table-toolbar flex flex-between">
          <div class="left" >
            <a-space>
              <a-input :style="{width:'160px'}"  v-model="formModel.title" placeholder="搜索字典名称" allow-clear />
              <a-range-picker v-model="formModel.createtime" :shortcuts="shortcuts" shortcuts-position="left" :style="{width:'250px'}" @change="fetchData"/>
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
        <div class="table-container">
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
            <template #title="{record,column}">
              <cell-tag-plus v-if="record[column.dataIndex]" :color="record.tagcolor" :value="record[column.dataIndex]"/>
            </template>
            <template #des="{record,column}">
              <a-typography-paragraph :ellipsis="{rows: 2,expandable: true}">{{ record[column.dataIndex] }}</a-typography-paragraph>
            </template>
            <template #createtime="{ record }">
              {{dayjs(record.createtime).format("YYYY-MM-DD HH:mm")}}
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
              <template v-if="record.data_from=='business'">
                <Icon icon="svgfont-bianji1" class="iconbtn" @click="handleEdit(record)" :size="18" color="rgb(var(--arcoblue-6))"></Icon>
                <a-divider direction="vertical" />
                <a-popconfirm content="您确定要删除吗?" @ok="handleDel(record)" position="tr">
                  <Icon icon="svgfont-icon7" class="iconbtn" :size="18" color="#ed6f6f"></Icon>
                </a-popconfirm>
              </template>
              <span v-else style="font-size: 12px" class="success">公共数据</span>
            </template>
          </a-table>
        </div>
      </div>
    </div>
  </div>
  </page-card>
    <!--表单-->
    <AddForm @register="registerModal" @success="handleData"/>
    <AddMenu @register="registerAddMenuModal" @success="handleAddMenu"/>
  </div>
</template>

<script lang="ts" setup>
  import { ref, reactive, watch, onMounted,nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  //api
  import { getList,upStatus,del} from '@/api/datacenter/dictionary';
  import { getmenuList,delMenuList, menuItem} from '@/api/datacenter/tabledata';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import { shortcuts, dayjs} from '@/utils/dayjs';
  //数据
  import { columns} from './data';
  //表单
  import { useModal } from '/@/components/Modal';
  import AddForm from './AddForm.vue';
  import AddMenu from './AddMenu.vue';
  import {Icon} from '@/components/Icon';
  import { Message ,Modal} from '@arco-design/web-vue';
  import { Pagination } from '@/types/global';
  import { useUserStore } from '@/store';
  import {tabletool,SizeProps,statusOptions} from '/@/components/tabletool';
  const userInfo = useUserStore();
  const [registerModal, { openModal }] = useModal();
  const [registerAddMenuModal, { openModal:openAddMenuModal }] = useModal();
 
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
  // const boxheight=document.documentElement.clientHeight;//页面高度
  type Column = TableColumnData & { checked?: true };
  const { loading, setLoading } = useLoading(true);
  const renderData = ref([]);
  const nenudata = ref<menuItem []>([]);
  const cloneColumns = ref<Column[]>([]);
  const isFullscreen = ref(false);
  const size = ref<SizeProps>('large');
   //查询字段
   const generateFormModel = () => {
    return {
      title: '',
      name: '',
      createtime: [],
      status: '',
    };
  };
  const handleTablename=ref("")//数据表
  const handleGroupId=ref(0)//分类id
  const formModel = ref(generateFormModel());
  const fetchData = async () => {
    setLoading(true);
    try {
      const data= await getList({page:pagination.current,pageSize:pagination.pageSize,...formModel.value,tablename:handleTablename.value,group_id:handleGroupId.value});
      renderData.value = data.items;
      pagination.current = data.page;
      pagination.total = data.total;
    } catch (err) {
      setLoading(false);
    } finally {
      setLoading(false);
    }
  };
  //获取左边数据
  const Menuitem=ref<menuItem>()
  const getmenudata=async()=>{
    nenudata.value = await getmenuList({})
    if(nenudata.value&&nenudata.value.length>0){
        selectedKeys.value=[nenudata.value[0].id]
        Menuitem.value=nenudata.value[0]
        handleTablename.value=nenudata.value[0].tablename
        handleGroupId.value=nenudata.value[0].id
         nextTick(()=>{
          fetchData();
         })
      }else{
        setLoading(false);
      }
  }
  //更新左边数据
  const handleAddMenu=()=>{
    getmenudata()
  }
  //选择菜单数据
  const selectedKeys=ref<number []>([])
  const handleClickMenu=(TableData:any)=>{
    renderData.value =[]
    pagination.current =1
    setLoading(true);
    selectedKeys.value=[TableData.id]
    handleTablename.value=TableData.tablename
    handleGroupId.value=TableData.id
    Menuitem.value=TableData
    fetchData();
  }
  //添加菜单数据
  const AddMenuData=(type:number)=>{
    if(type==2&&!Menuitem.value){
      Message.error("未选择编辑数据")
    }else if(type==2&&(Menuitem.value?.data_from=="common"||Menuitem.value?.business_id!=userInfo.$state.business_id)){
      Message.error("您没有编辑权限，只能编辑自己添加数据")
    }else{
      openAddMenuModal(true, {
        isUpdate: type==2?true:false,
        record:type==2?Menuitem:null
      });
    }
  }
  //删除菜单数据
  const delMenuData=()=>{
    if(!Menuitem.value){
      Message.error("未选择删除数据")
    }else if(Menuitem.value?.data_from=="common"||Menuitem.value?.business_id!=userInfo.$state.business_id){
      Message.error("该分类你没有删除权限")
    }else{
       Modal.warning({
        title: '您确定要删除内容吗？',
        content: '删除后内容将无法恢复请谨慎操作！',
        cancelText:"取消",
        okText:"删除",
        titleAlign:"start",
        hideCancel:false,
        onOk:(async()=>{
          Message.loading({content:"删除中",id:"upStatus",duration:0})
          await delMenuList({ids:selectedKeys.value})
          nextTick(()=>{
            Message.success({content:"删除成功",id:"upStatus",duration:200})
            getmenudata();
          })
          })
      });
    }
  }
  //组件挂载完成后执行的函数
  onMounted(()=>{
    getmenudata()
  })
  const reset = () => {
    formModel.value = generateFormModel();
    fetchData();
  };

  watch(
    () => columns.value,
    (val) => {
      cloneColumns.value = cloneDeep(val);
    },
    { deep: true, immediate: true }
  );
  //添加
  const createData=()=>{
    if(!selectedKeys.value||selectedKeys.value.length==0){
      Message.error("未选择字典数据")
    }else{
      openModal(true, {
        isUpdate: false,
        tablename:handleTablename.value,
        group_id:selectedKeys.value[0],
        record:null
      });
    }
  }
  //编辑数据
  const handleEdit=async(record:any)=>{
    openModal(true, {
      isUpdate: true,
      tablename:handleTablename.value,
      group_id:selectedKeys.value[0],
      record:record
    });
  }
  //更新数据
  const handleData=async()=>{
    fetchData();
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
      await upStatus({id:record.id,status:record.status, tablename:handleTablename.value});
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
      await del({ids:[record.id], tablename:handleTablename.value});
      fetchData();
      Message.success({content:"删除成功",id:"del",duration:2000})
    }catch (error) {
      Message.loading({content:"删除中",id:"del",duration:1})
    } 
}
</script>

<script lang="ts">
  export default {
    name: 'dictionary',
  };
</script>

  
<style lang="less" scoped>
  @import '@/assets/style/fulltable.less';
  .flexbox{
      display: flex;
      justify-content: space-between; // 替代原先的space-between布局方式
      height: 100%;
      .menu{
        height: 100%;
        width: 230px;
        margin-right: 10px;
        background-color: var(--color-neutral-1);
        padding: 10px 0px;
        .btn{
          text-align: center;
        }
        .tablebox{
          height: calc(100% - 36px);
          //分类标题背景色
          :deep(.arco-table-th){
            background-color: var(--color-neutral-3);
          }
          :deep(.arco-table-tr-checked .arco-table-td){
            color: rgb(var(--primary-6));
          }
          .titlebox{
            display: flex;
            .text{
              flex: 1;
            }
          }
        }
    }
    .content{
      flex: 1;
      height: 100%;
      min-width: 400px;
      padding: 10px 5px;
      background-color: var(--color-neutral-1);
    }
  }
  :deep(.arco-btn-size-medium){
    padding: 0 10px;
  }
  //搜索
  :deep(.arco-input-wrapper),:deep(.arco-picker),:deep(.arco-select-view-single){
     background-color: var(--color-bg-4);
  }
</style>
