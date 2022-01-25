<script setup>
import { onMounted } from 'vue';
import { useRouterStore } from '@/store/router.js'
// const useRouter = useRouterStore()
onMounted(async () => {
    // await useRouter.initRouter("首页内");
    // console.log(1, a)
})


</script>


<template>
    <div class="home-container">
        <el-row>
            <el-col :span="12"><div class="top-welcome">欢迎111</div></el-col>
            <el-col :span="12"><div class="top-welcome">欢迎</div></el-col>
        </el-row>
    </div>
</template>

<style lang="scss" scoped>

.home-container {
   padding: 10px;
   box-sizing: border-box;
}


</style>