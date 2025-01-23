<template>
  <div class="container" >
    <page-card breadcrumb :isFullscreen="isFullscreen">
      <div class="flex-page flexcolumn">
        <div class="custom-full-table">
        <div class="table-toolbar flex flex-between">
          <div class="left">
            <a-space>
              <a-input :style="{width:'180px'}"  v-model="formModel.name" placeholder="搜索名称、备注" allow-clear />
              <a-range-picker v-model="formModel.createtime" :shortcuts="shortcuts" shortcuts-position="left" :style="{width:'230px'}" @change="search"/>
              <a-select v-model="formModel.status" :options="statusOptions" placeholder="状态" :style="{width:'78px'}" @change="search"/>
              <a-button type="primary" @click="search">
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
            <tabletool :showbtn="['refresh','selectdensity','fullscreen']"
              @refresh="fetchData" @selectdensity="(data)=>size=data" @fullscreen="(data)=>isFullscreen=data">
                  <a-button type="primary" @click="handleAddDb">
                    <template #icon>
                      <icon-plus />
                    </template>
                    添加数据表
                  </a-button>
                  <a-button  @click="handleCodeTool">
                    <template #icon>
                      <icon-tool />
                    </template>
                    生成代码工具
                  </a-button>
            </tabletool>
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
            <template #createtime="{ record ,column}">
              {{dayjs(record[column.dataIndex]).format("YYYY-MM-DD HH:mm")}}
            </template>
            <template #fromtype="{ record ,column}">
              <span style="font-size: 12px" :class="record[column.dataIndex]==1?'primary':'success'">{{record[column.dataIndex]==1?"代码模板生成":"数据表生成"}}</span>
            </template>
            <template #codelocation="{ record ,column}">
              <span style="font-size: 12px" :class="record[column.dataIndex]=='adminDirName'?'warning':'success'">{{record[column.dataIndex]=='adminDirName'?"admin 端":"business端"}}</span>
            </template>
            <template #overtime="{ record }">
              {{record.overtime==0?"不限":record.overtime}}
            </template>
            <template #api_path="{ record }">
              <template v-if="record.is_install==1">
                <div class="default fontsizemin" v-if="record.tpl_type!='web'"> {{ record.api_path +"/"+record.api_filename}}</div>
                <div class="default fontsizemin" v-if="record.component!='/'&&record.tpl_type!='go'"> {{  "views/"+record.component}}</div>
              </template>
              <template v-else>
                <span style="font-size: 12px" class="warning">待生成</span>
              </template>
            </template>
            <template #status="{ record }">
              <a-switch v-model="record.status" :checked-value="0" :unchecked-value="1" @change="handleStatus(record)">
                <template #checked>展示</template>
                <template #unchecked>隐藏</template>
              </a-switch>
            </template>
            <template #operations="{ record }">
              <button-group v-if="record.fromtype==0" size="small"  >
                <a-button style="color: rgb(var(--arcoblue-6));"  @click="openCodeMaker(record)">生成代码</a-button>
                <a-popconfirm :content="`您确定要${record.is_install==1?'卸载':'删除'}数据库和代码吗?`" @ok="handleDel(record)">
                  <a-button style="color: red;">{{record.is_install==1?"卸载":"删除"}}</a-button>
                </a-popconfirm>
              </button-group>
              <div v-else>
                <a-popconfirm :content="`您确定要删除代码吗?`" @ok="handleDelCodetpl(record)">
                  <a-button style="color: red;" size="small">删除生成的代码</a-button>
                </a-popconfirm>
              </div>
            </template>
          </a-table>
        </div>
        <div class="tigbox">
          <div class="title">生成代码说明：</div>
          <div class="item">
            <div class="label">1.新增数据表时，请点击右上角的“更新数据表”进行更新</div>
          </div>
          <div class="item">
            <div class="label">2.生成数据列表不带分类，则直接选择“模板类型=仅数据列表”</div>
            <div class="text"></div>
          </div>
          <div class="item">
            <div class="label">3.生成列表数据时，如果数据有对应分类（如：产品分类）时“模板类型=数据关联分类”，产品列表数据表命名为：business_product_content 分类数据表命名为：business_product__cate</div>
            <div class="text"></div>
          </div>
        </div>
      </div>
      </div>
    </page-card>
    <!--表单-->
    <AddDb @register="registerModal"  @success="handleData"/>
    <MakeCode @register="registerMakeCodeModal"  @success="handleData"/>
  </div>
</template>

<script lang="ts" setup>
  import { ref, reactive, watch, onMounted,onUnmounted} from 'vue';
  import useLoading from '@/hooks/loading';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  //数据
  import { columns} from './data';
  //api
  import { getList,upStatus,del,delTplCode} from '@/api/developer/generatecode';
  import { useModal} from '/@/components/Modal';
  import AddDb from './AddDb.vue';
  import MakeCode from './MakeCode.vue';
  import { Message ,Button} from '@arco-design/web-vue';
  import { Pagination } from '@/types/global';
  import { useRouter } from 'vue-router'
  import { shortcuts, dayjs} from '@/utils/dayjs';
  import {tabletool,SizeProps,statusOptions} from '/@/components/tabletool';
  const router = useRouter();
  const ButtonGroup=Button.Group
  const [registerModal, { openModal }] = useModal();
  const [registerMakeCodeModal, { openModal:markeCode }] = useModal();
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
  //查询 
  const search = () => {
    fetchData();
  };
  const reset = () => {
    formModel.value = generateFormModel();
    fetchData();
  };
  //组件挂载完成后执行的函数
  const timer=ref()
  onMounted(()=>{
    fetchData();
  })
   //离开
onUnmounted(() => {
  timer.value && clearTimeout(timer.value);
})
  watch(
    () => columns.value,
    (val) => {
      cloneColumns.value = cloneDeep(val);
    },
    { deep: true, immediate: true }
  );
  //添加数据表
  const handleAddDb=async()=>{
    openModal(true, {
      isUpdate: false,
    });
  }
  //代码生成工具
  const handleCodeTool=()=>{
    markeCode(true, {
      isUpdate: false,
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
      await upStatus({id:record.id,status:record.status});
      Message.success({content:"更新状态成功",id:"upStatus",duration:2000})
    }catch (error) {
      Message.loading({content:"更新状态中",id:"upStatus",duration:1})
    } 
  }
  //删除数据
  const handleDel=async(record:any)=>{
    try {
      Message.loading({content:"删除中",id:"del",duration:0})
      await del({id:record.id,is_install:record.is_install});
      fetchData();
      Message.success({content:"删除成功",id:"del",duration:2000})
    }catch (error) {
      Message.loading({content:"删除中",id:"del",duration:1})
    } 
  }
  //打开代码生成器
  const openCodeMaker=(record:any)=>{
    router.push({name:"codemaker",query:{id:record.id}})
  }
  //删除代码模板
  const handleDelCodetpl=async(record:any)=>{
    try {
      Message.loading({content:"删除中",id:"del",duration:0})
      await delTplCode({id:record.id});
      fetchData();
      Message.success({content:"删除成功",id:"del",duration:2000})
    }catch (error) {
      Message.loading({content:"删除中",id:"del",duration:2})
    } 
  }
</script>

<script lang="ts">
  export default {
    name: 'generatecode',
  };
</script>

<style scoped lang="less">
   @import '@/assets/style/fulltable.less';
  //添加底部说明
  .tablebox{
    flex: 1;
  }
    //说明
  .tigbox{
    .title{
      font-size: 16px;
      font-weight: 600;
    }
    .item{
      display: flex;
      margin-top: 10px;
      .label{
        color: var(--color-neutral-10);
      }
      .text{
        flex:1;
        color: 	var(--color-neutral-6);
      }
    }
  }
</style>
