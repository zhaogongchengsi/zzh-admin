import { defineStore } from "pinia";
import {
  dataToTree,
  ParAndChildren,
  filePathCompile,
  separation,
  copyRouter,
} from "@/utils/asyncRoute.js";
import { GetRouter } from "@/api/router.js";
import pageRouter from "@/routers/index";
import routerComponent from "@/views/routerHandel.vue";
import notComponent from "@/components/NotComponent.vue";
const routerFile = import.meta.glob("../../views/**/*.{vue,js,ts,jsx,tsx}");
let routerfiles = filePathCompile(routerFile);

function replaceRouter(routerTree, fileTree) {
  return routerTree.map((router) => {
    let _component = fileTree[router.Component];
    if (_component) {
      router.Component = _component;
    } else {
      if (router.children && router.children.length > 0) {
        router.Component = routerComponent;
      } else {
        router.Component = notComponent;
      }
    }
    if (router.children && router.children.length > 0) {
      router.children = replaceRouter(router.children, fileTree);
    }
    return router;
  });
}

export const useRouterStore = defineStore("router", {
  state: () => {
    return {
      asyncRouters: [], // 最终的路由数据
      OriginlRoutData: [], // 原始路由数据
      root: [], // 根路由数据
      children: [],
    };
  },
  actions: {
    async initRouter(user) {
      try {
        const originlRouter = await GetRouter(); // 获取路由数据
        const { parents, children } = separation(originlRouter); // 分离根组件和非根组件
        const asyncRouters = ParAndChildren(parents, children);
        const routers = replaceRouter(asyncRouters, routerfiles); // 替换路由组件
        this.root = parents;
        this.children = children;
        this.OriginlRoutData = JSON.parse(JSON.stringify(asyncRouters));
        this.asyncRouters = routers;
        const _r = copyRouter(routers);
        pageRouter.addRoute({
          path: "/",
          name: "baselayou",
          component: () => import("@/views/Layou/index.vue"),
          children: [
            {
              path: "/home_user",
              name: "home_user",
              component: () => import("@/views/layou/HomeUser.vue"),
            },
            ..._r,
          ],
          redirect: "/home_user", // 重定向到用户首页
        });
        return true;
      } catch (e) {
        return false;
      }
    },
  },
});
