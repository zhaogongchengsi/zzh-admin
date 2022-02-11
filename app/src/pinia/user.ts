import { defineStore } from "pinia";
import { login } from "@/api/user";
import { user } from "@/types"
import router from "@/router"
export const useInfoStore = defineStore("user", {
  state: () => {
      const userinfo = localStorage.getItem("z_user")
      const token = localStorage.getItem("z_token")
      if (userinfo && token) {
        return {
          userinfo: JSON.parse(userinfo),
          token
        }
      } else {
        router.push({
          path: "/login",
          name: "login",
        })
        return {
          userinfo: {},
          token: '',
        };
      }
  },
  actions: {
   async loginState (userinfo:user):Promise<boolean> {
    try {
      const u = await login(userinfo)
      this.token = u.token;
      this.userinfo = u.user;
      localStorage.setItem("z_token", u.token)
      localStorage.setItem("z_user", JSON.stringify(u.user))
      return true;
    } catch (err) {
      console.error('login err:',err)
      return false;
    }
   }
  },
});
