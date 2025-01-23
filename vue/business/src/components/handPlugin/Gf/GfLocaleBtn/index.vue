<template>
  <div>
   <a-dropdown trigger="click" @select="changeLocale as any">
     <a-button size="mini" class="gf_hover_btn" >
      <template #icon>
        <icon-language :size="18" />
      </template>
    </a-button>
    <template #content>
      <a-doption
        v-for="item in locales"
        :key="item.value"
        :value="item.value"
      >
        <template #icon>
          <icon-check v-show="item.value === currentLocale" />
        </template>
        {{ item.label }}
      </a-doption>
  </template>
  </a-dropdown>
</div>
</template>

<script setup lang="ts">
  import { LOCALE_OPTIONS } from '@/locale';
  import useLocale from '@/hooks/locale';
  const { changeLocale, currentLocale } = useLocale();
  const locales = [...LOCALE_OPTIONS];

</script>
