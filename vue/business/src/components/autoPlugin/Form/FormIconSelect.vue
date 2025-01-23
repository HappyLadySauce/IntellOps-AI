<template>
    <Icon :icon="icon+(modelValue==val?'-select':'')" size="35" @click="selectIcon"></Icon>
</template>

<script lang="ts" setup>
import { Icon} from '@/components/Icon';
const props = defineProps({
  icon: String,
  val: Number,
  modelValue: {
    type: Number,
    required: true
  },
})
const emits = defineEmits(['update:modelValue'])
//选择图标
const selectIcon=()=>{
      emits('update:modelValue', props.val)
}
</script>
