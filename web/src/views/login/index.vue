<script setup>
import { reactive, onMounted, ref } from "vue"
import { GetVerifyCode, Login } from "@/api/index.js"
import { useRouter } from "vue-router"
import rules from "./validator.js"
import { ElMessage } from 'element-plus'
import { userStore } from '@/store/user.js'
import { useRouterStore } from '@/store/router.js'
const user = userStore() // pinia
const useR = useRouterStore()
const formLabelAlign = reactive({ useradmin: 'root',  password: '12345', captcha: ''})
const verifyCodd = ref({})
const from = ref(null)
const Router = useRouter()


async function getVerifyCode () {
     let vaCode = await GetVerifyCode()
    verifyCodd.value = vaCode
}

onMounted(() => {
   getVerifyCode()
})

const refreshCode = () => {
    getVerifyCode()
}

const onLogin = () => {
    from.value.validate(async (fromState) => {                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 
        if (fromState) {
            try {
                const loginState = await Login({...formLabelAlign, captchaId: verifyCodd.value.id })
                localStorage.setItem("z_user", JSON.stringify(loginState.user.user))
                user.$patch({
                    user: loginState.user.user,
                    token: loginState.token
                })
                await useR.initRouter("login")
                // const routerState = await store.dispatch("router/SetAsyncRouter")
                if (loginState.success) {
                    Router.push("/home_user")
                } else {
                    console.log('登录成功但是路由初始化失败');
                }
            } catch (e) {
                 getVerifyCode()
            }
        } else {
             ElMessage.error('用户信息不完整')
        }
    })
}

</script>

<template>
  <div class="login-container">
      <div class="login-title"> 穿越之超级管理系统 </div>
      <div class="login-box">
          <div class="login-img"> <img src="/images/login.svg" alt="登录" /> </div>
          <div class="login-from">
                <el-form
                    label-position="right"
                    label-width="100px"
                    :model="formLabelAlign"
                    :rules="rules"
                    ref="from"
                >
                    <el-form-item label="账号" prop="useradmin">
                        <el-input v-model="formLabelAlign.useradmin"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password">
                         <el-input v-model="formLabelAlign.password" type="password" ></el-input>
                    </el-form-item>
                    <el-form-item label="验证码" prop="captcha">
                        <div class="form-group">
                            <el-input v-model="formLabelAlign.captcha"></el-input>
                            <div class="form-code" @click="refreshCode"> <img :src="verifyCodd.b64s" alt="验证码" /> </div>
                        </div>
                    </el-form-item>
                </el-form>
                <div class="form-submit"> <el-button @click="onLogin" type="primary">登录</el-button> </div>
          </div>
      </div>
  </div>
</template>



<style scoped lang="scss">

.login-container {
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;

    .login-title {
        color: #333;
        font-size:30px;
        font-weight: bold;
    }

    .login-box {
        width: 1000px;
        height: 500px;
        border-radius: 5px;
        display: flex;
        .login-img {
            flex: 1;
            display: flex;
            justify-content: center;
            align-items: center;
            img {
                width: 100%;
                height: auto;
            }
        }

        .login-from {
            flex: 1;
            display: flex;
            align-items: center;
            flex-direction: column;
            justify-content: center;
        }
        
    }
}

.form-group {
    display: flex;
    align-items: center;
    .form-code {
        width: 200px;
        border: 1px solid var(--el-border-color-base);
        box-sizing: border-box;
        border-radius: 5px;
        margin-left: 15px;
        height: 40px;
        padding: 3px;
        box-sizing: border-box;
        cursor: pointer;
        img {
            width: 100%;
            height: 100%;
        }
    }
}

.form-submit {
    width: 100%;
    box-sizing: border-box;
    padding-left: 100px;
}

.form-submit:deep(.el-button) {
    width: 100%;
}

</style>