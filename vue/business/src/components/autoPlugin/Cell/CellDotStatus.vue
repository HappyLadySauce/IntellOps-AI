<template>
    <span v-if="!dictItem"></span>
    <a-tag v-else-if="dictItem.color === 'primary'" color="arcoblue">
      <GfDot type="primary" style="width: 5px; height: 5px" />
      <span style="margin-left: 5px">{{ dictItem.label }}</span>
    </a-tag>
    <a-tag v-else-if="dictItem.color === 'success'" color="green">
      <GfDot type="success" style="width: 5px; height: 5px" />
      <span style="margin-left: 5px">{{ dictItem.label }}</span>
    </a-tag>
    <a-tag v-else-if="dictItem.color === 'warning'" color="orange">
      <GfDot type="warning" style="width: 5px; height: 5px" />
      <span style="margin-left: 5px">{{ dictItem.label }}</span>
    </a-tag>
    <a-tag v-else-if="dictItem.color === 'danger'" color="red">
      <GfDot type="danger" style="width: 5px; height: 5px" />
      <span style="margin-left: 5px">{{ dictItem.label }}</span>
    </a-tag>
    <a-tag v-else>
      <GfDot type="info" style="width: 5px; height: 5px" />
      <span style="margin-left: 5px">{{ dictItem.label }}</span>
    </a-tag>
</template>

<script lang="ts" setup>
import { computed } from 'vue';
import type { LabelValueState } from '@/types/global'
//状态  'primary' | 'success' | 'warning' | 'danger' | 'info'
const props = defineProps({
  dict: {
    type: Array<LabelValueState>,
    required: true
  },
  value: {
    type: [Number, String],
    required: true
  }
})
const dictItem = computed(() =>
  props.dict.find((d) => d.value === String(props.value) || d.value === Number(props.value))
)
</script>
