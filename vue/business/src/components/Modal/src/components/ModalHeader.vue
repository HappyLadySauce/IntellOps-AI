<template>
    <div class="selftitle" @dblclick="handleDblclick">
      <div class="left">
        <BasicTitle :helpMessage="helpMessage"> {{ title }} </BasicTitle>
      </div>
      <div class="right">
        <span class="btnicon" @click="handleDblclick">
          <icon-fullscreen-exit  v-if="isFull" :style="{fontSize:'12px'}"/>
          <icon-fullscreen  v-else :style="{fontSize:'12px'}"/>
        </span>
      </div>
    </div>
</template>
<script lang="ts">
  import type { PropType } from 'vue';
  import { defineComponent,ref ,unref,watch} from 'vue';
  import { BasicTitle } from '/@/components/Basic';
  export default defineComponent({
    name: 'BasicModalHeader',
    components: {  BasicTitle  },
    props: {
      helpMessage: {
        type: [String, Array] as PropType<string | string[]>,
      },
      title: { type: String },
      fullscreenVal: {type: Boolean, default: false },
    },
    emits: ['dblclick','fullscreen','fullexit'],
    setup(props, { emit }) {
      const isFull=ref(props.fullscreenVal)
      function handleDblclick(e: Event){
        isFull.value=!isFull.value
        emit('dblclick', unref(isFull));
      }
      watch(() => props.fullscreenVal,(newValue) => {
        isFull.value=newValue
      })
      return {isFull,handleDblclick };
    },
  });
</script>
<style lang="less" scoped>
 .selftitle{
  display: flex;
  width: 100%;
  .left{
    flex: 1;
  }
  .right{
    margin-right: 22px;
    color: var(--color-text-1);
    cursor: pointer;
    height: 100%;
    .btnicon{
      position: relative;
      display: inline-block;
      cursor: pointer;
      margin-top: 2px;
      width: 20px;
      height: 20px;
      display: flex;
      justify-content: center;
      align-items: center;
      // &::before {
      //     top: 50%;
      //     left: 50%;
      //     width: 20px;
      //     height: 20px;
      //     transform: translate(-50%, -50%);
      // }
      // &::before {
      //     position: absolute;
      //     display: block;
      //     box-sizing: border-box;
      //     background-color: transparent;
      //     border-radius: var(--border-radius-circle);
      //     transition: background-color 0.1s cubic-bezier(0, 0, 1, 1);
      //     content: '';
      // }
      .arco-icon{
        position: relative;
      }
      &:hover{
        border-radius: 100%;
        background-color: var(--color-neutral-2);
        width: 20px;
        height: 20px;
      }
    }
  }
 }
</style>
