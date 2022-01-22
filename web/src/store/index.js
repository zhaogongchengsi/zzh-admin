import { createStore } from "vuex";
// import VuexPersistence from 'vuex-persist'
import { router } from "@/store/module/routers";

// const vuexLocal = new VuexPersistence({
//   storage: window.localStorage,
//   modules: ['router']
// })

export const store = createStore({
  modules: {
    router,
  },
  // plugins: [vuexLocal.plugin]
});
