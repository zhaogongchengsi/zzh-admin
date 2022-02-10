import { defineStore } from "pinia";

export const useMenuStore = defineStore("menu", {
  state: () => {
    return {
      collapsed: false,
      theme: "light"
    };
  },
  actions: {
   toggleCollapse () {
      this.collapsed = !this.collapsed;
   },
   setTheme (theme:string) {
    this.theme = theme;
    if (theme === "dark") {
      document.body.setAttribute('arco-theme', 'dark')
    } else if (theme === "light") {
      document.body.removeAttribute('arco-theme');
    }
   }
  },
});
