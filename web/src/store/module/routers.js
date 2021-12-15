import { dataToTree, copyRouter, filePathCompile } from "@/utils/asyncRoute.js";
import { GetRouter } from "@/api/router.js";
const routerFile = import.meta.glob("../../views/**/*.{vue,js,ts,jsx,tsx}");

function replaceRouter(routerTree, fileTree) {
  return routerTree.map((router) => {
    let _component = fileTree[router.component];
    if (_component) {
      router.component = _component;
    } else {
      delete router.component;
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
  },
  mutations: {
    setAsyncRouter(state, asyncRouters) {
      state.asyncRouters = asyncRouters;
    },
    setOriginRouter(state, OriginRouter) {
      state.OriginlRoutData = OriginRouter;
    },
  },
  actions: {
    // 从后台获取动态路由
    async SetAsyncRouter({ commit }) {
      const originData = await GetMenuData();
      commit("setOriginRouter", originData);
      /*
       * 请求过来的路由数据 树形化之后 把里面所需要的路由数据拷贝出来
       */
      let routerTree = copyRouter(originData);
      /*
       * 将文件内的组件取出来
       */
      let routerfiles = filePathCompile(routerFile);
      /*
       * 将树形路由数据中的组件替换成真实组件
       */
      let _routerTree = replaceRouter(routerTree, routerfiles);
      commit("setAsyncRouter", _routerTree);

      //!  路由替换

      return true;
    },

    async SetOriginRouterData({ commit }) {
      const originData = await GetMenuData();
      commit("setOriginRouter", originData);
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
