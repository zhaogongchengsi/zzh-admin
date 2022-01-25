import "@/assets/style/reset.css";
import "element-plus/dist/index.css";
import "@/assets/style/index.css";
/* ----- */
import { createApp } from "vue";
import App from "./App.vue";
import router from "@/routers/index.js";
import { store } from "@/store/index.js";
import BaseLayou from "@/components/BaseLayou/index.vue";
import { doorgod } from "./doorgod.js";
// import "./persistence.js";
import { createPinia } from "pinia";

const app = createApp(App);
app.component("BaseLayou", BaseLayou);
app.use(router);
doorgod(router);
app.use(store);
app.use(createPinia());
app.mount("#app");
