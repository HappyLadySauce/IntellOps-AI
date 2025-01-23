<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :ok-text="okText" :isPadding="false" :maskClosable="false" :loading="loading" width="1100px" :minHeight="600" :title="getTitle" @ok="handleSubmit">
    <div  :style="{'height':`100%`}">
       <div class="navtabar">
          <a-steps type="arrow" :current="currentindex">
            <a-step description="把开发完成的代码打包，菜单、数据表数据全部打包并压缩成zip包，到右边上传到代码仓库" @click="currentindex=1" style="cursor: pointer;">代码打包</a-step>
            <a-step description="把已经打包代码上传到代码仓，如果上传到公共仓请填写价格(到社区完善收款码)" @click="currentindex=2" style="cursor: pointer;">上传代码到仓库</a-step>
          </a-steps>
       </div>
       <div class="tabcontent">
        <div class="tabitem" v-if="currentindex==1">
          <div class="formbox">
            <a-form ref="formRef" :model="formData" auto-label-width >
              <a-row :gutter="16">
                <a-col :span="12">
                  <a-form-item field="name" label="包名" :rules="[{required:true,message:'请填写包名'}]" >
                      <a-input v-model="formData.name" placeholder="请填包名、插件标识（以字母开头、数字命名）" @input="filterName" allow-clear @change="handlePackNameChange"/>
                  </a-form-item>
                  <a-form-item field="title" label="名称" :rules="[{required:true,message:'请填写包名称'}]" >
                      <a-input v-model="formData.title" placeholder="请填包名称（标题）" allow-clear/>
                  </a-form-item>
                  <a-form-item field="version" label="版本号" :rules="[{required:true,message:'请填写版本号'}]" >
                      <a-input v-model="formData.version" placeholder="请填版本号" allow-clear/>
                  </a-form-item>
                  <a-form-item field="noVerifyAPIRoot" label="忽略请求验证" >
                    <a-select v-model="formData.noVerifyAPIRoot" placeholder="选择忽略接口请求验证" allow-search multiple allow-clear>
                        <a-option v-for="item in appDirlist" :value="item">{{ item }}</a-option>
                    </a-select>
                  </a-form-item>
                  <a-form-item field="packtables" label="数据表" style="margin-bottom: 10px;">
                    <a-select v-model="formData.packtables" placeholder="选择选择导入的数据表" allow-search multiple allow-clear>
                      <template #label="{ data }">
                          <span>{{data?.value}}</span>
                        </template>
                        <a-option v-for="item in tablelist" :value="item.name">{{ item.name+" "+item.title }}</a-option>
                    </a-select>
                  </a-form-item>
                  <a-form-item field="des" label="描述" style="margin-bottom: 10px;">
                      <a-textarea v-model="formData.des" placeholder="请填说明" :auto-size="{minRows:2,maxRows:3}" allow-clear/>
                  </a-form-item>
                  <a-form-item field="menujson" label="B端菜单" style="margin-bottom: 10px;">
                    <div class="menubox">
                      <div class="menutabar">
                          <a-radio-group type="button" v-model="bmenuboxbar">
                            <a-radio value="selectemenu">选择菜单</a-radio>
                            <a-radio value="editmenu">编辑菜单</a-radio>
                          </a-radio-group>
                        </div>
                        <div class="selectemenu" v-if="bmenuboxbar=='selectemenu'">
                            <a-tree-select
                              v-model="formData.businessmenutree" 
                              :allow-search="true"
                              :allow-clear="true"
                              :tree-checkable="true"
                              :data="BusinessMenutreeData"
                              :loading="loadingMenu"
                              placeholder="选择系统菜单"
                              :fieldNames="{
                                  key: 'id',
                                  title: 'title',
                                  children: 'children'
                                }"
                              @change="handleMarkeMenuBusiness"
                              ></a-tree-select>
                        </div>
                        <div class="parambox" v-else>
                          <CodeEditor v-model:value="formData.businessmenujson" :mode="modeValue" />
                        </div>
                    </div>
                  </a-form-item>
                  <a-form-item field="menujson" label="A端菜单" style="margin-bottom: 10px;">
                    <div class="menubox">
                      <div class="menutabar">
                          <a-radio-group type="button" v-model="amenuboxbar">
                            <a-radio value="selectemenu">选择菜单</a-radio>
                            <a-radio value="editmenu">编辑菜单</a-radio>
                          </a-radio-group>
                        </div>
                        <div class="selectemenu" v-if="amenuboxbar=='selectemenu'">
                            <a-tree-select
                              v-model="formData.adminmenutree" 
                              :allow-search="true"
                              :allow-clear="true"
                              :tree-checkable="true"
                              :data="AdminMenutreeData"
                              :loading="loadingMenu"
                              placeholder="选择系统菜单"
                              :fieldNames="{
                                  key: 'id',
                                  title: 'title',
                                  children: 'children'
                                }"
                              @change="handleMarkeMenuAdmin"
                              ></a-tree-select>
                        </div>
                        <div class="parambox" v-else>
                          <CodeEditor v-model:value="formData.adminmenujson" :mode="modeValue" />
                        </div>
                    </div>
                  </a-form-item>
                  <a-form-item field="installcover" label="安装方式" style="margin-bottom: 10px;">
                    <a-radio-group v-model="formData.installcover">
                        <a-radio :value="false">新模块安装</a-radio>
                        <a-radio :value="true">覆盖原有代码（如登录界面）</a-radio>
                      </a-radio-group>
                  </a-form-item>
                </a-col>
                <a-col :span="12">
                 <div class="borderbox" >
                     <span class="title">后端代码</span>
                     <div class="borderItem" style="margin-top: 10px;">
                       <div class="formValue">
                        <a-form-item field="appfolders" label="app项目" style="margin-bottom: 10px;">
                          <a-select v-model="formData.appfolders" placeholder="选择后端app文件夹下的目录" allow-search multiple allow-clear>
                              <a-option v-for="item in applist" :value="item">{{ item }}</a-option>
                          </a-select>
                        </a-form-item>
                        <a-form-item field="utilsfiles" label="utils扩展文件" style="margin-bottom: 10px;">
                          <a-select v-model="formData.utilsfiles" show-header-on-empty placeholder="选择后端utils/plugin工具函数" allow-search multiple allow-clear>
                              <a-option v-for="item in utilstool" :value="item">{{ item }}</a-option>
                          </a-select>
                        </a-form-item>
                       </div>
                     </div>
                  </div>
                  <div class="borderbox" style="margin-top: 10px;">
                     <span class="title">前端代码-B端</span>
                     <div class="borderItem">
                       <div class="tabheard">views模块代码：</div>
                       <div class="formValue">
                        <a-form-item field="viewsfolders" label="views目录" style="margin-bottom: 10px;">
                          <a-select v-model="formData.viewsfolders_business" show-header-on-empty placeholder="选择前端views文件夹的目录" allow-search multiple allow-clear>
                              <a-option v-for="item in vue_business.viewsfolders" :value="item">{{ item }}</a-option>
                          </a-select>
                        </a-form-item>
                        <a-form-item field="viewsfiles" label="views文件" style="margin-bottom: 10px;">
                          <a-select v-model="formData.viewsfiles_business" show-header-on-empty placeholder="选择前端views文件夹下的文件" allow-search multiple allow-clear>
                              <a-option v-for="item in vue_business.viewsfiles" :value="item">{{ item }}</a-option>
                          </a-select>
                        </a-form-item>
                       </div>
                     </div>
                     <div class="borderItem">
                       <div class="tabheard">components组件代码：</div>
                       <div class="formValue">
                        <a-form-item field="componentfolders" label="components目录" style="margin-bottom: 10px;">
                          <a-select v-model="formData.componentfolders_business" show-header-on-empty placeholder="选择前端components组件文件夹的目录" allow-search multiple allow-clear>
                              <a-option v-for="item in vue_business.componentfolders" :value="item">{{ item }}</a-option>
                          </a-select>
                        </a-form-item>
                        <a-form-item field="componentfiles" label="components文件" style="margin-bottom: 10px;">
                          <a-select v-model="formData.componentfiles_business" show-header-on-empty placeholder="选择前端components组件文件夹下的文件" allow-search multiple allow-clear>
                              <a-option v-for="item in vue_business.componentfiles" :value="item">{{ item }}</a-option>
                          </a-select>
                        </a-form-item>
                       </div>
                     </div>
                 </div>
                  <div class="borderbox" style="margin-top: 10px;">
                     <span class="title">前端代码-A端</span>
                     <div class="borderItem">
                       <div class="tabheard">views模块代码：</div>
                       <div class="formValue">
                        <a-form-item field="viewsfolders" label="views目录" style="margin-bottom: 10px;">
                          <a-select v-model="formData.viewsfolders_admin" show-header-on-empty placeholder="选择前端views文件夹的目录" allow-search multiple allow-clear>
                              <a-option v-for="item in vue_admin.viewsfolders" :value="item">{{ item }}</a-option>
                          </a-select>
                        </a-form-item>
                        <a-form-item field="viewsfiles" label="views文件" style="margin-bottom: 10px;">
                          <a-select v-model="formData.viewsfiles_admin" show-header-on-empty placeholder="选择前端views文件夹下的文件" allow-search multiple allow-clear>
                              <a-option v-for="item in vue_admin.viewsfiles" :value="item">{{ item }}</a-option>
                          </a-select>
                        </a-form-item>
                       </div>
                     </div>
                     <div class="borderItem">
                       <div class="tabheard">components组件代码：</div>
                       <div class="formValue">
                        <a-form-item field="componentfolders" label="components目录" style="margin-bottom: 10px;">
                          <a-select v-model="formData.componentfolders_admin" show-header-on-empty placeholder="选择前端components组件文件夹的目录" allow-search multiple allow-clear>
                              <a-option v-for="item in vue_admin.componentfolders" :value="item">{{ item }}</a-option>
                          </a-select>
                        </a-form-item>
                        <a-form-item field="componentfiles" label="components文件" style="margin-bottom: 10px;">
                          <a-select v-model="formData.componentfiles_admin" show-header-on-empty placeholder="选择前端components组件文件夹下的文件" allow-search multiple allow-clear>
                              <a-option v-for="item in vue_admin.componentfiles" :value="item">{{ item }}</a-option>
                          </a-select>
                        </a-form-item>
                       </div>
                     </div>
                 </div>
                </a-col>
              </a-row>
            </a-form>
          </div>
        </div>
        <div class="tabitem" v-if="currentindex==2">
          <div class="formbox">
            <a-form ref="formRef" :model="formData" auto-label-width >
              <a-row :gutter="16">
                  <a-col :span="12">
                    <a-form-item field="name" label="包目录名" :rules="{required:true,message:'请填写包目录名'}" >
                        <a-input v-model="formData.name" placeholder="请填包目录名（以字母开头、数字命名）" @input="filterName" allow-clear/>
                    </a-form-item>
                    <a-form-item field="title" label="名称" :rules="[{required:true,message:'请填写包名称'}]" >
                        <a-input v-model="formData.title" placeholder="请填包名称（标题）" allow-clear/>
                    </a-form-item>
                    <a-row :gutter="16">
                    <a-col :span="14"> 
                      <a-form-item field="version" label="版本号" :rules="[{required:true,message:'请填写版本号'}]" >
                          <a-input v-model="formData.version" placeholder="请填版本号" allow-clear/>
                      </a-form-item>
                    </a-col>
                    <a-col :span="10"> 
                      <a-form-item field="price" label="价格" :label-col-style="{flex:'auto'}">
                          <a-input-number v-model="formData.price" :min="0" placeholder="请填价格" allow-clear/>
                      </a-form-item>
                    </a-col>
                    </a-row>
                    <a-form-item field="cid" label="选择分类" validate-trigger="input" :rules="[{required:true,message:'请选择分类'}]">
                      <a-select v-model="formData.cid" :options="cateList" allow-search :field-names="{value: 'id', label: 'name'}" placeholder="请选择分类" allow-clear/>
                    </a-form-item>
                  <a-form-item field="des" label="描述" style="margin-bottom: 10px;">
                      <a-textarea v-model="formData.des" placeholder="请填说明" :auto-size="{minRows:3,maxRows:3}" allow-clear/>
                  </a-form-item>
                  </a-col>
                  <a-col :span="12">
                    <div class="borderbox" style="margin-top: 10px;">
                      <span class="title">选择代码上传</span>
                      <div class="borderItem upfilebox" style="margin-top: 10px;">
                           <a-upload draggable action="/" :show-file-list="false" accept=".zip" :custom-request="customRequest"/>
                           <div class="showfile" v-if="formData.fileurl">
                              <div class="icon"><Icon icon="svgfont-yasuowenjian" class="tbicon" :size="14" ></Icon></div>
                              <div class="name">{{ upfileName }}</div>
                              <div class="url"><a :href="formData.fileurl" target="_blank">附件加载地址</a></div>
                           </div>
                      </div>
                    </div>
                    <div class="dirtig">
                      复制打包代码位置：
                      <a-link @click="copyText(codepack)" style=" word-break: break-all;">{{ codepack }}</a-link>
                    </div>
                  </a-col>
                  <a-col :span="24" v-if="tabwarehouse=='public'">
                    <a-alert title="上传公共仓注意事项">
                      1.上传代码仓后到个人中心（<a href="https://goflys.cn/" target="_blank">gofly快速开发社区</a>）的代码仓插件管理编辑插件的介绍详情和上传封面图。<br>
                      2.如果有开发文档再到点到查看插件详情的右上角编辑开发文档。
                    </a-alert>
                  </a-col>
                </a-row>
              </a-form>
           </div>
        </div>
       </div>
    </div>
  </BasicModal>
  <!--登录-->
  <Login ref="loginRef" @ok="loginOk"></Login>
</template>
<script lang="ts">
  import { defineComponent, ref, computed} from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import useLoading from '@/hooks/loading';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  //api
  import { packCode,getPackdirs,getMenutreeData,menuTreeToJson,userUploadFile } from '@/api/developer/codestoreoption';
  import { getTables,TableItem } from '@/api/common';
  import { upPackToService,getCodeCate } from '@/api/developer/codestore';
  import { Message ,TreeNodeData} from '@arco-design/web-vue';
  import { cloneDeep } from 'lodash';
  import { CodeEditor, MODE } from '/@/components/CodeEditor';
  import type { RequestOption} from '@arco-design/web-vue/es/upload/interfaces';
  import {Icon} from '@/components/Icon';
  import { useClipboard } from '@vueuse/core';
  import Login from "./Login.vue"; 
  export default defineComponent({
    name: 'AddBook',
    components: { BasicModal,CodeEditor,Icon,Login },
    emits: ['success'],
    setup(_, { emit }) {
      const loginRef=ref()
      const BusinessMenutreeData = ref<TreeNodeData[]>([]);
      const AdminMenutreeData = ref<TreeNodeData[]>([]);
      const tablelist = ref<TableItem[]>([]);
      const applist = ref<string[]>([]);
      const appDirlist = ref<string[]>([]);
      const utilstool = ref<string[]>([]);
      const vue_business = ref<any>({componentfiles:[],componentfolders:[],viewsfiles:[],viewsfolders:[]});
      const vue_admin = ref<any>({componentfiles:[],componentfolders:[],viewsfiles:[],viewsfolders:[]});
      const loadingMenu = ref(true);
      const currentindex = ref(1);
      const bmenuboxbar = ref("selectemenu");
      const amenuboxbar = ref("selectemenu");
      const upfileName=ref("");
      const { copy } = useClipboard();
      //表单内容
      const formData = ref({
        name:"",
        title:"",
        version:"1.0.0",
        price:10,
        cid: 1,
        des:"",
        installcover:false,
        noVerifyAPIRoot:[],
        //B端
        viewsfiles_business:[],
        viewsfolders_business:[],
        componentfiles_business:[],
        componentfolders_business:[],
        //B端
        viewsfiles_admin:[],
        viewsfolders_admin:[],
        componentfiles_admin:[],
        componentfolders_admin:[],
        //后端
        appfolders:[],
        utilsfiles:[],
        //数据库
        packtables:[],
        businessmenujson:[],//B菜单
        adminmenujson:[],//A菜单
        businessmenutree:[],//B菜单选中
        adminmenutree:[],//A菜单选中
        fileurl:"",
      });
      const domian=ref("")
      const code_token=ref("")
      const cateList = ref<TreeNodeData[]>([]);
      const tabwarehouse=ref("")
      const codepack=ref("")
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          setModalProps({ confirmLoading: false })
            domian.value=data.record.domian
            code_token.value=data.record.code_token
            tabwarehouse.value=data.record.tabwarehouse
            //获取数据表
            tablelist.value = await getTables({});
            const packdir= await getPackdirs({});
            const MData= await getMenutreeData({});
            if(MData){
              BusinessMenutreeData.value= MData.business_menuList;
              AdminMenutreeData.value= MData.admin_menuList;
            }
            loadingMenu.value=false
            const catedataArr=await getCodeCate({baseurl:domian.value})
            cateList.value = catedataArr.catedata
            codepack.value = catedataArr.codepack
            if(packdir){
              applist.value =packdir.applist
              appDirlist.value =packdir.appDirlist
              utilstool.value =packdir.utilstool
              //前端
              vue_business.value =packdir.vue_business
              vue_admin.value =packdir.vue_admin
            }
      });
      const getTitle = computed(() => ( `打包上传代码到${tabwarehouse.value=="public"?"公共":"私有"}仓库` ));
      const okText = computed(() => ( currentindex.value==1?"打包代码":"上传代码到仓库"));
     //点击安装
     const { loading, setLoading } = useLoading();
     const formRef = ref<FormInstance>();
     const handleSubmit =async () => {
      try {
          const res = await formRef.value?.validate();
          if (!res) {
              if(currentindex.value==1){
                setLoading(true);
                Message.loading({content:"打包中...",id:"PackCode",duration:0})
                var savedata=cloneDeep(formData.value)
                if(savedata.packtables){
                  savedata=Object.assign({},savedata,{packtables:savedata.packtables.join()})
                }
                if(savedata.noVerifyAPIRoot&&savedata.noVerifyAPIRoot.length>0){
                  savedata=Object.assign({},savedata,{noVerifyAPIRoot:savedata.noVerifyAPIRoot.join()})
                }else{
                  savedata=Object.assign({},savedata,{noVerifyAPIRoot:""})
                }
                const res= await packCode(savedata);
                Message.success({content:"打包成功",id:"PackCode",duration:2000})
                closeModal()
                emit('success');
                setLoading(false);
              }else{//提交到代码仓
                if(!formData.value.fileurl){
                    Message.warning({content:"没有找到代码包",id:"PackCode",duration:2000})         
                }else{
                    if(tabwarehouse.value=="public"&&!code_token.value){//登录窗口
                      Message.warning({content:"上传代码到公共仓，请先登录",id:"PackCode",duration:2000})     
                      loginRef.value.showLogin(domian.value)
                      return;
                    }
                    setLoading(true);
                    Message.loading({content:"提交代码包到仓库中...",id:"PackCode",duration:0})
                    var savedata=cloneDeep(formData.value)
                    if(savedata.packtables){
                      savedata=Object.assign({},savedata,{packtables:savedata.packtables.join(),code_token:code_token.value})
                    }
                    await upPackToService(savedata,domian.value);
                    Message.success({content:"提交代码包到仓库成功",id:"PackCode",duration:2000})
                    closeModal()
                    emit('success');
                    setLoading(false);
                }
              }
          }
        } catch (error) {
          setLoading(false);
          Message.loading({content:currentindex.value==1?"打包中...":"提交代码包到仓库中...",id:"PackCode",duration:1})
        }
    }
      const modeValue = ref<MODE>(MODE.JSON);
      //菜单选择变化时同时修改菜单json
      const handleMarkeMenuBusiness=async()=>{
        const res= await menuTreeToJson({menu:formData.value.businessmenutree,formtable:"business"});
        if(res){
          formData.value.businessmenujson=res
        }else{
          formData.value.businessmenujson=[]
        }
      }
      const handleMarkeMenuAdmin=async()=>{
        const res= await menuTreeToJson({menu:formData.value.adminmenutree,formtable:"admin"});
        if(res){
          formData.value.adminmenujson=res
        }else{
          formData.value.adminmenujson=[]
        }
      }
      //上传附件
      const customRequest = (options: RequestOption) => {
          const controller = new AbortController();
            (async function requestWrap() {
              const {
                onProgress,
                onError,
                onSuccess,
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
                const filename=fileItem?.name||""
                const resdata = await userUploadFile({ name: 'file', file: fileItem.file as Blob, filename,data:{cid:0,domainurl:domian.value+"/common/upload/fileNov"}},onUploadProgress);
                //更新图片
                if(resdata){
                  upfileName.value=filename
                  formData.value.fileurl=resdata.url
                }
              } catch (error) {
                onError(error);
              }
            })();
            return {
              abort() {
                controller.abort();
              },
            };
      };
      //复制文本
      const copyText=async(text:any)=>{
        await copy(text);
         Message.success("复制成功");
      }
      //过滤报包名汉字
      const filterName=async(strValue:any)=>{
        var reg = /[\u4e00-\u9fa5]/g;   
        formData.value.name=strValue.replace(reg, "");   
      }
      //登录成功
      const loginOk=(val:any)=>{
        code_token.value=val
      }
      //包名填好后自动选择与包名相关数据表
      const handlePackNameChange=(val:any)=>{
        //1.选择数据库表
        if(formData.value.packtables.length==0){
          const gettablelist= tablelist.value.filter((item:any) => {
              return item.name.indexOf(val) > -1;
          });
          gettablelist.forEach((item)=>{
            formData.value.packtables.push(item.name)
          })
        }
        //2.选择app下代码
        formData.value.appfolders=[]
        const getapplist= applist.value.filter((item:any) => {
            return item.indexOf(val) > -1;
        });
        getapplist.forEach((item)=>{
          formData.value.appfolders.push(item)
        })
        //3.选择utils扩展文件代码
        formData.value.utilsfiles=[]
        const getutilstool= utilstool.value.filter((item:any) => {
            return item.indexOf(val) > -1;
        });
        getutilstool.forEach((item)=>{
          formData.value.utilsfiles.push(item)
        })
        //4.选择views下代码
        formData.value.viewsfolders_business=[]
        const getbviewsfolders= vue_business.value.viewsfolders.filter((item:any) => {
            return item.indexOf(val) > -1;
        });
        getbviewsfolders.forEach((item)=>{
          formData.value.viewsfolders_business.push(item)
        })
        //5.选择views下代码
        formData.value.viewsfolders_admin=[]
        const getAviewsfolders= vue_admin.value.viewsfolders.filter((item:any) => {
            return item.indexOf(val) > -1;
        });
        getAviewsfolders.forEach((item)=>{
          formData.value.viewsfolders_admin.push(item)
        })
      }
      return { 
        registerModal, 
        getTitle, 
        loading,
        currentindex,
        formData,
        tablelist,applist,utilstool,appDirlist,
        vue_business,vue_admin,
        handleSubmit,
        modeValue,
        bmenuboxbar,amenuboxbar,
        BusinessMenutreeData,AdminMenutreeData,
        handleMarkeMenuBusiness,handleMarkeMenuAdmin,
        loadingMenu,
        customRequest,
        upfileName,
        cateList,
        codepack,
        copyText,
        okText,
        formRef,
        filterName,
        loginOk,loginRef,
        tabwarehouse,
        handlePackNameChange,
      };
    },
  });
</script>
<style lang="less" scoped>
  @import '@/assets/style/formlayer.less';
  
  .navtabar{
     width: 100%;
     :deep(.arco-steps-item){
      height: 57px;
      &:not(:last-child)::after{
        border-top: 29px solid transparent;
        border-bottom: 28px solid transparent;
      }
     }
     :deep(.arco-steps-item){
      height: 57px;
      &:not(:first-child)::before{
        border-top: 29px solid transparent;
       border-bottom: 26px solid transparent;
      }
     }
  }
  .tabcontent{
    padding: 15px;
  }
  .borderbox{
    width: 100%;
    border: var(--color-neutral-4) solid 1px;
    padding: 10px;
    border-radius: 5px;
    position: relative;
    .title{
      background-color: var(--color-bg-2);
      padding: 0px 10px;
      position: absolute;
      top: 0%;
      left: 50%;
      transform: translate(-50%,-50%);
    }
    .borderItem{
      width: 100%;
      .tabheard{
        font-weight: 700;
        font-size: var(size-4);
        margin-bottom: 8px;
      }
    }
  }
  .savebtn{
    text-align: center;
    margin-top: 15px;
  }
  //菜单
  .menubox{
    width: 100%;
    .parambox{
        width: 100%;
        border: #e5e7eb 1px solid;
        height: 130px;
        overflow: auto;
      }
  }
 :deep(.arco-tree-select-popup){
  text-align: center;
 }
 .upfilebox{
   :deep(.arco-upload-drag){
    height: 148px;
   }
 }
 .showfile{
  display: flex;
  margin-top: 5px;
  background-color: rgb(var(--arcoblue-1));
  padding: 10px;
  border-radius: 3px;
  .icon{
      padding-right: 10px;  
  }
  .name{
    flex: 1;
  }
 }
</style>

