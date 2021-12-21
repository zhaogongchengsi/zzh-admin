<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import BaseMenus from "@/components/BaseMenus.vue"
import { useStore } from 'vuex'
import { DadLookSon} from '@/utils/index'
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
    console.log(currentRoute.children)
  }



  // const _r =  DadLookSon(
  //   function (data) {
  //       if (data.Name === name && data.Path === href) {
  //           return function(son) {
  //             if (data.ID === son.ParentId) {
  //               return {...son,Path:`${data.Path}/${son.Path}`}
  //             } else {
  //               return false
  //             }
  //           }
  //       }
  //     return false
  //   },
  //   rootData,
  //   routerChildren
  // )
  // console.log('_r', _r);
  // menuData.value = _r
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
        <div class="system_container">
          <BaseMenus :menu-list="menuData" :router="true" />
        </div>
      </el-aside>
          <el-main>
          <slot></slot>
      </el-main>
    </el-container>
</template>

<style lang="scss" scoped>
.system_container:deep(.el-menu) {
  border-right: 0;
}

</style>