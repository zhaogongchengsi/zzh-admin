<script setup>
import { ref, onMounted} from 'vue'
import BaseMenus from "@/components/BaseMenus.vue"
import { storeToRefs } from 'pinia'
import { userStore } from '@/store/user.js'
import { useRouterStore } from '@/store/router.js'
const userinfo = userStore()
const routerInfo = useRouterStore()
const { user } = storeToRefs(userinfo)
const menuData = ref([])
onMounted(() => {
  menuData.value = routerInfo.children
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