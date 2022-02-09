import "@/assets/style/reset.css";
import "element-plus/dist/index.css";
import "@/assets/style/index.css";
/* ----- */
import { createApp } from "vue";
import App from "./App.vue";
import router from "@/routers/index.js";
import BaseLayou from "@/components/BaseLayou/index.vue";
import { doorgod } from "./doorgod.js";
import { createPinia } from "pinia";
import { useRouterStore } from "@/store/router.js";

const app = createApp(App);
app.component("BaseLayou", BaseLayou);
app.use(router);
app.use(createPinia());
doorgod(router, useRouterStore);
app.mount("#app");
