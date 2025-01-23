<template>
    <div >
        <FileManage @register="registerModal" @success="selectImg"/>
    </div>
</template>
<script lang="ts">
import { defineComponent, onMounted, PropType } from 'vue'
import FileManage from '@/components/attachment/FileManage.vue';
import { useModal } from '@/components/Modal';

export default defineComponent({
    name:"am-link-editor",
    components:{
        FileManage
    },
    props:{
        defaultText:String,
        defaultLink:String,
        defaultjuptType:String,
        className:String,
        onLoad:Function as PropType<() => void>,
        onOk:Function as PropType<(item:any) => void>
    },
    setup(props){
        const [registerModal, { openModal }] = useModal();
        //选图片
        const selectImg=(item:any)=>{
            if (props.onOk) 
             props.onOk(item);
        }
        onMounted(() => {
            openModal(true, {
                    filetype:"image",
                    getnumber: "more",//one 单张
                    openfrom: "manage",
                 });
        })

        return {
            registerModal,
            selectImg,
        }
    }
})
</script>
