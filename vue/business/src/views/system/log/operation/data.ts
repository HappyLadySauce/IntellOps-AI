import { computed } from 'vue';
import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
export const columns = computed<TableColumnData[]>(() => [
  {
    title:  "ID",
    dataIndex: "id",
    ellipsis: true, 
    tooltip: true, 
    width: 80,
    align:"left",
  },
  {
    title: '操作时间',
    dataIndex: 'createtime',
    slotName: 'createtime',
    ellipsis: true, 
    tooltip: true, 
    width: 160,
    align:"center"
  },
  {
    title: '操作人',
    dataIndex: 'user',
    slotName: 'user',
    width: 150,
    ellipsis: true, 
    tooltip: true, 
    align:"left"
  },
  {
    title: '操作内容',
    dataIndex: 'des',
    ellipsis: true, 
    tooltip: true, 
    width: 160,
    align:"left"
  },
  {
    title: '请求方式',
    dataIndex: 'request_method',
    render: ({ record }) =>methodName(record),
    width: 90,
    align:"left"
  },
  {
    title: '请求URL',
    dataIndex: 'url',
    width: 250,
    ellipsis: true, 
    tooltip: true, 
    align:"left"
  },
  {
    title: '状态',
    dataIndex: 'status',
    slotName: 'status',
    width: 90,
    align:"center"
  },
  {
    title: '操作 IP',
    dataIndex: 'ip',
    ellipsis: true, 
    tooltip: true, 
    width: 110,
    align:"left"
  },
  {
    title: '登录地点',
    dataIndex: 'address',
    width: 160,
    render: ({ record,column }) =>{if(column.dataIndex){return record[column.dataIndex].replace("|0|","|")}},
    ellipsis: true, 
    tooltip: true, 
    align:"center"
  },
  {
    title: '耗时',
    dataIndex: 'latency',
    slotName: 'latency',
    width: 110,
    align:"center"
  },
  {
    title: '操作',
    dataIndex: 'operations',
    slotName: 'operations',
    fixed: 'right',
    width: 80,
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
      value: 200,
  },
  {
      label: "失败",
      value: 1,
  },
];
//请求方式
const methodName=(record:any):string=>{
   if(record.request_method=="GET"){
    return "读取数据"
   }else if(record.request_method=="POST"){
     return "提交数据"
   }else if(record.request_method=="PUT"){
    return "修改数据"
  }else if(record.request_method=="DELETE"){
    return "删除数据"
  }else{
     return "未知"
   }
}