<template>
    <div class="container" >
      <page-card breadcrumb scrollPage>
        <div class="codetopbar">
            <div class="cleft">
              <a-space>
                <ButtonGroup >
                  <a-button @click="handleTabwarehouse('public')" :type="(tabwarehouse=='public'?'primary':undefined)" ><Icon icon="svgfont-code" class="tbicon" :size="12" ></Icon>公共仓</a-button>
                  <a-button @click="handleTabwarehouse('private')" :type="(tabwarehouse=='private'?'primary':undefined)"><Icon icon="svgfont-yinsixieyi" class="tbicon" :size="12"  ></Icon>私有仓</a-button>
                </ButtonGroup>
                <a-popconfirm :type="confhouse.private.url?'success':'error'" position="rt" @ok="upprivateURl">
                  <template #content>
                      <div class="CofigContent">
                        <a-input :style="{width:'290px'}" v-model="confhouse.private.url" placeholder="填写您的所在的企业私有仓网网址" allow-clear />
                        <div class="tig">填写您的企业私有仓网址(http或者https开头)</div>
                      </div>
                  </template>
                  <a-link v-if="tabwarehouse=='private'">
                    <template #icon>
                      <icon-edit :size="18"/>
                    </template>
                    <span style="user-select: none;">配置企业私有仓网址</span>
                  </a-link>
                  <a-link v-if="tabwarehouse=='private'" icon href="//goflys.cn/codedetail?id=21"  target="_blank" style="margin-left: 15px;">
                    <span style="user-select: none;">查看私有仓部署说明</span>
                  </a-link>
                </a-popconfirm>
              </a-space>
            </div>
            <div class="cright">
            <a-tooltip content="点击获取框架最新版本">
              <a-link @click="handleAsyncVersion"><Icon icon="svgfont-banbengengxin1" class="tbicon" :size="16" ></Icon><span style="user-select: none;">框架版本更新<span style="font-size: 10px;">v{{ version }}</span></span></a-link>
            </a-tooltip> 
          </div>
        </div>
        <div class="one-line-bar">
            <a-radio-group v-model="formModel.cid" type="button" size="large" @change="handleChangeCate" >
              <a-radio value="all">全部</a-radio>
              <template v-for="item in catelist">
                <a-tooltip :content="item.remark">
                <a-radio  :value="item.id">{{item.name}}</a-radio>
                </a-tooltip> 
              </template>
            </a-radio-group>
        </div>
        <div class="one-line-bar">
          <a-space>
            <a-input :style="{width:'220px'}"  v-model="formModel.searchword" placeholder="搜索代码名称" @press-enter="search" allow-clear />
            <a-button type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
              查询
            </a-button>
            <a-button @click="reset">
              {{ $t('searchTable.form.reset') }}
            </a-button>
            <a-radio-group v-model="formModel.type" type="button" @change="handleIsInstall">
              <a-radio value="all">全部</a-radio>
              <a-radio value="card">已安装</a-radio>
            </a-radio-group>
            <a-tooltip :content="threeOperation(formModel.code_token,'点击可以退出登录','点击可以登录您的账号')">
              <a-button status="success" v-if="tabwarehouse=='public'" @click="handleLogin"><Icon :icon="threeOperation(formModel.code_token,'icon-export','svgfont-denglu1')" class="tbicon" :size="14" ></Icon>{{formModel.code_token?"退出登录":"登录账号"}}</a-button>
            </a-tooltip>
            <a-tooltip content="把你需要帮助需求发给其他人做">
              <a-button status="success" @click="handlePublishRequire"><Icon icon="svgfont-100wanxuqiufabu" class="tbicon" :size="14" ></Icon>需求发布</a-button>
            </a-tooltip>
            <a-tooltip content="对开发完成代码打包、打包好的zip上传代到码仓">
              <a-button status="danger" @click="handlePackCode"><Icon icon="svgfont-zipyasuowenjian" class="tbicon" :size="14" ></Icon>打包/上传仓库</a-button>
            </a-tooltip>
            <a-tooltip content="检测包名包名是否被占用">
              <a-button status="warning" @click="handleCheckPackageName"><Icon icon="svgfont-biaoshimoxing" class="tbicon" :size="14" ></Icon>检测包名</a-button>
            </a-tooltip>
            <a-upload
                accept=".zip"
              :show-file-list="false"
              :custom-request="handleInstallLLocalCode"
              @before-upload="beforeUpload"
            >
              <template #upload-button>
              <a-tooltip content="点击这里可以选择本地代码zip包直接安装">
                <a-button >安装本地包</a-button>
              </a-tooltip>
            </template>
            </a-upload>
          </a-space>
        </div>

        <a-table
           row-key="id"
          :loading="loading"
          :pagination="pagination"
          :columns="(cloneColumns as TableColumnData[])"
          :data="renderData"
          size="large"
          :default-expand-all-rows="true"
          ref="artable"
          @page-change="handlePaageChange" 
          @page-size-change="handlePaageSizeChange" 
        >
          <template #title="{record,column}">
              <a-tooltip content="查看代码包介绍和使用文档">
                <a-link v-if="tabwarehouse=='public'" :href="`https://goflys.cn/codedetail?id=${record.id}&token=${token_str}`" target="_blank" icon>{{ record[column.dataIndex] }}</a-link>
                <a-link v-else @click="handleShowPrivate(record)">{{ record[column.dataIndex] }}</a-link>
              </a-tooltip>
          </template>
          <template #name="{ record }">
           {{ record.name }}<span v-if="record.nickname" style="padding-left: 5px;color: var(--color-neutral-4);">{{ record.nickname }}</span>
          </template>
          <template #price="{record,column}">
            <a-badge  >
              <template #content>
                {{record.user_type_text}}
               </template>
              <span :style="{color: record[column.dataIndex]==0?`rgb(var(--green-6))`:`rgb(var(--orange-6))`}">{{ record[column.dataIndex]==0?"免费":record[column.dataIndex] }}</span>
            </a-badge>
          </template>
          <template #author="{record,column}">
            <span :style="{color: record[column.dataIndex]?`var(--color-neutral-10)`:`var(--color-neutral-4)`}">{{ record[column.dataIndex]?record[column.dataIndex]:"匿名"}}</span>
          </template>
          <template #version="{record,column}">
            <span :style="{color: record[column.dataIndex]?`var(--color-neutral-10)`:`var(--color-neutral-4)`}">{{ record[column.dataIndex]?record[column.dataIndex]:"--"}}</span>
          </template>
          <template #des="{ record ,column}">
            <a-typography-paragraph  class="mytypography" :ellipsis="{rows: 2,showTooltip:true}">{{ record[column.dataIndex] }}</a-typography-paragraph>
          </template>
          <template #createtime="{record,column}">
            {{dayjs(record[column.dataIndex]*1000).format("YYYY-MM-DD")}}
          </template>
          <template #status="{ record ,column}">
            <a-tag :color="statusFont(record[column.dataIndex],'color')" >{{ statusFont(record[column.dataIndex],'text') }}</a-tag>
          </template>
          <template #operations="{ record }">
            <div class="btnoperation">
              <ButtonGroup :disabled="canClickbtn(record)">
                <a-button v-if="!record.is_install" @click="canClickbtn(record)?null:handleInstallCode(record)" type="primary" status="success">
                  <a-link :hoverable="false">安装</a-link>
                </a-button>
                <a-popconfirm content="您确定要卸载吗?" @ok="handleUninstallCode(record)">
                  <a-button v-if="record.is_install" type="primary" status="danger">
                      <a-link  status="danger" :hoverable="false">卸载</a-link>
                  </a-button>
                </a-popconfirm>
                <a-dropdown trigger="click" >
                  <a-button   >
                    终端代码下载<icon-down style="margin-left: 5px;"/>
                  </a-button>
                  <template #content>
                      <div class="dropmenu">
                        <a-doption v-for="packcode in record.attachment" :title="`点击下载${packcode.title}`" @click="handleDownCode(packcode)">
                          <div class="dropItem" >
                            <div class="icon"><Icon icon="svgfont-yasuowenjian" class="tbicon" :size="14" ></Icon></div>
                            <div class="title">{{packcode.title}}</div>
                          </div>
                        </a-doption>
                        <div class="dropmenuempty" v-if="!record.attachment||record.attachment.length==0">
                            不含终端代码
                        </div>
                      </div>
                  </template>
                </a-dropdown>
              </ButtonGroup>
            </div>
          </template>
        </a-table>
      </page-card>
       <!--登录-->
       <Login ref="loginRef" @ok="loginOk"></Login>
       <RequireForm @register="registerRequireModal"  @success="fetchData"/>
       <PackUpCode @register="registerPackUpCodeModal"  />
       <CheckPackName  v-model="showCheckPack" :domian="confhouse[tabwarehouse].url"/>
       <AddOrder ref="addOrderRef"  />
       <ShowPrivateContent @register="registerPrivateModal" />
    </div>
  </template>
  
  <script lang="ts" setup>
    import { ref, reactive, watch, onMounted } from 'vue';
    import useLoading from '@/hooks/loading';
    import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
    import type { RequestOption} from '@arco-design/web-vue/es/upload/interfaces';
    import cloneDeep from 'lodash/cloneDeep';
    import dayjs from 'dayjs';
    //数据
    import { columns} from './data';
    //api
    import { getList,asyncVersion,upBaseCode,DataItem,getCodeCate,upPrivateHouse,checkIsPay} from '@/api/developer/codestore';
    import { downCode,installCode,uninstallCode,getInstallPack,installLocalCode } from '@/api/developer/codestoreoption';
    //表单
    import { useModal } from '/@/components/Modal';
    import {Icon} from '@/components/Icon';
    import { Message,Button,Modal} from '@arco-design/web-vue';
    import { Pagination } from '@/types/global';
    import { threeOperation } from '@/utils/is';
    import { cstoreIsLogin,getCstoreToken,clearCstoreToken } from '@/utils/auth';//登录判断
    import Login from "./Login.vue"; 
    import RequireForm from './RequireForm.vue';
    import PackUpCode from './PackUpCode.vue';
    import CheckPackName from './CheckPackName.vue';
    import ShowPrivateContent from './ShowPrivateContent.vue';
    import AddOrder from "./AddOrder.vue"; 
    var ButtonGroup=Button.Group
    const [registerRequireModal, { openModal:openRequireModal }] = useModal();
    const [registerPackUpCodeModal, { openModal:openPackUpCodeModal }] = useModal();
    const [registerPrivateModal, { openModal:openPrivateModal }] = useModal();
    const tabwarehouse = ref('public');
    const token_str = ref(getCstoreToken());
    //配置仓库地址
    // let baseRootUrl = import.meta.env.VITE_APP_ENV=="production"?"https://api.goflys.cn":"http://localhost:8110"
    let baseRootUrl = "https://api.goflys.cn"
    const confhouse= ref({
      public:{
          url:baseRootUrl,
      },
      //私有地址
      private:{
        url:"",
      }
    });
    //分类数据
    const catelist = ref<DataItem[]>([]);
    //分页
    const basePagination: Pagination = {
      current: 1,
      pageSize: 10,
    };
    const pagination = reactive({
      ...basePagination,
      showTotal:true,
      showPageSize:true,
    });
    type Column = TableColumnData & { checked?: true };
    const { loading, setLoading } = useLoading(true);
    const renderData = ref([]);
    const cloneColumns = ref<Column[]>([]);
     //查询字段
     const generateFormModel = () => {
      return {
        cid: 'all',
        type: 'all',
        innames: 'all',
        searchword: '',
        code_token:getCstoreToken()||"",//登录账号的token
      };
    };
    const formModel = ref(generateFormModel());
    const fetchData = async () => {
      setLoading(true);
      try {
        const data= await getList({page:pagination.current,pageSize:pagination.pageSize,...formModel.value},confhouse.value[tabwarehouse.value].url);
        renderData.value = data.items;
        pagination.current = data.page;
        pagination.total = data.total;
      } catch (err) {
      } finally {
        setLoading(false);
      }
    };
    //组件挂载完成后执行的函数
    onMounted(()=>{
      showCheckPack.value=false
      getCateList()
    })
    //获取分类
    const version=ref("")
   const getCateList=async()=>{
    catelist.value=[];
    try {
         const resdata=await getCodeCate({baseurl:confhouse.value[tabwarehouse.value].url})
        if(resdata){
          catelist.value=resdata.catedata
          version.value=resdata.version
          confhouse.value.private.url=resdata.companyPrivateHouse
          //初始化数据
          renderData.value=[]
          pagination.current =1
          pagination.total =0
          fetchData();
        }
      }catch (error) {
        renderData.value=[]
        pagination.current =1
        pagination.total =0
        Message.success({content:"",id:"asyncVersion",duration:10})
      } 
    }
   //安装本地包
   const beforeUpload=()=>{
      if(import.meta.env.VITE_APP_ENV=="production"){
        Message.warning({content:"生产环境不能安装插件，请在开发环境中使用！",id:"upkey"})
        return  false
       }
       return  true
   }
   const handleInstallLLocalCode = (options: RequestOption) => {
      const controller = new AbortController();
        (async function requestWrap() {
          const {
            onProgress,
            onError,
            fileItem,
          } = options;
          onProgress(20);
          const onUploadProgress = (event: ProgressEvent) => {
            let percent;
            if (event.total > 0) {
              percent = (event.loaded / event.total) * 100;
            }
            onProgress(parseInt(String(percent), 10), event);
          };
          try {
            //开始手动上传
            Message.loading({content:"上传本地代码中",id:"upkey",duration:0})
            const filename=fileItem?.name||""
            const resdata = await installLocalCode({ name: 'file', file: fileItem.file as Blob, filename,data:{}},onUploadProgress);
            if(resdata){
              if(resdata["code"]==0){
                  Message.success({content:resdata["message"],id:"upkey",duration:2000})
                  if (resdata){
                    Message.loading({content:"安装中",id:"upkey",duration:0})
                    await installCode({name:resdata.data});
                    Message.success({content:"安装成功",id:"upkey",duration:2000})
                  }else{
                    Message.error({content:"已经安装",id:"upkey",duration:2000})
                  }
              }else{
                Message.error({content:resdata["message"],id:"upkey",duration:2000})
              }
            }else{
              Message.error({content:resdata["message"],id:"upkey",duration:2000})
            }
          } catch (error) {
            onError(error);
            Message.error({content:"上传失败",id:"upkey",duration:2000})
          }
        })();
        return {
          abort() {
            controller.abort();
          },
        };
    };
   //查找
   const search = () => {
     pagination.current =1
      fetchData();
    };
    const reset = () => {
      formModel.value = generateFormModel();
      pagination.current =1
      fetchData();
    };
  
    watch(
      () => columns.value,
      (val) => {
        cloneColumns.value = cloneDeep(val);
      },
      { deep: true, immediate: true }
    );
    //分页
    const handlePaageChange = (page:any) => {
      pagination.current=page
      fetchData();
    }
    //分页总数
    const handlePaageSizeChange = (pageSize:any) => {
      pagination.pageSize=pageSize
      fetchData();
    }
    //检查代码是否有新的代码更新
    const id_prod =import.meta.env.VITE_APP_ENV=="production"
    const handleAsyncVersion=async()=>{
      try {
          Message.loading({content:"检查版本更新中",id:"asyncVersion",duration:0})
          const res= await asyncVersion({code_token:formModel.value.code_token},baseRootUrl);
         if(res){
          Message.success({content:"",id:"asyncVersion",duration:10})
           if(res.type==0){
            Modal.warning({
              title: res.title ,
              content:res.content,
            });
           }else{
            Modal.success({
              title: res.title ,
              content:res.content ,
              hideCancel:false,
              cancelText:"下次更新",
              okText:"立即更新",
              onOk:async()=>{
                if(id_prod){
                  Message.error({content:"Go属于编译语言，生产环境无法更新代码，请您在开发环境更新！",id:"asyncVersion",duration:2000})
                  return
                }
                try {
                    Message.loading({content:"更新中",id:"asyncVersion",duration:0})
                    await upBaseCode({code_token:formModel.value.code_token},baseRootUrl);
                    Message.success({content:"更新成功",id:"asyncVersion",duration:1000})
                }catch (error) {
                  Message.success({content:"",id:"asyncVersion",duration:10})
                } 
              },
            });
           }
         }else{
           Message.success({content:"",id:"asyncVersion",duration:10})
         }
      }catch (error) {
        Message.success({content:"",id:"asyncVersion",duration:10})
      } 
    }
  //切换仓库
   const handleTabwarehouse=(val:string)=>{
     tabwarehouse.value=val
     //切换仓库
     getCateList()
   }
  //切换分类
   const handleChangeCate=(val:any)=>{
    formModel.value.cid=val
     //切换分类
     pagination.current =1
     fetchData();
   }
  //获取是否安装状态
   const handleIsInstall=async(val:any)=>{
    if(val=="all"){
      formModel.value.innames="all"
    }else{
      const res= await getInstallPack({});
      formModel.value.innames=res
    }
    pagination.current =1
     fetchData();
   }
   //更新私有仓地址
   const upprivateURl=async()=>{
    try {
        Message.loading({content:"更新私有仓地址中",id:"upprivateURl",duration:0})
        await upPrivateHouse({companyPrivateHouse: confhouse.value.private.url});
        Message.success({content:"更新成功",id:"upprivateURl",duration:2000})
      }catch (error) {
        Message.success({content:"",id:"upprivateURl",duration:10})
      } 
   }
    //打包
    const handlePackCode=async()=>{
      openPackUpCodeModal(true, {
        isUpdate: false,
        record:{domian:confhouse.value[tabwarehouse.value].url,tabwarehouse:tabwarehouse.value,code_token:formModel.value.code_token}
      });
    }
    //安装
    const addOrderRef=ref()
    const handleInstallCode=async(record:any)=>{
      if(record.isPay){
        InstallPack(record)//下载安装
      }else{
        if(formModel.value.code_token){//购买代码
          const res= await checkIsPay({code_token:formModel.value.code_token,codestore_id:record.id},confhouse.value[tabwarehouse.value].url);
          if(res){
            InstallPack(record)//下载安装
          }else{
            addOrderRef.value.showModel({code_token:formModel.value.code_token,codestore_id:record.id,type:record.type,name:record.name,title:record.title,price:record.price,paycode:record.paycode},confhouse.value[tabwarehouse.value].url)
          }
        }else{//登录窗口
          loginRef.value.showLogin(confhouse.value[tabwarehouse.value].url)
        }
      }
    }
    //执行下载安装
    const InstallPack=(record:any)=>{
      Modal.warning({
        title: '温馨提示',
        content: `您确定要安装“${record.title}”吗？`,
        okText:"确定",
        hideCancel:false,
        titleAlign:"start",
        cancelText:"取消",
        onOk:(async()=>{
          try {
            Message.loading({content:"下载代码中",id:"install",duration:0})
            const downres= await downCode({code_token: formModel.value.code_token,name:record.name,id:record.id},confhouse.value[tabwarehouse.value].url);
            if (downres){
              Message.loading({content:"安装中",id:"install",duration:0})
              await installCode({name:record.name});
              Message.success({content:"安装成功",id:"install",duration:2000})
            }else{
              Message.error({content:"已经安装",id:"install",duration:2000})
            }
            fetchData();
          }catch (error) {
            Message.loading({content:"下载代码中",id:"install",duration:1})
          }
        })
      });
    }
    //卸载
    const handleUninstallCode=async(record:any)=>{
      try {
        Message.loading({content:"卸载中",id:"uninstall",duration:0})
        const res= await uninstallCode({name:record.name});
        Message.success({content:"卸载成功",id:"uninstall",duration:2000})
        fetchData();
      }catch (error) {
        Message.loading({content:"卸载中",id:"uninstall",duration:1})
      }
    }
    //下载客户终端代码
    const handleDownCode=(codeData:any)=>{
      if(codeData.url){
        const link = document.createElement('a') // 创建一个 a 标签用来模拟点击事件	
        link.style.display = 'none'
        link.href =codeData.url
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
      }else{
        Modal.warning({
          title: '温馨提示',
          content: '请您登录并检查是否已经购买再来下载，如果下载失败请联系客服。',
          okText:"知道了"
        });
      }
    }
    //登录账号
    const loginRef=ref()
    const handleLogin=()=>{
      if(cstoreIsLogin()){//退出登录
        clearCstoreToken()
        formModel.value.code_token=""
        Message.success({content:"退出登录成功",id:"PackCode"})
      }else{
        loginRef.value.showLogin(confhouse.value[tabwarehouse.value].url)
      }
    }
    //登录成功
    const loginOk=(token:string)=>{
      formModel.value.code_token=token
      token_str.value=token
    }
    //发布需求
    const handlePublishRequire=()=>{
      if(formModel.value.code_token){
        openRequireModal(true, {
          isUpdate: false,
          record:{domian:baseRootUrl,code_token:formModel.value.code_token}
        });
      }else{//登录窗口
        loginRef.value.showLogin(confhouse.value[tabwarehouse.value].url)
      }
    }
    //判断是否可以点击
    const canClickbtn=(record:any)=>{
      if(tabwarehouse.value=="private"){
        if(record.status==1||record.status==4){
          return false
        }else{
          return true
        }
      }else{
        if((record.type==0&&record.status==1)||(record.type==1&&record.status==4)){
          return false
        }else{
          return true
        }
      }
    }
    //检测包名是否可用
    const showCheckPack=ref(false)
    const handleCheckPackageName=()=>{
      showCheckPack.value=true
    }
    //查看私有包详情
    const handleShowPrivate=(record:any)=>{
      openPrivateModal(true, {
        isUpdate: false,
        record: {content:record.content,title:record.title}
      });
    }
    //状态
    const statusFont=(val:any,type:any)=>{
       let str="",color="";
       switch(val) {
          case 1:
          str="已发布"
          color="blue"
              break;
          case 2:
          str="待接单"
          color="orange"
              break;
          case 3:
          str="已接单"
          color="green"
              break;
          case 4:
          str="开发完成"
          color="arcoblue"
              break;
          case 5:
          str="审批失败"
          color="red"
              break;
          case 6:
          str="拒绝接单"
          color="red"
              break;
          default:
          str="--"
          color=""
      } 
      if(type=="color"){
         return color
      }else{
        return str
      }
    }
  </script>
  
  <script lang="ts">
    export default {
      name: 'codestore',
    };
  </script>
  
  <style scoped lang="less">
  .mytypography{
    width: 100%;
    margin-bottom:0;
  }
    .tbicon{
      margin-right: 5px;
    }
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
    .codetopbar{
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 5px;
    }
    .one-line-bar{
      padding-bottom: 2px;
      margin-bottom: 3px;
      overflow-x: auto;
      :deep(.arco-radio-button-content){
        white-space: nowrap; 
      }
    }
    .CofigContent{
      .tig{
        padding-top: 5px;
        color: var(--color-neutral-4);
      }
    }
    .dropmenu{
      width: 135px;
      .dropItem{
        display: flex;
     
      }
      .dropmenuempty{
        text-align: center;
      }
    }
  :deep(.arco-table .arco-table-cell) {
      padding: 9px 10px;
  }
  :deep(.arco-badge-custom-dot) {
    top: -10px;
    width: 63px;
    color: rgb(var(--green-5));
    background-color: rgb(var(--orange-1));
  }
  .btnoperation{
    text-align: right;
  }
  </style>
  