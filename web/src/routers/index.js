import { createRouter, createWebHistory } from "vue-router";
const routes = [
  {
    path: "/login",
    name: "login",
    component: () => import("@/views/login/index.vue"),
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
