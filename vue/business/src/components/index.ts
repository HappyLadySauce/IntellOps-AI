import { App } from 'vue';
import Chart from './chart/index.vue';
import {Icon} from '@/components/Icon';
import Breadcrumb from '@/components/autoPlugin/breadcrumb/index.vue';
import PageCard from '@/components/autoPlugin/pagecard/index.vue';
import CellAvatar from '@/components/autoPlugin/Cell/CellAvatar.vue';
import CellNodeAvatar from '@/components/autoPlugin/Cell/CellNodeAvatar.vue';
import CellGender from '@/components/autoPlugin/Cell/CellGender.vue';
import CellStatus from '@/components/autoPlugin/Cell/CellStatus.vue';
import DotStatus from '@/components/autoPlugin/Cell/CellDotStatus.vue';
import CellTag from '@/components/autoPlugin/Cell/CellTag.vue';
import CellTagPlus from '@/components/autoPlugin/Cell/CellTagPlus.vue';
import CellTags from '@/components/autoPlugin/Cell/CellTags.vue';
import CellImage from '@/components/autoPlugin/Cell/CellImage.vue';
import CellImages from '@/components/autoPlugin/Cell/CellImages.vue';
import CellCopy from '@/components/autoPlugin/Cell/CellCopy.vue';
import GfDot from '@/components/autoPlugin/GfDot/index';
//表单组件
import FormBelongTable from '@/components/autoPlugin/Form/FormBelongTable.vue';
import FormDicSelect from '@/components/autoPlugin/Form/FormDicSelect.vue';
import FormEditorBox from '@/components/autoPlugin/Form/FormEditorBox.vue';
import FormAudioBox from '@/components/autoPlugin/Form/FormAudioBox.vue';
import FormFileBox from '@/components/autoPlugin/Form/FormFileBox.vue';
import FormImageBox from '@/components/autoPlugin/Form/FormImageBox.vue';
import FormImagesBox from '@/components/autoPlugin/Form/FormImagesBox.vue';

export default {
  install(Vue: App) {
    Vue.component('Chart', Chart);
    Vue.component('Icon', Icon);
    Vue.component('Breadcrumb', Breadcrumb);
    Vue.component('PageCard', PageCard);
    Vue.component('CellAvatar', CellAvatar);
    Vue.component('CellNodeAvatar', CellNodeAvatar);
    Vue.component('CellGender', CellGender);
    Vue.component('CellStatus', CellStatus);
    Vue.component('DotStatus', DotStatus);
    Vue.component('CellTag', CellTag);
    Vue.component('CellTagPlus', CellTagPlus);
    Vue.component('CellTags', CellTags);
    Vue.component('CellImage', CellImage);
    Vue.component('CellImages', CellImages);
    Vue.component('CellCopy', CellCopy);
    Vue.component('GfDot', GfDot);
    //表单组件
    Vue.component('FormBelongTable', FormBelongTable);
    Vue.component('FormDicSelect', FormDicSelect);
    Vue.component('FormEditorBox', FormEditorBox);
    Vue.component('FormAudioBox', FormAudioBox);
    Vue.component('FormFileBox', FormFileBox);
    Vue.component('FormImageBox', FormImageBox);
    Vue.component('FormImagesBox', FormImagesBox);
  },
};
