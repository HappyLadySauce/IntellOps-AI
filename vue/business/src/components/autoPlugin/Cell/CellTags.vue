<template>
  <a-overflow-list v-if="data.length" :min="1">
    <template v-for="(item,index) in data">
      <a-tag v-if="item" size="small" :color="getColor(index)">
        {{ item }}
      </a-tag>
    </template>
    <template #overflow="{ number }">
      <a-popover :content-style="{ maxWidth: '300px', padding: '8px 12px' }">
        <a-tag color="arcoblue" size="small">+{{ number }}</a-tag>
        <template #content>
          <a-space wrap>
            <template v-for="(tag,i) in data.filter((i, n) => n >= data.length - number)">
            <a-tag :key="tag" v-if="tag" size="small" :color="getColor(i)">
              {{ tag }}
            </a-tag>
          </template>
          </a-space>
        </template>
      </a-popover>
    </template>
  </a-overflow-list>
</template>

<script lang="ts" setup>
withDefaults(defineProps<Props>(), {
  data: () => []
})

interface Props {
  data: string[]
}
//获取随机颜色
const getColor=(index:number)=>{
    const color=[ 'green','cyan','blue','arcoblue','purple','pinkpurple','magenta','gray','red','orangered','orange','gold','lime']
    if(index<=12){
       return color[index]
    }else{
        const randNum =  Math.floor(Math.random() * 11) ;
        return color[randNum]
    }
}
</script>

