export type SizeProps = 'mini' | 'small' | 'medium' | 'large';
import type { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
//状态
export const statusOptions :SelectOptionData[]= [
    {
        label: "全部",
        value: "",
    },
    {
        label: "正常",
        value: 0,
    },
    {
        label: "禁用",
        value: 1,
    },
    ];