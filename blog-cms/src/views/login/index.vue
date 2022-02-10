<template>
  <div :class="['login-container', {'login-container-dark': menuStore.theme}]">
    <toptool :menu="false" />
    <div class="login-from">
      <a-card hoverable :style="{ width: '1000px'}" :loading="false">
        <div class="from-box">
          <div class="from-img">1</div>
          <blog-form :model="userinfo" :style="{width:'50%', flex: '1'}" @submit="handleSubmit">
            <blog-form-item field="name" label="账号">
              <a-input v-model="userinfo.username" placeholder="请输入账号" />
            </blog-form-item>
            <blog-form-item field="post" label="密码">
              <a-input v-model="userinfo.password" placeholder="请输入密码" />
            </blog-form-item>
            <blog-form-item>
              <blog-button html-type="submit">登录</blog-button>
            </blog-form-item>
          </blog-form>
        </div>
      </a-card>
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { Ref } from 'vue'
import { ref  } from "vue"
import toptool from "@/components/TopTool.vue"
import { useMenuStore } from '@/pinia'
import { useInfoStore } from "@/pinia/user.ts"
import { useRouter } from 'vue-router'
const menuStore = useMenuStore()
const userStore = useInfoStore()
const router = useRouter()
interface user {
  username: string
  password: string
}
const userinfo : Ref<user> = ref({
  username: "",
  password: "",
})

const handleSubmit = async (data:user):Promise<void> => {
 const userStatus = await userStore.login(data)
 if (userStatus === true) {
   router.push({
      path: "/", 
      name: "home",
   })
 }
}
</script>

<style lang="scss" scoped>
  .login-container {
    width: 100%;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: space-between;
    // background-color: #ffffff;
    flex-direction: column;
    background: url(/bg1.jpg) no-repeat center;
  }

  .login-container-dark {
    background: var(--color-bg-2);
  }

  .login-from {
    padding: 30px;
    flex: 1;
    display: flex;
    align-items: center;
  }

  .from-box {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
    .from-img {
      width: 50%;
      height: 500px;
    }
  }

  .login-container:deep(.toptool-container) {
    width: 100%;
    justify-content: flex-end;
    border-bottom: 0;
    box-sizing: border-box;
    padding-top: 15px;
  }
</style>