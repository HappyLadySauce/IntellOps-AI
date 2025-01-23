import localeMessageBox from '@/components/message-box/locale/zh-CN';
import localeSettings from './zh-CN/settings';
import localeMenu from './zh-CN/menu';

//自动加载
const autoLocle = import.meta.glob('@/views/**/locale/zh-CN.ts', {
  eager: true,
});
function formatModules(_modules: any, result: {}) {
  Object.keys(_modules).forEach((key) => {
    const defaultModule = _modules[key].default;
    if (!defaultModule) return;
    result=Object.assign({},result,defaultModule);
  });
  return result;
}
export default {
  'navbar.docs': '文档中心',
  'navbar.action.locale': '切换为中文',
  'searchTable.operation.create': '新建',
  'searchTable.operation.export': '导出',
  'searchTable.form.reset': '重置',
  'footer.copyright': '版权所有',
  'system': '系统管理',
  'datacenter': '数据中心',
  'sys.copy.success': '复制成功',
  ...localeSettings,
  ...localeMenu,
  ...localeMessageBox,
  ...formatModules(autoLocle,{})
};
