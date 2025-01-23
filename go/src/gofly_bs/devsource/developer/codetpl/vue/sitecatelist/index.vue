<template>
  <div class="container" >
    <page-card breadcrumb :isFullscreen="isFullscreen">
      <a-split :style="{ width: '100%' ,height: '100%'}" v-model:size="splitsize">
        <template #resize-trigger>
          <div class="resizebox-line toggle-list">
            <div class="resizebox-line-icon"></div>
          </div>
        </template>
        <template #first>
          <div class="split-menu">
            <div class="list-header-wrapper">
              <div class="list-header">
                <span class="catalogue-title">分类管理</span>
                <div class="collapse-root">
                    <span class="header-operation " @click="()=>{onVisitemenu()}" :class="`collapse-toggle-${idupdown?'down':'up'}`"></span>
                </div>
                <a-tooltip content="新建分类">
                  <Icon class="addicon" @click="handleCreateCate(0)" icon="icon-plus-circle-fill" :size="29" color="#055fe6"></Icon>
                </a-tooltip>
              </div>
            </div>
              <div class="tablebox">
                <div class="scrollbarcontainer" >
                <!--菜单-->
                <a-tree
                    :blockNode="true"
                    :data="treeData"
                    :show-line="true"
                    ref="treeRef"
                    :selected-keys="treeId"
                    @select="handleSelectTree"
                  >
                  <template #switcher-icon="node,{ isLeaf }">
                      <IconDown v-if="!isLeaf" />
                      <template v-else>
                        <Icon class="yuandian" v-if="node.key!=0" icon="svgfont-yuandian2" :size="15"></Icon>
                        <Icon class="yuandian" v-else icon="svgfont-home" :size="15" ></Icon>
                      </template>
                  </template>
                  <template #title="record">
                    <a-tooltip v-if="record.name.length>=8" :content="record.name" background-color="#4e5969">
                      <span  class="titleBox">{{ record.name }}</span>
                    </a-tooltip>
                    <span v-else class="titleBox">{{ record.name }}</span>
                  </template>
                  <template #extra="nodeData">
                    <span v-if="nodeData.key!=0" class="page-tree-tool" :class="{'page-tree-tool-active':nodeData.showbtn}">
                        <a-popover position="br" trigger="click" :popup-visible="nodeData.showbtn"  @popup-visible-change="(visible)=>nodeData.showbtn=visible">
                          <Icon class="unicon page-tree-menu-button" icon="svgfont-gengduo" :size="13"></Icon>
                          <template #content>
                            <div class="tigmenu">
                              <div class="item" @click="handleEditCate(nodeData)"><Icon class="bicon" icon="svgfont-bianji1" :size="13"></Icon> 编辑</div>
                              <div class="item delete" @click="handleDelCate(nodeData.id)"><Icon class="bicon" icon="svgfont-icon7" :size="13"></Icon> 删除</div>
                            </div>
                          </template>
                        </a-popover>
                        <span class="optional-operation-wrapper">
                          <a-tooltip content="创建子数据">
                            <Icon class="unicon" @click="handleCreateCate(nodeData.id)" icon="svgfont-jia" :size="12"></Icon>
                          </a-tooltip>
                        </span>
                      </span>
                  </template>
                  </a-tree>
                </div>
              </div>
            </div>
        </template>
        <template #second >
          <div class="flex-page contentbox">
            <div class="custom-full-table">
              <div class="search-row" :class="{unhasesearchbtn:!showsearchdownbtn,hasesearchbtn:showsearchdownbtn}">
                <div class="search-field">
                  <div class="search-line search-drown" :class="{searchdowncss:searchdown}">
                    <div class="search-drown-box" ref="searchRef">
                      <a-space wrap>
<!--SearchHtml-->
                      </a-space>
                    </div>
                  </div>
                </div>
                <div class="search-btn">
                  <a-space>
                    <a-tooltip mini :content="searchdown?'收起查询条件':'展开查询条件'">
                      <a-button v-if="showsearchdownbtn" @click="searchdown=!searchdown">
                        <template #icon>
                          <icon-up v-if="searchdown"/>
                          <icon-down v-else/>
                        </template>
                      </a-button>
                    </a-tooltip>
                    <a-button type="primary" @click="handleSearch">
                        <template #icon>
                          <icon-search />
                        </template>
                        查询
                      </a-button>
                      <a-button @click="handleReset">
                        {{ $t('searchTable.form.reset') }}
                      </a-button>
                    </a-space>
                </div>
                <div class="search-option">
                  <tabletool :showbtn="['create','refresh','selectdensity','settingcolumn','fullscreen']" :columns="columns" :disabledColumnKeys="['id']"
                  @create="createData" @refresh="fetchData" @selectdensity="(data)=>size=data" @fullscreen="(data)=>isFullscreen=data" @settingcolumn="(data)=>cloneColumns=data"></tabletool>
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
                  @page-change="handlePageChange" 
                  @page-size-change="handlePageSizeChange" 
                >
<!--tableUIHtml-->
                  <template #operations="{ record }">
                    <Icon icon="svgfont-bianji1" class="iconbtn" @click="handleEdit(record)" :size="18" color="#0960bd"></Icon>
                    <a-divider direction="vertical" />
                    <a-popconfirm content="您确定要删除吗?" @ok="handleDel(record)">
                      <Icon icon="svgfont-icon7" class="iconbtn" :size="18" color="#ed6f6f"></Icon>
                    </a-popconfirm>
                  </template>
                </a-table>
              </div>
            </div>
          </div>
        </template>
      </a-split>
    </page-card>
    <!--表单-->
    <AddForm @register="registerModal" @success="handleData"/>
    <CateForm @register="registerCateModal" @success="getCatedata(1)"/>
  </div>
</template>

<script lang="ts" setup>
  import { ref, reactive, watch, onMounted,nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  //APi
  import { getList,upStatus,del} from './api';
  import { getTree,delCate} from './api/cate';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import { shortcuts, dayjs} from '@/utils/dayjs';
  //数据
  import { columns} from './data';
  //表单
  import { useModal } from '/@/components/Modal';
  import AddForm from './AddForm.vue';
  import CateForm from './CateForm.vue';
  import {Icon} from '@/components/Icon';
  import { Message,TreeNodeData,TreeInstance,Modal} from '@arco-design/web-vue';
  import { Pagination } from '@/types/global';
  import {tabletool,SizeProps} from '/@/components/tabletool';
  //分菜单
  const splitsize=ref("205px")
  const [registerModal, { openModal }] = useModal();
  const [registerCateModal, { openModal: openCateModal }] = useModal();
  const treeRef = ref<TreeInstance | null>(null);
  //选择菜单数据
  const treeData= ref<TreeNodeData []>([]);
  const treeId=ref<number []>([0])

  //选择文档-获取数据
  const clicdata=ref(0)
  const handleSelectTree=async(selectedKeys :any)=>{
      treeId.value=selectedKeys
      formModel.value.cid= selectedKeys[0];
      nextTick(()=>{
          fetchData();
      })
      //展开操作
      if(clicdata.value!=selectedKeys[0]){
        treeRef.value?.expandNode(selectedKeys,true)
        clicdata.value=selectedKeys[0]
      }else{
        treeRef.value?.expandNode(selectedKeys,false)
        clicdata.value=0
      }
  }
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
  const isFullscreen = ref(false);
  const size = ref<SizeProps>('large');
   //查询字段
   const generateFormModel = () => {
    return {
//{SearchField}
    };
  };
  const formModel = ref(generateFormModel());
  const fetchData = async () => {
    setLoading(true);
    try {
      const data= await getList({ page:pagination.current,pageSize:pagination.pageSize,...formModel.value});
      renderData.value = data.items;
      pagination.current = data.page;
      pagination.total = data.total;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
};
  //获取左边数据
  const getCatedata=async(type:number)=>{
    const catedata= await getTree({})
    const parntList_df : any=[{key: 0,name: "全部数据"}];
      if(catedata){
        treeData.value=parntList_df.concat(catedata)
      }else{
        treeData.value=parntList_df
      }
    if(type==1)return
     setLoading(false);
      if(treeData.value&&treeData.value.length>0){
        nextTick(()=>{
          setLoading(true);
          fetchData();
        })
      }
      
  }
  //组件挂载完成后执行的函数
  const searchRef=ref(null)
  const searchdown = ref(false);
  const showsearchdownbtn=ref(false)
  onMounted(()=>{
    getCatedata(0);
    searchbar()
    window.onresize = () => {
      searchbar()
    }
  })
 //更新搜索框按钮状态
 const searchbar=()=>{
  if(searchRef.value){
    const searchboxheight=searchRef.value["offsetHeight"]
    if(searchboxheight>40){
      showsearchdownbtn.value=true
    }else{
      showsearchdownbtn.value=false
    }
  }
 }
 //查找
 const handleSearch = () => {
    pagination.current=1
    fetchData();
  };
  //重置
  const handleReset = () => {
    pagination.current=1
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
    pagination.current=1
    fetchData();
  }
  //分页
  const handlePageChange = (page:any) => {
    pagination.current=page
    fetchData();
  }
  //分页总数
  const handlePageSizeChange = (pageSize:any) => {
    pagination.pageSize=pageSize
    fetchData();
  }
  //更新状态
  const handleStatus=async(record:any,field:string)=>{
    try {
      Message.loading({content:"更新状态中",id:"up",duration:0})
      await upStatus({id:record.id,[field]:record[field]});
      Message.success({content:"更新状态成功",id:"up",duration:2000})
    }catch (error) {
      record.status=record.status==1?0:1
      Message.loading({content:"更新状态中",id:"up",duration:1})
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
//分类菜单
const idupdown=ref(true)
function onVisitemenu(){
    idupdown.value=!idupdown.value
    treeRef.value?.expandAll(idupdown.value)
}
//添加分类
const handleCreateCate=(pid:any)=>{
  openCateModal(true, {
    isUpdate: false,
    record:{pid:pid}
  });
}
//编辑分类数据
const handleEditCate=(record:any)=>{
  openCateModal(true, {
    isUpdate: true,
    record:record
  });
}
//删除分类数据
const handleDelCate=async(id:any)=>{
  Modal.warning({
        title: '您确定要删除内容吗？',
        content: '删除后内容将无法恢复请谨慎操作！',
        cancelText:"取消",
        okText:"删除",
        titleAlign:"start",
        hideCancel:false,
        onOk:(async()=>{
              try {
                Message.loading({content:"删除中",id:"del",duration:0})
                const res= await delCate({ids:[id]});
                if(res){
                  getCatedata(0);
                  if( formModel.value.cid==id){
                    formModel.value.cid=0
                    treeId.value=[0]
                  }
                  Message.success({content:"删除成功",id:"del",duration:2000})
                }
              }catch (error) {
                Message.loading({content:"删除中",id:"del",duration:1})
              } 
          })
      });
}
</script>

<script lang="ts">
  export default {
    name: '{vuetplname}', // If you want the include property of keep-alive to take effect, you must name the component
  };
</script>

<style scoped lang="less">
  @import '@/assets/style/catelist.less';
  @import '@/assets/style/fulltable.less';
</style>
