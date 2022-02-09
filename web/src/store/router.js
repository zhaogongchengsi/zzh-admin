import { defineStore } from "pinia";
import { replaceRouter, RouterDataToTree } from "@/utils/asyncRoute.js";
import { GetRouter } from "@/api/router.js";

export const useRouterStore = defineStore("router", {
  state: () => {
    return {
      routers: [], // 根路由数据
      children: [],
      originlRouter: [],
    };
  },
  actions: {
    async initRouter(user) {
      try {
        let _router = await GetRouter(); // 获取路由数据
        this.originlRouter = RouterDataToTree(_router);
        localStorage.setItem("z_router", JSON.stringify(_router));
        const router = replaceRouter(_router);
        this.routers = router;
        return true;
      } catch (e) {
        console.error("err", e);
        return false;
      }
    },
  },
});
