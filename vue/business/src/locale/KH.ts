//柬埔寨语
//自动加载
const autoLocle = import.meta.glob('@/views/**/locale/KH.ts', {
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
  'navbar.docs': 'មជ្ឈមណ្ឌលឯកសារ',
  'navbar.action.locale': 'ប្តូរទៅជាភាសាចិន',
  'searchTable.operation.create': 'ថ្មី។',
  'searchTable.operation.export': 'នាំចេញ',
  'searchTable.form.reset': 'កំណត់ឡើងវិញ',
  'footer.copyright': 'រក្សាសិទ្ធិគ្រប់យ៉ាង',
  'system': 'ការគ្រប់គ្រងប្រព័ន្ធ',
  'datacenter': 'មជ្ឈមណ្ឌលទិន្នន័យ',
  ...formatModules(autoLocle,{})
};
