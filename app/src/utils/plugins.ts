import { h } from 'vue'
import { Icon } from '@arco-design/web-vue';
import IconFont from "@/components/IconFont.vue"

export default {
  install: (app, options) => {
    app.directive('dark', {
      mounted(el:Element,{value, arg}) {
        dark(el,value,arg)
      },
      beforeUpdate (el,{value,arg}) {
        dark(el,value,arg)
      }
    })
    app.component("icon-font", IconFont)
  }
}
function dark (el:Element, value:boolean | "light" | "dark", suffix:string) {
  const className = el.classList.item(0)
  if (value === true || value === "dark") {
    if (className) {
       el.classList.add(`${className}-${suffix || 'dark'}`)
    }
  } else if (value === "light") {
    if (className) {
       el.classList.remove(`${className}-${suffix || 'dark'}`)
    }
  }
}

