import { Post, confirmStatus, Get } from "@/utils/service"
import type { user } from "@/types/request"
import type { userInfo, state, VerificationCode } from "@/types/response"
import { Message } from '@arco-design/web-vue';
export async function login (user:user) :Promise<userInfo> {
  return new Promise<userInfo>((resolve, reject) =>{
    Post("user/login", user)
      .then((response) => {
        const state = confirmStatus(<state>response.data.state)
        if (state) {
          resolve(response.data.data)
          Message.success({
            content: '登录成功'
          })
        } else {
          reject(response.data.state)
          Message.error({
            content: response.data.state.message
          })
        }
      })
      .catch(() => reject)
  })
}

export async function getVerifCode () :Promise<VerificationCode>{
  return new Promise<VerificationCode>((resolve, reject) =>{
    Get("verify_code")
    .then((response) => {
      const state = confirmStatus(<state>response.data.state)
      if (state) {
          resolve(response.data.data)
      } else {
        reject(response.data.state)
      }
    }).catch(() => reject)
  })
}