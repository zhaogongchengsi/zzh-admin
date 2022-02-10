import { defineStore } from "pinia";

export const useInfoStore = defineStore("user", {
  state: () => {
    return {
      userinfo: {},
      token: '',
    };
  },
  actions: {
   async login (user) {
    localStorage.setItem("z_token", "123123908osdifoifhas09gfu23fi")
    return true;
   }
  },
});
