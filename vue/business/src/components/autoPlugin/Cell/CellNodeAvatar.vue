<!-- 背景框头像 -->
<template>
  <div class="avatar" :style="AvatarStyle">
    <a-avatar :size="size" class="icon" :image-url="GetFullPath(avatar)" v-if="avatar" />
    <a-avatar :size="size" class="icon" v-else> <icon-user /> </a-avatar>
    <div class="name" :style="NameStyle" v-if="name">{{ name }}</div>
  </div>
</template>

<script setup>
import { computed,  onBeforeMount } from "vue";
import { IconUser } from "@arco-design/web-vue/es/icon";
import { GetFullPath } from "@/utils/tool";

const props = defineProps({
  avatar: { type: String, default: "" },
  size: { type: Number, default: 22 },
  id: { type: String, default: "" },
  name: { type: Boolean,required: true, default: "" },
});

let AvatarStyle = computed(() => {
  let padding = props.name ? 4 : 0;
  return {
    height: props.size + padding * 2 + "px",
    borderRadius: (props.size + padding * 2) / 2 + "px",
    padding: padding + "px",
  };
});

let NameStyle = computed(() => {
  return {
    fontSize: props.size / 2 + 1 + "px",
  };
});

onBeforeMount(() => {
});
</script>

<style lang="less" scoped>
.avatar {
  background: var(--color-neutral-2);
  display: flex;
  align-items: center;
  width: fit-content;

  .icon {
    overflow: hidden;
    flex-shrink: 0;
  }

  .name {
    user-select: none;
    font-family: PingFangSC-Regular, PingFang SC;
    color: var(--color-neutral-10);
    margin: 0 4px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    max-width: 64px;
  }
}
</style>
