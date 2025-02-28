import dayjs from 'dayjs';
import  { successResponseWrap } from '@/utils/setup-mock';
export interface ContentDataRecord {
  x: string;
  y: number;
}

export function queryContentData() {
  const presetData = [58, 81, 53, 90, 64, 88, 49, 79];
  const getLineData = () => {
    const count = 8;
    return new Array(count).fill(0).map((el, idx) => ({
      x: dayjs()
        .day(idx - 2)
        .format('YYYY-MM-DD'),
      y: presetData[idx],
    }));
  };
  return successResponseWrap([...getLineData()]);
}

export interface PopularRecord {
  key: number;
  clickNumber: string;
  title: string;
  increases: number;
}

export function queryPopularList(params: { type: string }) {
  const textList = [
    {
      key: 1,
      clickNumber: '346.3w+',
      title: '经济日报：财政政策要精准提升…',
      increases: 35,
    },
    {
      key: 2,
      clickNumber: '324.2w+',
      title: '双12遇冷，消费者厌倦了电商平…',
      increases: 22,
    },
    {
      key: 3,
      clickNumber: '318.9w+',
      title: '致敬坚守战“疫”一线的社区工作…',
      increases: 9,
    },
    {
      key: 4,
      clickNumber: '257.9w+',
      title: '普高还是职高？家长们陷入选择…',
      increases: 17,
    },
    {
      key: 5,
      clickNumber: '124.2w+',
      title: '人民快评：没想到“浓眉大眼”的…',
      increases: 37,
    },
  ];
  return successResponseWrap([...textList]);
}
