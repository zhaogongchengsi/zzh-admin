<script setup>
import { MenuIcon } from '@heroicons/vue/solid'
import { onMounted, reactive, ref } from 'vue';
import { useStore } from 'vuex'
import Icons from '../../utils/Icons';
import BaseMenusVue from '../../components/BaseMenus.vue';
import { userStore } from '@/store/user.js'
import { useRouterStore } from '@/store/router.js'
import { storeToRefs } from 'pinia'
import { useRouter } from "vue-router"
const Router = useRouter()
const userinfo = userStore()
const routerInfo = useRouterStore()
// activeColor: "#1890ff"
// authorityId: "one"
// avatarImg: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
// baseColor: "#333"
// nickName: "系统用户"
// parentnodeid: "1"
// sideMode: "dark"
// userName: "admin"
// useradmin: "12345"
// baseTextColor
const { user } = storeToRefs(userinfo)
const { solids, outlines } = Icons;
const store = useStore()
const menuList = ref([])
const asideState = ref(false)


onMounted(async () => {
    let rootMenu = routerInfo.root
    console.log("11",routerInfo)
    menuList.value = rootMenu
})

const activeIndex = ref('1')
const handleSelect = (key, keyPath) => {
    console.log(key, keyPath)
}

const openAside = () => {
    asideState.value = !asideState.value
}

const onSelect = (index,path,item) => {
    console.log(index,path,item)
}
const logout = () => {
    localStorage.removeItem("z_router")
    localStorage.removeItem("z_user")
    localStorage.removeItem("token")
    Router.push("/login")
} 

</script>



<template>
    <div class="home-container">
            <el-container direction="vertical">
                <el-header>
                    <div class="header-container" :style="{backgroundColor:user.baseColor}">
                        <div class="header-left-logo">
                            <div class="header-logo-img"><img src="/images/logo.png" alt=""></div>
                            <span class="header-logo-title" :style="{color:user.baseTextColor}" >超级管理系统</span>
                        </div>
                        <div class="header-center-x" :style="{backgroundColor:user.baseColor}">
                            <BaseMenusVue 
                                :menu-list="menuList" 
                                mode="horizontal" 
                                :router="true" 
                                :background-color="user.baseColor" 
                                :active-text-color="user.activeColor" 
                                :text-color="user.baseTextColor"
                            />
                        </div>
                        <div class="header-right-x">
                            <el-dropdown>
                                <div class="header-avatar">
                                    <el-avatar
                                        :src="user.avatarImg"
                                        fit="contain"
                                    ></el-avatar>
                                    <span class="user-name"  :style="{color:user.baseTextColor}"> {{user.nickName}} </span>
                                </div>
                                <template #dropdown>
                                    <el-dropdown-menu>
                                        <el-dropdown-item>个人中心</el-dropdown-item>
                                        <el-dropdown-item>设置</el-dropdown-item>
                                        <el-dropdown-item>修改密码</el-dropdown-item>
                                        <el-dropdown-item :divided="true" @click="logout">注销登录</el-dropdown-item>
                                    </el-dropdown-menu>
                                </template>
                            </el-dropdown>
                        </div>
                    </div>
                </el-header>
                <el-main>
                    <el-scrollbar>
                        <div class="main-container">
                           <router-view />
                        </div>
                    </el-scrollbar>
                </el-main>
                <el-footer height="30px">
                    <div class="footer-container"><span> 假装这就是备案 </span></div>
                </el-footer>
            </el-container>
    </div>
</template>

<style lang="scss" scoped>
.main-container {
      --el-header-padding: 0 0px;
    --el-header-height: 60px;
}
.main-container:deep(.el-main) {
  height: calc(100vh - 60px - 30px);
  border-left: 1px solid var(--el-border-color-base);

}
.main-dom {
    height: 2000px;
}

.footer-container {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}

.header-left-logo {
    width: 200px;
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: space-evenly;
    .header-logo-img {
        width: 45px;
        padding: 10px;
        border:1px solid var(--el-border-color-base);
        border-radius: 50%;
        img {
            width: 100%;
            height: auto;
        }
    }

}

.router-breadcrumb {
    padding: 10px 5px;
    box-sizing: border-box;
    border-bottom: 1px solid var(--el-border-color-base);
}

.home-container:deep(.el-main) {
   padding: 0;
}

.home-container:deep(.el-header) {
    border-bottom: 1px solid var(--el-border-color-base);
      padding: 0 !important;
}

.home-container:deep(.el-footer) {
    border-top: 1px solid var(--el-border-color-base);
}

.aside-container {
    height: 100vh;
    background-color: #333;
    .aside-logo {
        text-align: center;
        height: 60px;
        display: flex;
        justify-content:center;
        align-items:center;
        .cms-logo {
            width: 30px;
            height: 30px;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            border: 1px solid var(--el-border-color-base);
            img {
                height: 50%;
                width: auto;
            }
        }

        .cms-title {
            padding: 0 15px;
            box-sizing: border-box;
            font-size: 13px;
            font-weight: 600;
            color: #fff;
        }
        
    }
}
 .header-center-x:deep(.el-menu--horizontal) {
        border-bottom: none !important;
}
.header-container {
    width: 100%;
    height: 60px;
    display: flex;
    justify-content: space-between;
    .header-left-x {
        height: inherit;
        display: flex;
        align-items:center;
        justify-content: center;
    }

    .header-center-x {
        flex: 1;
        padding: 0 50px;
    }

    .header-right-x {
        height: inherit;
        display: flex;
        align-items:center;
        justify-content: center;
        box-sizing: border-box;
        padding-right: 30px;
    }

    .header-avatar {
       display: flex;
       align-items: center;
       .user-name {
           margin-left: 10px;
       }
    }
}

.header-container:deep(.el-menu--horizontal) {
    height: 60px;
    box-sizing: border-box;
}

.main-container {
   height: calc(100vh - 90px);
//    border-top: 1px solid var(--el-border-color-base);
}
</style>