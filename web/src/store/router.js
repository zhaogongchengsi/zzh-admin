import { defineStore } from "pinia";
import {
  replaceRouter,
  ParAndChildren,
  filePathCompile,
  separation,
  copyRouter,
} from "@/utils/asyncRoute.js";
import { GetRouter } from "@/api/router.js";
import pageRouter from "@/routers/index";

const routerFile = import.meta.glob("../views/**/*.{vue,js,ts,jsx,tsx}");
let routerfiles = filePathCompile(routerFile);

// function replaceRouter(routerTree, fileTree) {
//   return routerTree.map((router) => {
//     let _component = fileTree[router.Component];
//     if (_component) {
//       router.Component = _component;
//     } else {
//       if (router.children && router.children.length > 0) {
//         router.Component = routerComponent;
//       } else {
//         router.Component = notComponent;
//       }
//     }
//     if (router.children && router.children.length > 0) {
//       router.children = replaceRouter(router.children, fileTree);
//     }
//     return router;
//   });
// }

export const useRouterStore = defineStore("router", {
  state: () => {
    const _router = JSON.parse(localStorage.getItem("z_router"));
    if (_router) {
      // const asyncRouters = replaceRouter(_router.originlRoutData, routerfiles);
      // const _r = copyRouter(asyncRouters);
      // pageRouter.addRoute({
      //   path: "/",
      //   name: "baselayou",
      //   component: () => import("@/views/Layou/index.vue"),
      //   children: [
      //     {
      //       path: "/home_user",
      //       name: "home_user",
      //       component: () => import("@/views/layou/HomeUser.vue"),
      //     },
      //     ..._r,
      //   ],
      //   redirect: "/home_user", // 重定向到用户首页
      // });
      // console.log(pageRouter.getRoutes());
      return {
        ..._router,
        isRouter: true,
      };
    }
    return {
      asyncRouters: [], // 最终的路由数据
      OriginlRoutData: [], // 原始路由数据
      root: [], // 根路由数据
      children: [],
      isRouter: false, // 是否获取到路由信息
    };
  },
  actions: {
    async initRouter(user) {
      try {
        const originlRouter = await GetRouter(); // 获取路由数据
        console.log(121, originlRouter);
        const { parents, children } = separation(originlRouter); // 分离根组件和非根组件
        const asyncRouters = ParAndChildren(parents, children);
        localStorage.setItem(
          "z_router",
          JSON.stringify({
            root: parents,
            children: children,
            originlRoutData: asyncRouters,
          })
        );

        const routers = replaceRouter(asyncRouters, routerfiles); // 替换路由组件
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
        this.isRouter = true;
        console.log("当前调用路由的函数", user);
        return true;
      } catch (e) {
        return false;
      }
    },
  },
});
