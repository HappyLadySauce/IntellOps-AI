<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :isPadding="false" popup-container="body" :loading="loading" width="998px" @height-change="onHeightChange" :minHeight="modelHeight" :title="getTitle" @ok="handleSubmit">
    <div class="addFormbox" :style="{'min-height':`${windHeight}px`}">
      <div class="tabs-header" >
        <div class="tabs-nav-wrap">
          <template v-for="iten in tapList" >
            <div class="tap_item" v-if="iten.id==1||(iten.id==2&&selectdata.filetype=='image')" :class="{item_active:activeKey==iten.id}" @click="selectTab(iten.id)">
                <div class="label">{{iten.name}}</div>
            </div>
          </template>
        </div>
        <div class="tabs-bar" :style="{top: `${(activeKey-1)*64}px`,height: `64px`}"></div>
      </div>
      <div class="tabs-content" >
        <div class="content_box">
            <!--我的附件-->
            <div class="myfile" v-show="activeKey==1"  :style="{height:`${windHeight}px`}">
              <div class="header_toopl">
               <div class="btnlist">
                <a-upload
                    :accept="upaccept"
                    :show-file-list="false"
                    :multiple="true"
                    :custom-request="customRequest"
                >
                  <template #upload-button>
                   <a-button >
                      <template #icon>
                        <icon-upload />
                      </template>直接上传  
                    </a-button>
                  </template>
                </a-upload>
                <a-button @click="onAddFolder"  style="margin-left: 10px;">
                   <template #icon>
                    <icon-folder-add />
                    </template> 新增文件夹
                </a-button>
               </div>
               <div class="search">
                <a-input-search v-model="searchword" :style="{width:'300px'}" @press-enter="onSearchMyfile" allow-clear @search="onSearchMyfile" placeholder="搜索文件名称"/>
               </div>
             </div>
             <!--目录菜单h: 22+12-->
            <div class="dirmun" v-if="dirmenu&&dirmenu.length>0">
              <span class="dirmun_item" v-for="item in dirmenu" @click="onDirmun(item)">
                <span class="dirmun_inner">{{item['title']}}</span> 
                <span class="dirmun_separator">&gt;</span>
              </span>
            </div>
             <div class="file_list" :style="{height:`${windHeight-(dirmenu.length>0?80:55)}px`}">
              <a-scrollbar  style="overflow: auto;" :style="{height:`${windHeight-(dirmenu.length>0?80:55)}px`}">
                <div class="file_list_box">
                    <div class="file_item" v-for="iten in myFileList">
                      <div>
                        <div :class="{imgbox:iten.type==0||iten.type==2||iten.type==3||iten.type==4,folderbox:iten.type==1}" @click="onSelectImg(iten)">
                          <!--图片-->
                          <div class="img_raw" v-if="iten.type==0" >
                            <div class="img_preview">
                              <img :src="GetFullPath(iten.url)" class="img_item"/>
                            </div>
                          </div>
                          <!--视频-->
                          <div class="img_raw" v-else-if="iten.type==2" >
                            <div class="img_preview videobox">
                              <img :src="iten.cover_url" class="img_item"/>
                              <icon-play-circle class="folder_video" :size="50" style="color:#ffffff"/>
                            </div>
                          </div>
                          <!--文件夹-->
                          <div class="folder_raw" v-else-if="iten.type==1" 
                            @dragover="allowDrop($event)"
                            @drop="drop($event,iten.id)"
                            @dragleave="dragleaves()"
                          >
                            <Icon icon="svgfont-wenjianjia-1" class="folder_preview" :forceFallback="true" :size="82" color="#ed6f6f"></Icon>
                          </div>
                          <!--其他文件-音频、文件、其他-->
                          <div class="img_raw" v-else >
                            <div class="img_preview">
                              <Icon :icon="getFileImg(iten)" :size="50" ></Icon>
                            </div>
                          </div>
                          <div class="edit_layer " :class="{l_active:selectList.indexOf(`my_${iten.id}`)>-1}" :draggable="iten.type==0||iten.type==2||iten.type==3||iten.type==4" @dragstart="drag($event,iten.id)">
                            <Icon icon="svgfont-guanbi" class="colse" v-if="iten.is_common!=1" :size="18" @click="onDel($event,iten.id,iten.type)" color="red"></Icon>
                            <Icon icon="svgfont-mark_sel" class="choose" :size="18" v-if="iten.type==0||iten.type==2||iten.type==3||iten.type==4"></Icon>
                          </div>
                        </div>
                        <div class="titlebox">
                          <div class="edit_title" title="点击可编辑文件名称">
                            <a-input v-model="iten.title" size="mini" placeholder="输入名称" :readonly="iten.is_common==1" style=" padding: 0;" @change="onChangFile" @pressEnter="(e)=>{onUpFile(iten.id,e)}" @blur="(e)=>{onUpFile(iten.id,e)}" :bordered="false"/> 
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
               </a-scrollbar>
              </div>
            </div>
            <!--图片库-->
            <div class="hcontent"  v-show="activeKey==2" :style="{height:`${windHeight}px`}">
              <FileLibrary ref="fileLibraryRef" @ok="selectLibrary"></FileLibrary>
            </div>
        </div>
      </div>
    </div>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, reactive,toRefs,computed} from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import useLoading from '@/hooks/loading';
  //api
  import { getMyFiles,upImgPid,delFile,delDir,saveFile,FileItem,PictureItem} from '@/api/datacenter/attachment';
  import { userUploadApi} from '@/api/datacenter/upfile';
  import { Message ,Modal} from '@arco-design/web-vue';
  import type { RequestOption} from '@arco-design/web-vue/es/upload/interfaces';
  import FileLibrary from './FileLibrary.vue';
  import {Icon} from '@/components/Icon';
  import { GetFullPath } from "@/utils/tool";
  export default defineComponent({
    name: 'FileManage',
    components: { BasicModal,FileLibrary,Icon},
    emits: ['success',"onfresh"],
    setup(_, { emit }) {
      //判断是否存在编辑器
      const { loading, setLoading } = useLoading();
      const activeKey= ref(1);
      const openfrom= ref("use");
      const tapList= ref([
          {id:1,name:"我的文件"},
        ]);
      const modelHeight= ref(620);
      const windHeight= ref(620);
      //附件管理
      const arraynct: string[] = [] // 定义字符串数组
      const pictureitem: PictureItem[] = [] // 定义整数型数组
      const file_item: FileItem[] = [] // 定义整数型数组
      const pagedata=reactive({
        activeKey:1,
        myFileList:file_item,//我的文件
        searchword:"",
        typeid:0,//素材图/插图
        pid:0,//当前层级id 0=父级
        upLoading:false,//上传文件
        upaccept:"image/*",
        selectdata:{getnumber:"one",exdata:"",filetype:"image" },//选返回值
        selectList:arraynct,//选择
        dirmenu:[],//我的文件目录菜单
        //分页
        pagenum:1,
        pagetotall:0,
        picture:pictureitem,
        //允许上传文件大小
        fileMaxSize:1,//1M
        //窗口
        modaltitle:"选择图片",
        modaldes:"(只能添加jpg,jpeg,bmp,png,gif, 大小不超过1MB。)",//备注
        tapList:[
          {id:1,name:"我的文件"},
        ],
      })
     const MaxSizeVideo= window?.globalConfig.MaxSizeVideo
     const MaxSizeImage= window?.globalConfig.MaxSizeImage
     const param=ref();
      //打开也初始化
      const modelTime=ref("")
      const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
          activeKey.value=1
          setModalProps({ confirmLoading: false });
          pagedata.selectList=[]
          pagedata.pid=0
          param.value=data?.param|| ""//图片带的参数
          pagedata.selectdata.getnumber=data?.getnumber|| "one"//获取多张还是单张-默认单张
          pagedata.selectdata.filetype=data?.filetype|| "image" //获取附件类型-默认图片
          openfrom.value=data?.openfrom|| "user" //获取附件类型-默认图片
          pagedata.selectdata.exdata=data
          if(pagedata.selectdata.filetype=="video"){
            pagedata.fileMaxSize=MaxSizeVideo
            pagedata.modaltitle="选择视频"
            modelTime.value="视频"
            pagedata.upaccept="video/*"
            pagedata.modaldes=`(只能添加mp4, 大小不超过${MaxSizeVideo}MB。)`
        }else if(pagedata.selectdata.filetype=="audio"){
            pagedata.fileMaxSize=MaxSizeImage
            pagedata.modaltitle="选择音频"
            modelTime.value="音频"
            pagedata.upaccept="audio/*"
            pagedata.modaldes=`(只能添加音频类文件, 大小不超过${MaxSizeImage}MB。)`
        }else if(pagedata.selectdata.filetype=="file"){
            pagedata.fileMaxSize=MaxSizeImage
            pagedata.modaltitle="选择文件"
            modelTime.value="文件"
            pagedata.upaccept=".zip,.rar,.xls,.xlsx,.pdf,.doc,.docx"
            pagedata.modaldes=`(只能添加.zip,.rar,.xls,.xlsx,.pdf,.doc,.docx, 大小不超过${MaxSizeImage}MB。)`
        }else{
            pagedata.fileMaxSize=MaxSizeImage
            pagedata.modaltitle="选择图片"
            modelTime.value="图片"
            pagedata.upaccept="image/*"
            pagedata.modaldes=`(只能添加jpg,jpeg,bmp,png,gif, 大小不超过${MaxSizeImage}MB。)`
        }
        getMyfile("")
      });
      const getTitle = computed(() => (modelTime.value+'管理'));
      //添加文件夹
      const onAddFolder=()=>{
         onSaveData({
            title:"新建文件夹",
            type:1,
            pid:pagedata.pid,
         })
      }
      //保存数据
      const onSaveData=async(data:object)=> {
        try {
          Message.loading({ content: '提交中...', id:"saveData",duration:0});
            const resultdata = await saveFile(data);
            if(resultdata){
              pagedata.myFileList=resultdata.concat(pagedata.myFileList)
              Message.success({ content: '提交成功！', id:"saveData", duration: 2000 });
            }else if(resultdata==0){
              Message.success({ content: '已提交！', id:"saveData", duration: 2000 });
            }
          } catch {
            Message.clear();
          }
      }
      //上传附件=图片
      const customRequest = (options: RequestOption) => {
          Message.loading({content:"上传中",id:"upStatus",duration:0})
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
              const filename=fileItem?.name||""
              const resdata = await userUploadApi({ name: 'file', file: fileItem.file as Blob, filename,data:{pid:pagedata.pid,filetype:pagedata.selectdata.filetype}},onUploadProgress) as any;
              //更新图片
              if(resdata){
                if(resdata["code"]==0){
                    if(resdata["message"]=="文件已上传"){
                      pagedata.myFileList=pagedata.myFileList.filter((item)=>resdata["data"].id!=item.id)
                    }
                    pagedata.myFileList=[resdata["data"]].concat(pagedata.myFileList)
                    Message.success({content:"上传成功",id:"upStatus",duration:2000})
                    emit("onfresh",true)
                }else{
                  Message.clear()
                  Modal.error({
                    title: '上传文件失败',
                    titleAlign:"start",
                    content: resdata["message"]
                  });
                }
              }else{
                Message.error({content:resdata["message"],id:"upStatus",duration:2000})
              }
              Message.clear()
            } catch (error) {
              onError(error);
              Message.error({content:"上传失败",id:"upStatus",duration:2000})
            }
          })();
          return {
            abort() {
              controller.abort();
            },
          };
      };
       //监听高度
       const onHeightChange=(val:any)=>{
        windHeight.value=val
      }
      //选择侧边菜单
      const fileLibraryRef=ref()
      const selectTab=(id:number)=>{
        activeKey.value=id
        if(id==2){
          fileLibraryRef.value.initGetPictureCate(pagedata.selectdata.getnumber)
        }
      }
      //选择图片库图片返回
      const selectLibrary=(imgs:any,picture:any)=>{
        pagedata.picture=picture
        pagedata.selectList=pagedata.selectList.filter((item)=>item.indexOf("mt_")==-1)
        imgs.forEach((item:any)=>{
           if(pagedata.selectList.indexOf(item)==-1){
            pagedata.selectList.push(item)
           }
        })
    }
      //最后选择图片
      interface VideoItem {
        url: string;
        cover_url: string;
      }
      //选择图片-点击确定
     const handleSubmit = async () => {
         try {
            if(pagedata.selectList&&pagedata.selectList.length>0){
              const img_list: string[] = [] // 定义字符串数组
              const video_list: VideoItem[] = [] // 定义字符串数组
              let onefileinfo
              pagedata.selectList.forEach((item)=>{
                const id_arr=item.split("_")
                if(pagedata.selectdata.filetype=="image"){
                  if(id_arr[0]=="my"){//去我的附件找
                    const myfile= pagedata.myFileList.find((data)=>data.id==parseInt(id_arr[1]))
                    if(myfile){
                      onefileinfo=myfile
                      img_list.push(myfile.url)
                    }
                  }else if(id_arr[0]=="mt"){//去图库找
                    const pic= pagedata.picture.find((data)=>data.id==parseInt(id_arr[1]))
                    if(pic){
                      onefileinfo=pic
                      img_list.push(pic.url)
                    }
                  }
                }else if(pagedata.selectdata.filetype=="file"||pagedata.selectdata.filetype=="audio"){
                   const myfile= pagedata.myFileList.find((data)=>data.id==parseInt(id_arr[1]))
                    if(myfile){
                      onefileinfo=myfile
                      img_list.push(myfile.url)
                    }
                }else{
                    const myfile= pagedata.myFileList.find((data)=>data.id==parseInt(id_arr[1]))
                    if(myfile){
                      onefileinfo=myfile
                      video_list.push({"url":myfile.url,"cover_url":myfile.cover_url})
                    }
                }
              })
              if(pagedata.selectdata.filetype=="image"){//图片
                if(pagedata.selectdata.getnumber=="one"){
                  emit('success', {param:param.value,filetype:pagedata.selectdata.filetype,type:pagedata.selectdata.getnumber,url:img_list[0],list:[],data:onefileinfo});
                }else{
                  emit('success', {param:param.value,filetype:pagedata.selectdata.filetype,type:pagedata.selectdata.getnumber,url:"",list:img_list,data:pagedata.selectdata.exdata});
                }
              }else if(pagedata.selectdata.filetype=="file"){//文件
                if(pagedata.selectdata.getnumber=="one"){
                  emit('success', {param:param.value,filetype:pagedata.selectdata.filetype,type:pagedata.selectdata.getnumber,url:img_list[0],list:[],data:onefileinfo});
                }else{
                  emit('success', {param:param.value,filetype:pagedata.selectdata.filetype,type:pagedata.selectdata.getnumber,url:"",list:img_list,data:pagedata.selectdata.exdata});
                }
              }else if(pagedata.selectdata.filetype=="audio"){//音频
                if(pagedata.selectdata.getnumber=="one"){
                  emit('success', {param:param.value,filetype:pagedata.selectdata.filetype,type:pagedata.selectdata.getnumber,url:img_list[0],list:[],data:onefileinfo});
                }else{
                  emit('success', {param:param.value,filetype:pagedata.selectdata.filetype,type:pagedata.selectdata.getnumber,url:"",list:img_list,data:pagedata.selectdata.exdata});
                }
              }else{//视频
                if(pagedata.selectdata.getnumber=="one"){
                  emit('success', {param:param.value,filetype:pagedata.selectdata.filetype,type:pagedata.selectdata.getnumber,...video_list[0],list:[],data:onefileinfo});
                }else{
                  emit('success', {param:param.value,filetype:pagedata.selectdata.filetype,type:pagedata.selectdata.getnumber,url:"",cover_url:"",list:video_list,data:pagedata.selectdata.exdata});
                }
              }
              closeModal()
            }else{
              if(openfrom.value=="use"){
                Message.error(pagedata.selectdata.filetype=="image"?"请选择图片":"请选择视频");
              }else{
                closeModal()
              }
            }
          } catch (error:any) {
            Message.error(error.String())
          }
      };
      //搜索我的文件
      const onSearchMyfile=()=>{
        if(pagedata.searchword){
          getMyfile(pagedata.searchword)
        }else{
          getMyfile("")
        }
      }
     //获取-我的文件
     const getMyfile=async(searchword:any) =>{
        setLoading(true);
        const getdbdata = await getMyFiles({searchword:searchword,pid:pagedata.pid,filetype:pagedata.selectdata.filetype});
        if(getdbdata){
          pagedata.myFileList=getdbdata.items
          const findgallery= tapList.value.find((item)=>item.id==2)
          if(getdbdata.hasegallery&&!findgallery){
            tapList.value.push({id:2,name:"图片库"})
          }
          if(getdbdata.dirmenu.length>0){
            let parntList : any=[{id: 0,title: "默认文件夹",pid: 0}];
            pagedata.dirmenu = parntList.concat(getdbdata.dirmenu)
          }else{
            pagedata.dirmenu=[]
          }
        }
        setLoading(false);
      }
       //选择文件
       function onSelectImg(item:any){
        if(item.type==0||item.type==3||item.type==4){//是图片//或图片库
          var id_str=`my_${item.id}`
          if(pagedata.selectList.indexOf(id_str)>-1){
            pagedata.selectList=pagedata.selectList.filter((id)=>id!=id_str)
          }else{
            if(pagedata.selectdata.getnumber=="one"){
              pagedata.selectList=[id_str]
            }else{
              pagedata.selectList.push(id_str)
            }
          }
        }else if(item.type==2){
          var id_str=`my_${item.id}`
          if(pagedata.selectList.indexOf(id_str)>-1){
            pagedata.selectList=pagedata.selectList.filter((id)=>id!=id_str)
          }else{
            if(pagedata.selectdata.getnumber=="one"){
              pagedata.selectList=[id_str]
            }else{
              pagedata.selectList.push(id_str)
            }
          }
        }else if(item.type==1){//是文件夹
          pagedata.pid=item.id
          pagedata.searchword=""
          getMyfile("")
        }
      }
       //拖拽
      /**
       * 1当将元素或文本选择拖动到有效放置目标（每几百毫秒）上时，会触发此事件
       *
       * @param {Object} event event对象
       */
       const allowDrop = (event:any) => {
        //阻止浏览器的默认事件
        event.preventDefault()

       }
          /**
       * 2当在有效放置目标上放置元素或选择文本时触发此事件
       *
       * @param {Object} event event对象
       */
      const drop = (event:any,pid:any) => {
        const imgid=event.dataTransfer.getData('imgid')
        //处理移动
        onUpImgPid(imgid,pid)
      }
        /**
       * 3 当拖动的元素或文本选择离开有效的放置目标时，会触发此事件
       *
       * @param {Object} event event对象
       */
       const dragleaves = () => {
        /* 删除提示组件 */
      }
      //4.移动前
      const drag=(event:any,id:any)=>{
        event.dataTransfer.setData('imgid', id)
      }
      //提交文件目录移动
       const onUpImgPid=async(imgid:any,pid:any) =>{
        Message.loading({ content: '移动中...', id:"upImgPid",duration:0});
          const resultdata = await upImgPid({imgid:imgid,pid:pid});
          if(resultdata){
            pagedata.myFileList=pagedata.myFileList.filter((item)=>imgid!=item.id)
            Message.success({ content: '移动成功！', id:"upImgPid", duration: 2000 });
          }else if(resultdata==0){
            Message.clear()
          } else{
            Message.error({ content: '移动失败', id:"upImgPid", duration: 2000 });
          }
      }
    //点击目录菜单
    const onDirmun=(item:any)=>{
        if(pagedata.pid!=item.id){
          pagedata.pid=item.id
          pagedata.searchword=""
          getMyfile("")
        }
      }
      //删除附件
      const onDel=(event:any,id:any,type:any)=>{
          //阻止事件冒泡到父元素
          event.stopPropagation();
          Modal.warning({
              title: '您确定删除吗？',
              content: '删除后内容将无法恢复请谨慎操作！',
              cancelText:"取消",
              okText:"删除",
              titleAlign:"start",
              hideCancel:false,
              onOk:(async()=>{
                  Message.loading({content:"删除中",id:"delFile",duration:0})
                  var resultdata: any
                  if(type==1){
                    resultdata = await delDir({id:id});
                  }else{
                    resultdata = await delFile({ids:[id]});
                  }
                  if(resultdata){
                    pagedata.myFileList=pagedata.myFileList.filter((item)=>id!=item.id)
                    Message.success({content:"删除成功",id:"delFile", duration: 2})
                    emit("onfresh",true)
                  }else {
                    Message.error({content:"删除失败",id:"delFile", duration: 2})
                  }
                })
            });
      }
    //是否修改文件名称
    const ischangname=ref(false)
    const onChangFile=()=>{
      ischangname.value=true
    }
    //更新附件名称
    const onUpFile=async(id:any,ev:any)=>{
        const title=ev.target._value
        if(ischangname.value){
          ischangname.value=false
          Message.loading({ content: '更新中...', id:"upFile",duration:0});
          const resultdata = await saveFile({id:id,title:title});
          if(resultdata){
            Message.success({ content: '更新成功！', id:"upFile", duration: 2000 });
          }else if(resultdata==0){
            Message.clear()
          } else{
            Message.error({ content: '更新失败', id:"upFile", duration: 2000 });
          }
        }
      }
      // 获取文件图标，如果是图片就显示图片
      /** 文件类型图标 Map 映射 */
      interface FileExtendNameIconMap {
        [key: string]: string
      }
       const FileIcon: FileExtendNameIconMap = {
        mp3: 'svgfont-yinpin',
        mp4: 'svgfont-shipin1',
        dir: 'svgfont-wenjianjia-1',
        ppt: 'file-ppt',
        doc: 'svgfont-wendang-docx_doc',
        docx: 'svgfont-wps',
        xls: 'svgfont-XLS',
        xlsx: 'svgfont-XLS',
        txt: 'svgfont-wenbenwendang-txt',
        rar: 'svgfont-yasuowenjian',
        zip: 'svgfont-zipyasuo',
        html: 'svgfont-HTMLtubiao',
        css: 'svgfont-css',
        js: 'svgfont-js',
        other: 'svgfont-qitawenjian' // 未知文件
      }
      /** WPS、OffOfficeTypesice文件类型 */
      const OfficeTypes = ['txt','ppt', 'pptx', 'doc', 'docx', 'xls', 'xlsx','zip','rar','html','css','js']
      /** 音频类型 */
       const AudioTypes = ['mp3','ram','wav']
      const getFileImg =(data:any) => {
        const extension = GetFileExtension(data)
        if (!Object.keys(FileIcon).includes(extension)) {
        if(AudioTypes.includes(extension)){
          return FileIcon['mp3']
          }else if(OfficeTypes.includes(extension)){
            return FileIcon['othetxtr']
          }
          return FileIcon['other']
        }
        return FileIcon[extension]
      }
      /** 获取文件后缀 */
     const GetFileExtension=(data:any)=>{
        const name = data.name.toLowerCase()
        if(data.type?.toString()=="1"){
          return "dir"
        }else if(name){
          const filearr=name.split(".")
          if(filearr.length>=2){
            return filearr[filearr.length-1]
          }
        }
        return "other"
      }
      return { 
        ...toRefs(pagedata),
        GetFullPath,
        registerModal, 
        getTitle, 
        handleSubmit,
        loading,
        tapList,
        activeKey,selectLibrary,
        customRequest,
        modelHeight,
        onHeightChange,windHeight,selectTab,
        fileLibraryRef,onAddFolder,onSearchMyfile,
        //拖拽
        allowDrop,drop,dragleaves,drag,
        onSelectImg,onDel,onChangFile,onUpFile,onDirmun,
        getFileImg,
      };
    },
  });
</script>
<style lang="less" scoped>
  @import '@/assets/style/formlayer.less';
  //我的文件
  .myfile{
    padding: 10px;
    .header_toopl{
      height: 35px;
      display: flex;
      .btnlist{
        flex:1;
      }
      .seach{

      }
    }
     //我的-目录
     .dirmun{
          height: 22px;
          line-height: 22px;
          overflow: hidden;
          margin: 6px 0px;
          .dirmun_item{
            display: inline-flex;
            align-items: center;
            &:last-child .dirmun_inner {
                color: rgba(0,0,0,.38);
                cursor: default;
            }
            &:last-child .dirmun_separator {
                display: none;
            }
            .dirmun_inner{
              color: #557ce1;
              cursor: pointer;
              font-size: 13px;
            }
            .dirmun_separator{
              padding: 0 4px;
              font-size: 14px;
              color: rgba(0,0,0,.38);
            }
          }
        }
     //我的文件列表
     .file_list{
          .file_list_box{
            padding: 0;
            margin: 0;
            overflow: hidden;
            .file_item{
              float: left;
              width: 84px;
              list-style: none;
              margin-top: 20px;
              margin-right: 23px;
              //图片
              .imgbox{
                cursor: pointer;
                position: relative;
                height: 100%;
                .img_raw{
                  position: relative;
                  display: flex;
                  justify-content: center;
                  align-items: center;
                  height: 84px;
                  background-color: #f3f3f3;
                  overflow: hidden;
                  border: 1px solid #e7e7eb;
                  .img_preview{
                    box-sizing: border-box;
                    text-align: center;
                    font-size: 0;
                    line-height: 0;
                    width: 100%;
                    .img_item{
                      display: inline-block;
                      width: auto;
                      height: auto;
                      max-width: 100%;
                      max-height: 100%;
                      vertical-align: middle;
                    }
                  }
                      //视频
                    .videobox{
                      position: relative;
                      width: 100%;
                        height: 100%;
                      .folder_video{
                        position: absolute;
                        top: 18px;
                        left: 16px;
                      }
                    }
                }
                .edit_layer{
                  display: none;
                  position: absolute;
                  width: 100%;
                  height: 100%;
                  border: 1px solid #5874d8;
                  left: 0;
                  top: 0;
                  box-sizing: border-box;
                  .colse{
                    width: 14px;
                    height: 14px;
                    position: absolute;
                    right: -5px;
                    top: -5px;
                    cursor: pointer;
                  }
                  .choose{
                    display: none;
                    position: absolute;
                    right: 0;
                    bottom: 0;
                    content: "";
                    width: 28px;
                    height: 24px;
                  }
                }
                //选中
                .l_active{
                  display: block;
                  .colse{
                    display: none;
                  }
                  .choose{
                      display: block;
                    }
                }
                  //浮框
                &:hover{
                    .edit_layer{
                      display: block;
                      .colse{
                          display: block;
                      }
                    }
                  }
              }
              //文件夹
              .folderbox{
                cursor: pointer;
                position: relative;
                height: 100%;
                //文件夹
                .folder_raw{
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 84px;
                    cursor: pointer;
                    .folder_preview{

                    }
                }
                .edit_layer{
                  display: none;
                  position: absolute;
                  width: 100%;
                  height: 100%;
                  left: 0;
                  top: 0;
                  box-sizing: border-box;
                  .colse{
                    width: 14px;
                    height: 14px;
                    position: absolute;
                    top: 13px;
                    right: -2px;
                    cursor: pointer;
                  }
                }
                //选中
                .l_active{
                  display: block;
                  .colse{
                    display: none;
                  }
                  .choose{
                      display: block;
                    }
                }
                  //浮框
                &:hover{
                    .edit_layer{
                      display: block;
                      .colse{
                          display: block;
                      }
                    }
                  }
              }
              //名称
              .titlebox{
                height: 24px;
                margin-top: 1px;
                font-size: 13px;
                color: #333;
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
                text-align: center;
                .edit_title{
                  display: inline-block;
                  position: relative;
                  vertical-align: middle;
                  pointer-events: auto;
                  &:hover{
                    // border: 1px solid #5874d8;
                  }
                  :deep(.arco-input-wrapper .arco-input.arco-input-size-mini){
                    // width: 100%;
                    // height: 100%;
                    // padding: 0;
                    border: 0;
                    text-align: center;
                  }
                }
              }
            }
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
            cursor: pointer;
            opacity: 0;
            transition: opacity 0.1s cubic-bezier(0, 0, 1, 1);
        }
      }
      .upload-picture-card{
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        .upload-picture-card-text{
           text-align: center;
           color:  var(--color-neutral-5);
        }
      }
    }
  }
</style>

