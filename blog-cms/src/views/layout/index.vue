<script lang="ts" setup>
import TopTool from "./TopTool.vue"

</script>

<template>
  <div class="app-container">
    <blog-layout style="height: 100vh;">
      <blog-layout-sider theme="dark">Sider</blog-layout-sider>
      <blog-layout>
        <div class="app-content">
          <blog-layout-header> <TopTool /></blog-layout-header>
          <blog-layout-content>
            <router-view />
          </blog-layout-content>
        </div>
      </blog-layout>
    </blog-layout>
  </div>
</template>

<style lang="scss" scoped>
.app-container {
  --header-tools-height: 30px;
  --container-bg-color: #fff;
  width: 100%;
  height: 100vh;
  box-sizing: border-box;
  background-color: var(--color-menu-dark-bg);
}
.app-content {
  background-color: #fff;
  height: 100%;
}
</style>
