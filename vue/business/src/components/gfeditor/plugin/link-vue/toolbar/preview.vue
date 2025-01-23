<template>
    <div class="data-link-preview" data-element="ui">
        <a-tooltip :content="openTitle">
            <a
            class="data-icon data-icon-link data-link-preview-open"
            :href="href"
            target="_blank"
            >
                {{href}}
            </a>
        </a-tooltip>
        <div v-if="!readonly" class="data-link-op">
            <a-tooltip :content="editTitle" v-if="!isMobile">
                <a class="data-icon data-icon-edit" @click="onEdit" />
            </a-tooltip>
            <a v-if="isMobile" class="data-icon data-icon-edit" @click="onEdit" />
            <a-tooltip :content="removeTitle" v-if="!isMobile">
                <a class="data-icon data-icon-unlink" @click="onRemove" />
            </a-tooltip>
            <a v-if="isMobile" class="data-icon data-icon-unlink" @click="onRemove" />
        </div>
    </div>
</template>
<script lang="ts">
import { defineComponent, onMounted, PropType } from 'vue'
import { LanguageInterface, isMobile } from '@/components/gfeditor/emain'
import { Tooltip } from  '@arco-design/web-vue';

export default defineComponent({
    name:"am-link-preview",
    components:{
        ATooltip:Tooltip
    },
    props:{
        language:{
            type:Object as PropType<LanguageInterface>,
            required:true
        } as const,
        readonly: {
            type: Boolean,
            required:true
        } as const,
        href:String,
        className:String,
        onEdit:Function as PropType<(event:MouseEvent) => void>,
        onRemove:Function as PropType<(event:MouseEvent) => void>,
        onLoad:Function as PropType<() => void>,
    },
    setup(props){
        const openTitle = props.language.get<string>('link', 'link_open')
        const editTitle = props.language.get<string>('link', 'link_edit')
        const removeTitle = props.language.get<string>('link', 'link_remove')

        onMounted(() => {
            if(props.onLoad) props.onLoad()
        })

        return {
            openTitle,
            editTitle,
            removeTitle,
            isMobile
        }
    }
})
</script>