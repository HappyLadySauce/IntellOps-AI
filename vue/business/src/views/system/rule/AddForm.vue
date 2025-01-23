<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :loading="loading" helpMessage="这里是动态添加后台系统菜单、路由，访问的路由名称，路由指向的vue代码路径。" width="980px" :minHeight="500" :title="getTitle" @ok="handleSubmit">
    <a-form ref="formRef" :model="formData" auto-label-width>
     <a-form-item field="type" label="菜单类型" style="margin-bottom:15px;">
        <a-radio-group type="button" v-model="formData.type" @change="handleChangeType">
          <a-radio :value="0">目录</a-radio>
          <a-radio :value="1">菜单</a-radio>
          <a-radio :value="2">按钮/权限</a-radio>
        </a-radio-group>
      </a-form-item>
      <a-row :gutter="16">
        <a-col :span="12">
          <a-form-item field="title" :label="formData.type==2?'权限名称':'菜单名称'" validate-trigger="input" :rules="[{required:true,message:formData.type==2?'请填写菜单名称':'请填写权限名称'}]" style="margin-bottom:15px;">
            <a-input-group style="width: 100%;">
              <a-input v-model="formData.title" :placeholder="formData.type==2?'输入权限名称':'输入写菜单名称'" allow-clear/>
              <a-select v-if="formData.type==2" @change="handleShortcutAuth" :options="shortcutAuthList" allow-clear :style="{width:'90px'}" placeholder="快捷" />
            </a-input-group>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="locale" label="多语言标识" validate-trigger="input" style="margin-bottom:15px;">
            <a-input v-model="formData.locale" placeholder="输入写菜单名称(i18n的标识)" allow-clear/>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="上级菜单" field="pid" style="margin-bottom:15px;" :rules="[{required:formData.type==2?true:false,message:'请选择上级菜单'}]">
            <a-tree-select placeholder="选择上级菜单" :data="parntList" 
             :fieldNames="{
                key: 'id',
                title: 'title',
                children: 'children'
              }"
              v-model="formData.pid" @change="getRouteList">
            </a-tree-select>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="排序" field="weigh" style="margin-bottom:15px;">
            <a-input-number v-model="formData.weigh" placeholder="请填排序" />
          </a-form-item>
        </a-col>
        <a-col :span="12"  v-if="formData.type!=2">
          <a-form-item label="图标" field="icon" style="margin-bottom:15px;">
            <a-input-search v-model="formData.icon" placeholder="选择图标/填写"  search-button allow-clear>
              <template  v-if="formData.icon" #prefix>
                <Icon :icon="formData.icon"></Icon>
              </template>
                <template #button-icon>
                  <a-popover position="bl" trigger="click">
                    <icon-apps :size="23"/>
                    <template #content>
                       <IconPicker v-model="formData.icon"></IconPicker>
                    </template>
                  </a-popover>
                </template>
              </a-input-search>
          </a-form-item>
        </a-col>
        <a-col :span="12"  v-if="formData.type!=2">
          <a-form-item field="routename" label="路由名称" validate-trigger="input" :rules="[{required:true,message:'请填写路由名称'},{match:/^[A-Za-z0-9]+$/,message:'路由必须由字母组成'},{validator:handleChekHaseRule}]" >
            <a-input v-model="formData.routename" placeholder="前端路由名称（name）" @blur="handleRoutename" allow-clear/>
          </a-form-item>
        </a-col>
        <a-col :span="12"  v-if="formData.type!=2">
          <a-form-item field="routepath" label="路由地址" validate-trigger="input" :rules="[{required:true,message:'请填写路由地址'}]" >
            <a-input v-model="formData.routepath" placeholder="前端路由地址（path）建议与路由名称一致" allow-clear/>
          </a-form-item>
        </a-col>
        <a-col :span="12" v-if="formData.type==0">
          <a-form-item field="redirect" label="重定向地址" validate-trigger="input" :rules="[{required:true,message:'请填写重定向地址'}]">
            <a-input v-model="formData.redirect" placeholder="前端重定向地址（redirect）" allow-clear/>
          </a-form-item>
        </a-col>
        <a-col :span="12" v-if="formData.type==1">
          <a-form-item field="component" label="组件路径" validate-trigger="input" :rules="[{required:true,message:'请填写组件路径'}]">
            <a-input v-model="formData.component" placeholder="vue前端组件路径（component）" allow-clear/>
          </a-form-item>
        </a-col>
        <a-col :span="12" v-if="formData.type==2">
          <a-form-item field="path" label="接口路径" validate-trigger="input" :rules="[{required:formData.permission?false:true,message:'请填写数据请求地址'}]">
            <a-select v-model="formData.path" placeholder="输入后端api请求地址（用于接口权限验证）" allow-search allow-clear allow-create>
                <a-option v-for="item in routerlist" :value="item.path">{{ item.path }}</a-option>
            </a-select>
            <template #help>
              后端接口请求路径，用于接口请求身份验证(即RBAC权限)
              <a-tooltip content="刷新路由选项数据">
                <icon-refresh @click="getRouteList" :spin="getRouteLoading" style="margin-left: 5px;color: rgb(var(--arcoblue-6));" class="iconbtnc"/>
              </a-tooltip>
            </template>
          </a-form-item>
        </a-col>
        <a-col :span="12" v-if="formData.type==2">
          <a-form-item field="permission" label="权限标识" validate-trigger="input" :rules="[{required:formData.path?false:true,message:'请填写权限标识'}]">
            <a-input v-model="formData.permission" :placeholder="ppermission" allow-clear/>
            <template #help>
              用于前端页面按钮显示或隐藏控制
              <a-tooltip :content="ppermissiontigcode" mini>
                <icon-info-circle />
              </a-tooltip>
              <a-tooltip content="复制标识代码粘贴到前端节点中" v-if="formData.permission" mini>
                <icon-copy @click="handleCopyAuth" style="margin-left: 8px;color: rgb(var(--arcoblue-6));" class="iconbtnc"/>
              </a-tooltip>
            </template>
          </a-form-item>
        </a-col>
        <a-col :span="12" v-if="formData.type==2">
          <a-form-item field="des" label="描述" >
            <a-textarea v-model="formData.des" placeholder="输入权限描述" allow-clear/>
          </a-form-item>
        </a-col>
        <a-col :span="12" v-if="formData.type!=2">
          <a-form-item field="isext" label="是否外链" style="margin-bottom:5px;">
              <a-switch v-model="formData.isext" :checked-value="1" :unchecked-value="0" checked-text="是" unchecked-text="否" />
              <span class="formrighttig">在浏览器新标签打开</span>
          </a-form-item>
        </a-col>
        <a-col :span="12"  v-if="formData.type!=2">
          <a-form-item field="keepalive" label="是否缓存" style="margin-bottom:5px;">
            <a-switch v-model="formData.keepalive" :checked-value="1" :unchecked-value="0" checked-text="是" unchecked-text="否" />
            <span class="formrighttig">在页签模式生效,页面name和路由name保持一致</span>
          </a-form-item>
        </a-col>
        <a-col :span="12"  v-if="formData.type!=2">
          <a-form-item field="hideinmenu" label="菜单隐藏" style="margin-bottom:5px;" >
            <a-switch v-model="formData.hideinmenu" :checked-value="1" :unchecked-value="0" checked-text="是" unchecked-text="否" />
            <span class="formrighttig">是否在左侧菜单中隐藏该项</span>
          </a-form-item>
        </a-col>
        <a-col :span="12"  v-if="formData.type!=2">
          <a-form-item field="hidechildreninmenu" label="显示单项" style="margin-bottom:5px;" >
            <a-switch v-model="formData.hidechildreninmenu" :checked-value="1" :unchecked-value="0" checked-text="是" unchecked-text="否" />
            <span class="formrighttig">hideChildrenInMenu强制在左侧菜单中显示单项</span>
          </a-form-item>
        </a-col>
        <a-col :span="12"  v-if="formData.type!=2">
          <a-form-item field="noaffix" label="添加tab中" style="margin-bottom:5px;">
            <a-switch v-model="formData.noaffix" :checked-value="1" :unchecked-value="0" checked-text="是" unchecked-text="否" />
            <span class="formrighttig">如果设置为true标签将不会添加到tab-bar中</span>
          </a-form-item>
        </a-col>
        <a-col :span="12"  v-if="formData.type!=2">
          <a-form-item field="activemenu" label="高亮菜单" style="margin-bottom:5px;">
            <a-switch v-model="formData.activemenu" :checked-value="1" :unchecked-value="0" checked-text="是" unchecked-text="否" />
            <span class="formrighttig">高亮设置的菜单项</span>
          </a-form-item>
        </a-col>
        <a-col :span="12"  v-if="formData.type==1">
          <a-form-item field="onlypage" label="独立页面" style="margin-bottom:5px;">
            <a-switch v-model="formData.onlypage" :checked-value="1" :unchecked-value="0" checked-text="是" unchecked-text="否" />
            <span class="formrighttig">不需layout和登录，如登录页、数据大屏</span>
          </a-form-item>
        </a-col>
        <a-col :span="12"  v-if="formData.type!=2">
          <a-form-item field="requiresauth" label="是否登录鉴权" style="margin-bottom:5px;">
            <a-switch v-model="formData.requiresauth" :checked-value="1" :unchecked-value="0" checked-text="是" unchecked-text="否" />
            <span class="formrighttig">前端是否需要登录鉴权</span>
          </a-form-item>
        </a-col>
      </a-row>
    </a-form>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref} from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import { useI18n } from 'vue-i18n';
  import { cloneDeep } from 'lodash-es';
  //api
  import { getContent,getParent,getRoutes,save,RuleItem } from '@/api/system/rule';
  import {  checkHaseRule } from '@/api/common';
  import { IconPicker ,Icon} from '@/components/Icon';
  import { Message } from '@arco-design/web-vue';
  import { CopyText } from "@/utils/tool";
  export default defineComponent({
    name: 'AddBook',
    components: { BasicModal,IconPicker,Icon },
    emits: ['success'],
    setup(_, { emit }) {
      const { t } = useI18n();
      const isUpdate = ref(false);
      const parntList = ref<RuleItem[]>([]);
      const parntMdList = ref<RuleItem[]>([]);
      const routerlist = ref<any[]>([]);
      const shortcutAuthList = ref<any[]>([{value:'getList',label:'查看'},{value:'save',label:'添加/编辑'},{value:'del',label:'删除'},{value:'upStatus',label:'状态'},{value:'getContent',label:'详情'}]);
      const getRouteLoading = ref<boolean>(false);
        var mdpid=0
      //表单
      const formRef = ref<FormInstance>();
      const ppermission = ref<string>(`权限标识(前端dom中添加v-permission="'add'")`);
      const ppermissiontigcode = ref<string>(`<a-button v-permission="'add'">使用示例</a-button>`);
      //表单字段
      const basedata={
            id:0,
            title: '',
            des: '',
            locale: '',
            weigh: 1,
            type: 0,
            pid:0,
            icon:"",
            routepath:"",//path
            routename:"",//name
            component:"",//外部layar、import
            redirect:"",//重定向
            path:"",//权限-接口路径
            permission:"",//权限标识
            isext:0,//是否外链
            keepalive:0,//是否缓存 0=否1=是
            requiresauth:1,//是否需要登录鉴权 默认false
            hideinmenu:0,//是否在左侧菜单中隐藏该项
            hidechildreninmenu:0,//强制在左侧菜单中显示单项
            activemenu:0,//高亮设置的菜单项
            noaffix:0,//如果设置为true，标签将不会添加到tab-bar中
            onlypage:0,//独立页面，如登录页、数据大屏
        }
      const formData = ref(basedata)
      const m_component=ref("")
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          mdpid=0
          formRef.value?.resetFields()
          setModalProps({ confirmLoading: false });
          isUpdate.value = !!data?.isUpdate;
          const resultdata = await getParent({id:unref(isUpdate)?data.record.id:0});
          const parntList_df : any=[{id: 0,title: "一级菜单",pid: 0,locale:""}];
          if(resultdata&&resultdata.tree){
              parntList.value=parntList_df.concat(resultdata.tree)
              parntMdList.value=resultdata.list
          }else{
            parntList.value=parntList_df
          }
          getRouteList()
          if (unref(isUpdate)) {
            m_component.value=data.record.component
            formData.value=cloneDeep(data.record)
            setLoading(true);
            const mewdata = await getContent({id:data.record.id});
            formData.value=Object.assign({},formData.value,mewdata)
            setLoading(false);
          }else{
            m_component.value=""
            formData.value=cloneDeep(basedata)
            //判断添加子集
            if(data.record){
              formData.value.pid=data.record.id
              formData.value.type=data.record.type+1
            }
          }
      });
      //获取权限路由数据
      const getRouteList=async()=>{
        getRouteLoading.value=true
        const routeslist= await getRoutes({});
        const ruledata=parntMdList.value.find((item)=>item.id == formData.value.pid)
        if(ruledata){
          var routeStr=ruledata.routepath
          if(ruledata.pid!=0){
            const fdata=parntMdList.value.find((item)=>item.id == ruledata.pid)
            if(fdata){
              if(routeStr.indexOf("/")>-1){
                routeStr=fdata.routepath+routeStr
              }else{
                routeStr=fdata.routepath+"/"+routeStr
              }
            }
          }
          routerlist.value = routeslist.filter((item:any) => {
              return item.path.indexOf(routeStr) > -1;
          });
          if(routerlist.value.length==0){
            routerlist.value = routeslist.filter((item:any) => {
                return item.path.indexOf(ruledata.routepath) > -1;
            });
          }
          //根据菜单找不到直接给全部
          if(routerlist.value.length==0){
            routerlist.value=routeslist
          }
        }else{
           routerlist.value=routeslist
        }
        setTimeout(()=>{
          getRouteLoading.value=false
        },500)
      }
      const getTitle = computed(() => (!unref(isUpdate) ? '新增系统菜单' : '编辑系统菜单'));
     //点击确认
     const { loading, setLoading } = useLoading();
     const handleSubmit = async () => {
      try {
          const res = await formRef.value?.validate();
          if (!res) {
            setLoading(true);
            if(formData.value.type==0&&formData.value.component==""){
              formData.value.component='LAYOUT'
            }
            if(formData.value.type==0&&formData.value.pid==0&&formData.value.routepath&&formData.value.routepath.substr(0,1)!="/"){
              formData.value.routepath='/'+formData.value.routepath
            }
            Message.loading({content:"更新中",id:"upStatus",duration:0})
            await save(unref(formData));
            Message.success({content:"更新成功",id:"upStatus",duration:2000})
            closeModal()
            emit('success');
            setLoading(false);
          }
        } catch (error) {
          setLoading(false);
          Message.loading({content:"更新中",id:"upStatus",duration:1})
        }
      };
      //切换菜单类型
      const handleChangeType=(value:any)=>{
         if(value==0){
          formData.value.component="LAYOUT"
         }else if(value==1){
          formData.value.component=m_component.value
          formData.value.redirect=""
         }else if(value==2){
          formData.value.component=""
          formData.value.redirect=""
        }
        if(value==2){
            parntList.value=parntList.value.filter((item)=>item.id!=0)
            if(mdpid==0)
             mdpid=parntList.value[0].id
            if(formData.value.pid==0){
              formData.value.pid=mdpid
            }else{
              mdpid=formData.value.pid
            }
        }else{
          parntList.value=parntList.value.filter((item)=>item.id!=0)
          const parntList_df : any=[{id: 0,title: "一级菜单",pid: 0,locale:""}];
          parntList.value=parntList_df.concat(parntList.value)
          if(value==1){
            if(mdpid==0)
             mdpid=parntList.value[0].id
            if(formData.value.pid==0){
              formData.value.pid=mdpid
            }else{
              mdpid=formData.value.pid
            }
          }else{
            formData.value.pid=0
          }
        }
      }
      //检查是否存在路由
      const handleChekHaseRule=(value:any, cb:any)=>{
        return new Promise(async(resolve) => {
            if(value){
              const resData = await checkHaseRule({id:formData.value.id,routename:value,codelocation:"busDirName"});
              if(!resData){
                cb("该路由已被占用请换一个")
              }
            }
            resolve(true)
          })
      }
      //填写路由名时填写path
      const handleRoutename=()=>{
          if(formData.value.routename){
            formData.value.routepath="/"+formData.value.routename
          }
      }
      //权限快捷选择输入
      const quickData :any={"getList":"view","save":"add","del":"del","upStatus":"status","getContent":"details"}
      const handleShortcutAuth=(value:any)=>{
        formData.value.title=shortcutAuthList.value.find((item)=>item.value==value).label
        formData.value.permission=quickData[value]
        const routesdata=routerlist.value.find((item:any) => {
            return item.path.indexOf(value) > -1;
        });
        if(routesdata){
          formData.value.path=routesdata.path
        }
      }
      //复制前端权限识别代码
      const handleCopyAuth=()=>{
         CopyText(`v-permission="'${formData.value.permission}'"`,t('sys.copy.success'))
      }
     
      return { 
        registerModal, 
        getTitle, 
        handleSubmit,
        formRef,
        loading,
        formData,
        parntList,
        handleChangeType,handleChekHaseRule,handleRoutename,
        ppermission,routerlist,getRouteLoading,getRouteList,
        handleShortcutAuth,shortcutAuthList,handleCopyAuth,ppermissiontigcode,
      };
    },
  });
</script>
<style lang="less" scoped>
.formrighttig{
  padding-left: 10px;
  font-size: 12px;
  color: var(--color-neutral-4);
}
</style>
