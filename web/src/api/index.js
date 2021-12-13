
import {Get, Post} from "@/service/index.js"
import { ElMessage } from 'element-plus'
import { JudgeRequestStatus } from '@/utils'
export function GetVerifyCode() {
    return new Promise(function(resolve, reject) {
        Get("/verify_code").then(res => {
            JudgeRequestStatus(res.state)(
                () => {
                    resolve(res.data)
                },
                () => {
                    ElMessage.error(res.message)
                    reject({})
                }
            )
        }).catch(reject)
    })
}

export function Login (user) {
    return new Promise(function(resolve, reject) {
        Post("/user/login", user).then(res => {
            JudgeRequestStatus(res.state)(
                () => {
                    ElMessage({
                        message: '登录成功',
                        type: 'success',
                    })
                    localStorage.setItem('token', res.data.token)
                    resolve({
                        success: true,
                        user: res.data
                    })
                },
                () => {
                    ElMessage.error('登录失败，用户名或密码错误')
                    reject({})
                }
            )
        }).catch(reject)
    })
}