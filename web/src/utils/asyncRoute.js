import routerComponent from "@/views/routerHandel.vue";
import notComponent from "@/components/NotComponent.vue";
import pageRouter from "@/routers/index";
const routerFile = import.meta.glob("../views/**/*.{vue,js,ts,jsx,tsx}");
let routerfiles = filePathCompile(routerFile);

export function ParAndChildren(parents, children) {
  let _parents = JSON.parse(JSON.stringify(parents));
  let _children = JSON.parse(JSON.stringify(children));

  pasParAndChildren(_parents, _children);

  function pasParAndChildren(par, chi) {
    par.map((pitem) => {
      chi.map((citem, ci) => {
        if (citem.parent_id === pitem.ID) {
          // _children.splice(ci, 1);
          pasParAndChildren([citem], _children);
          if (pitem.children) {
            pitem.children.push(citem);
          } else {
            pitem.children = [citem];
          }
        }
      });
    });
  }

  return _parents;
}

export function separation(data) {
  let parents = data.filter((p) => p.parent_id === p.ID || p.parent_id === 0),
    children = data.filter((c) => c.parent_id !== c.ID);
  return {
    parents,
    children,
  };
}

//  替换文件路径
export function filePathCompile(files) {
  let newPaths = {};
  for (const file in files) {
    const pathreg = /\.\/([^.]+).*/i;
    const path = file.match(pathreg);
    newPaths[path[1]] = files[file];
  }
  return newPaths;
}

export function replaceComponent(routerTree) {
  return routerTree.map((router) => {
    let _component = routerfiles[router.component];
    if (_component) {
      router.component = _component;
    } else {
      if (router.children && router.children.length > 0) {
        let rootCom = routerfiles[router.component + "/index"];
        if (rootCom) {
          router.component = rootCom;
        } else {
          router.component = routerComponent;
        }
      } else {
        router.component = notComponent;
      }
    }
    if (router.children && router.children.length > 0) {
      router.children = replaceComponent(router.children);
    }
    return router;
  });
}

export function RouterDataToTree(data) {
  const { parents, children } = separation(data); // 分离根组件和非根组件
  return ParAndChildren(parents, children);
}

export function replaceRouter(originlRouter) {
  const asyncRouters = RouterDataToTree(originlRouter);
  const routers = replaceComponent(asyncRouters); // 替换路由组件
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
      ...routers,
    ],
    redirect: "/home_user", // 重定向到用户首页
  });
  return routers;
}
