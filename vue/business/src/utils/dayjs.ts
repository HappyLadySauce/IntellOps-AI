import  { type ShortcutType } from '@arco-design/web-vue'
import {computed} from 'vue';
import dayjs from 'dayjs'
export const  shortcuts = computed<ShortcutType[]>(() => {
    return [
      {
        label: '今天',
        value: (): Date[] => [dayjs().startOf('day').toDate(), dayjs().toDate()]
      },
      {
        label: '昨天',
        value: (): Date[] => [
          dayjs().subtract(1, 'day').startOf('day').toDate(),
          dayjs().subtract(1, 'day').endOf('day').toDate()
        ]
      },
      {
        label: '本周',
        value: (): Date[] => [dayjs().startOf('week').add(1, 'day').toDate(), dayjs().toDate()]
      },
      {
        label: '本月',
        value: (): Date[] => [dayjs().startOf('month').toDate(), dayjs().toDate()]
      },
      {
        label: '本年',
        value: (): Date[] => [dayjs().startOf('year').toDate(), dayjs().toDate()]
      }
    ]
  })

export {dayjs}