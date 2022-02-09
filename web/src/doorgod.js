let asyncRouterFlag = false;
const whiteList = ["login"];

async function initRouter(router, pinia) {
  const token = localStorage.getItem("token");
  if (token) {
    await pinia.initRouter();
  } else {
    router.push("/login");
  }
}

export function doorgod(router, pinia) {
  router.beforeEach(async function (to, from) {
    if (!whiteList.includes(to.name)) {
      if (!asyncRouterFlag) {
        asyncRouterFlag = true;
        const useRouter = pinia();
        await initRouter(router, useRouter);
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
