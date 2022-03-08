import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  { 
    path: "/", 
    name: "home",
    component: () => import("@/views/layout/index.vue"),
    redirect:'/home',
    children: [
      {
        path: "/article_list",
        name: "articleList",
        component: () => import("@/views/article/ArticleList.vue"),
      },
      {
        path: "/article_create",
        name: "articleCreate",
        component: () => import("@/views/article/ArticleCreate.vue"),
      },
      {
        path: "/home",
        name: "home",
        component: () => import("@/views/home/index.vue"),
      },
      {
        path: "/media_admin",
        name: "media_admin",
        component: () => import("@/views/MediaLib/MedAdm.vue"),
      },
      {
        path: "/media_upload",
        name: "media_upload",
        component: () => import("@/views/MediaLib/UpMed.vue"),
      },
      {
        path: "/tags",
        name: "tags",
        component: () => import("@/views/article/Tags.vue"),
      },
      {
        path: "/article_type",
        name: "article_type",
        component: () => import("@/views/article/CreateType.vue"),
      }
    ]
  },
  { 
    path: "/login", 
    name: "login",
    component: () => import("@/views/login/index.vue")
  },
];

const router = createRouter({
  // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
  history: createWebHashHistory(),
  routes, // `routes: routes` 的缩写
});

router.beforeEach((to,from) => {
  if (to.name !== 'login') {
    const token = localStorage.getItem('z_token');
    if (!token) {
      return { path: '/login',}
    }
  }
  return true;
})

export default router;
