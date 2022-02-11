import { Post, confirmStatus, Get } from "@/utils/service"
import type { user } from "@/types/request"
import type { userInfo, state, VerificationCode } from "@/types/response"
export async function login (user:user) :Promise<userInfo> {
       console.log(user)
  return new Promise<userInfo>((resolve, reject) =>{
    Post("user/login", user)
      .then((response) => {
        console.log(response);
        const state = confirmStatus(<state>response.data.state)
        if (state) {
          resolve(response.data.data)
        } else {
          reject(response.data.state)
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