import { computed } from 'vue';
import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  export const columns = computed<TableColumnData[]>(() => [
    {
      title: '生成方式',
      dataIndex: 'fromtype',
      slotName:"fromtype",
      width: 120,
    },
    {
      title: '说明/备注',
      dataIndex: 'comment',
      width: 220,
      align:"center"
    },
    {
      title: '表名称',
      dataIndex: 'tablename',
      width: 200,
    },
    {
      title: '代码位置',
      dataIndex: 'codelocation',
      slotName:"codelocation",
      width: 120,
      align:"center"
    },
    {
      title: '代码路径',
      dataIndex: 'api_path',
      slotName:"api_path",
      width: 250,
      align:"center"
    },
    {
      title: '引擎',
      dataIndex: 'engine',
      width: 100,
      align:"center"
    },
    {
      title: '编码',
      dataIndex: 'collation',
      width: 200,
      align:"center"
    },
    // {
    //   title: '记录数',
    //   dataIndex: 'table_rows',
    //   width: 100,
    //   align:"center"
    // },
    // {
    //   title: '自增索引',
    //   dataIndex: 'auto_increment',
    //   width: 100,
    //   align:"center"
    // },
    {
      title: '状态',
      dataIndex: 'status',
      slotName: 'status',
      width: 100,
      align:"center"
    },
    {
      title: '更新时间',
      dataIndex: 'createtime',
      slotName: 'createtime',
      width: 160,
      align:"center"
    },
    // {
    //   title: '创建时间',
    //   dataIndex: 'createtime',
    //   align:"center"
    // },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
      fixed: 'right',
      width: 180,
      align:"center"
    },
  ]);
  //生成代码-表单字段
  export const columnsfiles = computed<TableColumnData[]>(() => [
    {
      title: '选择',
      dataIndex: 'isform',
      slotName: 'isform',
      width: 60,
    },
    {
      title: '字段名称',
      dataIndex: 'name',
      slotName: 'name',
      width: 140,
    },
    {
      title: '字段',
      dataIndex: 'field',
      width: 120,
    },
    {
      title: '必填',
      dataIndex: 'required',
      slotName: 'required',
      width: 60,
    },
    {
      title: '表单类型',
      dataIndex: 'formtype',
      slotName: 'formtype',
      width: 150,
    },
    {
      title: '关联数据',
      dataIndex: 'datatable',
      slotName: 'datatable',
      width: 220,
    },
    {
      title: '关联显示字段',
      dataIndex: 'datatablename',
      slotName: 'datatablename',
      width: 160,
    },
    {
      title: '表单布局',
      dataIndex: 'gridwidth',
      slotName: 'gridwidth',
      width: 110,
      align:'center'
    },
  ]);
  //列表
  export const columnslist = computed<TableColumnData[]>(() => [
    {
      title: '选择',
      dataIndex: 'islist',
      slotName: 'islist',
      width: 60,
    },
    {
      title: '字段名称',
      dataIndex: 'name',
      slotName: 'name',
    },
    {
      title: '字段',
      dataIndex: 'field',
    },
    {
      title: '筛选排序',
      dataIndex: 'isorder',
      slotName: 'isorder',
      align:"center",
      width: 88,
    },
    {
      title: '对齐方向',
      dataIndex: 'align',
      slotName: 'align',
      width: 138,
    },
    {
      title: '宽度',
      dataIndex: 'width',
      slotName: 'width',
      width: 120,
    },
    {
      title: '显示组件/默认空',
      dataIndex: 'show_ui',
      slotName: 'show_ui',
      width: 152,
    },
  ]);
  //搜索
  export const columnsseach = computed<TableColumnData[]>(() => [
    {
      title: '选择',
      dataIndex: 'issearch',
      slotName: 'issearch',
      width: 60,
    },
    {
      title: '名称',
      dataIndex: 'name',
      width: 136,
    },
    {
      title: '查询方式',
      dataIndex: 'searchway',
      slotName: 'searchway',
      width: 138,
    },
    {
      title: '查询类型',
      dataIndex: 'searchtype',
      slotName: 'searchtype',
      width: 150,
    },
    {
      title: '搜索框宽',
      dataIndex: 'searchwidth',
      slotName: 'searchwidth',
      width: 100,
    },
  ]);