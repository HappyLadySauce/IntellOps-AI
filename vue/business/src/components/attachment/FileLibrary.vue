<template>
     <div class="file_Library">
        <div class="header_toopl">
            <div class="btnlist">
                <a-radio-group v-model="typeid" type="button" size="large" @change="changeMeterType">
                <a-radio :value="0">素材图</a-radio>
                <a-radio :value="1">插图</a-radio>
                </a-radio-group>
            </div>
            <div class="search">
                <a-input-search v-model="searchword" :style="{width:'300px'}" @press-enter="onSearchMatter" allow-clear @search="onSearchMatter" placeholder="搜索素材名称"/>
            </div>
        </div>
        <!--分类-->
        <div class="catelist" v-if="typeid==0">
            <a-scrollbar style="height:52px;overflow: auto;width: 100%;" >
                <ul class="cateul">
                    <li class="cate_item " :class="{select:cid==0}" @click="selectCate(0)">全部</li>
                    <template v-for="item in picturecate">
                    <li class="cate_item " v-if="item.type==0||item.type==2" :class="{select:item.id==cid}" @click="selectCate(item.id)">{{item.name}}</li>
                    </template>
                </ul>
            </a-scrollbar>
        </div>
        <div class="catelist" v-if="typeid==1">
            <a-scrollbar style="height:52px;overflow: auto;width: 100%;" >
             <ul class="cateul">
                <li class="cate_item " :class="{select:cid==0}" @click="selectCate(0)">全部</li>
                <template v-for="item in picturecate">
                  <li class="cate_item " v-if="item.type==1||item.type==2" :class="{select:item.id==cid}" @click="selectCate(item.id)">{{item.name}}</li>
                </template>
            </ul>
        </a-scrollbar>
        </div>
        <!--图片库-->
        <a-spin :loading="loading" tip="加载中..."  style=" width: 100%;" >
        <div class="file_list" >
            <a-scrollbar height="435px" v-if="picture&&picture.length>0">
            <div class="file_list_box">
                <div class="file_item" v-for="item in picture" >
                <div   >
                    <div class="imgbox" @click="onSelectImg(item)">
                    <!--图片-->
                    <div class="img_raw">
                        <div class="img_preview">
                          <img :src="GetFullPath(item.url)" class="img_item"/>
                        </div>
                    </div>
                    <div class="edit_layer "  :class="{l_active:selectList.indexOf(`mt_${item.id}`)>-1}" >
                        <Icon icon="svgfont-mark_sel" class="choose" :size="18" ></Icon>
                    </div>
                    </div>
                </div>
                </div>
            </div>
            <div class="mypagination" v-if="pagetotall&&pagetotall>14">
                <a-pagination v-model:current="pagenum" :total="pagetotall"  @change="changePage" show-total show-jumper />
            </div>
            </a-scrollbar>
            <!--空-->
            <div class="emptylist" v-if="!loading&&(!picture||picture.length==0)" :style="{height:`435px`}">
                <div class="emptybox">
                    <a-empty   description="暂时没有该类素材，去看看其他类型吧！"/>
                </div>
            </div>
        </div>
      </a-spin>
    </div>
</template>
<script lang="ts">
import { defineComponent,reactive, toRefs, ref, unref} from 'vue';
import { BasicModal } from '/@/components/Modal';
import useLoading from '@/hooks/loading';
import { GetFullPath } from "@/utils/tool";
//api
import {get_picture,getPictureCate,CateItem ,PictureItem } from '@/api/datacenter/attachment';
import {Icon} from '@/components/Icon';
export default defineComponent({
    name: 'FileLibrary',
    components: { BasicModal ,Icon},
    emits: ['ok'],
    setup(_, { emit }) {
        const { loading, setLoading } = useLoading(true);
        const arraynct: string[] = [] // 定义字符串数组
        const catitem: CateItem[] = [] // 定义整数型数组
        const pictureitem: PictureItem[] = [] // 定义整数型数组
        const pagedata=reactive({
            searchword:"",
            typeid:0,//素材图/插图
            getnumber:"one",//是否多选 或者 单选
            selectList:arraynct,//选择
            //图库
            picturecate:catitem,
            picture:pictureitem,
            cid:0,
            //分页
            pagenum:1,
            pagetotall:0,
        })
     //更新图片类型
     function changeMeterType(){
        pagedata.searchword=""
        pagedata.cid=0
        fetchData("")
      }
    //获取分类=初始化
    const initGetPictureCate=async(getnumber:string)=>{
        pagedata.selectList=[]
        pagedata.getnumber=getnumber
        pagedata.searchword=''
         const getdbdata = await getPictureCate({});
          if(getdbdata){
            pagedata.picturecate=getdbdata
            pagedata.cid=0
            fetchData("")
          }
    }
    //获取列表数据
    const fetchData = async (searchword:string) => {
      setLoading(true);
      try {
        const getdbdata= await get_picture({searchword:searchword,type:pagedata.typeid,cid:pagedata.cid,page:pagedata.pagenum});
        if(getdbdata){
            pagedata.picture=getdbdata.items
            pagedata.pagenum=getdbdata.page
            pagedata.pagetotall=getdbdata.total
        }
      } catch (err) {
        // you can report use errorHandler or other
      } finally {
        setLoading(false);
      }
    };
    //图库搜索
    const onSearchMatter=()=>{
        if(pagedata.searchword){
            fetchData(pagedata.searchword)
        }else{
            fetchData("")
        }
    }
    //选择图片
    const onSelectImg=(item:any)=>{
        var id_str=`mt_${item.id}`
          if(pagedata.selectList.indexOf(id_str)>-1){
            pagedata.selectList=pagedata.selectList.filter((id)=>id!=id_str)
          }else{
            if(pagedata.getnumber=="one"){
              pagedata.selectList=[id_str]
            }else{
              pagedata.selectList.push(id_str)
            }
          }
          emit('ok',pagedata.selectList,pagedata.picture);
    }
     //选择分类
     const selectCate=(id:number)=>{
        pagedata.cid=id
        fetchData("")
      }
    //分页
    const changePage=(page:number)=>{
        pagedata.pagenum=page
        fetchData("")
      }
    return { 
        ...toRefs(pagedata),
        loading,
        changeMeterType,onSearchMatter,
        onSelectImg,selectCate,changePage,
        initGetPictureCate,GetFullPath,
    };
    },
});
</script>
<style lang="less" scoped>
 //图片库
 .file_Library{
    padding: 10px;
        .header_toopl{
          height: 35px;
          display: flex;
          .btnlist{
            padding-right: 20px;
          }
          .seach{

          }
        }
        //分类
        .catelist{
          width: 100%;
          height: 50px;
          line-height: 50px;
          overflow: hidden;
          .cateul{
            padding: 0;
            margin: 0;
            transition: none 0s ease 0s;
            white-space: nowrap;
            .cate_item{
              // float: left;
              display: inline;
              margin-right: 28px;
              color: #343434;
              font-size: 12px;
              cursor: pointer;
              user-select: none;
              list-style: none;
              &.select,&:hover{
                  color: #557ce1;
              }
            }
          }
        }
        //文件列表
        .file_list{
          min-height: 450px;
          .file_list_box{
            padding: 0;
            margin: 0;
            overflow: hidden;
            .file_item{
              float: left;
              width: 116px;
              list-style: none;
              margin-top: 20px;
              margin-right: 34px;
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
                  height: 116px;
                  background-color: #f3f3f3;
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
                  .choose{
                    display: none;
                    position: absolute;
                    right: 0;
                    bottom: 0;
                    width: 28px;
                    height: 24px;
                  }
                }
                //选中
                .l_active{
                  display: block;
                  .choose{
                      display: block;
                    }
                }
                  //浮框
                &:hover{
                    .edit_layer{
                      display: block;
                    }
                  }
              }
            }
          }
        }
        //分页
        .mypagination{
          text-align: center;
          margin-top: 15px;
          margin-bottom: 10px;
        }
      }
</style>
  
  