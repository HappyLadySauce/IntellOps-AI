import localeMessageBox from '@/components/message-box/locale/en-US';
import localeSettings from './en-US/settings';
import localeMenu from './en-US/menu';
//自动加载 views下的语言包
const autoLocle = import.meta.glob('@/views/**/locale/en-US.ts', {
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
  'navbar.docs': 'Docs',
  'navbar.action.locale': 'Switch to English',
  'searchTable.operation.create': 'Create',
  'searchTable.operation.export': 'Export',
  'searchTable.form.reset': 'Reset',
  'footer.copyright': 'Copyright',
  'system': 'System',
  'datacenter': 'data center',
  'sys.copy.success': 'replicate successful',
  ...localeSettings,
  ...localeMenu,
  ...localeMessageBox,
  ...formatModules(autoLocle,{})
};
