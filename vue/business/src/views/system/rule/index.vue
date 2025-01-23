<template>
  <div class="container">
    <page-card breadcrumb scrollPage :isFullscreen="isFullscreen">
      <a-row style="margin-bottom: 10px">
        <a-col :span="16">
           <div>
            此项菜单适用于开发时，方便业务端添加新业务项，开发完成即可在admin端关闭。也可以直接在admin端的业务管理添加。拖拽可修改同级顺序，跨不同层级请进入编辑页面修改。
           </div>
        </a-col>
        <a-col
          :span="8"
           style="text-align: right;"
        >
          <tabletool :showbtn="['create','refresh','selectdensity','fullscreen']"
           @create="createData" @refresh="fetchData" @selectdensity="(data)=>size=data" @fullscreen="(data)=>isFullscreen=data">
            <a-button type="outline" @click="handleshowAll">
              {{ isShowAll?"仅看菜单":"展开权限"  }}
            </a-button>
            <a-button class="gf_hover_btn-border" @click="handeHindAll" >
              <template #icon><icon-caret-down v-if="dragPidHind"/><icon-caret-up v-else/></template>
            </a-button>
          </tabletool>
        </a-col>
      </a-row>
      <a-table
         row-key="id"
        :loading="loading"
        :pagination="false"
        :columns="(cloneColumns as TableColumnData[])"
        :data="renderData"
        :bordered="{wrapper:true,cell:true}"
        :size="size"
        :default-expand-all-rows="true"
        :scroll="{x:'100%'}"
        ref="artable"
        :draggable="{ type: 'handle', width: 38 }"
        @change="handleChange" 
      >
        <template #title="{ record }">
          <span v-html="record.spacer" style="padding-right: 5px;color: var(--color-neutral-4);"></span>{{ record.title }}
        </template>
        <template #icon="{ record }">
          <Icon :icon="record.icon" :size="20"></Icon>
        </template>
        <template #component="{ record }">
           <span v-if="record.type==0" style="color: rgb(var(--arcoblue-7));">{{ record.component }}</span>
           <a-tooltip v-else-if="record.type==1" content="前端组件路径">
             <span  style="color: var(--color-neutral-10);">{{ record.component }}</span>
           </a-tooltip>
           <a-tooltip  v-else content="后端接口路径">
             <span style="color:rgb(var(--orange-6));">{{ record.path }}</span>
          </a-tooltip>
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
          <a-space>
            <Icon icon="icon-plus-circle-fill" class="iconbtn" :size="18" color="rgb(var(--arcoblue-6))" @click="handleAddSubmenu(record)" v-if="record.type!=2"></Icon>
            <Icon icon="svgfont-bianji1" class="iconbtn" @click="handleEdit(record)" :size="17" color="rgb(var(--arcoblue-6))"></Icon>
            <a-popconfirm content="您确定要删除吗?" @ok="handleDel(record)" position="tr">
              <Icon icon="svgfont-icon7" class="iconbtn" :size="17" color="rgb(var(--red-6))"></Icon>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </page-card>
    <!--表单-->
    <AddForm @register="registerModal"  @success="handleData"/>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted, watch, nextTick } from 'vue';
  import useLoading from '@/hooks/loading';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import dayjs from 'dayjs';
  //数据
  import { columns} from './data';
  import { getList,upStatus,del} from '@/api/system/rule';
  import { tableWeigh } from '@/api/common';
  //表单
  import { useModal } from '/@/components/Modal';
  import AddForm from './AddForm.vue';
  import {Icon} from '@/components/Icon';
  import { Message } from '@arco-design/web-vue';
  //排序
  import {tabletool,SizeProps} from '/@/components/tabletool';
  const [registerModal, { openModal }] = useModal();
  type Column = TableColumnData & { checked?: true };
  const { loading, setLoading } = useLoading(true);
  const renderData = ref([]);
  const cloneColumns = ref<Column[]>([]);
  const size = ref<SizeProps>('large');
  const isFullscreen = ref(false);
  const artable=ref()
  const pids=ref<number []>([])
  const fetchData = async () => {
    setLoading(true);
    try {
      const data= await getList({});
      renderData.value = data;
      pids.value=[]
      getMenuId(data)//获取存在子菜单的id
      nextTick(()=>{
        artable.value.expand(pids.value)
      })
    } catch (err) {
      console.error("错误",err)
    } finally {
      setLoading(false);
    }
  };
  //获取菜单id
  const getMenuId=(data:any)=>{
      for (let item of data) {
        if (item.type != 2&&item.children&&childrenIsNemu(item.children)) { //等于2的时候把第3层children-设为-undefined
          pids.value.push(item.id)
        } 
        if(item.children) {
          // children若不为空数组，则继续 递归调用 本方法
          getMenuId(item.children);
        }
      }
  }
  //判断子类是否存在菜单
  const childrenIsNemu=(data:any):Boolean=>{
    for (let item of data) {
      if(item.type!=2) {
        return true
      }
    }
    return false
  }
  //全部展开和父级展开
  const isShowAll=ref(false)
  const handleshowAll=()=>{
    isShowAll.value=! isShowAll.value
    if (isShowAll.value){
      artable.value.expandAll(true)
    }else{
      artable.value.expandAll(false)
      nextTick(()=>{
        artable.value.expand(pids.value)
      })
    }
  }

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
      await del({ids:[record.id]});
      fetchData();
      Message.success({content:"删除成功",id:"del",duration:2000})
    }catch (error) {
      Message.loading({content:"删除中",id:"del",duration:1})
    } 
  }
    //排序拖拽
    const handleChange = (_data:any,extra:any,currentData:any) => {
        if(extra.type=="drag"){
          renderData.value = _data
         const finddata= FindDragDataList(_data,extra.dragTarget.pid)
         var weighList :any[]=[]
         finddata.forEach((item:any,index:number) => {
          weighList.push({"id":item.id,"weigh":index+1})
         });
          orrdersave(weighList,extra.dragTarget.id,extra.dragTarget.pid)
        }
      }
    //获取拖拽数据所在的同级数据
    const FindDragDataList=(renderData:any,pid:any)=> {
        // 每次进来使用find遍历一次
        let res = renderData.filter((item:any) => item.pid == pid);
        if (res.length>0) {
          return res;
        } else {
          for (let i = 0; i < renderData.length; i++) {
            if (renderData[i].children instanceof Array && renderData[i].children.length > 0) {
              res = FindDragDataList(renderData[i].children, pid);
              if (res.length>0) return res;
            }
          }
          
          return [];
        }
      };
    //提交排序
    const orrdersave=async(weighList:any,id:any,pid:any)=>{
      if(!id){
         return
      }
      try {
          Message.loading({content:"更新排序中",id:"tableWeigh",duration:0})
          await tableWeigh({weighList:weighList,id:id,pid:pid,tableanme:"business_auth_rule"});
          Message.success({content:"更新排序成功",id:"tableWeigh",duration:2000})
      }catch (error) {
         Message.loading({content:"更新排序中",id:"tableWeigh",duration:2})
      }
    }
  //初始化排序
  onMounted(() => {
    fetchData();
  })
  //添加子菜单
  const handleAddSubmenu=(record:any)=>{
    openModal(true, {
      isUpdate: false,
      record:record
    });
  }
  //单元格 hover 进入时触发
  const dragPidHind=ref(false)
  const handeHindAll=()=>{
    dragPidHind.value=! dragPidHind.value
    if(dragPidHind.value){
      artable.value.expandAll(false)
    }else{
      nextTick(()=>{
        artable.value.expand(pids.value)
      })
    }
  }
</script>

<script lang="ts">
  export default {
    name: 'rule',//页签是名字和路由name相同则缓存生效
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
