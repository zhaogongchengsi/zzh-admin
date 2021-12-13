import { createRouter, createWebHistory } from 'vue-router'



const routes = [
  { 
      path: '/login', 
      component: () => import("@/views/login/index.vue")
  },
  {
      path: '/',
      name: "home",
      component: () => import("@/views/home/index.vue"),
      children:[
        {
          path: 'menus',
          name: "Menus",
          component: () => import("@/views/system/Menus.vue")
        }
      ]
  },
  {
      path: '/:catchAll(.*)',
      name: "error",
      component: () => import("@/views/error/index.vue")
  },

]


const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router;