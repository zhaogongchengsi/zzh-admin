import router from "@/routers/index.js";
import { store } from "@/store/index.js";

const whiteList = ["login"];
let asyncRouterFlag = 0;

export async function RouterDoorGod(to, from, next) {
  const _token = localStorage.getItem("token");

  if (whiteList.indexOf(to.name) > -1) {
    if (_token) {
      if (!asyncRouterFlag) {
        asyncRouterFlag++;
        await store.dispatch("router/SetAsyncRouter");
      }
      next();
    } else {
      next();
    }
  } else {
    if (_token) {
      if (!asyncRouterFlag) {
        asyncRouterFlag++;
        await store.dispatch("router/SetAsyncRouter");
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
}

router.beforeEach(RouterDoorGod);
