import { createApp } from "vue";
import App from "./App.vue";
import "@/assets/style/reset.css";
import router from "@/routers/index.js";
import "element-plus/dist/index.css";
import "@/assets/style/index.css";
import { store } from "@/store/index.js";

createApp(App).use(router).use(store).mount("#app");
