<script lang="ts" setup>
import TopTool from "../../components/TopTool.vue";
import Menus from "./Menus.vue";
import { useMenuStore } from "@/pinia";
const menuStore = useMenuStore();
</script>

<template>
  <div class="app-container">
    <blog-layout style="height: 100vh; border-radius: 5px;">
      <blog-layout-sider
        :theme="menuStore.theme"
        collapsible
        :collapsed="menuStore.collapsed"
        :hide-trigger="true"
      >
        <Menus />
      </blog-layout-sider>
      <blog-layout>
        <div :class="['app-content', {'app-content-dark': menuStore.theme==='dark'}]">
          <blog-layout-header>
            <TopTool />
          </blog-layout-header>
          <blog-layout-content>
            <div :class="['app-router', {'app-router-dark': menuStore.theme}]">
              <router-view />
            </div>
          </blog-layout-content>
        </div>
      </blog-layout>
    </blog-layout>
  </div>
</template>

<style lang="scss" scoped>
.app-container {
  --header-tools-height: 50px;
  --container-bg-color: #fff;
  --sider-width: 200px;
  width: 100%;
  height: 100vh;
  box-sizing: border-box;
  background-color: var(--color-menu-dark-bg);
  border-radius: 5px;
  overflow: hidden;
}

.app-router {
  border-top: 1px solid var(--color-neutral-3);
  height: calc(100vh - 52px);
  overflow: auto;
  &::-webkit-scrollbar {
    /*滚动条整体样式*/
    width: 5px; /*高宽分别对应横竖滚动条的尺寸*/
    height: 1px;
  }
  &::-webkit-scrollbar-thumb {
    /*滚动条里面小方块*/
    border-radius: 5px;
    box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2);
    background: #535353;
  }
  &::-webkit-scrollbar-track {
    /*滚动条里面轨道*/
    box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2);
    border-radius: 5px;
    background: #ededed;
  }
}
.app-content {
  background: var(--color-menu-light-bg);
  height: 100%;
}

.app-content-dark {
  background: var(--color-menu-dark-bg);
  box-shadow: 0 2px 5px 0 rgb(0 0 0 / 8%);
}
</style>
