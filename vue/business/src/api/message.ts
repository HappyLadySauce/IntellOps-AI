import { defHttp } from '@/utils/http';
enum Api {
  messageList = '/common/message/getList',
  messageRead = '/common/message/read',
  chatList = '/common/chat/get_list',
}
export interface MessageRecord {
  id: number;
  type: string;
  title: string;
  subTitle: string;
  avatar?: string;
  content: string;
  time: string;
  status: 0 | 1;
  messageType?: number;
}
export interface responseData {
  items: MessageRecord[];
  total: number;
}
//获取消息列表
export function queryMessageList() {
  return defHttp.get<responseData>({ url:Api.messageList}, { errorMessageMode: 'message' });
}

interface MessageStatus {
  ids: number[];
}
//标记消息为已读
export type MessageListType = MessageRecord[];
export function setMessageStatus(params: MessageStatus) {
  return defHttp.post<MessageListType>({ url:Api.messageRead, params:params}, { errorMessageMode: 'message' });
}

export interface ChatRecord {
  id: number;
  username: string;
  content: string;
  time: string;
  isCollect: boolean;
}
//聊天信息
export function queryChatList() {
  return defHttp.get<ChatRecord[]>({ url:Api.chatList}, { errorMessageMode: 'message' });
}
