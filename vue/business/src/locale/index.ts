import { createI18n } from 'vue-i18n';
import en from './en-US';
import cn from './zh-CN';
import la from './LA';//老挝语
import kh from './KH';//柬埔寨语

export const LOCALE_OPTIONS = [
  { label: '中文', value: 'zh-CN' },
  { label: 'English', value: 'en-US' },
  { label: 'ລາວ', value: 'LA' },//老挝语
  { label: 'កម្ពុជា។', value: 'KH' },//柬埔寨语
];
const defaultLocale = localStorage.getItem('arco-locale') || 'zh-CN';

const i18n = createI18n({
  locale: defaultLocale,
  fallbackLocale: 'zh-CN',//默认中文
  allowComposition: true,
  messages: {
    'en-US': en,
    'zh-CN': cn,
    'LA': la,
    'KH': kh,
  },
});

export default i18n;
