import { reactive } from 'vue'
import { routerCompile } from '@/utils/routerToMenus.js'


const routerMod = {
    menuRouter : reactive({
        menus: []
    }),
    setMenus (newMenus) {
        this.menuRouter.menus = newMenus
    },
    getMenus () {

        // routerCompile(this.menuRouter.menus)

        return this.menuRouter.menus
    }
}


export default routerMod;

