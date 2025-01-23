import {
	Plugin,
} from '@/components/gfeditor/emain';
import { createApp } from 'vue';
import ArcoVue from '@arco-design/web-vue';
import AmEditor from './editor.vue';
export default class extends Plugin {
	static get pluginName() {
		return 'imgattachment';
	}
	init() {
	}
	async execute() {
		const mountNode = document.createElement("body");
		const vm = createApp(AmEditor, {
			onLoad: () => {
			},
			onOk: (item:any) => {
				for(  let img of  item.list ){
					this.editor.command.execute("image", 'done', img);
				}
			},
		});
		vm.use(ArcoVue)
		vm.mount(mountNode);
		return;
	}


	destroy() {
		const editor = this.editor;
	}
}
