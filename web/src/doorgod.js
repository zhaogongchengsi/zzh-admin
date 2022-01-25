import { useRouterStore } from "@/store/router.js";
const whiteList = ["login"];
let asyncRouterFlag = 0;
export function doorgod(router) {
  router.beforeEach(function (to, from, next) {
    const _token = localStorage.getItem("token");
    if (whiteList.includes(to.name)) {
      // 当在白名单内
      if (_token) {
        if (!asyncRouterFlag) {
          asyncRouterFlag++;
          useRouterStore();
        }
        next();
      } else {
        next();
      }
    } else {
      if (_token) {
        if (!asyncRouterFlag) {
          asyncRouterFlag++;
          useRouterStore();
          next({ path: "/home_user" });
        } else {
          if (to.matched.length) {
            next();
          } else {
            next({ path: "/login" });
          }
        }
      }
      // 不在白名单中并且未登陆的时候
      if (!_token) {
        next({
          name: "login",
        });
      }
    }
  });
}
