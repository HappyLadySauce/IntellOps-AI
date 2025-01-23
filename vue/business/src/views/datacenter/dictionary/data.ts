import { computed } from 'vue';
import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  export const columns = computed<TableColumnData[]>(() => [
    {
      title: '排序',
      dataIndex: 'weigh',
      width:80,
      align:"center"
    },
    {
      title: 'ID',
      dataIndex: 'id',
      width:80,
      align:"center"
    },
    {
      title: '字典名称',
      dataIndex: 'keyname',
      slotName: 'title',
      width:150,
      align:"center"
    },
    {
      title: '字典项值',
      dataIndex: 'keyvalue',
      width:150,
      align:"center"
    },
    {
      title: '字典描述',
      dataIndex: 'des',
      slotName: 'des',
      width:180,
      align:"center"
    },
    {
      title: '状态',
      dataIndex: 'status',
      slotName: 'status',
      width:100,
      align:"center"
    },
    {
      title: '创建时间',
      dataIndex: 'createtime',
      slotName: 'createtime',
      width: 160,
      align:"center"
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
      fixed: 'right',
      width: 100,
      align:"center"
    },
  ]);
