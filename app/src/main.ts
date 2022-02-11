import { createApp } from 'vue'
import App from './App.vue'
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
import { createPinia } from 'pinia'
import router from './router';
import "./assets/reset.css"
// 编辑器插件
import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';
import Prism from 'prismjs';
// 自定义插件
import dark from '@/utils/plugins'
VueMarkdownEditor.use(vuepressTheme, {
  Prism,
});
const app = createApp(App)
app.use(ArcoVue, {
  componentPrefix: 'blog'
});
app.use(router)
app.use(createPinia())
app.use(VueMarkdownEditor);
app.use(dark, {
  icon: 'https://at.alicdn.com/t/font_3178565_i7epnvlht5.js'
})
app.mount('#app')
