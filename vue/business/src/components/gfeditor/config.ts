// import md5 from 'md5'
// import { getToken } from '@/utils/auth';
//展示不用这个配置
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
import File, { FileComponent, FileUploader } from "@/components/gfeditor/plugin/file";
import LineHeight from "@/components/gfeditor/plugin/line-height";
import Link from "@/components/gfeditor/plugin/link-vue";
import Hr from "@/components/gfeditor/plugin/hr";
import PaintFormat from "@/components/gfeditor/plugin/paintformat";
import Quote from "@/components/gfeditor/plugin/quote";
import Redo from "@/components/gfeditor/plugin/redo";
import Removeformat from "@/components/gfeditor/plugin/removeformat";
import Strikethrough from "@/components/gfeditor/plugin/strikethrough";
import Underline from "@/components/gfeditor/plugin/underline";
import Undo from "@/components/gfeditor/plugin/undo";
import Unorderedlist from "@/components/gfeditor/plugin/unorderedlist";
import Orderedlist from "@/components/gfeditor/plugin/orderedlist";
import {
  ToolbarPlugin,
  ToolbarComponent,
  fontFamilyDefaultData,
} from '@/components/gfeditor/toolbar';

//获取配置中的上传文件路径

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
  File,
  FileUploader,
  LineHeight,
  Link,
  Hr,
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
];

export const cards: Array<CardEntry> = [
  ToolbarComponent,
  ImageComponent,
  FileComponent,
];
let engine: EngineInterface | null = null;

export const onLoad = (e: EngineInterface) => {
  engine = e;
};
const DOMAIN =import.meta.env.VITE_APP_ENV=="production"? window?.globalConfig.Upload_url: window?.globalConfig.Upload_url_dev;
export const pluginConfig: { [key: string]: PluginOptions } = {
  [Italic.pluginName]: {
    // 默认为 _ 下划线，这里修改为单个 * 号
    markdown: "*",
  },
  //加载图片
  [Image.pluginName]: {
    onBeforeRender: (status: string, url: string) => {
      if (url.startsWith("data:image/")) return url;
      return url + `?_t=${Date.parse(new Date().toString())/1000}`;
    },
  },
  [ImageUploader.pluginName]: {
    file: {
      action: `${DOMAIN}/upload/image`,
      headers:()=>{
        // const token = getToken();
        // const API_SECRET =ApiSecret; //从config读取的配置-打包后可以修改-变化
        const timestamp: number = Date.parse(new Date().toString())/1000;
        return {
          "verify-time":timestamp,
        // "verify-encrypt":md5(API_SECRET+timestamp),Authorization:token,
      }
      },
    },
    remote: {
      action: `${DOMAIN}/upload/image`,
    },
    isRemote: (src: string) => src.indexOf(DOMAIN) < 0,
  },
  //附件
  [FileUploader.pluginName]: {
    action: `${DOMAIN}/upload/file`,
    headers:()=>{
      // const token = getToken();
      // const API_SECRET =ApiSecret; //从config读取的配置-打包后可以修改-变化
      const timestamp: number = Date.parse(new Date().toString())/1000;
      return {
        "verify-time":timestamp,
        // "verify-encrypt":md5(API_SECRET+timestamp),Authorization:token
      }
    },
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
