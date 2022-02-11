import { createApp } from 'vue'
import App from './App.vue'
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
import { createPinia } from 'pinia'
import router from './router';
import "./assets/reset.css"

const app = createApp(App)
app.use(ArcoVue, {
  componentPrefix: 'blog'
});
app.use(router)
app.use(createPinia())
app.mount('#app')