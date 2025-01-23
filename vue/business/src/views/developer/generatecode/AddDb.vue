<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" helpMessage="选择数据库已经建好的数据表结构生成代码，框架不提供数据表创建工具，请使用Navicat数据库管理工具创建，这工具很好用，比自己后台搞一套高效多了。" :loading="loading" width="1000px" :minHeight="modelHeight" :title="getTitle" @ok="handleSubmit">
    <div class="addFormbox" :style="{'min-height':`${modelHeight}px`}">
      <div class="table-content">
        <a-table
         row-key="name"
        :loading="loading"
        :pagination="false"
        :columns="(cloneColumns as TableColumnData[])"
        :data="dblist"
        :row-selection="{type: 'checkbox',showCheckedAll: true,onlyCurrent: false}"
        v-model:selectedKeys="selectIds"
        @selection-change="handleSelect"
        :scroll=" {y: 557}"
        :bordered="{wrapper:true,cell:true}"
        :default-expand-all-rows="true"
        ref="artable"
      >
      <template #name-filter="{ filterValue, setFilterValue, handleFilterConfirm, handleFilterReset}">
        <div class="custom-filter">
          <a-space direction="vertical">
            <a-input :model-value="filterValue[0]" @input="(value)=>setFilterValue([value])" @press-enter="handleFilterConfirm" placeholder="搜索表名称" allow-clear/>
             <div class="custom-filter-footer">
              <a-button @click="handleFilterConfirm" type="primary">查找</a-button>
              <a-button @click="handleFilterReset">重置</a-button>
            </div>
          </a-space>
        </div>
      </template>
      </a-table>
      </div>
    </div>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref,  h} from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import useLoading from '@/hooks/loading';
  //api
  import { upCodeTable,checkedHaseTable} from '@/api/developer/generatecode';
  import { getTables } from '@/api/common';
  import { Message,Modal } from '@arco-design/web-vue';
  import { IconPicker ,Icon} from '@/components/Icon';
  import { IconSearch } from '@arco-design/web-vue/es/icon';
  export default defineComponent({
    name: 'AddDb',
    components: { BasicModal,IconPicker,Icon },
    emits: ['success'],
    setup(_, { emit }) {
      const modelHeight= ref(600);
      const dblist= ref([]);
      //编辑器
      const [registerModal, {  closeModal }] = useModalInner(async (data) => {
        setLoading(true);
        dblist.value= await getTables({});
        setLoading(false);
      });
      //选择表
      const selectIds=ref([])
      const tablenames=ref<string []>([])
      const handleSelect=(ids:any)=>{
        tablenames.value=ids
      }
     //点击添加
     const { loading, setLoading } = useLoading();
     const handleSubmit = async () => {
      if(tablenames.value.length==0){
        Message.warning({content:"请选择生成代码的数据表",id:"checked",duration:2000})
        return
      }
      try {
         Message.loading({content:"检查中...",id:"checked",duration:0})
         const data= await checkedHaseTable({tablenames:tablenames.value});
         Message.loading({content:"检查中...",id:"checked",duration:1})
         if(data.length>0){
          Modal.warning({
              title: '温馨提示',
              content: `您选择的“${data.join(",")}”数据表已经在生成列表中，您需要重新添加吗？`,
              okText:"确定",
              hideCancel:false,
              titleAlign:"start",
              cancelText:"取消",
              onOk:(async()=>{
                addTables()
              })
            });
         }else{
          addTables()
         }
        } catch (error) {
          Message.loading({content:"检查中...",id:"checked",duration:1})
        }
      };
      //添加执行添加表
      const addTables=async()=>{
        try {
            selectIds.value=[]
            Message.loading({content:"添加中",id:"upStatus",duration:0})
            await upCodeTable({tablenames:tablenames.value});
            Message.success({content:"添加成功",id:"upStatus",duration:2000})
            closeModal()
            emit('success');
        } catch (error) {
          Message.loading({content:"添加中",id:"upStatus",duration:1})
        }
      }
      return { 
        registerModal, 
        getTitle:"选择数据表添加到生成代码列表", 
        handleSubmit,
        loading,
        modelHeight,
        dblist,
        handleSelect,
        selectIds,
        cloneColumns:[
          {
            title: '表名称',
            dataIndex: 'name',
            filterable: {
              filter: (value:any, record:any) => record.name.includes(value),
              slotName: 'name-filter',
              icon: () => h(IconSearch)
            }
          },
          {
            title: '表备注',
            dataIndex: 'title',
            align:"center"
          },
        ]
      };
    },
  });
</script>
<style lang="less" scoped>
  .content{
      width: 100%;
  }
  .custom-filter {
    padding: 20px;
    background: var(--color-bg-5);
    border: 1px solid var(--color-neutral-3);
    border-radius: var(--border-radius-medium);
    box-shadow: 0 2px 5px rgb(0 0 0 / 10%);
  }
  .custom-filter-footer {
    display: flex;
    justify-content: space-between;
  }
</style>

