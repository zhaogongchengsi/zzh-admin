<template>
  <div :class="['login-container', {'login-container-dark': menuStore.theme}]">
    <toptool :menu="false" />
    <div class="login-from">
      <blog-spin :loading="islodin" dot >
        <a-card hoverable :style="{ width: '1000px'}">
          <div class="from-box">
              <div class="from-img">1</div>
              <blog-form :model="userinfo" :style="{width:'50%', flex: '1'}" @submit="handleSubmit">
                <blog-form-item field="name" label="账号">
                  <a-input v-model="userinfo.useradmin" placeholder="请输入账号" />
                </blog-form-item>
                <blog-form-item field="post" label="密码">
                  <a-input v-model="userinfo.password" placeholder="请输入密码" />
                </blog-form-item>
                <blog-form-item field="captcha" label="验证码">
                  <div class="code-box">
                    <a-input v-model="userinfo.captcha"  allow-clear placeholder="请输入验证码" />
                    <img :src="code.b64s" @click="getcode" />
                  </div>
                </blog-form-item>
                <blog-form-item>
                  <blog-button html-type="submit">
                    <a-space size="medium">
                      <span>登录</span>
                    </a-space>
                  </blog-button>
                </blog-form-item>
              </blog-form>
            </div>
        </a-card>
      </blog-spin>
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { Ref } from 'vue'
import { ref, reactive, onMounted,toRef } from "vue"
import toptool from "@/components/TopTool.vue"
import { useMenuStore } from '@/pinia'
import { useInfoStore } from "@/pinia/user"
import { useRouter } from 'vue-router'
import { user } from '@/types/request'
import { VerificationCode } from "@/types/response"
import { getVerifCode } from "@/api/user"
const menuStore = useMenuStore()
const userStore = useInfoStore()
const router = useRouter()
const userinfo : Ref<user> = ref({
  useradmin: "root",
  password: "12345",
  captcha: "",
  captchaId: ""
})
const code: Ref<VerificationCode>=ref({
  id: "",
  b64s: "",
})
const islodin = ref(false)
const getcode = async () => {
  const res = await getVerifCode()
  code.value = res
  userinfo.value.captchaId = res.id
}

onMounted(() => {
  getcode()
})

const handleSubmit = async (data:user):Promise<void> => {
  islodin.value = true
 const userStatus = await userStore.loginState({...userinfo.value})
 if (userStatus === true) {
   router.push({
      path: "/", 
      name: "home",
   })

 } else {
   getcode()
 }
 islodin.value = false
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

  .code-box {
    display:flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    img {
      margin-left: 30px;
      height: 32px;
      cursor: pointer;
    }
  }
</style>