import { createApp } from 'vue'
import App from './App.vue'
import "@/assets/style/reset.css"
import router from '@/routers/index.js'
import 'element-plus/dist/index.css'
import '@/assets/style/index.css'


createApp(App).use(router).mount('#app')
