<template>
    <a-tooltip :placement="placement || 'bottom'" :visible="(!!title || !!hotkeyText) && !isMobile ? visible : false"
    :content-class="(!!title || !!hotkeyText) && !isMobile ? '' : 'prompt-popover-hide'"
    :arrow-class="(!!title || !!hotkeyText) && !isMobile ? '' : 'prompt-popover-hide'"
    :arrow-style="{display:(!!title || !!hotkeyText) && !isMobile ? '' : 'none'}"
    >
        <template #content>
            <div v-if="!!title" class="toolbar-tooltip-title">{{title}}</div>
            <div v-if="!!hotkeyText" class="toolbar-tooltip-hotkey" v-html="hotkeyText"></div>
        </template>

        <button
        :class="['toolbar-button',className,{'toolbar-button-active': active,'toolbar-button-disabled':disabled}]"
        ref="element"
        @click="triggerClick"
        @mousedown="triggerMouseDown"
        @mouseenter="triggerMouseEnter"
        @mouseleave="triggerMouseLeave">
            <slot name="icon">
                <span v-if="iconIsHtml" v-html="icon"></span>
                <span v-if="!iconIsHtml && icon" :class="`data-icon data-icon-${icon}`" />
            </slot>
            <slot>{{typeof content === 'function' ? content(engine) : content}}</slot>
        </button>
    </a-tooltip>
</template>
<script lang="ts">
import { defineComponent, ref} from "vue"
import { Tooltip} from  '@arco-design/web-vue';
import { formatHotkey, isMobile } from '@/components/gfeditor/emain'
import { autoGetHotkey } from "../utils"
import { buttonProps } from "../types"

export default defineComponent({
    name:"am-button",
    components:{
        ATooltip:Tooltip
    },
    props:buttonProps,
    setup(props){
        const element = ref<HTMLButtonElement | undefined>()
        let hotkey = props.hotkey;
        //默认获取插件的热键
        if (props.engine && (hotkey === true || hotkey === undefined)) {
            hotkey = autoGetHotkey(
                props.engine,
                props.command && !Array.isArray(props.command) ? props.command.name : props.name,
            );
        }
        if (typeof hotkey === 'string' && hotkey !== '') {
            hotkey = formatHotkey(hotkey)
        }

        const visible = ref(false)

        return {
            iconIsHtml:/^<.*>/.test((props.icon || "").trim()),
            isMobile,
            visible,
            hotkeyText:hotkey,
            element
        }
    },
    data(){
        return {
            visible:false
        }
    },
    methods:{
        triggerMouseDown(event:MouseEvent){
            event.preventDefault();
            if (this.disabled) return;
            if (this.onMouseDown) this.onMouseDown(event, this.engine);
            this.visible = false
        },
        triggerMouseEnter(event:MouseEvent){
            if(this.onMouseEnter) this.onMouseEnter(event, this.engine)
            this.visible = true
        },
        triggerMouseLeave(event:MouseEvent){
            if(this.onMouseLevel) this.onMouseLevel(event, this.engine)
            this.visible = false
        },
        triggerClick(event: MouseEvent){
            const nodeName = (event.target as Node).nodeName;
            if (nodeName !== 'INPUT' && nodeName !== 'TEXTAREA')
                event.preventDefault();
            if (this.disabled) return;
            if (this.onClick && this.onClick(event, this.engine) === false) return;
            if (this.autoExecute !== false) {
                let commandName = this.name;
                let commandArgs = [];
                if (this.command) {
                    if (!Array.isArray(this.command)) {
                        commandName = this.command.name;
                        commandArgs = this.command.args;
                    } else {
                        commandArgs = this.command;
                    }
                }
				if(this.engine)
                this.engine.command.execute(commandName, ...commandArgs);
            }
        }
    }
})
</script>
<style css>
.editor-toolbar .toolbar-button {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: auto;
    min-width: 26px;
    margin: 0;
    text-align: center;
    padding: 0 7px;
    background-color: transparent;
    border: 1px solid transparent;
    border-radius: 3px 3px;
    font-size: 16px;
    cursor: pointer;
    color: var(--color-neutral-8);
    outline: none;
}

.editor-toolbar.editor-toolbar-popup .toolbar-button {
    min-width: 24px;
    line-height: 24px;
    border-radius: 4px;
}

.editor-toolbar:not(.editor-toolbar-mobile) .toolbar-button {
    padding: 0px;
    margin: 0 1px;
}

.editor-toolbar:not(.editor-toolbar-mobile) .toolbar-button:hover {
    border: 1px solid transparent;
    background-color: var(--color-neutral-2);
}

.editor-toolbar:not(.editor-toolbar-mobile) .toolbar-button:active,.editor-toolbar .toolbar-button-active,.editor-toolbar:not(.editor-toolbar-mobile) .toolbar-button-active:hover {
    background-color: #e8e8e8;
    border: 1px solid transparent;
}

.editor-toolbar .toolbar-button-disabled,.editor-toolbar:not(.editor-toolbar-mobile) .toolbar-button-disabled:hover {
    background-color: transparent;
    border: 1px solid transparent;
    box-shadow: none;
    color: #000000;
    opacity: 0.25;
    cursor: not-allowed;
}
.data-icon-arrow {
    position: absolute;
    top: calc(100% / 2 - 2px);
    width: 8px;
    height: 8px;
    background-image: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4NCjxzdmcgd2lkdGg9IjhweCIgaGVpZ2h0PSI1cHgiIHZpZXdCb3g9IjAgMCA4IDUiIHZlcnNpb249IjEuMSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayI+DQogICAgPCEtLSBHZW5lcmF0b3I6IFNrZXRjaCA1Mi41ICg2NzQ2OSkgLSBodHRwOi8vd3d3LmJvaGVtaWFuY29kaW5nLmNvbS9za2V0Y2ggLS0+DQogICAgPHRpdGxlPkdyb3VwIENvcHkgNjwvdGl0bGU+DQogICAgPGRlc2M+Q3JlYXRlZCB3aXRoIFNrZXRjaC48L2Rlc2M+DQogICAgPGcgaWQ9IlN5bWJvbHMiIHN0cm9rZT0ibm9uZSIgc3Ryb2tlLXdpZHRoPSIxIiBmaWxsPSJub25lIiBmaWxsLXJ1bGU9ImV2ZW5vZGQiIG9wYWNpdHk9IjAuNDUiPg0KICAgICAgICA8ZyBpZD0idG9vbGJhciIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTMyOC4wMDAwMDAsIC0xOC4wMDAwMDApIj4NCiAgICAgICAgICAgIDxnIGlkPSJwYXJhZ3JhcGgtc3R5bGUiPg0KICAgICAgICAgICAgICAgIDxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDIyNi4wMDAwMDAsIDQuMDAwMDAwKSI+DQogICAgICAgICAgICAgICAgICAgIDxnIGlkPSJHcm91cC1Db3B5LTYiIHRyYW5zZm9ybT0idHJhbnNsYXRlKDEwMi4wMDAwMDAsIDEyLjAwMDAwMCkiPg0KICAgICAgICAgICAgICAgICAgICAgICAgPHJlY3QgaWQ9IlJlY3RhbmdsZSIgeD0iMCIgeT0iMCIgd2lkdGg9IjgiIGhlaWdodD0iOCI+PC9yZWN0Pg0KICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD0iTTAuNTk2MDkzNzUsMi41NTcwMzEyNSBMMy43NDUzMTI1LDYuMzc4MTI1IEMzLjg3NzM0Mzc1LDYuNTI1NzgxMjUgNC4xMDg1OTM3NSw2LjUyNTc4MTI1IDQuMjQwNjI1LDYuMzc4MTI1IEw3LjQwNTQ2ODc1LDIuNTU3MDMxMjUgQzcuNTk2MDkzNzUsMi4zNDI5Njg3NSA3LjQ0NDUzMTI1LDIuMDAzOTA2MjUgNy4xNTc4MTI1LDIuMDAzOTA2MjUgTDAuODQ0NTMxMjUsMi4wMDM5MDYyNSBDMC41NTcwMzEyNSwyLjAwMzkwNjI1IDAuNDA0Njg3NSwyLjM0Mjk2ODc1IDAuNTk2MDkzNzUsMi41NTcwMzEyNSBaIiBpZD0iU2hhcGUiIGZpbGw9IiMwMDAwMDAiIGZpbGwtcnVsZT0ibm9uemVybyI+PC9wYXRoPg0KICAgICAgICAgICAgICAgICAgICA8L2c+DQogICAgICAgICAgICAgICAgPC9nPg0KICAgICAgICAgICAgPC9nPg0KICAgICAgICA8L2c+DQogICAgPC9nPg0KPC9zdmc+);
    background-repeat: no-repeat;
    transition: all 0.25s cubic-bezier(0.3, 1.2, 0.2, 1);
}
.prompt-popover-hide {
    display: none;
}
</style>
