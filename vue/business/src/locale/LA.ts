//老挝
//自动加载
const autoLocle = import.meta.glob('@/views/**/locale/LA.ts', {
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
  'navbar.docs': 'ສູນເອກະສານ',
  'navbar.action.locale': 'ປ່ຽນເປັນພາສາຈີນ',
  'searchTable.operation.create': 'ໃຫມ່',
  'searchTable.operation.export': 'ສົ່ງອອກ',
  'searchTable.form.reset': 'ຕັ້ງຄືນໃຫມ່',
  'footer.copyright': 'ສະຫງວນໄວ້',
  'system': 'ການຄຸ້ມຄອງລະບົບ',
  'datacenter': 'ສູນຂໍ້ມູນ',
  ...formatModules(autoLocle,{})
};
