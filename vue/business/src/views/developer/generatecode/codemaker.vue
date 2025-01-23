<template>
    <div class="page-container">
       <a-card class="general-card onelineCard" ref="onelineCardRef" style="height: 100%;">
           <div class="maker-rawer">
               <a-split :style="{height:(cardboxHeight)+'px',width: '100%',overflow:'hidden'}"
                     v-model:size="atrrSize"
                     :min="`${cardboxWidtht-600}px`"
                     :max="`${cardboxWidtht-350}px`"
                   >
                 <template #resize-trigger-icon>
                   <span class="drag-icon-center" @click.stop="handleChangeAtrr">
                     <icon-caret-left class="selfsplit-trigger-icon" v-if="openAtrr"/>
                     <icon-caret-right class="selfsplit-trigger-icon" v-else/>
                   </span>
                 </template>
                   <template #first>
                     <!--表单字段-->
                     <div class="maker-content flex flex-flow flex-between">
                       <div class="header header-table">
                           <div class="title">字段设置</div>
                       </div>
                       <div class="content flex flex-flow flex-between" >
                          <div class="table-box" >
                               <div class="title">1.表单字段(拖拽<icon-drag-arrow />对表单字段排序、勾选填写字段、修改字段名称、可勾选必填、选择字段输入框类型、关联数据表及中心数据、选择布局)</div>
                               <div class="table-main">
                                   <a-table :columns="columnsfiles" :data="field_list" @change="handleChangeDragField" :pagination="false" size="small"
                                   :draggable="{ type: 'handle', width: 40 }" :scroll="{ y: `${(cardboxHeight-50)/2-30}px` }" :scrollbar="true">
                                       <template #drag-handle-icon>
                                           <icon-drag-arrow />
                                       </template>
                                       <template #isform="{ rowIndex }">
                                           <a-checkbox v-model="field_list[rowIndex].isform" />
                                       </template>
                                       <template #required="{ rowIndex }">
                                           <a-checkbox v-model="field_list[rowIndex].required" />
                                       </template>
                                       <template #name="{ rowIndex }">
                                          <a-input v-model="field_list[rowIndex].name" placeholder="填写字段名称" @input="handleinputForm(field_list[rowIndex])"/>
                                       </template>
                                       <template #formtype="{ record }">
                                           <a-select v-model="record.formtype" placeholder="请选择" @change="handleUpshowui(record)">
                                               <a-option value="text">文本框</a-option>
                                               <a-option value="number">数字本文</a-option>
                                               <a-option value="textarea">文本域</a-option>
                                               <a-option value="editor">富文本</a-option>
                                               <a-option value="select">下拉框</a-option>
                                               <a-option value="switch">开关Switch</a-option>
                                               <a-option value="radio">单选框</a-option>
                                               <a-option value="checkbox">复选框</a-option>
                                               <a-option value="belongto">关联表字段</a-option>
                                               <a-option value="belongDic" >关联字典数据</a-option>
                                               <a-option value="date">日期控件</a-option>
                                               <a-option value="datetime">日期时间控件</a-option>
                                               <a-option value="time">时间选择器</a-option>
                                               <a-option value="colorpicker">颜色选择器</a-option>
                                               <a-option value="image">单图上传</a-option>
                                               <a-option value="images">多图上传</a-option>
                                               <a-option value="audio">音频上传</a-option>
                                               <a-option value="file">单文件上传</a-option>
                                               <a-option value="files" disabled>多文件上传</a-option>
                                           </a-select>
                                       </template>
                                       <template #datatable="{ record }">
                                           <a-select v-model="record.datatable" @change="getTableField" v-if="record.formtype=='belongto'" placeholder="请选择" allow-clear allow-search>
                                               <template v-for="item in tablelist">
                                                   <a-tooltip :content="item.title">
                                                       <a-option :value="item.name" :title="item.name">{{ item.name }}</a-option>
                                                   </a-tooltip>
                                               </template>
                                           </a-select>
                                           <a-select v-model="record.dic_group_id" v-if="record.formtype=='belongDic'" placeholder="请选择字典数据"  allow-clear allow-search>
                                             <a-option v-for="item in dictablelist" :value="item.id" >{{ item.title }}</a-option>
                                           </a-select>
                                       </template>
                                       <template #datatablename="{ record }">
                                           <a-select v-model="record.datatablename" v-if="record.formtype=='belongto'&&record.datatable!=''"
                                           @popup-visible-change="(val:any)=>{val?getTableField(record.datatable):null}" placeholder="请选择" allow-clear allow-search>
                                               <template v-for="item in fieldList[record.datatable]">
                                                 <a-option :value="item['value']" >{{ item["value"] }}</a-option>
                                               </template>
                                           </a-select>
                                       </template>
                                       <template #gridwidth="{ record }">
                                           <a-space>
                                               <div class="flex flex-flow">
                                                   <a-tooltip content="单列" position="top" mini>
                                                     <div class="gridbox" ><FormIconSelect icon="svgfont-bujurongqi-buju1" :val="24" v-model="record.gridwidth"/></div>
                                                   </a-tooltip>
                                                   <a-tooltip content="双列" position="top" mini>
                                                     <div class="gridbox"><FormIconSelect icon="svgfont-bujurongqi-buju2" :val="12" v-model="record.gridwidth"/></div> 
                                                   </a-tooltip>
                                               </div>
                                               <div class="flex flex-flow">
                                                   <a-tooltip content="三列" position="top" mini>
                                                    <div class="gridbox"><FormIconSelect icon="svgfont-bujurongqi-buju3" :val="8" v-model="record.gridwidth"/></div>
                                                   </a-tooltip>
                                                   <a-tooltip content="四列" position="top" mini>
                                                    <div class="gridbox"><FormIconSelect icon="svgfont-bujurongqi-buju4" :val="6" v-model="record.gridwidth"/></div>
                                                   </a-tooltip>
                                               </div>
                                           </a-space>
                                       </template>
                                   </a-table>
                               </div>
                          </div>
                          <div class="table-box">
                               <div class="title">2.列表显示字段(拖拽<icon-drag-arrow />对表单字段排序、勾选显示字段、修改字段名称,勾选排序字段、对齐方式、列宽度、定义显示组件)</div>
                               <div class="table-main">
                                   <a-table :columns="columnslist" :data="list" @change="handleChangeDragList" :pagination="false" size="small"
                                   :draggable="{ type: 'handle', width: 40 }" :scroll="{ y: `${(cardboxHeight-50)/2-30}px` }" :scrollbar="true">
                                       <template #drag-handle-icon>
                                           <icon-drag-arrow />
                                       </template>
                                       <template #name="{ rowIndex }">
                                          <a-input v-model="list[rowIndex].name" placeholder="填写字段名称" @input="handleinputForm(list[rowIndex])"/>
                                       </template>
                                       <template #islist="{ rowIndex }">
                                           <a-checkbox v-model="list[rowIndex].islist" />
                                       </template>
                                       <template #isorder="{ rowIndex }">
                                           <a-checkbox v-model="list[rowIndex].isorder" />
                                       </template>
                                       <template #align="{ record }">
                                           <a-select v-model="record.align" placeholder="选择对齐方式">
                                               <a-option value="left">左边对齐</a-option>
                                               <a-option value="center">居中对齐</a-option>
                                               <a-option value="right">右边对齐</a-option>
                                           </a-select>
                                       </template>
                                       <template #width="{ rowIndex }">
                                          <a-input-number v-model="list[rowIndex].width" placeholder="空或0表示不设置"/>
                                       </template>
                                       <template #show_ui="{ record }">
                                        <a-select v-model="record.show_ui" placeholder="选择组件插槽" allow-clear>
                                            <a-option value="des">多行文本</a-option>
                                            <a-option value="cellcopy">复制文本</a-option>
                                            <a-option value="image">单图片</a-option>
                                            <a-option value="images">多图/相册</a-option>
                                            <a-option value="tag">单标签</a-option>
                                            <a-option value="tags">多标签</a-option>
                                            <a-option value="cellstatus">状态icon</a-option>
                                            <a-option value="switchstatus">状态开关</a-option>
                                            <a-option value="dotstatus">状态(dot)</a-option>
                                            <a-option value="link">URL网址</a-option>
                                            <a-option value="gender">性别</a-option>
                                            <a-option value="avatar">头像名(无背景)</a-option>
                                            <a-option value="nodeavatar">头像名(背景)</a-option>
                                            <a-option value="date">日期</a-option>
                                            <a-option value="datetime">日期+时间</a-option>
                                            <a-option value="color">颜色</a-option>
                                            <a-option value="dic">字典数据展示</a-option>
                                        </a-select>
                                       </template>
                                   </a-table>
                               </div>
                          </div>
                       </div>
                     </div>
                   </template>
                   <template #second >
                       <!--属性-->
                       <div class="maker-content">
                       <div class="header">
                           <a-tabs class="htabs" v-model:active-key="TabsIndex"  :hide-content="false">
                               <a-tab-pane :key="1" title="基本设置" />
                               <a-tab-pane :key="2" title="搜索字段" />
                               <a-tab-pane :key="3" title="模板类型" />
                           </a-tabs>
                       </div>
                       <div class="content" >
                           <a-scrollbar style="overflow: auto;" :style="{height:`${cardboxHeight-40}px`}">
                           <div class="formbox" v-if="codedata">
                               <a-form ref="formRef" :model="codedata" auto-label-width  >
                                   <div class="form-rawer" v-show="TabsIndex==1">
                                       <a-form-item field="comment" label="表名:" validate-trigger="input" style="margin-bottom:10px;">
                                           <span class="textbox">{{ codedata.tablename }}</span>
                                       </a-form-item>
                                       <a-form-item label="代码生成位置" field="codelocation" style="margin-bottom:15px;" :rules="[{required:true,message:'请选择底阿妈生成位置'}]">
                                           <a-radio-group v-model="codedata.codelocation" @change="handleGetMenuList">
                                               <a-radio value="busDirName">business(B端)</a-radio>
                                               <a-radio value="adminDirName" :disabled="!haseadmin">admin(A端)</a-radio>
                                           </a-radio-group>
                                       </a-form-item>
                                       <a-form-item label="上级菜单" field="pid" style="margin-bottom:15px;" :rules="[{required:true,message:'请填上级菜单'}]">
                                           <a-tree-select placeholder="选择上级菜单" :data="parntList" 
                                           @change="handleChangeMenu"
                                           :fieldNames="{ key: 'id',title: 'title',children: 'children'}" v-model="codedata.pid">
                                           </a-tree-select>
                                           <template #extra>
                                             <div v-if="codedata.codelocation=='busDirName'">若绑定到新的目录请到<a-link href="/system/rule">菜单管理(dev)</a-link>添加新目录</div>
                                             <div v-else>若绑定到新的目录请到admin后台的“菜单管理”添加新目录</div>
                                           </template>
                                       </a-form-item>
                                       <a-form-item label="菜单图标" field="icon" style="margin-bottom:15px;">
                                           <a-input-search v-model="codedata.icon" placeholder="选择图标/填写"  search-button>
                                               <template  v-if="codedata.icon" #prefix>
                                               <Icon :icon="codedata.icon"  ></Icon>
                                               </template>
                                               <template #button-icon>
                                                   <a-popover position="br" trigger="click">
                                                   <icon-apps :size="23"/>
                                                   <template #content>
                                                       <IconPicker @change="handleIcon"></IconPicker>
                                                   </template>
                                                   </a-popover>
                                               </template>
                                           </a-input-search>
                                       </a-form-item>
                                       <a-form-item field="rule_name" label="菜单名称:" validate-trigger="input" :rules="[{required:true,message:'请填写菜单名称'}]" >
                                           <a-input v-model="codedata.rule_name" placeholder="请填菜单名称" />
                                       </a-form-item>
                                       <a-form-item field="routename" label="路由name" validate-trigger="input" :rules="[{required:true,message:'请填写路由路由（name）'},{match:/^[A-Za-z0-9]+$/,message:'路由必须由字母数字组成'},{validator:handleChekHaseRule}]" >
                                           <a-input v-model="codedata.routename" placeholder="请填路由（name）" />
                                       </a-form-item>
                                       <a-form-item field="routepath" label="路由Path" validate-trigger="input" :rules="[{required:true,message:'请填写路由路由（path）'}]" >
                                           <a-input v-model="codedata.routepath" placeholder="请填路由（path）" />
                                           <template #extra>
                                           <div>如果是根目录或访问路由不加上级目录在名称前加/</div>
                                           </template>
                                       </a-form-item>
                                       <a-form-item field="component" label="前端组件路径" validate-trigger="input" :rules="[{required:true,message:'请填写前端组件路径'}]" >
                                           <a-input v-model="codedata.component" placeholder="请填前端组件路径" />
                                           <template #extra>
                                              <div>前端代码位置<span class="textbox">src/views{{codedata.component}}.vue</span></div>
                                           </template>
                                       </a-form-item>
                                       <a-form-item label="后台代码路径" :content-flex="false" :merge-props="false" :rules="[{required:true,message:'请填后台代码路径径'}]">
                                           <a-row >
                                               <a-col :span="13">
                                                   <a-form-item field="api_path"  validate-trigger="input" no-style>
                                                       <a-input v-model="codedata.api_path" placeholder="后台代码路径" />
                                                   </a-form-item>
                                               </a-col>
                                               <a-col :span="1">
                                                 <div class="flex-all-center" style="height: 32px;">/</div>
                                               </a-col>
                                               <a-col :span="10">
                                                   <a-form-item field="api_filename" validate-trigger="input" no-style>
                                                       <a-input v-model="codedata.api_filename"  placeholder="代码文件名称" />
                                                   </a-form-item>
                                               </a-col>
                                           </a-row>
                                           <template #extra>
                                              <div>后台代码在<span class="textbox">{{ codedata.api_path +"/"+codedata.api_filename}}</span>文件下</div>
                                           </template>
                                       </a-form-item>
                                   </div>
                                   <div class="form-rawer" v-show="TabsIndex==2">
                                       <div class="seachtable">
                                           <a-table :columns="columnsseach" :data="search_list" @change="handleChangeDragSearch" :pagination="false" size="small"
                                           :draggable="{ type: 'handle', width: 40 }" :scroll="{ y: `${cardboxHeight-70}px` }" :scrollbar="true">
                                               <template #drag-handle-icon>
                                                   <icon-drag-arrow />
                                               </template>
                                               <template #name="{ rowIndex }">
                                               <a-input v-model="search_list[rowIndex].name" placeholder="填写字段名称"/>
                                               </template>
                                               <template #issearch="{ rowIndex }">
                                                   <a-checkbox v-model="search_list[rowIndex].issearch" />
                                               </template>
                                               <template #searchway="{ record,rowIndex }">
                                                   <a-select v-model="record.searchway" placeholder="请选择" @change="handlesearchway(record.searchway,rowIndex)">
                                                       <a-option value="=">=</a-option>
                                                       <a-option value="!=">!=</a-option>
                                                       <a-option value=">">></a-option>
                                                       <a-option value=">=">>=</a-option>
                                                       <a-option value="<">{{ "<" }}</a-option>
                                                       <a-option value="<=">{{ "<=" }}</a-option>
                                                       <a-option value="like">Like</a-option>
                                                       <a-option value="set">Set</a-option>
                                                       <a-option value="in">In</a-option>
                                                       <a-option value="between">Between</a-option>
                                                   </a-select>
                                               </template>
                                               <template #searchtype="{ rowIndex }">
                                                  <a-select v-model="search_list[rowIndex].searchtype" placeholder="请选择">
                                                       <a-option value="text" :disabled="search_list[rowIndex].searchway=='between'||search_list[rowIndex].searchway=='in'||search_list[rowIndex].searchway=='set'">文本框</a-option>
                                                       <a-option value="number" :disabled="search_list[rowIndex].searchway=='between'||search_list[rowIndex].searchway=='in'||search_list[rowIndex].searchway=='set'">数字输入框</a-option>
                                                       <a-option value="select" :disabled="search_list[rowIndex].searchway=='between'||search_list[rowIndex].searchway=='in'">下拉框</a-option>
                                                       <a-option value="date" :disabled="search_list[rowIndex].searchway=='between'||search_list[rowIndex].searchway=='in'||search_list[rowIndex].searchway=='set'">日期选择</a-option>
                                                       <a-option value="daterange" :disabled="search_list[rowIndex].searchway!='between'">日期范围选择</a-option>
                                                       <a-option value="selectmore" :disabled="search_list[rowIndex].searchway!='in'">多选下拉框</a-option>
                                                       <a-option value="dic" :disabled="search_list[rowIndex].searchway=='between'||search_list[rowIndex].searchway=='in'||search_list[rowIndex].searchway=='set'">字典数据</a-option>
                                                       <a-option value="belongto" :disabled="search_list[rowIndex].searchway=='between'||search_list[rowIndex].searchway=='in'||search_list[rowIndex].searchway=='set'">关联表数据</a-option>
                                                   </a-select>
                                               </template>
                                               <template #searchwidth="{ rowIndex }">
                                                   <a-input-number v-model="search_list[rowIndex].searchwidth" placeholder="填写框宽"/>
                                               </template>
                                           </a-table>
                                       </div>
                                   </div>
                                   <div class="form-rawer"  v-show="TabsIndex==3">
                                       <a-radio-group  v-model="codedata.tpl_type">
                                           <div class="tpllist">
                                               <template v-for="item in tpllist"  >
                                                   <div class="tpl-item" @click="codedata.tpl_type=item.type">
                                                       <div class="tpl-wrap">
                                                           <div class="tpl-img">
                                                               <a-carousel
                                                                   indicator-type="never"
                                                                   :style="{
                                                                       width: '100%',
                                                                       height: '180px',
                                                                   }"
                                                               >
                                                                   <a-carousel-item v-for="image in item.images">
                                                                       <a-image
                                                                           :src="image"
                                                                           fit="fill"
                                                                           width="100%"
                                                                           height="100%"
                                                                       />
                                                                   </a-carousel-item>
                                                               </a-carousel>
                                                           </div>
                                                           <div class="tlp-option">
                                                               <div class="checkbox">
                                                                   <a-radio :value="item.type"><div class="title">{{ item.title }}</div></a-radio>
                                                               </div>
                                                               <div class="catetable" v-if="item.type==codedata.tpl_type&&(item.type.indexOf('cate')>-1)">
                                                                   <a-form-item field="cate_tablename" label="关联分类数据" :rules="[{required:true,message:'请选择关联分类数据数据表'}]" >
                                                                       <a-select v-model="codedata.cate_tablename" placeholder="请选择关联分类数据表" allow-clear>
                                                                           <template v-for="item in tablelist">
                                                                               <a-tooltip :content="item.title">
                                                                                   <a-option :value="item.name" >{{ item.name }}</a-option>
                                                                               </a-tooltip>
                                                                           </template>
                                                                       </a-select>
                                                                   </a-form-item>
                                                               </div>
                                                           </div>
                                                       </div>
                                                   </div>
                                               </template>
                                           </div>
                                       </a-radio-group>
                                   </div>
                               </a-form>
                           </div>
                         </a-scrollbar>
                       </div>
                     </div>
                   </template>
                 </a-split>
           </div>
       </a-card>
       <div class="actions">
           <a-space>
               <a-button @click="onReback">返回列表</a-button>
               <a-button @click="handleUninstall" type="primary" status="danger" v-if="codedata.is_install==1" :loading="uninstallloading">卸载代码</a-button>
               <a-button type="primary" :loading="loading" @click="onSubmitClick">{{ codedata.is_install?"重新生成":"立即生成" }}</a-button>
           </a-space>
       </div>
   </div>
   </template>
   <script setup lang="ts">
   import { onMounted, ref, onUnmounted,toRefs,computed, nextTick } from 'vue'
   import { IconPicker ,Icon} from '@/components/Icon';
   import {Message} from '@arco-design/web-vue';
   import { FormInstance } from '@arco-design/web-vue/es/form';
   import useLoading from '@/hooks/loading';
   // import { cloneDeep, join } from 'lodash-es';
   //api
   import { getContent,CodedataItem,uninstallcode,CodeListItem,save,getdbfield,TreeItem,getDicTable,getMenuParent} from '@/api/developer/generatecode';
   import { getTables,TableItem } from'@/api/common';
   import { RuleItem } from '@/api/system/rule';
   import {  checkHaseRule } from '@/api/common';
   import { columnsfiles,columnslist,columnsseach } from './data';
   import { useRoute, useRouter } from "vue-router";
   import { useAppStore} from '@/store';
   import { inArray} from '@/utils/is';
   import FormIconSelect from '@/components/autoPlugin/Form/FormIconSelect.vue';
   const appStore = useAppStore();
   const route = useRoute(), router = useRouter();
   const emits = defineEmits([])
   const props = defineProps({
   })
   const {} = toRefs(props)
   const RightWeidth=400;//左边属性宽
   const windowWidth = (document.documentElement.clientWidth || document.body.clientWidth)-220;
   const atrrSize = ref(`${windowWidth-RightWeidth}px`);
   const onelineCardRef = ref(null);
   const cardboxHeight = computed(() => (onelineCardRef.value ? onelineCardRef.value["$el"]["offsetHeight"] : 300));
   const cardboxWidtht=ref(windowWidth)
   const TabsIndex=ref(1)
   const codedata=ref<CodedataItem>({
       id:0,
       tablename:"",
       codelocation:"busDirName",
       comment:"",
       pid:0,
       rule_id:0,
       rule_name:"",
       icon:"",
       is_install:0,
       routepath:"",
       routename:"",
       component:"",
       api_path:"",
       api_filename:"",
       tpl_type:"list",
       cate_tablename:"",
   })
   const id_prod =import.meta.env.VITE_APP_ENV=="production"
   const parntList = ref<RuleItem[]>([]);
   const tablelist = ref<TableItem[]>([]);
   const dictablelist = ref<TableItem[]>([]);
   const modelroot=ref("")
   onMounted(() => {
       initData()
   })
   const initData=async()=>{
       await getpageData()
       tablelist.value = await getTables({});
       dictablelist.value = await getDicTable({});
       nextTick(()=>{
           cardboxWidtht.value=onelineCardRef.value ? onelineCardRef.value["$el"]["offsetWidth"] : 300
           atrrSize.value=(cardboxWidtht.value-RightWeidth)+"px"
       })
       //根据列表数据表获取分类表
       if(tablelist.value){
           const catetable=codedata.value.tablename+"_cate"
           if(tablelist.value.find((item)=>item.name==catetable)){
               codedata.value.cate_tablename=catetable
           }
       }
   }
    //离开
   onUnmounted(() => {
       appStore.footer=true
   })
   //获取页面数据
   const field_list = ref<CodeListItem[]>([]);
   const list = ref<CodeListItem[]>([]);
   const search_list = ref<CodeListItem[]>([]);
   const mpid=ref(0)
   const haseadmin=ref(false)
   const getpageData=async()=>{
      const mdata= await getContent({id:route.query.id})
      codedata.value=mdata.data
      haseadmin.value=mdata.haseadmin
      mpid.value=mdata.data.pid
      field_list.value=mdata.field_list
      list.value=mdata.list
      search_list.value=mdata.search_list
      handleGetMenuList()
   }
   //获取菜单列表
   const handleGetMenuList=async()=>{
       const resultdata = await getMenuParent({codelocation:codedata.value.codelocation});
       const parntList_df : any=[{id: 0,title: "一级菜单",pid: 0,locale:""}];
       if(resultdata.list){
           parntList.value=parntList_df.concat(resultdata.list)
       }else{
           parntList.value=[]
       }
       if(inArray(resultdata.ids,mpid.value)&&!inArray(resultdata.ids,codedata.value.pid)){
           codedata.value.pid=mpid.value
       }else if(!inArray(resultdata.ids,codedata.value.pid)){
           codedata.value.pid=0
       }
       if(!codedata.value.routename){//如果已经保存就不自动处理菜单
        const tanme_arr=codedata.value.tablename.split("_")
        if(tanme_arr&&tanme_arr.length==3){
            //父级菜单
            const ruledatas=parntList.value.find((item:any)=>item["routepath"]=="/"+tanme_arr[1])
            if(ruledatas){
                   codedata.value.pid=ruledatas.id
            }
        }
         handleChangeMenu(codedata.value.pid)
       }
   }
   const openAtrr=ref(false)
   const handleChangeAtrr=()=>{
       openAtrr.value=! openAtrr.value
       if(openAtrr.value){
         atrrSize.value=(cardboxWidtht.value-4)+"px"
       }else{
         atrrSize.value=(cardboxWidtht.value-RightWeidth)+"px"
       }
   }
   //获取表字段
   const fieldList = ref<TreeItem[][]>([]);
   const getTableField=async(value:any)=>{
       if(value){
           fieldList.value[value] = await getdbfield({tablename:value});
       }else{
           fieldList.value =[]
       }
   }
   //返回
   const onReback=()=>{
       router.go(-1)
   }
   //生成表单
   const formRef = ref<FormInstance>();
   const { loading, setLoading } = useLoading();
   const onSubmitClick=async()=>{
       if(id_prod){
           Message.error({content:"Go属于编译语言，生产环境无法修改代码，请您在开发环境使用！",id:"upStatus",duration:2000})
           return
       }
       try {
           const res = await formRef.value?.validate();
           if (!res) {
               setLoading(true);
               Message.loading({content:"生成中",id:"upStatus",duration:0})
               const data= await save({codedata:codedata.value,field_list:field_list.value,list:list.value,search_list:search_list.value});
               Message.success({content:"生成成功",id:"upStatus",duration:2000})
               setLoading(false);
           }
       } catch (error) {
           console.error(error)
           setLoading(false);
           Message.loading({content:"生成中",id:"upStatus",duration:1})
       }
   
   }
   //菜单改变
   const handleChangeMenu=(val:any)=>{
       const ruledata=parntList.value.find((item)=>item.id==val)
       if(ruledata){
           modelroot.value=ruledata['routepath']
       }
       WatchChangPath(codedata.value)
   }
   //修改路由及文件路径
   const WatchChangPath=(data:any)=>{
       const tanme_arr=data.tablename.split("_")
       if(!modelroot.value){//一级菜单
           codedata.value.routepath=`${tanme_arr[1]}`
           codedata.value.routename=`${tanme_arr[1]}`
           codedata.value.component=`${tanme_arr[1]}/index`
           codedata.value.api_path=`${codedata.value.codelocation=="busDirName"?"business":"admin"}/${tanme_arr[1]}`
           codedata.value.api_filename="index.go"
           return
       }
       const modle_name=(modelroot.value.split("/")[1]).trim()
       if(tanme_arr&&tanme_arr.length==3){
           codedata.value.routepath=`${tanme_arr[2]}`
           codedata.value.routename=`${tanme_arr[2]}`
           codedata.value.component=`${modelroot.value?modle_name:tanme_arr[1]}/${tanme_arr[2]}/index`
           codedata.value.api_path=`${codedata.value.codelocation=="busDirName"?"business":"admin"}/${modelroot.value?modle_name:tanme_arr[1]}`
           codedata.value.api_filename=tanme_arr[2]+".go"
       }else{
           codedata.value.routepath=`${tanme_arr[1]}`
           codedata.value.routename=`${tanme_arr[1]}`
           codedata.value.component=`${modelroot.value?modle_name:'xx'}/${tanme_arr[1]}/index`
           codedata.value.api_path=`${codedata.value.codelocation=="busDirName"?"business":"admin"}/${modelroot.value?modle_name:'xx'}`
           codedata.value.api_filename=tanme_arr[1]+".go"
       }
   }
   //卸载
   const uninstallloading=ref(false)
   const handleUninstall=async()=>{
       if(id_prod){
           Message.error({content:"Go属于编译语言，生产环境无法修改代码，请您在开发环境使用！",id:"upStatus",duration:2000})
           return
       }
       try {
           uninstallloading.value=true
           Message.loading({content:"卸载中",id:"upStatus",duration:0})
           await uninstallcode({id:route.query.id});
           uninstallloading.value=false
           codedata.value.is_install=0
           Message.success({content:"卸载成功",id:"upStatus",duration:2000})
       }catch (error) {
           uninstallloading.value=false
           Message.loading({content:"卸载中",id:"upStatus",duration:1})
       } 
   }
   //选择图标
   const handleIcon=(icon:any)=>{
       codedata.value.icon=icon
   }
   //拖拽排序
   const handleChangeDragField = (_data:any) => {
       _data.forEach((item:any,index:any)=>{
           item['field_weigh']=index+1
       })
       field_list.value=_data
   }
   const handleChangeDragList = (_data:any) => {
       _data.forEach((item:any,index:any)=>{
           item['list_weigh']=index+1
       })
       list.value=_data
   }
   const handleChangeDragSearch = (_data:any) => {
        _data.forEach((item:any,index:any)=>{
           item['search_weigh']=index+1
       })
       search_list.value=_data
   }
   //表单修改名称
   const handleinputForm=(data:any)=>{
       var fielddata = field_list.value.filter((item) => {
           return item.id == data.id;
       });
       fielddata[0].name = data.name;
       //列表
       var listdata = list.value.filter((item) => {
           return item.id == data.id;
       });
       listdata[0].name = data.name;
       var searchdata = search_list.value.filter((item) => {
           return item.id == data.id;
       });
       searchdata[0].name = data.name;
   }
   //设置显示列表展示Ui
   const handleUpshowui=(data:any)=>{
    var listdata = list.value.filter((item) => {
        return item.id == data.id;
    });
    var searchdata = search_list.value.filter((item) => {
        return item.id == data.id;
    });
    listdata[0].show_ui="";
    searchdata[0].searchway = "=";
    searchdata[0].searchtype = "text";
    if(data.formtype=="belongDic"){
        listdata[0].show_ui="dic";
        searchdata[0].searchtype="dic";
    }else if(data.formtype=="belongto"){
        searchdata[0].searchway = "belongto";
    }else if(data.formtype=="textarea"||data.formtype=="editor"){
        listdata[0].show_ui="des";
    }else if(data.formtype=="text"){
        searchdata[0].searchway = "like";
    }else if(data.formtype=="number"){
        searchdata[0].searchtype = "number";
    }else if(data.formtype=="switch"){
       listdata[0].show_ui="switchstatus";
       searchdata[0].searchtype = "select";
    }else if(data.formtype=="radio"){
       listdata[0].show_ui="tag";
       searchdata[0].searchtype = "select";
    }else if(data.formtype=="checkbox"){
       listdata[0].show_ui="tags";
       searchdata[0].searchway = "in";
       searchdata[0].searchtype = "selectmore";
    }else if(data.formtype=="date"){
       listdata[0].show_ui="date";
       searchdata[0].searchtype = "date";
    }else if(data.formtype=="datetime"){
       listdata[0].show_ui="datetime";
       searchdata[0].searchway = "between";
       searchdata[0].searchtype = "daterange";
    }else if(data.formtype=="colorpicker"){
       listdata[0].show_ui="color";
    }else if(data.formtype=="image"){
       listdata[0].show_ui="image";
    }else if(data.formtype=="images"){
       listdata[0].show_ui="images";
    }else if(data.formtype=="file"){
       listdata[0].show_ui="cellcopy";
    }
   }
   const GetAssetsFile = (url: string) => {
     return new URL(`../../../assets/images/${url}`, import.meta.url).href;
   };
   //模板列表
   const contentlist=ref(GetAssetsFile("contentlist.png"))
   const contentcatelist=ref(GetAssetsFile("contentcatelist.png"))
   const sitecatelist=ref(GetAssetsFile("sitecatelist.png"))
   const tpllist=ref([
       {
           "title":"简单数据列表,没有富文本，适合简单数据列表管理（不带分类）",
           "type":"list",
           "images":[
           contentlist.value,
           ],
       },
       {
           "title":"富文本数据列表,带有富文本，适合丰富数据列表管理（不带分类）",
           "type":"contentlist",
           "images":[
           contentlist.value
           ],
       },
       {
           "title":"富文本数据带有分类数据,带有富文本，右上角管理分类数据，适合丰富数据列表管理（带分类）",
           "type":"contentcatelist",
           "images":[
             contentcatelist.value
           ],
       },
       {
           "title":"富文本数据带有分类数据,带有富文本，左侧有分类菜单数据，适合丰富数据列表管理（带分类）",
           "type":"sitecatelist",
           "images":[
           sitecatelist.value
           ],
       },
   ])
   //搜索条件自动赋值
   const handlesearchway=(val:any,rowIndex:any)=>{
       if(val=="between"){
           search_list.value[rowIndex].searchtype="daterange"
           search_list.value[rowIndex].searchwidth=230
       }else if(val=="in"){
           search_list.value[rowIndex].searchtype="selectmore"
           search_list.value[rowIndex].searchwidth=150
       }else if(val=="set"){
           search_list.value[rowIndex].searchtype="select"
           search_list.value[rowIndex].searchwidth=120
       }else{
           search_list.value[rowIndex].searchtype="text"
           search_list.value[rowIndex].searchwidth=120
       }
   }
    //检查是否存在路由
    const handleChekHaseRule=(value:any, cb:any)=>{
        return new Promise(async(resolve) => {
        if(value){
            const resData = await checkHaseRule({id:codedata.value.rule_id,codelocation:codedata.value.codelocation,routename:value});
            if(!resData){
            cb("该路由已被占用请换一个")
            }
        }
        resolve(true)
        })
    }
   </script>
   
   <style lang="less" scoped>
   .gridbox{
       height: 20px;
       overflow: hidden;
       display: flex;
       align-items: center;
       user-select: none;
       cursor: pointer;
   }
   .page-container{
       border-radius: 4px;
       padding: 10px;
       padding-bottom: 20px;
       height: 100%;
       .maker-rawer{
           height: 100%;
           .maker-content{
               height: 100%;
               width: 100%;
               .header{
                   height: 40px;
                   .htabs{
                       // padding-top: 10px;
                   }
               }
               .header-table{
                   border-bottom: var(--color-neutral-3) solid 1px;
                   display: flex;
                   align-items:center;
                   .title{
                       padding-left: 15px;
                       color:var(--color-neutral-8);
                       font-size: 16px;
                   }
               }
               .content{
                   flex: 1;
                   width: 100%;
                   .table-box{
                       flex: 1;
                       padding: 0px 10px;
                       .title{
                           padding-left: 15px;
                           height: 30px;
                           display: flex;
                           align-items:center;
                       }
                       .table-main{
   
                       }
                   }
                   .formbox{
                       padding:  10px;
                       .form-rawer{
                           .textbox{
                               color: var(--color-neutral-6);
                           }
                       }
                   }
                   //模板列表
                   .tpllist{
                       width: 100%;
                       height: 100%;
                       .tpl-item{
                           cursor: pointer;
                           margin-bottom: 10px;
                           .tpl-wrap{
                               width: calc(100% - 2px);
                               height: calc(100% - 2px);
                               border: 1px solid #f5f5fa;
                               border-radius: 3px;
                               user-select: none;
                               overflow: hidden;
                               padding: 10px;
                               .tpl-img{
                                   height: 180px;
                                   border-radius: 3px;
                                   overflow: hidden;
                               }
                               .tlp-option{
                                   margin-top: 10px;
                                   .catetable{
                                       margin-top: 10px;
                                       align-items: center;
                                       .label{
                                           padding-right: 5px;
                                       }
                                       .text{
                                           flex:1;
                                       }
                                   }
                               }
                           }
                       }
                   }
               }
           }
       }
       //提交按钮
       .actions{
           position: fixed;
           left: 0;
           right: 0;
           bottom: 0;
           height: 50px;
           padding: 10px;
           background: var(--color-bg-2);
           text-align: right;
       }
   }
   //设置tab
   :deep(.arco-tabs-nav-tab-list){
       width: 100%;
       display: flex;
       justify-content: center;
   }
   :deep(.arco-tabs-nav-type-line .arco-tabs-tab){
       flex: 1;
       text-align: center;
       align-self: center;
       justify-content: center;
   }
   //设置car
   :deep(.arco-card-size-medium .arco-card-body){
       height: 100%;
       padding:0px ;
   }
   :deep(.arco-carousel-item-current){
       text-align: center;
   }
   :deep(.arco-split-trigger-icon-wrapper){
       user-select: none;
       background-color:var(--color-neutral-1);
   }
   :deep(.arco-table-cell){
       padding: 5px 10px;
   }
    //移动图标
    .drag-icon-center{
         border-radius: 2px 0 0 2px;
         cursor: pointer;
         width: 5px;
         height: 40px;
         display:flex;
         user-select: none;
         align-items:center;
         justify-content:center;
         color:#ffffff;
         font-size: 14px;
         background-color: rgb(var(--arcoblue-6));
       }
       .selfsplit-trigger-icon {
           display: inline-block;
           margin: -3px;
       }
   </style>
   