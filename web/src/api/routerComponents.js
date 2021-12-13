import { Get } from "@/service/index.js"
import { dataToTree, copyRouter, filePathCompile } from "@/utils/asyncRoute.js"
const routerFile = import.meta.glob("../views/**/*.{vue,js,ts,jsx,tsx}");



export function GetRouter () {
    return new Promise(function (resolve, reject) {
        Get('/menu/get_menu')
        .then(function (response) {
            /*
             * 请求过来的路由数据 树形化之后 把里面所需要的路由数据拷贝出来 
             */
         let routerTree = copyRouter(dataToTree(response.data))
            /*
             * 将文件内的组件取出来 
             */
         let routerfiles = filePathCompile(routerFile)
            /*
             * 将树形路由数据中的组件替换成真实组件 
             */
         let _routerTree = replaceRouter(routerTree, routerfiles)
            resolve(_routerTree)
        })
        .catch(function (error) {
            reject(error)
        })
    })
}

function replaceRouter (routerTree, fileTree) {
  return  routerTree.map((router) => {
        let _component = fileTree[router.component]
        if (_component) {
            router.component = _component
        } else {
            delete router.component
        }
        if (router.children && router.children.length > 0) {
            router.children = replaceRouter(router.children, fileTree)
        }
        return router
    })
}

