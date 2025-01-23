<template>
  <ul class="gi-option">
    <slot></slot>
  </ul>
</template>

<script setup lang="ts">
defineOptions({ name: 'GiOption' })
</script>

<style lang="less" scoped>
.gi-option {
  width: 100%;
  min-width: 120px;
  padding: 0;
  margin: 0;
}
</style>
