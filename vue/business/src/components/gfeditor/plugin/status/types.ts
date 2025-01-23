import { PluginOptions } from '@/components/gfeditor/emain';

export interface StatusOptions extends PluginOptions {
	hotkey?: string | Array<string>;
	colors?: {
		background: string;
		color: string;
	}[];
}
