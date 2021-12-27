import "@/assets/style/reset.css";
import "element-plus/dist/index.css";
import "@/assets/style/index.css";
/* ----- */
import { createApp } from "vue";
import App from "./App.vue";
import router from "@/routers/index.js";
import { store } from "@/store/index.js";
import BaseLayou from "@/components/BaseLayou/index.vue";

import "./persistence.js" 

const app = createApp(App);
app.component("BaseLayou", BaseLayou);
app.use(router);
app.use(store);
app.mount("#app");
