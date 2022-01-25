<script setup>
import { getuserinfo } from "@/api/user.js"
import { onMounted } from 'vue';

onMounted(async () => {
    // const a = await getuserinfo()
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