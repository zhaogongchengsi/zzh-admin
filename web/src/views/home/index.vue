<script setup>
import { MenuIcon } from '@heroicons/vue/solid'
import { onMounted, reactive, ref } from 'vue';
import { useStore } from 'vuex'
import BaseMenus from "@/components/BaseMenus.vue"
const store = useStore()

const menuList = ref([])
const asideState = ref(false)
const asideWidth = ref("200px")
const topNavs = ref([
    {
        Path:"/system",
        Label:"系统管理",
        Icon: "DesktopComputerIcon",
        router: "/"
    }
])

onMounted(async () => {
    menuList.value = store.state.router.OriginlRoutData
    console.log('menuList', menuList.value);
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

</script>



<template>
    <div class="home-container">
        <el-container>
            <el-aside :width="asideWidth">
                <div class="aside-container">
                    <div class="aside-logo">
                        <div class="cms-logo"><img src="/images/logo.png" alt="logo" /></div>
                        <div class="cms-title">穿越之超级管理系统</div>
                    </div>
                    <BaseMenus :menu-list="menuList" background-color="#333" text-color="#fff" />
                </div>
            </el-aside>
            <el-container>
                <el-header>
                    <div class="header-container">
                        <div class="header-left-x"> <MenuIcon class="h-10 w-10 text-Gray-900 cursor-pointer" /> </div>
                        <div class="header-center-x">
                            <BaseMenus :menu-list="topNavs" text-color="#333" mode="horizontal" @select="onSelect" />
                        </div>
                        <div class="header-right-x">
                            <el-dropdown>
                                <div class="header-avatar">
                                    <el-avatar
                                        src="/images/logo.png"
                                        fit="scale-down"
                                    ></el-avatar>
                                    <span class="user-name"> 超级管理员 </span>
                                </div>
                                <template #dropdown>
                                    <el-dropdown-menu>
                                        <el-dropdown-item>Action 1</el-dropdown-item>
                                        <el-dropdown-item>Action 2</el-dropdown-item>
                                        <el-dropdown-item>Action 3</el-dropdown-item>
                                        <el-dropdown-item>Action 4</el-dropdown-item>
                                    </el-dropdown-menu>
                                </template>
                            </el-dropdown>
                        </div>
                    </div>
                </el-header>
                <el-main>
                    <div class="main-container">
                        <div class="router-breadcrumb">
                            <el-breadcrumb separator="/">
                                <el-breadcrumb-item :to="{ path: '/' }">homepage</el-breadcrumb-item>
                                <el-breadcrumb-item>promotion list</el-breadcrumb-item>
                                <el-breadcrumb-item>promotion detail</el-breadcrumb-item>
                            </el-breadcrumb>
                        </div>
                        <router-view />
                    </div>
                </el-main>
                <el-footer>
                    <span> 假装这就是备案 </span>
                </el-footer>
            </el-container>
        </el-container>
    </div>
</template>

<style lang="scss" scoped>

.router-breadcrumb {
    padding: 10px 5px;
    box-sizing: border-box;
    border-bottom: 1px solid var(--el-border-color-base);
}

.home-container:deep(.el-main) {
   padding: 0;
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

.header-container {
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

.main-container {
   border-top: 1px solid var(--el-border-color-base);
}
</style>