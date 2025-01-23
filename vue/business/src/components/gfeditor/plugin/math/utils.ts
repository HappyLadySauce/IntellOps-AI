import { EditorInterface } from '@/components/gfeditor/emain';

export const getLocales = <T extends string | {} = string>(
	editor: EditorInterface,
) => {
	return editor.language.get<{ [key: string]: T }>('math');
};
