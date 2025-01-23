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
										<a-col :span="12">
											<a-form-item field="title" label="标题" :rules="[{required:true,message:'请填写标题'}]" >
													<a-input v-model="formData.title" placeholder="输入标题" />
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="cid" label="分类"  >
													<FormBelongTable v-model="formData.cid" placeholder="请选分类" tablename="createcode_product_cate" showfield="name"/>
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="num" label="库存"  >
													<a-input-number v-model="formData.num" placeholder="输入库存" />
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="userType" label="用户类型"  >
													<FormDicSelect v-model="formData.userType" placeholder="请选用户类型" dicgroupid="2"/>
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="sex" label="性别"  >
													<a-radio-group v-model="formData.sex" >
														<a-radio :value="0" >未知</a-radio>
														<a-radio :value="1" >男</a-radio>
														<a-radio :value="2" >女</a-radio>
													</a-radio-group>
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="status" label="状态"  >
													<a-radio-group v-model="formData.status" >
														<a-radio :value="0" >正常</a-radio>
														<a-radio :value="1" >隐藏</a-radio>
													</a-radio-group>
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="price" label="价格" :rules="[{required:true,message:'请填写价格'}]" >
													<a-input-number v-model="formData.price" placeholder="输入价格" />
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="workerway" label="出行方式"  >
													<a-checkbox-group v-model="formData.workerway" >
														<a-checkbox value="car" >汽车</a-checkbox>
														<a-checkbox value="bus" >公交</a-checkbox>
														<a-checkbox value="air" >飞机</a-checkbox>
													</a-checkbox-group>
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="image" label="单张图"  >
													<FormImageBox v-model="formData.image" placeholder="请选单张图"/>
											</a-form-item>
										</a-col>
										<a-col :span="12">
											<a-form-item field="likeColor" label="喜欢颜色"  >
													<a-color-picker showPreset v-model="formData.likeColor"/>
											</a-form-item>
										</a-col>
										<a-col :span="24">
											<a-form-item field="images" label="多张图"  >
													<FormImagesBox v-model="formData.images" placeholder="请选多张图"/>
											</a-form-item>
										</a-col>
										<a-col :span="24">
											<a-form-item field="moreimgs" label="图组"  >
													<FormImagesBox v-model="formData.moreimgs" placeholder="请选图组"/>
											</a-form-item>
										</a-col>
										<a-col :span="24">
											<a-form-item field="record_audio" label="记录录音"  >
													<FormAudioBox v-model="formData.record_audio" placeholder="请选记录录音"/>
											</a-form-item>
										</a-col>
										<a-col :span="24">
											<a-form-item field="file" label="附件"  >
													<FormFileBox v-model="formData.file" placeholder="请选附件"/>
											</a-form-item>
										</a-col>
										<a-col :span="24">
											<a-form-item field="des" label="描述"  >
													<a-textarea v-model="formData.des" placeholder="输入描述" :auto-size="{minRows:3,maxRows:5}"/>
											</a-form-item>
										</a-col>
                    </a-row>
                </div>
              </a-scrollbar>
              <!--编辑器-->

							<FormEditorBox v-model="formData.content" placeholder="请编辑内容详情" :winHeights="windHeight" :activeKey="activeKey" :subnum="2" />
							<FormEditorBox v-model="formData.contenttow" placeholder="请编辑注意内容" :winHeights="windHeight" :activeKey="activeKey" :subnum="3" />             
          </div>
        </a-form>
      </div>
    </div>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, computed, unref} from 'vue';
  import { BasicModal, useModalInner} from '/@/components/Modal';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import { cloneDeep } from 'lodash-es';
  //api
  import { save,getContent } from './api';
  import { Message} from '@arco-design/web-vue';
  export default defineComponent({
    name: 'AddForm',
    components: { BasicModal},
    emits: ['success'],
    setup(_, { emit }) {
      const visibleimage=ref(false);
      //判断是否存在编辑器
      const isEditor=ref(true);
      const isUpdate = ref(false);
      const activeKey= ref(1);
      const modelHeight= ref(620);
      const windHeight= ref(620);
      //表单
      const formRef = ref<FormInstance>();
      //表单字段
      const basedata={
            id:0,
            title:"",
            cid:"",
            num:0,
            userType:"",
            sex:0,
            status:0,
            price:null,
            workerway:[],
            image:"",
            likeColor:"",
            images:"",
            moreimgs:"",
            record_audio:"",
            file:"",
            des:"",
            content:"",
            contenttow:"",

        }
      const formData = ref(basedata)
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          formRef.value?.resetFields()
          activeKey.value=1
          setModalProps({ confirmLoading: false });
          isUpdate.value = !!data?.isUpdate;
          if (unref(isUpdate)) {
            formData.value=cloneDeep(data.record)
            setLoading(true);
            const mewdata = await getContent({id:data.record.id});
            formData.value=Object.assign({},formData.value,mewdata)
            setLoading(false);
          }else{
            formData.value=cloneDeep(basedata)
          }
      });
      const getTitle = computed(() => (!unref(isUpdate) ? '新增数据' : '编辑数据'));
     //提交数据
     const { loading, setLoading } = useLoading();
     const handleSubmit = async () => {
      try {
          const res = await formRef.value?.validate();
          if (!res) {
            setLoading(true);
            Message.loading({content:"提交中",id:"submit",duration:0})
            let savedata=cloneDeep(unref(formData))
            await save(savedata);
            Message.success({content:"提交成功",id:"submit",duration:2000})
            closeModal()
            emit('success');
            setLoading(false);
          }
        } catch (error) {
          setLoading(false);
          Message.loading({content:"提交中",id:"submit",duration:1})
        }
      };
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
        tapList:[
          {id:1,name:"基础内容"},
					{id:2,name:"内容详情"},
					{id:3,name:"注意内容"},
        ],
        activeKey,
        modelHeight,
        onHeightChange,windHeight,
        isEditor,
        visibleimage,
      };
    },
  });
</script>
<style lang="less" scoped>
  @import '@/assets/style/formlayer.less';
</style>

