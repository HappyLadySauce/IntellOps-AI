<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" :loading="loading" width="1000px" @height-change="onHeightChange" :minHeight="modelHeight" :title="getTitle" @ok="handleSubmit">
    <div class="addFormbox" :style="{'min-height':`${windHeight}px`}">
      <div class="tabs-header" v-if="isEditor">
        <div class="tabs-nav-wrap">
            <div class="tap_item" v-for="iten in tapList" :class="{item_active:activeKey==iten.id}" @click="()=>{activeKey=iten.id}">
                <div class="label">{{iten.name}}</div>
            </div>
        </div>
        <div class="tabs-bar" :style="{top: `${(activeKey-1)*64}px`,height: `64px`}"></div>
      </div>
      <div class="tabs-content" :class="{addpadding:!isEditor}">
        <a-form ref="formRef" :model="formData" auto-label-width>
          <div class="content_box">
              <!--基础信息-->
              <a-scrollbar v-show="activeKey==1" style="overflow: auto;" :style="{height:`${windHeight}px`}">
                <div class="besecontent" >
                  <a-row :gutter="16">
										<a-col :span="24">
											<a-form-item field="title" label="标题" :rules="[{required:true,message:'请填写标题'}]" >
													<a-input v-model="formData.title" placeholder="请填标题" allow-clear/>
											</a-form-item>
										</a-col>
										<a-col :span="24">
											<a-form-item field="des" label="说明" :rules="[{required:true,message:'请填写说明'}]" >
													<a-textarea v-model="formData.des" placeholder="请填说明" :auto-size="{minRows:3,maxRows:5}" allow-clear/>
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="name" label="限定包名" :rules="[{required:true,message:'请填限定包名'}]" >
													<a-input v-model="formData.name" placeholder="请填限定包名,安装是代码包名(以字母、数字命名)" allow-clear/>
											</a-form-item>
										</a-col>
                    <a-col :span="12">
											<a-form-item field="price" label="预算金额" :rules="[{required:true,message:'请填价格'}]" >
													<a-input v-model="formData.price" placeholder="请填价格" allow-clear/>
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="customer" label="联系人" :rules="[{required:true,message:'请填写联系人'}]" >
													<a-input v-model="formData.customer" placeholder="请填联系人" allow-clear/>
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="mobile" label="联系电话" :rules="[{required:true,message:'请填写联系电话'}]" >
													<a-input v-model="formData.mobile" placeholder="请填联系电话" allow-clear/>
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="wx" label="微信" >
													<a-input v-model="formData.wx" placeholder="请填微信" allow-clear/>
											</a-form-item>
										</a-col>
                  </a-row>
                </div>
              </a-scrollbar>
              <!--高级信息-->
              <div class="hcontent"  v-show="activeKey==2" :style="{height:`${windHeight}px`}">
                <Editor :minHeight="windHeight" ref="editorRef" @updata="handleEditUpdta"/>
              </div>
          </div>
        </a-form>
      </div>
    </div>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref} from 'vue';
  import { BasicModal, useModalInner,useModal } from '/@/components/Modal';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import { cloneDeep } from 'lodash-es';
  //api
  import { requirement } from '@/api/developer/requireForm';
  import { Message ,TreeNodeData} from '@arco-design/web-vue';
  import Editor from "@/components/Editor/Main.vue"; // @ is an alias to /src
  export default defineComponent({
    name: 'AddBook',
    components: { BasicModal,Editor },
    emits: ['success'],
    setup(_, { emit }) {
      const [registerFileModal, { openModal:openFileModal }] = useModal();
      const visibleimage=ref(false);
      const cateList = ref<TreeNodeData[]>([]);
      //判断是否存在编辑器
      const isEditor=ref(true);
      const isUpdate = ref(false);
      const activeKey= ref(1);
      const modelHeight= ref(620);
      const windHeight= ref(620);
      //表单
      const formRef = ref<FormInstance>();
      //表单字段
      const basedata=ref({
            id:0,
            cid:0,
            content:"",
            code_token:"",
            title:"",
            name:"",
            des:"",
            customer:"",
            mobile:"",
            wx:"",
            price:"",
            attachment:[],
            createtime:0,
        })
      const formData = ref(basedata.value)
      //编辑器
      const editorRef = ref();
      const domian = ref("");
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          formRef.value?.resetFields()
          activeKey.value=1
          setModalProps({ confirmLoading: false })
          isUpdate.value = !!data?.isUpdate;
          formData.value.code_token= data.record.code_token
          domian.value=data.record.domian
          if(editorRef.value){
              editorRef.value.setVal("")
            }
      });
      const getTitle = computed(() => (!unref(isUpdate) ? '发布需求' : '编辑数据'));
     //点击确认
     const { loading, setLoading } = useLoading();
     const handleSubmit = async () => {
      try {
          const res = await formRef.value?.validate();
          if (!res) {
            setLoading(true);
            Message.loading({content:"提交中",id:"requirement",duration:0})
            let savedata=cloneDeep(unref(formData))
            await requirement(savedata,domian.value);
            Message.success({content:"提交成功",id:"requirement",duration:2000})
            closeModal()
            emit('success');
            setLoading(false);
          }
        } catch (error) {
          setLoading(false);
          Message.loading({content:"提交中",id:"requirement",duration:1})
        }
      };
      //编辑器返回数据
      const handleEditUpdta=(val:string)=>{
         formData.value["content"]=val
      }
       //监听高度
       const onHeightChange=(val:any)=>{
        windHeight.value=val
      }
      return { 
        registerModal, 
        getTitle, 
        handleSubmit,
        formRef,
        loading,
        formData,
        OYoptions:[
          { label: '否', value: 0 },
          { label: '是', value: 1 },
        ],
        SHoptions:[
            { label: '正常', value: 0 },
            { label: '禁用', value: 1 },
        ],
        tapList:[
          {id:1,name:"需求信息"},
          {id:2,name:"需求内容"},
        ],
        activeKey,
        handleEditUpdta,
        modelHeight,
        editorRef,
        onHeightChange,windHeight,
        isEditor,cateList,
        registerFileModal,visibleimage,
      };
    },
  });
</script>
<style lang="less" scoped>
  @import '@/assets/style/formlayer.less';
  
  .addpadding{
    padding: 0px 20px;
  }
  .upfilezip{
    display: flex;
    .upbtn{
      padding-right: 10px;
    }
    .showfile{
      flex: 1;
      height: 32px;
      line-height: 32px;
      a{
        text-decoration: none;
      }
    }
  }
  //上传图片
  .upimagebox{
    display: flex;
    .imagebtn{
      position: relative;
      width: 160px;
      height: 90px;
      background-color: var(--color-neutral-1);
      border-radius: 4px;
      overflow: hidden;
      -ms-flex-negative: 0;
      flex-shrink: 0;
      //预览
      .upload-show-picture{
        position: relative;
        box-sizing: border-box;
        width: 100%;
        height: 100%;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
        img{
          height: 100%;
        }
        &:hover{
          .upload-show-picture-mask{
             opacity: 1;
          }
        }
        .upload-show-picture-mask{
            position: absolute;
            top: 0;
            right: 0;
            bottom: 0;
            left: 0;
            color: var(--color-white);
            font-size: 16px;
            line-height: 90px;
            text-align: center;
            background: rgba(0, 0, 0, 0.5);
            opacity: 0;
            transition: opacity 0.1s cubic-bezier(0, 0, 1, 1);
            .opbtn{
              cursor: pointer;
            }
        }
      }
      .upload-picture-card{
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        user-select: none;
        cursor: pointer;
        .upload-picture-card-text{
           text-align: center;
           color:  var(--color-neutral-5);
        }
      }
    }
  }
</style>

