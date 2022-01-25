import {
  replaceRouter,
  filePathCompile,
  copyRouter,
} from "@/utils/asyncRoute.js";
// import pageRouter from "@/routers/index";
const routerFile = import.meta.glob("./views/**/*.{vue,js,ts,jsx,tsx}");
let routerfiles = filePathCompile(routerFile);
const whiteList = ["login"];
let asyncRouterFlag = false;

function initRouter(router) {
  const _router = JSON.parse(localStorage.getItem("z_router"));
  const token = localStorage.getItem("token");
  if (_router && token) {
    const asyncRouters = replaceRouter(_router.originlRoutData, routerfiles);
    const _r = copyRouter(asyncRouters);
    router.addRoute({
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
  } else {
    router.push("/login");
  }
}

export function doorgod(router) {
  router.beforeEach(async function (to, from) {
    if (!whiteList.includes(to.name)) {
      if (!asyncRouterFlag) {
        asyncRouterFlag = true;
        initRouter(router);
        return {
          path: "home_user",
        };
      }

      return true;
    } else {
      return true;
    }
  });
}
