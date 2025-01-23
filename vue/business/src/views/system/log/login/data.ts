import { computed } from 'vue';
import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  export const columns = computed<TableColumnData[]>(() => [
    {
      title:  "ID",
      dataIndex: "id",
      align:"left",
      fixed: 'left',
      ellipsis: true, 
      tooltip: true, 
      width: 80,
    },
    {
      title: '登录时间',
      dataIndex: 'createtime',
      slotName: 'createtime',
      ellipsis: true, 
      tooltip: true, 
      width: 160,
      align:"left"
    },
    {
      title: '用户昵称',
      dataIndex: 'user',
      slotName: 'user',
      width: 160,
      ellipsis: true, 
      tooltip: true, 
      align:"left"
    },
    {
      title: '登录行为',
      dataIndex: 'des',
      width: 160,
      ellipsis: true, 
      tooltip: true, 
      align:"left"
    },
    {
      title: '状态',
      dataIndex: 'status',
      slotName: 'status',
      width: 100,
      align:"center"
    },
    {
      title: '登录 IP',
      dataIndex: 'ip',
      width: 100,
      ellipsis: true, 
      tooltip: true, 
      align:"left"
    },
    {
      title: '登录地点',
      dataIndex: 'address',
      render: ({ record,column }) =>{if(column.dataIndex){return record[column.dataIndex].replace("|0|","|")}},
      width: 150,
      ellipsis: true, 
      tooltip: true, 
      align:"left"
    },
    {
      title: '客户端信息',
      dataIndex: 'user_agent',
      slotName: 'user_agent',
      width: 100,
      align:"center"
    },
  ]);
  //状态
  export const statusOptions :any[]= [
    {
        label: "全部",
        value: "",
    },
    {
        label: "成功",
        value: 0,
    },
    {
        label: "失败",
        value: 1,
    },
    ];