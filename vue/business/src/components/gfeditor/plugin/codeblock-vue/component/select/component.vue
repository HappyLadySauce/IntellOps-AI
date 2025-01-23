<template>
    <a-select
    allow-search
    size="small"
    style="min-width: 128px"
    :default-value="defaultValue"
    :popup-container="getContainer"
    @change="hangSelect"
    >
        <a-option
        v-for="item in modeDatas"
        :label="item.name"
        :value="item.value"
        >
            {{item.name}}
        </a-option>
    </a-select>
</template>
<script lang="ts">
import { defineComponent, PropType } from 'vue'
import { Select } from  '@arco-design/web-vue';
export default defineComponent({
    name:"gf-codeblock-select",
    components:{
        ASelect:Select,
        AOption:Select.Option,
    },
    props:{
        modeDatas: Array as PropType<Array<{value: string, syntax: string, name: string}>>,
        defaultValue:String,
        getContainer:HTMLElement,
        onSelect:Function
    },
    setup(props){
       const hangSelect=(value:any)=>{
            if (props.onSelect) props.onSelect(value);
        }
        return {
            hangSelect
        }
    }
})
</script>