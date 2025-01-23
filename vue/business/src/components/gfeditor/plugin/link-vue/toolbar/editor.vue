<template>
    <a-config-provider :auto-insert-space-in-button=false>
        <div data-element="ui" :class="['data-link-editor', className]">
            <p>{{textTitle}}</p>
            <p>
                <a-input
                class="data-link-input"
                v-model="text"
                :placeholder="textPlaceholder"
                />
            </p>
            <p>{{linkTitle}}</p>
            <p>
                <a-input
                ref="linkRef"
                class="data-link-input"
                v-model="link"
                :placeholder="linkPlaceholder"
                />
            </p>
            <p>{{juptTypeTitle}}</p>
            <p>
            <a-radio-group v-model="juptType"  >
                <a-radio value="_self">此窗口打开</a-radio>
                <a-radio value="_blank">打开新页面</a-radio>
            </a-radio-group>
            </p>
            <p>
            <a-button
            class="data-link-button"
            @click="onOk(text, link,juptType)"
            :disabled="link.trim() === ''"
            >
                {{buttonTitle}}
            </a-button>
            </p>
        </div>
    </a-config-provider>
</template>
<script lang="ts">
import { defineComponent, onMounted, PropType, ref } from 'vue'
import { LanguageInterface } from '@/components/gfeditor/emain'
import {ConfigProvider,Button,Input,Radio} from '@arco-design/web-vue';

export default defineComponent({
    name:"am-link-editor",
    components:{
        AConfigProvider:ConfigProvider,
        AInput:Input,
        AButton:Button,
        ARadio:Radio,
        ARadioGroup:Radio.Group,
    },
    props:{
        language:{
            type:Object as PropType<LanguageInterface>,
            required:true
        } as const,
        defaultText:String,
        defaultLink:String,
        defaultjuptType:String,
        className:String,
        onLoad:Function as PropType<() => void>,
        onOk:Function as PropType<(text: string, link: string,juptType:string) => void>
    },
    setup(props){
        const text = ref(props.defaultText)
        const link = ref(props.defaultLink)
        const juptType = ref(props.defaultjuptType)
        const linkRef = ref<HTMLElement | null>(null)
        const textTitle = props.language.get<string>('link', 'text')
        const textPlaceholder = props.language.get<string>('link', 'text_placeholder')

        const linkTitle = props.language.get<string>('link', 'link')
        const juptTypeTitle = props.language.get<string>('link', 'juptType')
        const linkPlaceholder = props.language.get<string>('link', 'link_placeholder')

        const buttonTitle = props.language.get<string>('link', 'ok_button')


        onMounted(() => {
			if(linkRef.value) linkRef.value.focus();
            setTimeout(() => {
                if (props.onLoad) props.onLoad();
            }, 200);
        })

        return {
            text,
            link,
            textTitle,
            textPlaceholder,
            linkTitle,
            juptTypeTitle,
            juptType,
            linkPlaceholder,
            buttonTitle,
            linkRef
        }
    }
})
</script>
