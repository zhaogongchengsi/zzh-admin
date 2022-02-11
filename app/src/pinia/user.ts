import { defineStore } from "pinia";
import { login } from "@/api/user";
import { user } from "@/types"
export const useInfoStore = defineStore("user", {
  state: () => {
    return {
      userinfo: {},
      token: '',
    };
  },
  actions: {
   async loginState (userinfo:user):Promise<boolean> {
    try {
      const u = await login(userinfo)
      this.token = u.token;
      this.userinfo = u.user;
      localStorage.setItem("z_token", u.token)
      return true;
    } catch (err) {
      console.error('login err:',err)
      return false;
    }
   }
  },
});
