import { computed } from 'vue';
import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  export const columns = computed<TableColumnData[]>(() => [
    {
      title: 'ID',
      dataIndex: 'id',
      width: 80,
      align:'center'
    },
    {
      title: '用户名',
      dataIndex: 'name',
      slotName : 'name',
      width: 180,
    },
    {
      title: '账号',
      dataIndex: 'username',
      width:100,
    },
    {
      title: '角色',
      dataIndex: 'rolename',
      width: 150,
    },
    {
      title: '部门',
      dataIndex: 'depname',
      width: 150,
    },
    {
      title: '邮箱',
      dataIndex: 'email',
      slotName: 'email',
      width: 180,
    },
    {
      title: '备注',
      dataIndex: 'remark',
      slotName: 'remark',
      width: 220,
      align:"center"
    },
    {
      title: '状态',
      dataIndex: 'status',
      slotName: 'status',
      width: 150,
      align:"center"
    },
    {
      title: '创建时间',
      dataIndex: 'createtime',
      width: 180,
      align:"center"
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
      fixed: 'right',
      width: 120,
      align:"center"
    },
  ]);