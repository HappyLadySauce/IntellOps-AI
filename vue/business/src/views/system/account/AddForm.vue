<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :loading="loading" helpMessage="编辑和修改菜单" width="900px" :minHeight="500" :title="getTitle" @ok="handleSubmit">
    <a-form ref="formRef" class="form-content" :model="formData" auto-label-width>
      <a-row :gutter="16">
        <a-col :span="12">
          <a-form-item field="name" label="用户名" validate-trigger="input" :rules="[{required:true,message:'请填写用户名'}]"  >
            <a-input v-model="formData.name" placeholder="请填用户名" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="nickname" label="昵称" validate-trigger="input"  >
            <a-input v-model="formData.nickname" placeholder="请填写昵称" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="选择角色" field="roleid"  :rules="[{required:true,message:'请选择角色'}]">
            <a-tree-select placeholder="选择角色" :data="roleList" 
            :fieldNames="{
                key: 'id',
                title: 'name',
                children: 'children'
              }"
              multiple   v-model="formData.roleid">
            </a-tree-select>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="dept_id" label="选择部门"  >
            <a-tree-select placeholder="选择部门" :data="deptList" 
            :fieldNames="{
                key: 'id',
                title: 'name',
                children: 'children'
              }"  v-model="formData.dept_id">
            </a-tree-select>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="username" label="登录账号" :rules="usernameRules"
          :validate-trigger="['change', 'blur']">
            <a-input v-model="formData.username" placeholder="请填登录账号" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="password" label="登录密码" >
            <a-input v-model="formData.password" placeholder="登录密码(不修改则为空，默认密码123456)" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="mobile" label="手机号码"  >
            <a-input  v-model="formData.mobile" placeholder="请填手机号码" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="tel" label="座机" >
            <a-input v-model="formData.tel" placeholder="请填座机" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="email" label="邮箱"  >
            <a-input v-model="formData.email" placeholder="请填邮箱" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="address" label="地址"  >
            <a-input v-model="formData.address" placeholder="请填地址" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="city" label="城市"  >
            <a-input v-model="formData.city" placeholder="请填城市" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="company" label="公司"  >
            <a-input v-model="formData.company" placeholder="请填公司" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="remark" label="备注"  >
            <a-textarea v-model="formData.remark" placeholder="请填备注" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="avatar" label="头像"  >
            <FormImageBox v-model="formData.avatar" placeholder="请选择图片"/>
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
  import { getRole, save,DataItem,isAccountexist} from '@/api/system/account';
  import { getParent } from '@/api/system/dept';
  import { IconPicker ,Icon} from '@/components/Icon';
  import { Message } from '@arco-design/web-vue';
  export default defineComponent({
    name: 'AddBook',
    components: { BasicModal,IconPicker,Icon },
    emits: ['success'],
    setup(_, { emit }) {
      const { t } = useI18n();
      const isUpdate = ref(false);
      const roleList = ref<DataItem[]>([]);
      const deptList = ref<DataItem[]>([]);
      //表单
      const formRef = ref<FormInstance>();
      //表单字段
      const basedata={
            id:0,
            name: '',
            nickname: '',
            dept_id: 0,
            roleid: [],
            username: "",
            password: "",
            avatar:"",
            tel:"",
            mobile:"",//手机
            email:"",//邮箱
            address:"",//地址
            city:"",//城市
            remark:"",//备注
            company:"",//公司
        }
      const formData = ref(basedata)
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          formRef.value?.resetFields()
          setModalProps({ confirmLoading: false });
          const resultdata = await getRole();
          if(resultdata){
            roleList.value=resultdata
          }else{
            roleList.value=[]
          }
          const deptdata = await getParent();
          const parntList_df : any=[{id: 0,name: "未选部门",pid: 0}];
          if(deptdata){
            deptList.value=parntList_df.concat(deptdata)
          }else{
            deptList.value=parntList_df
          }
          isUpdate.value = !!data?.isUpdate;
          if (unref(isUpdate)) {
            formData.value=cloneDeep(data.record)
          }else{
            formData.value=cloneDeep(basedata)
          }
      });
      const getTitle = computed(() => (!unref(isUpdate) ? '新增管理账号' : '编辑管理账号'));
     //点击确认
     const { loading, setLoading } = useLoading();
     const handleSubmit = async () => {
      try {
          const res = await formRef.value?.validate();
          if (!res) {
            setLoading(true);
            if(!unref(isUpdate)&&formData.value.password==""){
              formData.value.password="123456"
            }
            Message.loading({content:"提交中",id:"save",duration:0})
            await save(unref(formData));
            Message.success({content:"提交成功",id:"save",duration:2000})
            closeModal()
            emit('success');
            setLoading(false);
          }
        } catch (error) {
          setLoading(false);
          Message.error({content:"提交失败",id:"save",duration:1})
        }
      };
      //验证账号唯一性
      const usernameRules = [{
        validator: (value:any, cb:any) => {
          return new Promise(async(resolve) => {
            if(!value){
              cb('请填写登录账号')
            }else{
              let sdata={username:value}
              if(formData.value.id>0){
                sdata=Object.assign({},sdata,{id:formData.value.id}) as any
              }
              const resData = await isAccountexist(sdata);
              if(resData.code==1){
                cb(resData.message)
              }
            }
            resolve(true)
          })
        }
      }];
      return { 
        registerModal, 
        getTitle, 
        handleSubmit,
        formRef,
        loading,
        formData,
        roleList,
        deptList,
        t,
        OYoptions:[
          { label: '否', value: 0 },
          { label: '是', value: 1 },
        ],
        usernameRules,
      };
    },
  });
</script>
<style lang="less" scoped>
.arco-upload-list-item{
  margin: 0px;
}
.form-content{
  padding: 10px;
}
</style>