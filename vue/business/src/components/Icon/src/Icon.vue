<template>
  <SvgIcon
    :size="size"
    :icon="icon"
    v-if="isSvgIcon"
    :spin="spin"
    :color="color"
  />
  <component v-else :is="getIconRef" :spin="spin" :size="size" :style="getWrapStyle"></component>
</template>
<script lang="ts">
  import type { PropType } from 'vue';
  import {
    defineComponent,
    ref,
    computed,
    CSSProperties,
  } from 'vue';
  import SvgIcon from './SvgIcon.vue';
  import { isString } from '/@/utils/is';
  import { propTypes } from '/@/utils/propTypes';

  const SVG_END_WITH_FLAG = 'svg';
  export default defineComponent({
    name: 'Icon',
    components: { SvgIcon },
    props: {
      // icon name
      icon: propTypes.string,
      // icon color
      color: propTypes.string,
      // icon size
      size: {
        type: [String, Number] as PropType<string | number>,
        default: 18,
      },
      spin: propTypes.bool.def(false),
      prefix: propTypes.string.def(''),
    },
    setup(props) {
      const elRef = ref<ElRef>(null);

      const isSvgIcon = computed(() => props.icon?.startsWith(SVG_END_WITH_FLAG));
      const getIconRef = computed(() => `${props.prefix ? props.prefix + ':' : ''}${props.icon}`);

      const getWrapStyle = computed((): CSSProperties => {
        const { size, color } = props;
        let fs = size;
        if (isString(size)) {
          fs = parseInt(size, 10);
        }
        return {
          fontSize: `${fs}px`,
          color: color,
          display: 'inline-flex',
        };
      });

      return { elRef, getWrapStyle, isSvgIcon,getIconRef };
    },
  });
</script>
<style lang="less">
  .app-iconify {
    display: inline-block;
    &-spin {
      svg {
        animation: loadingCircle 1s infinite linear;
      }
    }
  }
  span.iconify {
    display: block;
    min-width: 1em;
    min-height: 1em;
    border-radius: 100%;
  }
</style>
