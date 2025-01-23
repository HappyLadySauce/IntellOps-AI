import md5 from 'md5'
import { getToken } from '@/utils/auth';
import {
  PluginEntry,
  CardEntry,
  PluginOptions,
  NodeInterface,
  RangeInterface,
  EngineInterface,
} from '@/components/gfeditor/emain';
//引入插件 
import Alignment from "@/components/gfeditor/plugin/alignment";
import Backcolor from "@/components/gfeditor/plugin/backcolor";
import Bold from "@/components/gfeditor/plugin/bold";
import Code from "@/components/gfeditor/plugin/code";
import Fontcolor from "@/components/gfeditor/plugin/fontcolor";
import Fontfamily from "@/components/gfeditor/plugin/fontfamily";
import Fontsize from "@/components/gfeditor/plugin/fontsize";
import Heading from "@/components/gfeditor/plugin/heading";
import Indent from "@/components/gfeditor/plugin/indent";
import Italic from "@/components/gfeditor/plugin/italic";
import Image, { ImageComponent, ImageUploader } from "@/components/gfeditor/plugin/image";
import imgattachment from "@/components/gfeditor/plugin/imgattachment";
import File, { FileComponent, FileUploader } from "@/components/gfeditor/plugin/file";
import Video, { VideoComponent, VideoUploader } from "@/components/gfeditor/plugin/video";
import Table, { TableComponent } from "@/components/gfeditor/plugin/table";
import LineHeight from "@/components/gfeditor/plugin/line-height";
import Link from "@/components/gfeditor/plugin/link-vue";
import Hr , { HrComponent }from "@/components/gfeditor/plugin/hr";
import Codeblock, { CodeBlockComponent } from  "@/components/gfeditor/plugin/codeblock-vue";
import PaintFormat from "@/components/gfeditor/plugin/paintformat";
import Quote from "@/components/gfeditor/plugin/quote";
import Redo from "@/components/gfeditor/plugin/redo";
import Removeformat from "@/components/gfeditor/plugin/removeformat";
import Strikethrough from "@/components/gfeditor/plugin/strikethrough";
import Underline from "@/components/gfeditor/plugin/underline";
import Undo from "@/components/gfeditor/plugin/undo";
import Unorderedlist from "@/components/gfeditor/plugin/unorderedlist";
import Orderedlist from "@/components/gfeditor/plugin/orderedlist";
import Math, { MathComponent } from  "@/components/gfeditor/plugin/math";
import Status, { StatusComponent } from  "@/components/gfeditor/plugin/status";
import {
  ToolbarPlugin,
  ToolbarComponent,
  fontFamilyDefaultData,
} from '@/components/gfeditor/toolbar';
import { GetFullPath } from "@/utils/tool";
export const plugins: Array<PluginEntry> = [
  Alignment,
  Backcolor,
  Bold,
  Code,
  Fontcolor,
  Fontfamily,
  Fontsize,
  Heading,
  Indent,
  Italic,
  Image,
  ImageUploader,
  imgattachment,
  File,
  FileUploader,
  Video,
  VideoUploader,
  Table,
  LineHeight,
  Link,
  Hr,
  Codeblock,
  PaintFormat,
  Quote,
  Redo,
  Removeformat,
  Strikethrough,
  Undo,
  Underline,
  Orderedlist,
  Unorderedlist,
  ToolbarPlugin,
  Math,
  Status,
];

export const cards: Array<CardEntry> = [
  HrComponent,
  CodeBlockComponent,
  ToolbarComponent,
  ImageComponent,
  VideoComponent,
  TableComponent,
  FileComponent,
  MathComponent,
  StatusComponent,
];
let engine: EngineInterface | null = null;
export const onLoad = (e: EngineInterface) => {
  engine = e;
};
//获取配置中的上传文件路径
const DOMAIN =import.meta.env.VITE_APP_ENV=="production"? window?.globalConfig.Upload_url: window?.globalConfig.Upload_url_dev;
const Root_url = import.meta.env.VITE_APP_ENV=="production"? window?.globalConfig.Root_url: window?.globalConfig.Root_url_dev;//前缀
export const pluginConfig: { [key: string]: PluginOptions } = {
  [Italic.pluginName]: {
    // 默认为 _ 下划线，这里修改为单个 * 号
    markdown: "*",
  },
  [Math.pluginName]: {
    action: `${Root_url}/common/api/latex`,
    headers:()=>{
      const token = getToken();
      const timestamp: number = Date.parse(new Date().toString())/1000;
      const passstr=md5(import.meta.env.VITE_ENCRYPT+timestamp) as any
      return {"apiverify":window.btoa(passstr+"#"+timestamp),Authorization:token}
    },
    parse: (res: any) => {
      if (res.code==0) return { result: true, data: res.data};
      return { result: false };
    },
  },
  //加载图片
  [Image.pluginName]: {
    onBeforeRender: (status: string, url: string) => {
      if (url.startsWith("data:image/")) return url;
      return GetFullPath(url) ;
    },
  },
  //上传图片
  [ImageUploader.pluginName]: {
    file: {
      action: `${Root_url}/datacenter/upfile/upload`,
      data:{
        filetype:"image"
      },
      headers:()=>{
        const token = getToken();
        const timestamp: number = Date.parse(new Date().toString())/1000;
        const passstr=md5(import.meta.env.VITE_ENCRYPT+timestamp) as any
        return {"apiverify":window.btoa(passstr+"#"+timestamp),Authorization:token}
      },
    },
    remote: {
      action: `${DOMAIN}/upload/thirdImage`,
      headers:()=>{
        const token = getToken();
        const timestamp: number = Date.parse(new Date().toString())/1000;
        const passstr=md5(import.meta.env.VITE_ENCRYPT+timestamp) as any
        return {"apiverify":window.btoa(passstr+"#"+timestamp),Authorization:token}
      },
    },
    isRemote: (src: string) => src.indexOf(DOMAIN) < 0,
  },
  //附件
  [FileUploader.pluginName]: {
    action: `${DOMAIN}/upload/file`,
    headers:()=>{
      const token = getToken();
      const timestamp: number = Date.parse(new Date().toString())/1000;
      const passstr=md5(import.meta.env.VITE_ENCRYPT+timestamp) as any
      return {"apiverify":window.btoa(passstr+"#"+timestamp),Authorization:token}
    },
  },
  //加载视频
  [Video.pluginName]: {
    onBeforeRender: (status: string, url: string) => {
      return GetFullPath(url);
    },
    showTitle:false,
  },
  //上传视频
  [VideoUploader.pluginName]: {
    action: `${Root_url}/datacenter/upfile/upload`,
    data:{
      filetype:"video"
    },
    headers:()=>{
      const token = getToken();
      const timestamp: number = Date.parse(new Date().toString())/1000;
      const passstr=md5(import.meta.env.VITE_ENCRYPT+timestamp) as any
      return {"apiverify":window.btoa(passstr+"#"+timestamp),Authorization:token}
    },
    limitSize: 1024 * 1024 * 50,
  },
  [Fontfamily.pluginName]: {
    //配置粘贴后需要过滤的字体
    filter: (fontfamily: string) => {
      const item = fontFamilyDefaultData.find((item) =>
        fontfamily
          .split(",")
          .some(
            (name) =>
              item.value
                .toLowerCase()
                .indexOf(name.replace(/"/, "").toLowerCase()) > -1
          )
      );
      return item ? item.value : false;
    },
  },
  [Fontsize.pluginName]: {
    //配置粘贴后需要过滤的字体大小
    filter: (fontSize: string) => {
      return (
        [
          "12px",
          "13px",
          "14px",
          "15px",
          "16px",
          "19px",
          "22px",
          "24px",
          "29px",
          "32px",
          "40px",
          "48px",
        ].indexOf(fontSize) > -1
      );
    },
  },
  [LineHeight.pluginName]: {
    //配置粘贴后需要过滤的行高
    filter: (lineHeight: string) => {
      if (lineHeight === "14px") return "1";
      if (lineHeight === "16px") return "1.15";
      if (lineHeight === "21px") return "1.5";
      if (lineHeight === "28px") return "2";
      if (lineHeight === "35px") return "2.5";
      if (lineHeight === "42px") return "3";
      // 不满足条件就移除掉
      return ["1", "1.15", "1.5", "2", "2.5", "3"].indexOf(lineHeight) > -1;
    },
  },

};
