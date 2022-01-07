import {
  dataToTree,
  ParAndChildren,
  filePathCompile,
  separation,
  copyRouter,
} from "@/utils/asyncRoute.js";
import { GetRouter } from "@/api/router.js";
import pageRouter from "@/routers/index";
import NotComponent from "@/components/NotComponent.vue";
const routerFile = import.meta.glob("../../views/**/*.{vue,js,ts,jsx,tsx}");
let routerfiles = filePathCompile(routerFile);

function replaceRouter(routerTree, fileTree) {
  return routerTree.map((router) => {
    let _component = fileTree[router.Component];
    if (_component) {
      router.Component = _component;
    } else {
      router.Component = router.Component;
    }
    if (router.children && router.children.length > 0) {
      router.children = replaceRouter(router.children, fileTree);
    }
    return router;
  });
}

async function GetMenuData() {
  const originlRouter = await GetRouter();
  return dataToTree(originlRouter);
}

export const router = {
  namespaced: true,
  state: {
    asyncRouters: [], // 最终的路由数据
    OriginlRoutData: [], // 原始路由数据
    root: [], // 根路由数据
    children: [],
  },
  mutations: {
    setAsyncRouter(state, asyncRouters) {
      state.asyncRouters = asyncRouters;
    },
    setOriginRouter(state, OriginRouter) {
      state.OriginlRoutData = OriginRouter;
    },
    setRootData(state, RootData) {
      state.root = RootData;
    },
    setchildren(state, children) {
      state.children = children;
    },
  },
  actions: {
    // 从后台获取动态路由
    async SetAsyncRouter({ commit }) {
      try {
        const originlRouter = await GetRouter(); // 获取路由数据
        const { parents, children } = separation(originlRouter); // 分离根组件和非根组件
        const asyncRouters = ParAndChildren(parents, children);
        const routers = replaceRouter(asyncRouters, routerfiles); // 替换路由组件
        commit("setRootData", parents); // 提交根路由
        commit("setchildren", children); // 提交子路由
        commit("setOriginRouter", JSON.parse(JSON.stringify(asyncRouters))); // 备份路由信息 供菜单渲染使用
        commit("setAsyncRouter", routers); // 提交生成后的路由信息
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

    async SetOriginRouterData({ commit }) {
      const originlRouter = await GetRouter(); // 获取路由数据
      const { parents, children } = separation(originlRouter); // 分离根组件和非根组件
      const asyncRouters = ParAndChildren(parents, children);
      commit("setOriginRouter", asyncRouters);
    },
  },
  getters: {
    // 获取动态路由
    getAsyncRouter(state) {
      return state.asyncRouters;
    },

    getOriginRouter(state) {
      return state.OriginlRoutData;
    },
  },
};
