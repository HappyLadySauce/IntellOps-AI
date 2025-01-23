import { computed } from 'vue';
import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  export const columns = computed<TableColumnData[]>(() => [
    {
      title: '菜单/权限名称',
      dataIndex: 'title',
      slotName: 'title',
      width: 220,
    },
    {
      title: '图标',
      dataIndex: 'icon',
      slotName: 'icon',
      width: 70,
      align:"center"
    },
    {
      title: '前端权限标识',
      dataIndex: 'permission',
      slotName: 'permission',
      width: 120,
    },
    {
      title: '前端组件/权限路由',
      dataIndex: 'component',
      slotName: 'component',
      width: 390,
    },
    {
      title: '排序',
      dataIndex: 'weigh',
      slotName: 'weigh',
      width: 100,
      align:"center"
    },
    {
      title: '状态',
      dataIndex: 'status',
      slotName: 'status',
      width: 90,
      align:"center"
    },
    {
      title: '创建时间',
      dataIndex: 'createtime',
      slotName: 'createtime',
      width: 120,
      align:"center"
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
      fixed: 'right',
      width: 90,
      align:"center"
    },
  ]);