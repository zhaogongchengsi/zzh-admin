<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import BaseMenus from "@/components/BaseMenus.vue"
import { useStore } from 'vuex'
import { DadLookSon} from '@/utils/index'
import { storeToRefs } from 'pinia'
import { userStore } from '@/store/user.js'
const userinfo = userStore()
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

const store = useStore()
const router = useRouter()
const menuData = ref([])




onMounted(() => {
  const { fullPath, href, name } = router.currentRoute.value
  const routerChildren = store.state.router.OriginlRoutData

  const currentRoute = routerChildren.find(r => {
    return r.Path === href && r.Name === name
  })

  if (currentRoute) {
    menuData.value = currentRoute.children
  }

})


const props = defineProps({
  asideWidth: {
    type: String,
    default: "200px"
  }
})

</script>

<template>
    <el-container>
      <el-aside :width="props.asideWidth">
        <div class="system_container" :style="{backgroundColor:user.baseColor}">
          <BaseMenus 
            :menu-list="menuData" 
            :router="true" 
            :background-color="user.baseColor" 
            :active-text-color="user.activeColor" 
            :text-color="user.baseTextColor"
          />
        </div>
      </el-aside>
          <el-main>
          <slot></slot>
      </el-main>
    </el-container>
</template>

<style lang="scss" scoped>
.system_container {
  height: 100%;
}
.system_container:deep(.el-menu) {
  border-right: 0;
}

</style>