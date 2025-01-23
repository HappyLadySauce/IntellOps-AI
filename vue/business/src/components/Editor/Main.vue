<template>
    <div class="editor-wrapper" :style="{height: `${minHeight}px`}">
      <gf-toolbar v-if="engine" :engine="engine" :items="items" />
      <div class="editor-container" :style="{height: `${minHeight-37}px`}">
        <div class="editor-content">
          <div ref="container" @click="handleFocus"></div>
        </div>
      </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, onUnmounted, ref,nextTick } from "vue";
import { IconSync } from '@arco-design/web-vue/es/icon';
import Engine, {EngineInterface,isMobile}  from '@/components/gfeditor/emain';
import GfToolbar from '@/components/gfeditor/toolbar';
import { cards, plugins, pluginConfig, onLoad } from "./config";
export default defineComponent({
  name: "engine-edit",
  components: {
    GfToolbar,
    IconSync,
  },
  props:{
    docData: {
      type: String,
      default: ""
    },
    placeholder: {
      type: String,
      default: "开始编辑"
    },
    minHeight: {
      type: Number,
      default: 320
    },
    valueTyle: {//获取内容内getValue=获取原生值，getHtml=用于展示html
      type: String,
      default: "getHtml"
    },
  },
  emits: ['updata',"doclick"],
  data() {
    // toolbar 配置项
    return {
      items: [
        ["collapse"],
        ["undo", "redo", "paintformat", "removeformat"],
        ["heading", "fontfamily", "fontsize"],
        ["fontcolor", "backcolor"],
        ["alignment","unorderedlist", "orderedlist", "indent", "line-height"],
        ["link", "quote","bold", "italic", "strikethrough", "underline"],
      ],
    };
  },
  setup(props, { emit }) {
    // 编辑器容器
    const container = ref<HTMLElement | null>(null);
    // 编辑器引擎
    const engine = ref<EngineInterface | null>(null);
    // 当前所有协作用户
    const members = ref([]);
    // 默认设置为当前在加载中
    const loading = ref(true);
    const upLoading = ref(false);
    onMounted(() => {
      // 容器加载后实例化编辑器引擎
      if (container.value) {
        //实例化引擎
        const engineInstance = new Engine(container.value, {
          // 启用的插件
          plugins,
          // 启用的卡片
          cards,
          autoPrepend:false,
          lazyRender:false,
          placeholder:props.placeholder,
          // 所有的卡片配置
          config: pluginConfig,
        });
        onLoad(engineInstance);
        engineInstance.setHtml("", () => {
          loading.value = false;
        });
        //只用编辑编辑器后才监听修改
        engineInstance.on("focus", () => {
            // 监听编辑器值改变事件
            engineInstance.on("change", () => {
              if(props.valueTyle=="getValue"){
                emit('updata',engineInstance.getValue());
              }else{
                emit('updata',engineInstance.getHtml());
              }
            });
        });
        engine.value = engineInstance;
      }
        //监听键盘事件
        document.addEventListener("keydown", (e)=> {
          if(e.keyCode == 83 && (navigator.platform.match('Mac') ? e.metaKey : e.ctrlKey)){
                e.preventDefault();
            }
        })
    });
    onUnmounted(() => {
      if (engine.value) engine.value.destroy();
    });
    //初始化值
    const setVal=(val:string)=>{
      if (engine.value) {
        if(props.valueTyle=="getValue"){
            engine.value.setValue(val, () => {
              loading.value = false;
            });
        }else{
          engine.value.setHtml(val, () => {
            loading.value = false;
          });
        }
      }
    }
    //锁定
    const handleFocus=()=>{
      emit("doclick")
    }
    return {
      loading,
      isMobile,
      container,
      engine,
      members,
      upLoading,
      setVal,
      handleFocus,
    };
  },
});
</script>
<style lang="less" scoped>
.editor-wrapper {
  width: 100%;
  height: 100%;
  overflow: hidden;
}
.editor-toolbar {
  width: 100%;
  background-color: var(--color-menu-light-bg);
  box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.02);
  z-index: 210;
}

.editor-container {
  width: 100%;
  height: 100%;
  margin: 0 auto;
  overflow: auto;
  position: relative;
}

.editor-content {
  position: relative;
  width: 100%;
  height: 100%;
  margin: 0 auto;
  background-color: var(---color-bg-1);
}

.editor-content .am-engine {
  padding: 20px 30px 20px 30px;
  background-color: var(--color-menu-light-bg);
  color:	var(--color-neutral-10);
  height: 100%;
  width: 100%;
}
.am-engine-view h1, .am-engine h1, .am-engine-view h2, .am-engine h2, .am-engine-view h3, .am-engine h3, .am-engine-view h4, .am-engine h4, .am-engine-view h5, .am-engine h5, .am-engine-view h6, .am-engine h6{
  color:	var(--color-neutral-10);
}
.editor-toolbar-group{
  border-left: 1px solid var(--color-neutral-2);
}
</style>
