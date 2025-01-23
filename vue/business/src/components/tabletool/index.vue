<template>
  <a-space>
  <a-button v-permission="'add'" type="primary" @click="handleCreate" v-if="inArray(showbtn,'create')">
    <template #icon>
      <icon-plus />
    </template>
    {{ $t('searchTable.operation.create') }} 
  </a-button>
  <slot></slot>
  <a-tooltip v-permission="'import'" :content="$t('searchTable.actions.import')" v-if="inArray(showbtn,'import')">
    <a-button v-permission="'import'" class="gf_hover_btn-border" @click="handleImport">
      <template #icon>
        <icon-upload />
      </template>
    </a-button>
  </a-tooltip>
  <a-tooltip :content="$t('searchTable.actions.export')" v-if="inArray(showbtn,'export')">
    <a-button class="gf_hover_btn-border" @click="handleRxport" >
      <template #icon><icon-download /></template>
    </a-button>
  </a-tooltip>
  <a-tooltip :content="$t('searchTable.actions.refresh')" v-if="inArray(showbtn,'refresh')">
    <a-button class="gf_hover_btn-border" @click="handleRefresh" >
      <template #icon><icon-refresh /></template>
    </a-button>
  </a-tooltip>
  <a-dropdown @select="handleSelectDensity" v-if="inArray(showbtn,'selectdensity')">
    <a-tooltip :content="$t('searchTable.actions.density')">
      <a-button class="gf_hover_btn-border" >
        <template #icon><icon-line-height size="18" /></template>
      </a-button>
    </a-tooltip>
    <template #content>
      <a-doption
        v-for="item in densityList"
        :key="item.value"
        :value="item.value"
        :class="{ active: item.value === size }"
      >
        <span>{{ item.name }}</span>
      </a-doption>
    </template>
  </a-dropdown>
  <!--列设置-->
  <a-popover v-if="inArray(showbtn,'settingcolumn')" trigger="click" position="br"
      :content-style="{ minWidth: '120px', padding: '6px 8px 10px' }">
      <a-tooltip :content="$t('searchTable.actions.columnSetting')">
        <a-button class="gf_hover_btn-border">
          <template #icon>
            <icon-settings />
          </template>
        </a-button>
      </a-tooltip>
      <template #content>
        <div class="gf-table-draggable">
          <VueDraggable v-model="settingColumnList">
            <div v-for="item in settingColumnList" :key="item.title" class="drag-item">
              <div class="drag-item-move"><icon-drag-dot-vertical /></div>
              <a-checkbox v-model:model-value="item.show" :disabled="item.disabled">{{ item.title }}</a-checkbox>
            </div>
          </VueDraggable>
        </div>
        <a-divider :margin="6" />
        <a-row justify="center">
          <a-button type="primary" size="mini" long @click="resetSettingColumns">
            <template #icon><icon-refresh /></template>
            <template #default>{{ $t("searchTable.form.reset") }}</template>
          </a-button>
        </a-row>
      </template>
    </a-popover>

   <a-tooltip :content="isFullscreen?$t('searchTable.actions.exit'):$t('searchTable.actions.fullscreen')" v-if="inArray(showbtn,'fullscreen')">
    <a-button class="gf_hover_btn-border" @click="handleFullscreen" >
      <template #icon>
        <icon-fullscreen v-if="!isFullscreen" />
        <icon-fullscreen-exit v-else />
      </template>
    </a-button>
  </a-tooltip>
  <ImportDrawer v-model="showImportDrawer" ><!--导入数据配置--></ImportDrawer>
</a-space>
</template>

<script lang="ts" setup>
  import { ref, computed,watch } from 'vue';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import { VueDraggable } from 'vue-draggable-plus'
  import { useI18n } from 'vue-i18n';
  import { SizeProps } from './data';
  import { inArray } from '@/utils/is';
  import  ImportDrawer from './ImportDrawer.vue';
  const showImportDrawer=ref(false)
  const { t } = useI18n();
  const props = defineProps({
    showbtn: {
      type:  Array,
      required: true
    },
    columns: {
      type:  Array,
      required: false
    },
    disabledColumnKeys: {//禁止操作
      type:  Array,
      default: []
    },
  })
  const emits = defineEmits(['create','export','refresh','selectdensity','fullscreen',"settingcolumn"])
  const isFullscreen = ref(false);
  const size = ref<SizeProps>('large');
  const densityList = computed(() => [
    {
      name: t('searchTable.size.mini'),
      value: 'mini',
    },
    {
      name: t('searchTable.size.small'),
      value: 'small',
    },
    {
      name: t('searchTable.size.medium'),
      value: 'medium',
    },
    {
      name: t('searchTable.size.large'),
      value: 'large',
    },
  ]);
  const handleCreate=()=>{
    emits("create")
  }
  //导入
  const handleImport=()=>{
    showImportDrawer.value=true
  }
  //导出
  const handleRxport=()=>{
    emits("export")
  }
  const handleRefresh=()=>{
    emits("refresh")
  }
  //列表密度
  const handleSelectDensity = (
    val: string | number | Record<string, any> | undefined,
    e: Event
  ) => {
    size.value = val as SizeProps;
    emits("selectdensity",size.value)
  };
  const handleFullscreen=()=>{
    isFullscreen.value = !isFullscreen.value
    emits("fullscreen",isFullscreen.value)
  }
  //列设置
  type SettingColumnItem = { title: string, key: string, show: boolean, disabled: boolean }
  const settingColumnList = ref<SettingColumnItem[]>([])
  // 重置配置列
  const resetSettingColumns = () => {
    settingColumnList.value = []
    if (props.columns) {
      const arr = props.columns as TableColumnData[]
      arr.forEach((item:any) => {
        settingColumnList.value.push({
          key: item.dataIndex || (typeof item.title === 'string' ? item.title : ''),
          title: typeof item.title === 'string' ? item.title : '',
          show: item.show ?? true,
          disabled: props.disabledColumnKeys.includes(
            item.dataIndex || (typeof item.title === 'string' ? item.title : '')
          )
        })
      })
    }
  }
  // 排序和过滤可显示的列数据
const filterShowcolumns = () => {
  if (!props.columns) return
    const arr = props.columns as TableColumnData[]
    // 显示的dataIndex
    const showDataIndexs = settingColumnList.value
      .filter((i) => i.show)
      .map((i) => i.key || (typeof i.title === 'string' ? i.title : ''))
    // 显示的columns数据
    const filterColumns = arr.filter((i) =>
      showDataIndexs.includes(i.dataIndex || (typeof i.title === 'string' ? i.title : ''))
    )
    const sortedColumns: TableColumnData[] = []
    settingColumnList.value.forEach((i) => {
      filterColumns.forEach((j) => {
        if (i.key === (j.dataIndex || j.title)) {
          sortedColumns.push(j)
        }
      })
    })
    emits("settingcolumn",sortedColumns)
  }
  watch(
    () => props.columns,
    () => {
      resetSettingColumns()
    },
    { immediate: true, deep: true}//如果我们的需求是需要在初始化绑定数据的时候就执行，就可以将watch函数的immediate属性值设为true,反之设为false
  )
  //监听列表变化
  watch(
    () => settingColumnList.value,
    () => {
      filterShowcolumns()
    },
    { deep: true}//如果我们的需求是需要在初始化绑定数据的时候就执行，就可以将watch函数的immediate属性值设为true,反之设为false
  )
</script>

<style scoped lang="less">
.gf-table-draggable{
  padding: 1px 5px 1px 0; 
  max-height: 250px;
  box-sizing: border-box;
  overflow: hidden;
  overflow-y: auto;
  .drag-item {
    display: flex;
    align-items: center;
    cursor: pointer;
    &:hover {
      background-color: var(--color-fill-2);
    }
    .drag-item-move{
      padding-left: 2px;
      padding-right: 2px;
      cursor: move;
    }
    :deep(.arco-checkbox) {
      width: 100%;
      font-size: 12px;
      .arco-checkbox-icon {
        width: 14px;
        height: 14px;
      }
    }
  }
}
</style>
