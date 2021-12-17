<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import BaseMenus from "@/components/BaseMenus.vue"
const router = useRouter()
const menuData = ref([])




onMounted(() => {
  console.log("当前路由",router.currentRoute.value)
  console.log("全局组件获取的路由数据",router.getRoutes())
  const originRouter = router.getRoutes()
  const currentRouter = router.currentRoute.value

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
        <BaseMenus :menu-list="menuData" />
      </el-aside>
          <el-main>
          <slot></slot>
      </el-main>
    </el-container>
</template>

