import { createStore } from "vuex";

import { router } from "@/store/module/routers";

export const store = createStore({
  modules: {
    router,
  },
});
