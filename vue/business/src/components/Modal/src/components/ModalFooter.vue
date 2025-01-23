<template>
  <div>
    <slot name="insertFooter"></slot>
    <a-button  @click="handleCancel" v-if="showCancelBtn">
      {{ cancelText }}
    </a-button>
    <slot name="centerFooter"></slot>
    <a-button
      @click="handleOk"
      :loading="confirmLoading"
      v-if="showOkBtn"
      :type="okType"
      style="margin-left: 8px;"
    >
      {{ okText }}
    </a-button>
    <slot name="appendFooter"></slot>
  </div>
</template>
<script lang="ts">
  import { defineComponent } from 'vue';

  import { basicProps } from '../props';
  export default defineComponent({
    name: 'BasicModalFooter',
    props: basicProps,
    emits: ['ok', 'cancel'],
    setup(_, { emit }) {
      function handleOk(e: Event) {
        emit('ok', e);
      }
      function handleCancel(e: Event) {
        emit('cancel', e);
      }

      return { handleOk, handleCancel };
    },
  });
</script>
