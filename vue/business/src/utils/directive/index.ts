// index.ts
import { Resize } from './resizeObserver'

// 自定义指令集合
const directives: any = {
  Resize
}

export default {
  install(app: any) {
    Object.keys(directives).forEach((key) => {
      app.directive(key, directives[key])
    })
  }
}