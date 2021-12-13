import { GetRouter } from '@/api/routerComponents.js'
import { useRouter } from 'vue-router'
const router = useRouter()

export async function initRouter () {
    const _router = await GetRouter()
    router.addRoute({
        name: 'home',
        path: '/',
        component: () => import('@/views/home/index.vue'),
        children: _router,
    });
}