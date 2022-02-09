import { createApp } from 'vue'
import App from './App.vue'
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
import { createPinia } from 'pinia'
import router from '@/router/index.js';

const app = createApp(App)
app.use(ArcoVue, {
  // 用于改变使用组件时的前缀名称
  componentPrefix: 'blog'
});
app.use(router)
app.use(createPinia())
app.mount('#app')