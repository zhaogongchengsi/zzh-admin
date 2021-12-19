import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/login",
    component: () => import("@/views/login/index.vue"),
  },
  {
    path: "/",
    name: "BaseLayou",
    component: () => import("@/views/Layou/index.vue"),
    children: [
      {
        path: "system",
        name: "System",
        component: () => import("@/views/system/index.vue"),
        children: [
          {
            path: "menus",
            name: "Menus",
            component: () => import("@/views/system/Menus.vue"),
          },
          {
            path: "users",
            name: "Users",
            component: () => import("@/views/system/Users.vue"),
          },
        ],
      },
    ],
  },
  {
    path: "/:catchAll(.*)",
    name: "error",
    component: () => import("@/views/error/index.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
