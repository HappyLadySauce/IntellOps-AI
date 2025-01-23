import { computed } from 'vue';
import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
export const columns = computed<TableColumnData[]>(() => [
     {
       title:  '代码名称',
       dataIndex: 'title',
       slotName: 'title',
       width: 260,
     },
     {
       title:  '说明',
       dataIndex: 'des',
       slotName:"des",
       width: 400,
     },
     {
      title:  '作者',
      dataIndex: 'author',
      slotName: 'author',
      width: 100,
    },
    {
      title:  '价格',
      dataIndex: 'price',
      slotName: 'price',
      width: 90,
     },
     {
      title:  '下载',
      dataIndex: 'download',
      width: 80,
     },
     {
      title:  '版本',
      dataIndex: 'version',
      slotName: 'version',
      width: 80,
     },
     {
       title:  '状态',
       dataIndex: 'status',
       slotName: 'status',
       width: 90,
     },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
      width: 160,
      fixed: 'right',
      align:'center',
    },
  ]);
