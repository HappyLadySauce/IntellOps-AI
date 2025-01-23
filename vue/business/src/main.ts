import { createApp } from 'vue';
import ArcoVue from '@arco-design/web-vue';
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import globalComponents from '@/components';
import router from './router';
import store from './store';
import i18n from './locale';
import directive from './directive';
import App from './App.vue';
import directives from '@/utils/directive/index'//检测元素大小变化
import '@arco-design/web-vue/dist/arco.less';
import '@/assets/style/global.less';
import '@/assets/style/gfui.less';
const app = createApp(App);
app.use(ArcoVue, {});
app.use(ArcoVueIcon);

app.use(router);
app.use(store);
app.use(i18n);
app.use(globalComponents);
app.use(directive);
app.use(directives);
// 去掉Vue warn警告
app.config.warnHandler = () => null;

app.mount('#app');
