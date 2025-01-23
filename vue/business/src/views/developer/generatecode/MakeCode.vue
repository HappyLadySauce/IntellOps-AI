<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" :loading="loading" width="900px" :minHeight="modelHeight" :title="getTitle" @ok="handleSubmit">
      <div class="table-content">
        <a-form ref="formRef" :model="formData" auto-label-width>
          <a-row :gutter="16">
            <a-col :span="12">
              <a-form-item label="代码生成位置" field="codelocation" style="margin-bottom:15px;">
                  <a-radio-group v-model="formData.codelocation" @change="handleGetMenuList">
                      <a-radio value="busDirName">business(B端)</a-radio>
                      <a-radio value="adminDirName" :disabled="!haseadmin">admin(A端)</a-radio>
                  </a-radio-group>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item field="codetpl" label="代码模板" :rules="[{required:true,message:'请选择生成代码模板'}]" >
                <a-select v-model="formData.codetpl" placeholder="选择生成代码模板" @change="handdleCodetpl">
                  <a-option value="web">仅前端代码</a-option>
                  <a-option value="go">仅后端代码</a-option>
                  <a-option value="goweb">后端和前端代码</a-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="12" v-if="formData.codetpl!='go'">
              <a-form-item field="rule_name" label="菜单名称" validate-trigger="input" extra="请求后台页面router路由">
                <a-input v-model="formData.rule_name" placeholder="填生成的菜单名称（不需要菜单就不创建）" :max-length="50" allow-clear show-word-limit />
              </a-form-item>
            </a-col>
            <a-col :span="12" v-if="formData.rule_name">
              <a-form-item field="pid" label="左侧菜单" validate-trigger="input" extra="生成系统菜单即在左边路由中，若菜单有特殊可在系统设置">
                <a-tree-select placeholder="选择上级菜单" :data="parntList" @change="handleChangeMenu"
                  :fieldNames="{ key: 'id',title: 'title',children: 'children'}" v-model="formData.pid">
                  </a-tree-select>
              </a-form-item>
            </a-col>
            <a-col :span="12" v-if="formData.rule_name">
              <a-form-item field="routename" label="路由名" :validate-trigger="['input','change', 'blur']" :rules="[{required:true,message:'请填请求路由名称'},{match:/^[A-Za-z0-9]+$/,message:'路由必须由字母数字组成'},{validator:handleChekHaseRule}]" extra="vue-router路由routename的名称">
                <a-input v-model="formData.routename" placeholder="填生成的前端请求路由名称" :max-length="50" allow-clear show-word-limit @blur="handleChangeRoutename"/>
              </a-form-item>
            </a-col>
            <a-col :span="12"  v-if="formData.rule_name">
              <a-form-item label="菜单图标" field="icon" style="margin-bottom:15px;" extra="选择器没有找到图标可以去arco和自己添加的阿里icon复制">
                  <a-input v-model="formData.icon" placeholder="选择图标/填写"  >
                      <template  v-if="formData.icon" #prefix>
                        <Icon :icon="formData.icon"></Icon>
                      </template>
                  </a-input>
                  <a-popover position="br" trigger="click">
                    <a-button type="primary">选择</a-button>
                    <template #content>
                      <IconPicker v-model="formData.icon"></IconPicker>
                    </template>
                  </a-popover>
              </a-form-item>
            </a-col>
            <a-col :span="12" v-if="formData.codetpl!='go'">
              <a-form-item field="viewsdir" label="前端父目录" validate-trigger="input" extra="是生成代码views中那个子目录里,不选则直接放在views中" >
                <a-select v-model="formData.viewsdir" placeholder="选择生成前端代码放在views目录那个位置" allow-clear allow-create>
                  <a-option v-for="item in views_dir" :value="item">{{item}}</a-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="12" v-if="formData.codetpl!='go'">
              <a-form-item field="vuedir" label="前端新业务" validate-trigger="input" :rules="[{required:true,message:'请填文件名称'},{match:/^[A-Za-z]+$/,message:'必须由字母组成'}]" extra="这个新增的代码目录，里面存放index.vue等文件">
                <a-input v-model="formData.vuedir" placeholder="请填创建文件名" :max-length="50" allow-clear show-word-limit @blur="handleGoPath"/>
              </a-form-item>
            </a-col>
            <a-col :span="12" v-if="formData.codetpl!='web'">
              <a-form-item field="goappdir" label="后端模块目录" validate-trigger="input" extra="后端模块目录,如果要创建新模块就不用选(空)" >
                <a-select v-model="formData.goappdir" placeholder="选择代码放在app目录中模块(如business)" allow-clear allow-create>
                  <a-option v-for="item in go_app_dir" :value="item">{{item}}</a-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="12" v-if="formData.codetpl!='web'">
              <a-form-item field="godir" label="后端类目录" validate-trigger="input" :rules="[{required:true,message:'请填后端目录'},{match:/^[A-Za-z]+$/,message:'必须由字母组成'}]" extra="选择的目录不存在可以创建新目录" >
                <a-select v-model="formData.godir" placeholder="选择或创建新代码文件所在目录" allow-clear allow-create>
                  <template v-for="item in go_file_dir">
                   <a-option v-if="item.bdir==formData.goappdir" :value="item.value" :label="item.value">{{item.title}}</a-option>
                  </template>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="12" v-if="formData.codetpl!='web'">
              <a-form-item field="gofilename"  label="go文件目录"  validate-trigger="input" :extra="`go文件名称默认index，文件为${formData.gofilename}.go`" :rules="[{match:/^[A-Za-z]+$/,message:'文件名必由字母组成'}]" >
                <a-input v-model="formData.gofilename" placeholder="填生成go文件名"  allow-clear/>
              </a-form-item>
            </a-col>
          </a-row>
        </a-form>
        <a-alert type="warning">
          <template #title>温馨提示</template>
          1.你可以使用该工具快速生成开发代码文件，帮助您快速构架好编辑代码文件。<br>
          2.创建的文件会在你选择的创建下，没有选择则后端在app目录下，前端在src/views目录下。<br>
          3.卸载请在列表数据找到生成方式为“代码模板生成”找到刚生产数据，在操作栏点击删除即可。
        </a-alert>
      </div>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref} from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { Message} from '@arco-design/web-vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  //api
  import { getMenuParent,getGoVueDir,markeTplCode} from '@/api/developer/generatecode';
  import { RuleItem } from '@/api/system/rule';
  import {  checkHaseRule } from '@/api/common';
  import { IconPicker ,Icon} from '@/components/Icon';
  import { inArray} from '@/utils/is';
  export default defineComponent({
    name: 'MakeCode',
    components: { BasicModal,IconPicker,Icon },
    emits: ['success'],
    setup(_, { emit }) {
      const modelHeight= ref(520);
      const sysdir=ref(["business","common","datacenter","matter","system","user","developer"])//除去系统目录
      const formData=ref({
        rule_name:"",
        icon:"",
        codelocation:"busDirName",
        pid:0,
        component:"",
        routename:"",
        routepath:"",
        codetpl:"web",
        viewsdir:"",//前端现有目录
        vuedir:"",//前端目录以为前端有几个文件组成
        goappdir:"business",//后端App根目录
        godir:"",//后端目录
        gofilename:"index",
      })
      const parntList = ref<RuleItem[]>([]);
      const haseadmin=ref(false)
      const go_app_dir=ref<string []>([])
      const go_file_dir=ref<any []>([])
      const views_dir=ref<string []>([])
      const { loading, setLoading } = useLoading();
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
        setLoading(true);
        handleGetMenuList()
        setLoading(false);
      });
      //提交数据
      const formRef = ref<FormInstance>();
      const handleSubmit=async()=>{
          try {
            const res = await formRef.value?.validate();
            if (!res) {
              setLoading(true);
              Message.loading({content:"生成中",id:"upStatus",duration:2000})
              formData.value.routename=formData.value.vuedir//路由由前端目录决定
              formData.value.routepath="/"+formData.value.routename
              if(formData.value.viewsdir){
                formData.value.component=`${formData.value.viewsdir}/${formData.value.routename}/index`
              }else{
                formData.value.component=`${formData.value.routename}/index`
              }
              if(formData.value.codetpl=="go"){
                formData.value.rule_name=""
              }
              await markeTplCode(formData.value);
              Message.success({content:"生成成功",id:"upStatus",duration:2000})
              closeModal()
              emit('success');
              setLoading(false);
            }
          } catch (error) {
            setLoading(false);
            Message.loading({content:"生成中",id:"upStatus",duration:1})
          }
      }
      //获取菜单列表-切换生成代码端
      const handleGetMenuList=async()=>{
        //回复字段值
        formData.value.viewsdir=""
        if(formData.value.codelocation=="busDirName"){
          formData.value.goappdir="business"
        }else if(formData.value.codelocation=="adminDirName"){
          formData.value.goappdir="admin"
        }else{
          formData.value.goappdir=""
        }
        const resultdata = await getMenuParent({codelocation:formData.value.codelocation});
        const parntList_df : any=[{id: 0,title: "一级菜单",pid: 0,locale:""}];
        if(resultdata.list){
            parntList.value=parntList_df.concat(resultdata.list)
        }else{
            parntList.value=[]
        }
        if(!inArray(resultdata.ids,formData.value.pid)){
          formData.value.pid=0
       }
        handleChangeMenu(formData.value.pid)
        const markedata = await getGoVueDir({codelocation:formData.value.codelocation});
        haseadmin.value=markedata.haseadmin
        go_app_dir.value=markedata.go_app_dir
        views_dir.value=markedata.views_dir
        if(markedata.go_file_dir){
          markedata.go_file_dir.forEach((item:any) => {
            const dir_arr=item.split("/")
            if(dir_arr&&dir_arr.length>1&&!inArray(sysdir.value,dir_arr[1])){
              go_file_dir.value.push({bdir:dir_arr[0],value:dir_arr[1],title:item})
            }
          });
        }
      }
      //选择菜单改变
      const handleChangeMenu=(val:any)=>{
          formData.value.viewsdir=""
          if(formData.value.codelocation=="busDirName"){
            formData.value.goappdir="business"
          }else if(formData.value.codelocation=="adminDirName"){
            formData.value.goappdir="admin"
          }else{
            formData.value.goappdir=""
          }
          const ruledata=parntList.value.find((item)=>item.id==val)
          if(ruledata&&ruledata["routepath"]){
            const modle_name=(ruledata["routepath"].split("/")[1]).trim()
            if(modle_name){
              if(inArray(views_dir.value,modle_name))
                formData.value.viewsdir=modle_name
                //处理go
                if(formData.value.codetpl=='goweb'){
                    formData.value.godir=modle_name
                }
              if(inArray(go_app_dir.value,modle_name))
                formData.value.goappdir=modle_name
              }
          }
      }
      //检查是否存在路由
      const handleChekHaseRule=(value:any, cb:any)=>{
        return new Promise(async(resolve) => {
            if(value){
              const resData = await checkHaseRule({id:0,codelocation:formData.value.codelocation,routename:value});
              if(!resData){
                cb("该路由已被占用请换一个")
              }
            }
            resolve(true)
          })
      }
      //路由名称改变
      const handleChangeRoutename=()=>{
        formData.value.vuedir=formData.value.routename
        if(formData.value.codetpl=='goweb'){
          formData.value.gofilename=formData.value.routename
        }
      }
      //选择生成模板类型
      const handdleCodetpl=()=>{
        if(formData.value.codetpl=="go"){
          formData.value.rule_name=""
        }
      }
      //更新GO文件
      const handleGoPath=()=>{
      }
      return { 
        registerModal, 
        getTitle:"本工具用于不依靠数据表结构生成代码模板", 
        loading,
        modelHeight,
        handleSubmit,formRef,
        formData,handleChangeMenu,parntList,handleGetMenuList,
        haseadmin,go_app_dir,views_dir,go_file_dir,handleGoPath,
        handleChangeRoutename,handleChekHaseRule,handdleCodetpl,
      };
    },
  });
</script>
<style lang="less" scoped>
  .table-content {
    padding: 20px;
  }
</style>

