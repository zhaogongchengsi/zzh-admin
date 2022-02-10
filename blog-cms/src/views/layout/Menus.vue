<script setup lang="ts">
import { ref } from "vue"
import {
  IconBulb,
} from '@arco-design/web-vue/es/icon';
import { useMenuStore } from '@/pinia'
const menuStore = useMenuStore()
const collapsed = ref(false)


</script>


<template>
  <div>
    <div class="logo-container">
      <div class="logo-img">
        <img src="/logo.svg" alt="logo" />
      </div>
      <div class="logo-title">
      <blog-typography-title :heading="6" :ellipsis="false">
        超级管理平台
      </blog-typography-title>
      </div>
    </div>
    <blog-menu
      :style="{ width: '200px', borderRadius: '4px' }"
      :theme="menuStore.theme" 
      :collapsed="menuStore.collapsed"
      :default-open-keys="['0']"
      :default-selected-keys="['0_2']"
    >
      <blog-sub-menu key="1">
        <template #icon><icon-bulb></icon-bulb></template>
        <template #title>Navigation 1</template>
        <blog-menu-item key="2_0">Menu 1</blog-menu-item>
        <blog-menu-item key="2_1">Menu 2</blog-menu-item>
      </blog-sub-menu>
      <blog-sub-menu key="2">
        <template #icon><icon-bulb></icon-bulb></template>
        <template #title>Navigation 2</template>
        <blog-menu-item key="2_0">Menu 1</blog-menu-item>
        <blog-menu-item key="2_1">Menu 2</blog-menu-item>
      </blog-sub-menu>
      <blog-sub-menu key="3">
        <template #icon><icon-bulb></icon-bulb></template>
        <template #title>Navigation 3</template>
        <blog-menu-item key="2_0">Menu 1</blog-menu-item>
        <blog-menu-item key="2_1">Menu 2</blog-menu-item>
        <blog-sub-menu key="2_2" title="Navigation 4">
          <blog-menu-item key="2_2_0">Menu 1</blog-menu-item>
          <blog-menu-item key="2_2_1">Menu 2</blog-menu-item>
        </blog-sub-menu>
      </blog-sub-menu>
    </blog-menu>
  </div>
</template>

<style scoped lang="scss">
.menu-demo {
  box-sizing: border-box;
  width: 100%;
  padding: 40px;
  background-color: var(--color-neutral-2);
}

.logo-container {
  height: var(--header-tools-height);
  display: flex;
  align-items: center;
  .logo-img {
    width: 48px;
    height: var(--header-tools-height);
    flex-shrink: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    img {
      width: 70%;
      height: 70%;
    }
  }

  .logo-title {
    width: calc(200px - 48px);
    height: var(--header-tools-height);
    flex-shrink: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 18px;
    color: #fff;
  }
}
</style>