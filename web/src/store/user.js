import { Login } from "@/api/index.js";
import { defineStore } from "pinia";

export const userStore = defineStore("main", {
  state: () => {
    return {
      user: {},
      token: "",
    };
  },
  actions: {
    async login(user) {
      try {
        const loginState = await Login(user);
        this.user = loginState.user.user;
        this.token = loginState.user.token;
        return loginState;
      } catch (err) {
        console.error("登录错误", err);
      }
    },
  },
});
